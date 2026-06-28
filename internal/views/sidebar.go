package views

import (
	"fmt"

	"homepage/internal/database"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func GetRecentBlogPosts(contents []database.Content, count int) []database.Content {
	var blogPosts []database.Content
	for _, content := range contents {
		if content.ContentType == "blog" {
			blogPosts = append(blogPosts, content)
			if len(blogPosts) == count {
				break
			}
		}
	}
	return blogPosts
}

func Sidebar(contents []database.Content) g.Node {
	posts := GetRecentBlogPosts(contents, 5)
	items := make([]g.Node, 0, len(posts))
	for _, content := range posts {
		items = append(items, h.Li(
			Button(content.Title, "blue", fmt.Sprintf("/content/view?id=%d", content.ID)),
		))
	}
	return h.Aside(h.Class("sidebar"), g.Attr("aria-label", "Recent blog posts"),
		h.Div(h.Class("sidebar-inner"),
			h.H2(h.Class("sidebar-title"), g.Text("Recent Posts")),
			h.Ul(h.Class("sidebar-list"), g.Group(items)),
		),
	)
}
