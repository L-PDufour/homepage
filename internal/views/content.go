package views

import (
	"strconv"

	"homepage/internal/database"
	"homepage/internal/models"
	"homepage/internal/utils"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func ContentList(props models.ContentProps) g.Node {
	items := make([]g.Node, 0, len(props.Content))
	for _, content := range props.Content {
		items = append(items, contentItem(content, props))
	}

	var bottomActions g.Node
	if len(props.Content) > 1 && props.IsAdmin {
		bottomActions = h.Div(h.Class("content-actions"), NewContentButton())
	}

	return h.Div(h.Class("content-page"),
		h.Div(h.Class("content-panel"),
			h.Div(h.Class("content-panel-inner"),
				h.Div(h.Class("content-list"), g.Group(items)),
				g.If(bottomActions != nil, bottomActions),
			),
		),
	)
}

func contentItem(content database.Content, props models.ContentProps) g.Node {
	var body g.Node

	htmlContent, err := utils.GetHTMLContent(content.Content.String)
	switch {
	case err != nil:
		body = h.P(h.Class("content-error"), g.Text("Error rendering content."))
	case len(htmlContent) > 300 && len(props.Content) > 1:
		body = g.Raw(utils.TruncateMarkdown(htmlContent, 300))
	default:
		var adminActions g.Node
		if props.IsAdmin {
			adminActions = h.Div(h.Class("content-admin-actions not-prose"),
				EditButton(content.ID),
				DeleteButton(content.ID),
			)
		}
		body = g.Group([]g.Node{
			g.Raw(htmlContent),
			g.If(adminActions != nil, adminActions),
		})
	}

	var readMore g.Node
	if len(props.Content) > 1 {
		readMore = h.Div(h.Class("content-footer"), ReadMoreButton(content.ID))
	}

	return h.Div(
		h.ID("content-"+strconv.Itoa(int(content.ID))),
		h.Class("content-card"),
		h.Div(h.Class("content-card-body"),
			h.Div(h.Class("prose"), body),
		),
		g.If(readMore != nil, readMore),
	)
}
