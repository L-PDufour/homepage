package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"homepage/internal/handler"
)

func TestGetBlogPost(t *testing.T) {
	h := &handler.Handler{}

	tests := []struct {
		name           string
		id             string
		expectedStatus int
		expectedBody   string
	}{
		{"Valid ID", "1", http.StatusOK, "Blog post content"},
		{"Invalid ID", "999", http.StatusNotFound, "Blog post not found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/blog/"+tt.id, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			mux := http.NewServeMux()
			mux.HandleFunc("/blog/", h.GetBlogPost)

			mux.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}
