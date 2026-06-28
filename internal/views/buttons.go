package views

import (
	"strconv"

	g "maragu.dev/gomponents"
	ds "maragu.dev/gomponents-datastar"
	h "maragu.dev/gomponents/html"
)

// Button is a plain navigation link styled as a button.
func Button(text, color, action string) g.Node {
	return h.A(
		h.Href(action),
		h.Class("btn btn--"+color),
		g.Text(text),
	)
}

// ReadMoreButton navigates to the full content view — plain anchor.
func ReadMoreButton(contentID int64) g.Node {
	id := strconv.FormatInt(contentID, 10)
	return h.A(
		h.Href("/content?id="+id),
		h.Class("link-btn link-btn--green not-prose"),
		g.Text("Agrandir"),
	)
}

// EditButton navigates to the update form — plain anchor.
func EditButton(contentID int64) g.Node {
	id := strconv.FormatInt(contentID, 10)
	return h.A(
		h.Href("/content/update?id="+id),
		h.Class("link-btn link-btn--blue not-prose"),
		g.Text("Edit"),
	)
}

// NewContentButton navigates to the new content form — plain anchor.
func NewContentButton() g.Node {
	return h.A(
		h.Href("/content/new"),
		h.Class("link-btn link-btn--blue"),
		g.Text("Add"),
	)
}

func FormButton() g.Node {
	return g.Group([]g.Node{
		h.Button(h.Type("submit"), h.Class("link-btn link-btn--blue"), g.Text("Submit")),
		h.A(h.Href("javascript:history.back()"), g.Text("Cancel")),
	})
}

// DeleteButton stays on Datastar — it patches/removes a fragment in place, not a page nav.
func DeleteButton(contentID int64) g.Node {
	id := strconv.FormatInt(contentID, 10)
	return h.Button(
		ds.On("click",
			"confirm('Are you sure you want to delete this content?') && @delete('/content?id="+id+"')"),
		h.Class("link-btn link-btn--red"),
		g.Text("Delete"),
	)
}
