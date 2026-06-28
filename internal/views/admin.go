package views

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Adminpage() g.Node {
	return h.Doctype(
		h.HTML(
			h.Body(
				h.H1(g.Text("Admin Page")),
				h.P(g.Text("This is a protected admin page.")),
				h.A(h.Href("/bio"), g.Text("Back to Home")),
			),
		),
	)
}

func AdminAuthPage() g.Node {
	return h.Doctype(
		h.HTML(
			h.Body(
				h.H1(g.Text("Authentication Successful")),
				h.P(g.Text("You have been successfully authenticated. You can now access the admin page.")),
				h.A(h.Href("/admin"), g.Text("Go to Admin Page")),
				h.Script(g.Raw(`
					setTimeout(function() {
						window.location.href = "/admin";
					}, 3000);
				`)),
			),
		),
	)
}
