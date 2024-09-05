package handler

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"homepage/internal/blog"
	"homepage/internal/database"
	"homepage/internal/markdown"
	"homepage/internal/views"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Handler struct {
	DB          *database.Queries
	MD          markdown.MarkdownConverter
	BlogService *blog.BlogService
}
type AuthenticatedUser struct {
	Email   string
	IsAdmin bool
}

var authenticatedUsers = sync.Map{}
var isProduction = os.Getenv("GO_ENV") == "production"

func NewHandler(db *database.Queries, md markdown.MarkdownConverter, blogService *blog.BlogService) *Handler {
	return &Handler{
		DB:          db,
		MD:          md,
		BlogService: blogService,
	}
}

func (h *Handler) BlogPage(w http.ResponseWriter, r *http.Request) {
	component := views.Blog()
	component.Render(r.Context(), w)
}

func (h *Handler) GetBlogPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.BlogService.ListBlogPosts(r.Context())
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	// Check if the request is for the sidebar (recent posts)
	if r.URL.Query().Get("recent") == "true" {
		component := views.RecentPosts(posts)
		component.Render(r.Context(), w)
	} else {
		// For the main content area
		component := views.BlogPostList(posts)
		component.Render(r.Context(), w)
	}
}

func (h *Handler) GetBlogPost(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid blog ID", http.StatusBadRequest)
		return
	}

	post, htmlContent, err := h.BlogService.GetBlogPost(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Error fetching post", http.StatusInternalServerError)
		return
	}

	component := views.BlogContent(post, htmlContent)
	component.Render(r.Context(), w)
}

func (h *Handler) NewBlogPostForm(w http.ResponseWriter, r *http.Request) {
	component := views.BlogPostForm()
	component.Render(r.Context(), w)
}

// This handler needs to change to store Markdown content
func (h *Handler) CreateBlogPost(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	content := r.FormValue("content") // This is now Markdown content

	post, err := h.BlogService.CreateBlogPost(r.Context(), title, content)
	if err != nil {
		http.Error(w, "Error creating post", http.StatusInternalServerError)
		return
	}

	// Convert the Markdown content to HTML for display
	htmlContent, err := h.BlogService.MarkdownService.ConvertAndSanitize(post.Content)
	if err != nil {
		http.Error(w, "Error processing post content", http.StatusInternalServerError)
		return
	}

	component := views.BlogContent(post, htmlContent)
	component.Render(r.Context(), w)
}

func (h *Handler) Admin(w http.ResponseWriter, r *http.Request) {
	user, err := getAuthenticatedUser(r)
	log.Println("admin")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if !user.IsAdmin {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	views.Adminpage().Render(context.Background(), w)
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	user, err := getAuthenticatedUser(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	views.Homepage(user.IsAdmin, user.Email).Render(context.Background(), w)
}
func getAuthenticatedUser(r *http.Request) (*AuthenticatedUser, error) {
	log.Println("getAuth")
	if isProduction {
		return verifyCloudflareJWT(r)
	}
	// For local development, we'll simulate Cloudflare Access

	return simulateCloudflareAccess(r)
}

func verifyCloudflareJWT(r *http.Request) (*AuthenticatedUser, error) {
	teamDomain := os.Getenv("CF_TEAM_DOMAIN")
	ctx := context.Background()

	log.Printf("Using team domain: %s", teamDomain)

	providerURL := fmt.Sprintf("https://%s.cloudflareaccess.com", teamDomain)
	log.Printf("OIDC Provider URL: %s", providerURL)

	provider, err := oidc.NewProvider(ctx, providerURL)
	if err != nil {
		log.Printf("Error creating OIDC provider: %v", err)
		return nil, fmt.Errorf("failed to create OIDC provider: %v", err)
	}

	cookie, err := r.Cookie("CF_Authorization")
	if err != nil {
		log.Printf("No CF_Authorization cookie found: %v", err)
		return nil, fmt.Errorf("no CF_Authorization cookie found: %v", err)
	}

	log.Printf("CF_Authorization Cookie: %s", cookie.Value)

	verifier := provider.Verifier(&oidc.Config{ClientID: os.Getenv("CF_AUD_TAG")})
	token, err := verifier.Verify(ctx, cookie.Value)
	if err != nil {
		log.Printf("Failed to verify token: %v", err)
		return nil, fmt.Errorf("failed to verify token: %v", err)
	}

	var claims struct {
		Email string `json:"email"`
	}

	if err := token.Claims(&claims); err != nil {
		log.Printf("Failed to parse claims: %v", err)
		return nil, fmt.Errorf("failed to parse claims: %v", err)
	}

	log.Printf("Claims: %+v", claims)

	if user, ok := authenticatedUsers.Load(claims.Email); ok {
		log.Printf("User found in cache: %s", claims.Email)
		return user.(*AuthenticatedUser), nil
	}

	isAdmin := len(claims.Email) > 10 && claims.Email[len(claims.Email)-10:] == "@gmail.com"
	user := &AuthenticatedUser{
		Email:   claims.Email,
		IsAdmin: isAdmin,
	}

	authenticatedUsers.Store(claims.Email, user)
	log.Printf("New user authenticated and stored: %s, Admin: %v", claims.Email, isAdmin)

	return user, nil
}

func simulateCloudflareAccess(r *http.Request) (*AuthenticatedUser, error) {
	email := r.Header.Get("X-Simulated-Cloudflare-Access-Email")
	log.Printf("Simulated email: %s", email)
	if email == "" {
		return nil, fmt.Errorf("no authentication information provided")
	}

	if user, ok := authenticatedUsers.Load(email); ok {
		return user.(*AuthenticatedUser), nil
	}

	isAdmin := len(email) > 10 && email[len(email)-10:] == "@gmail.com"
	user := &AuthenticatedUser{
		Email:   email,
		IsAdmin: isAdmin,
	}

	authenticatedUsers.Store(email, user)
	return user, nil
}
