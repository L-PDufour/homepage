package views

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

var skills = []string{"Go", "C", "TypeScript", "Datastar", "PostgreSQL", "Nix", "Linux"}
var softSkills = []string{"Communication", "Travail d'équipe", "Résolution de problèmes"}
var interests = []string{"Ébénisterie", "Électronique", "Informatique"}

type cvJob struct {
	Title   string
	Period  string
	Body    []string
	Striped bool
}

var jobs = []cvJob{
	{
		Title:  "Développeur — Amérifor",
		Period: "2025 - Aujourd'hui",
		Body:   []string{"Développeur, responsable de l'ensemble du cycle de développement : conception, implémentation, déploiement et maintenance des applications."},
	},
	{
		Title:   "Centre de pédiatrie sociale de Québec",
		Period:  "10/2023 - Aujourd'hui",
		Body:    []string{"Participation bénévole au développement d'une application web pour l'inscription et la gestion des bénévoles pour la guignolée."},
		Striped: true,
	},
	{
		Title:  "Postes Canada",
		Period: "2013 - 2023",
		Body:   []string{"Gérer, acheminer le courrier et les colis en plus d'assurer un bon service à la clientèle."},
	},
	{
		Title:   "Agent de recherche",
		Period:  "2010",
		Body:    []string{"Analyser et rédiger des données statistiques et des rapports de recherche."},
		Striped: true,
	},
}

type cvEducation struct {
	Title   string
	Period  string
	Bullets []string
	Striped bool
}

var education = []cvEducation{
	{
		Title:  "Certificat en informatique — Université Laval",
		Period: "01/2025 - Aujourd'hui",
	},
	{
		Title:  "Boot.dev",
		Period: "2024",
		Bullets: []string{
			"Achèvement du cursus principal sur le développement backend",
			"Pratique de divers concepts passant de la programmation orientée objet aux bases de données",
			"Réalisation de projets pratiques intégrant les concepts appris",
		},
		Striped: true,
	},
	{
		Title:  "École 42",
		Period: "02/2023 - Aujourd'hui",
		Bullets: []string{
			"Introduction à la programmation en C",
			"Acquisition de compétences en algorithmique et en programmation système",
		},
	},
	{
		Title:  "Baccalauréat en développement social et analyse des problèmes sociaux — UQAR",
		Period: "2010",
		Bullets: []string{
			"Développement de compétences analytiques et de recherche",
			"Rédaction de rapports et présentation de résultats de recherche",
		},
		Striped: true,
	},
}

func tagList(items []string) g.Node {
	tags := make([]g.Node, 0, len(items))
	for _, item := range items {
		tags = append(tags, h.Span(h.Class("cv-tag"), g.Text(item)))
	}
	return h.Div(h.Class("cv-tag-list"), g.Group(tags))
}

func ResumeTemplate() g.Node {
	return Base(
		h.Div(h.ID("printArea"), h.Class("cv-page"),
			h.Div(h.Class("cv-banner"),
				h.H1(h.Class("cv-name"), g.Text("Léon-Pierre Dufour")),
				h.Section(h.Class("cv-summary"),
					h.P(g.Text("Développeur passionné, seul développeur au sein d'une PME où je conçois, implémente et maintiens des solutions logicielles de bout en bout. Autodidacte rigoureux, j'aime explorer en profondeur les technologies que j'utilise, du backend Go jusqu'à l'administration système avec Nix.")),
				),
			),
			h.Div(h.Class("cv-body"),
				h.Div(h.Class("cv-sidebar"),
					h.Header(h.Class("cv-contact"),
						h.P(h.Class("cv-label"), g.Text("Courriel")),
						h.P(h.Class("cv-value"), h.A(h.Href("mailto:leon@lpdufour.xyz"), g.Text("leon@lpdufour.xyz"))),
						h.P(h.Class("cv-label"), g.Text("Site")),
						h.P(h.Class("cv-value"), h.A(h.Href("https://dev.lpdufour.xyz"), h.Target("_blank"), g.Text("dev.lpdufour.xyz"))),
						h.P(h.Class("cv-label"), g.Text("Github")),
						h.P(h.Class("cv-value"), h.A(h.Href("https://github.com/l-pdufour"), h.Target("_blank"), g.Text("github.com/l-pdufour"))),
						h.P(h.Class("cv-label"), g.Text("Linkedin")),
						h.P(h.Class("cv-value"), h.A(h.Href("https://linkedin.com/in/l-pdufour"), h.Target("_blank"), g.Text("linkedin.com/in/l-pdufour"))),
					),
					h.Section(h.Class("cv-section"),
						h.H2(h.Class("cv-section-title"), g.Text("Compétences techniques")),
						tagList(skills),
					),
					h.Section(h.Class("cv-section"),
						h.H2(h.Class("cv-section-title"), g.Text("Compétences transversales")),
						tagList(softSkills),
					),
					h.Section(h.Class("cv-section"),
						h.H2(h.Class("cv-section-title"), g.Text("Intérêts")),
						tagList(interests),
					),
				),
				h.Div(h.Class("cv-main"),
					h.Section(h.Class("cv-section"),
						h.H2(h.Class("cv-section-title cv-section-title--lg"), g.Text("Expérience professionnelle")),
						g.Group(renderJobs(jobs)),
					),
					h.Section(h.Class("cv-section"),
						h.H2(h.Class("cv-section-title cv-section-title--lg"), g.Text("Formation")),
						g.Group(renderEducation(education)),
					),
				),
			),
			h.Div(h.Class("cv-print-bar print-hide"),
				h.Button(
					g.Attr("onclick", "window.print()"),
					h.Class("cv-print-btn"),
					g.Text("Imprimer en PDF"),
				),
			),
		),
	)
}

func renderJobs(items []cvJob) []g.Node {
	nodes := make([]g.Node, 0, len(items))
	for _, j := range items {
		class := "cv-entry"
		if j.Striped {
			class += " cv-entry--striped"
		}
		bodyNodes := make([]g.Node, 0, len(j.Body))
		for _, b := range j.Body {
			bodyNodes = append(bodyNodes, h.P(h.Class("cv-entry-body"), g.Text(b)))
		}
		nodes = append(nodes, h.Div(h.Class(class),
			h.H3(h.Class("cv-entry-title"), g.Text(j.Title)),
			h.P(h.Class("cv-entry-period"), g.Text(j.Period)),
			g.Group(bodyNodes),
		))
	}
	return nodes
}

func renderEducation(items []cvEducation) []g.Node {
	nodes := make([]g.Node, 0, len(items))
	for _, e := range items {
		class := "cv-entry"
		if e.Striped {
			class += " cv-entry--striped"
		}
		var bullets g.Node
		if len(e.Bullets) > 0 {
			items := make([]g.Node, 0, len(e.Bullets))
			for _, b := range e.Bullets {
				items = append(items, h.Li(g.Text(b)))
			}
			bullets = h.Ul(h.Class("cv-entry-list"), g.Group(items))
		}
		nodes = append(nodes, h.Div(h.Class(class),
			h.H3(h.Class("cv-entry-title"), g.Text(e.Title)),
			h.P(h.Class("cv-entry-period"), g.Text(e.Period)),
			g.If(bullets != nil, bullets),
		))
	}
	return nodes
}
