package views

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Plan() g.Node {
	return Base(
		h.Div(h.Class("content-page"),
			h.Div(h.Class("content-panel"),
				h.Div(h.Class("content-card"),
					h.Div(h.Class("prose")),
					h.P(g.Text("À venir")),
				),
			),
		),
	)
}
