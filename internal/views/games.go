package views

import (
	"homepage/internal/games"

	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func GamesPage(list []games.Game) g.Node {
	items := make([]g.Node, 0, len(list))
	for _, game := range list {
		items = append(items, h.Div(h.Class("content-card"),
			h.Div(h.Class("content-card-body"),
				h.Div(h.Class("prose"),
					h.H2(h.A(h.Href(game.URL()), g.Text(game.Title))),
					h.P(g.Text(game.Description)),
				),
			),
		))
	}

	return Base(h.Div(h.Class("content-page"),
		h.Div(h.Class("content-panel"),
			h.Div(h.Class("content-panel-inner"),
				h.Div(h.Class("content-list"), g.Group(items)),
			),
		),
	))
}

func GamePage(game games.Game) g.Node {
	return Base(h.Div(h.Class("content-page"),
		h.Div(h.Class("content-panel"),
			h.Div(h.Class("prose"),
				h.H1(g.Text(game.Title)),
				h.P(g.Text(game.Description)),
			),
			gameEmbed(game),
		),
	))
}

func gameEmbed(game games.Game) g.Node {
	switch game.Runtime {
	case games.Emscripten:
		return emscriptenEmbed(game.AssetsPath())
	default:
		return h.P(h.Class("content-error"), g.Text("Ce jeu n'est pas disponible."))
	}
}

func emscriptenEmbed(basePath string) g.Node {
	return g.Group([]g.Node{
		h.Div(h.Class("game-frame"),
			h.Canvas(
				h.Class("game-canvas"),
				h.ID("canvas"),
				g.Attr("oncontextmenu", "event.preventDefault()"),
				h.TabIndex("-1"),
			),
		),
		h.Script(g.Raw(`
			var canvasElement = document.getElementById('canvas');
			canvasElement.addEventListener('webglcontextlost', (e) => {
				alert('WebGL context lost. You will need to reload the page.');
				e.preventDefault();
			}, false);

			var Module = {
				canvas: canvasElement,
				locateFile: (path) => '` + basePath + `/' + path,
				totalDependencies: 0,
				monitorRunDependencies(left) {
					this.totalDependencies = Math.max(this.totalDependencies, left);
				}
			};

			window.onerror = (event) => {
				console.error('Emscripten runtime error:', event);
			};
		`)),
		h.Script(h.Type("text/javascript"), h.Src(basePath+"/game.js"), g.Attr("async")),
	})
}
