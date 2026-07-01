package views

import (
	"strconv"

	"homepage/internal/database"

	g "maragu.dev/gomponents"
	ds "maragu.dev/gomponents-datastar"
	h "maragu.dev/gomponents/html"
)

func contentTypeOptions() g.Node {
	return g.Group([]g.Node{
		h.Option(h.Value(""), g.Text("Select Content Type")),
		h.Option(h.Value("blog"), g.Text("Blog Post")),
		h.Option(h.Value("project"), g.Text("Project")),
		h.Option(h.Value("bio"), g.Text("Bio")),
	})
}

func NewForm() g.Node {
	return h.Form(
		ds.Signals(map[string]any{
			"type":    "",
			"title":   "",
			"content": "",
		}),
		h.Class("content-form"),
		h.Select(ds.Bind("type"), h.Name("type"), h.Class("form-input"), h.Required(), contentTypeOptions()),
		h.Input(ds.Bind("title"), h.Type("text"), h.Name("title"), h.Placeholder("Title"), h.Class("form-input"), h.Required()),
		h.Textarea(ds.Bind("content"), h.Name("markdown"), h.Rows("10"), h.Placeholder("Content"), h.Class("form-input"), h.Required()),
		h.Input(h.Type("url"), h.Name("image_url"), h.Placeholder("Image URL"), h.Class("form-input")),
		h.Input(h.Type("url"), h.Name("link"), h.Placeholder("Link (for projects)"), h.Class("form-input")),
		NewFormButton(),
	)
}

func EditForm(content database.Content) g.Node {
	id := strconv.Itoa(int(content.ID))
	return h.Form(
		ds.Signals(map[string]any{
			"type":    string(content.ContentType),
			"title":   content.Title,
			"content": content.Content.String,
		}),
		h.Class("content-form"),
		h.Select(ds.Bind("type"), h.Name("type"), h.Class("form-input"), h.Required(), contentTypeOptions()),
		h.Input(h.Type("hidden"), h.Name("id"), h.Value(id)),
		h.Input(ds.Bind("title"), h.Name("title"), h.Class("form-input")),
		h.Textarea(ds.Bind("content"), h.Name("markdown"), h.Rows("10"), h.Class("form-input")),
		FormButton(id),
	)
}
