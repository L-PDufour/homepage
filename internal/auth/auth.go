package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cloudflare/cloudflare-go"
	"github.com/coreos/go-oidc/v3/oidc"
)

var (
	CfTeamDomain = os.Getenv("CF_TEAM_DOMAIN")
	cfAudTag     = os.Getenv("CF_AUD_TAG")
	IsProduction = os.Getenv("GO_ENV") == "production"
	cfAPIKey     = os.Getenv("CF_API_KEY")
	adminEmail   = os.Getenv("ADMIN_EMAIL")
)

type AuthenticatedUser struct {
	Email   string
	IsAdmin bool
}

type Authenticator struct {
	verifier      *oidc.IDTokenVerifier
	cfAPI         *cloudflare.API
	cloudflareCtx context.Context
}

func NewAuthenticator() (*Authenticator, error) {
	var cfAPI *cloudflare.API = nil
	var err error = nil
	if IsProduction {
		cfAPI, err = cloudflare.NewWithAPIToken(cfAPIKey)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize Cloudflare API client: %v", err)
		}
	}
	cfCtx := context.Background()

	if !IsProduction {
		return &Authenticator{verifier: nil, cfAPI: cfAPI, cloudflareCtx: cfCtx}, nil
	}

	oidcCtx := context.Background()
	providerURL := fmt.Sprintf("https://%s.cloudflareaccess.com", CfTeamDomain)
	provider, err := oidc.NewProvider(oidcCtx, providerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create OIDC provider: %v", err)
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID: cfAudTag,
	})

	return &Authenticator{
		verifier:      verifier,
		cfAPI:         cfAPI,
		cloudflareCtx: cfCtx,
	}, nil
}

func (a *Authenticator) VerifyToken(r *http.Request) (*AuthenticatedUser, error) {
	cookie, err := r.Cookie("CF_Authorization")
	if err != nil {
		log.Printf("Failed to retrieve CF_Authorization cookie: %v", err)
		return nil, fmt.Errorf("no CF_Authorization cookie found")
	}

	log.Printf("CF_Authorization cookie found: %s", cookie.Value)

	token, err := a.verifier.Verify(r.Context(), cookie.Value)
	if err != nil {
		log.Printf("Failed to verify token: %v", err)
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	var claims struct {
		Email string `json:"email"`
	}
	if err := token.Claims(&claims); err != nil {
		log.Printf("Failed to parse claims: %v", err)
		return nil, fmt.Errorf("failed to parse claims: %v", err)
	}

	isAdmin := claims.Email == adminEmail
	return &AuthenticatedUser{
		Email:   claims.Email,
		IsAdmin: isAdmin,
	}, nil
}

// RedirectToLogin sends the user to the Cloudflare Access login page

// mockVerifyToken is a mock function used in non-production environments for token verification
func (a *Authenticator) MockVerifyToken(r *http.Request) (*AuthenticatedUser, error) {
	// In non-production, simulate an admin user
	email := adminEmail
	isAdmin := true

	// Create the user object
	user := &AuthenticatedUser{
		Email:   email,
		IsAdmin: isAdmin,
	}

	// Set the user in the request context
	ctx := context.WithValue(r.Context(), "user", user)
	*r = *r.WithContext(ctx) // Update the request context with the authenticated user

	return user, nil
}
