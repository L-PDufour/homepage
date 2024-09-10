package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
)

var (
	cfTeamDomain = os.Getenv("CF_TEAM_DOMAIN")
	cfAudTag     = os.Getenv("CF_AUD_TAG")
	isProduction = os.Getenv("GO_ENV") == "production"
)

type AuthenticatedUser struct {
	Email   string
	IsAdmin bool
}

type Authenticator struct {
	verifier *oidc.IDTokenVerifier
}

func NewAuthenticator() (*Authenticator, error) {
	if !isProduction {
		log.Println("Running in development mode, mock authenticator will be used.")
		// Return an Authenticator with a nil verifier for development
		return &Authenticator{verifier: nil}, nil
	}

	ctx := context.Background()
	providerURL := fmt.Sprintf("https://%s.cloudflareaccess.com", cfTeamDomain)
	provider, err := oidc.NewProvider(ctx, providerURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create OIDC provider: %v", err)
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID: cfAudTag,
	})

	return &Authenticator{verifier: verifier}, nil
}

func (a *Authenticator) VerifyToken(r *http.Request) (*AuthenticatedUser, error) {
	if !isProduction {
		return a.mockVerifyToken(r)
	}

	log.Println("Pas ici")
	cookie, err := r.Cookie("CF_Authorization")
	if err != nil {
		return nil, fmt.Errorf("no CF_Authorization cookie found")
	}

	token, err := a.verifier.Verify(r.Context(), cookie.Value)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	var claims struct {
		Email string `json:"email"`
	}
	if err := token.Claims(&claims); err != nil {
		return nil, fmt.Errorf("failed to parse claims: %v", err)
	}

	isAdmin := claims.Email == "leonpierre.dufour@gmail.com" // Replace with your admin email or logic

	return &AuthenticatedUser{
		Email:   claims.Email,
		IsAdmin: isAdmin,
	}, nil
}

func (a *Authenticator) mockVerifyToken(r *http.Request) (*AuthenticatedUser, error) {
	email := "leonpierre.dufour@gmail.com" // Your admin email logic
	isAdmin := true

	user := &AuthenticatedUser{
		Email:   email,
		IsAdmin: isAdmin,
	}

	// Set the user in the request context
	ctx := context.WithValue(r.Context(), "user", user)
	*r = *r.WithContext(ctx) // Update the request context

	return user, nil
}
