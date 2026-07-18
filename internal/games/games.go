// Package games is the registry of WebAssembly games served by the site.
//
// Each game lives under assets/games/<slug>/ and ships the files produced by
// its build. For Emscripten builds that means the glue script game.js plus
// game.wasm (and game.data when the build embeds a filesystem); the glue
// script fetches its sibling files relative to its own URL, so no extra
// routing is needed. To add a new game: drop its build output in a new
// asset folder and append an entry to All.
package games

// Runtime identifies which loader a game needs in the browser.
type Runtime int

const (
	// Emscripten games are bootstrapped by their game.js glue script.
	Emscripten Runtime = iota
)

type Game struct {
	Slug        string
	Title       string
	Description string
	Runtime     Runtime
}

// AssetsPath is the URL prefix the game's files are served from.
func (g Game) AssetsPath() string {
	return "/assets/games/" + g.Slug
}

// URL is the page the game is playable on.
func (g Game) URL() string {
	return "/games/" + g.Slug
}

var All = []Game{
	{
		Slug:        "chip8",
		Title:       "CHIP-8",
		Description: "Émulateur CHIP-8 compilé en WebAssembly avec Emscripten.",
		Runtime:     Emscripten,
	},
}

func BySlug(slug string) (Game, bool) {
	for _, g := range All {
		if g.Slug == slug {
			return g, true
		}
	}
	return Game{}, false
}
