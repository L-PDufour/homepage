package views

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Chip8Page() g.Node {
	return h.HTML(
		h.Lang("en"),
		h.Head(
			h.Meta(h.Charset("utf-8")),
		),
		h.Body(
			g.Group([]g.Node{Chip8Canvas()}),
		),
	)
}

func Chip8Canvas() g.Node {
	return g.Group([]g.Node{
		h.Div(
			h.Class("emscripten_border"),
			h.Canvas(
				h.Class("emscripten"),
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
				totalDependencies: 0,
				monitorRunDependencies(left) {
					this.totalDependencies = Math.max(this.totalDependencies, left);
				}
			};

			window.onerror = (event) => {
				console.error('Emscripten runtime error:', event);
			};
		`)),
		h.Script(h.Type("text/javascript"), h.Src("/assets/chip/game.js"), g.Attr("async")),
	})
}
