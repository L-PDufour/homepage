-- +goose Up
-- +goose StatementBegin
UPDATE content
SET markdown =
'## Dev.lpdufour.xyz

J''ai commencé ce site web après ma première année d''apprentissage de la programmation. J''avais envie d''essayer des technologies web modernes et différentes afin de créer une expérience utilisateur simple mais efficace. Le code source est disponible sur [GitHub](https://github.com/L-PDufour/homepage) et voici un [article](/content?id=19) où j''explore quelques concepts plus en détails.

### Technologies explorées :
- **Backend :** *Go* (bibliothèque standard), me permettant de comprendre les bases d''un serveur web.
- **Frontend :**
  - *Templ* et *Tailwind CSS* pour découvrir le rendu côté serveur et le styling moderne.
  - *HTMX* pour expérimenter avec l''interactivité côté client de manière simple.
- **Déploiement :** Initiation à *Docker* pour la conteneurisation et à *Nix* pour la gestion de l''environnement.'
WHERE content_type = 'project'::content_type AND title = 'website';

UPDATE content
SET markdown =
'## Probono

Au courant de la dernière année, j''ai eu l''opportunité de contribuer au développement d''une application web pour l''inscription et la gestion des bénévoles de la Guignolée du Centre de pédiatrie sociale de Québec. Ce projet collaboratif a été une excellente occasion d''appliquer mes connaissances dans un contexte réel.

### Technologies découvertes :
- **Backend :** Initiation à *NestJS* et *Swagger*, découvrant ainsi les frameworks backend modernes.
- **Frontend :** Premiers pas avec *React* et *Typescript* pour comprendre le développement d''interfaces utilisateur.
- **Déploiement :** Découverte de *Docker* et de ses principes de base.

Ce projet, toujours en développement, marque mes débuts dans un environnement de développement collaboratif. Il m''a permis d''appliquer mes connaissances théoriques à un cas pratique, tout en apprenant de mes collègues plus expérimentés.'
WHERE content_type = 'project'::content_type AND title = 'probono';

UPDATE content
SET markdown =
'## Divers

Mon profil [GitHub](https://github.com/L-PDufour/) témoigne de mon parcours d''apprentissage. Il regroupe divers projets, exercices et tutoriels que j''ai réalisés pour développer mes compétences.

### Langages explorés et apprentissages :
- **C** : Découverte des bases de la programmation à travers de petits exercices algorithmiques.
- **Go** : Apprentissage de la création d''APIs simples et de serveurs web basiques.
- **Python** : Initiation à la programmation orientée objet via de petits projets guidés.'
WHERE content_type = 'project'::content_type AND title = 'github';

UPDATE content
SET markdown =
'## À propos de moi


Bonjour, je m''appelle Léon-Pierre Dufour. Je suis un apprenti développeur enthousiaste, en pleine reconversion professionnelle. Depuis plus d''un an, je me plonge dans le monde passionnant du développement logiciel, avec un intérêt particulier pour Go, le développement web, et l''environnement Linux.

En-dehors du code, ma passion est de comprendre le fonctionnement des objets du quotidien. Cette curiosité me pousse à créer et à expérimenter avec divers projets pratiques :

- Fabrication d''une guitare électrique
- Projets d''électronique, comme la fabrication d''un clavier split
- Rénovation domiciliaire et ébénisterie

Ces expériences pratiques nourrissent ma créativité et renforcent ma capacité à résoudre des problèmes, des compétences que j''applique dans mon apprentissage de la programmation. Ma curiosité me pousse également à explorer et à personnaliser mes outils de travail, notamment Neovim et NixOS, ce qui me permet de mieux comprendre l''environnement de développement.


## Compétences en développement

Depuis le début de mon parcours en programmation, j''ai exploré et travaillé avec :

- Programmation en Go (découverte des concepts de base et création de petites applications)
- Développement web (HTML, CSS, JavaScript, introduction à React)
- Environnement Linux et initiation à NixOS
- Notions de base en gestion de version avec Git
- Découverte des méthodologies Agile

## Expérience professionnelle

Depuis environ un an, je contribue bénévolement au développement d''une plateforme d''inscription en ligne pour un OBNL. Ce projet me permet d''appliquer mes nouvelles connaissances en développement web dans un contexte réel et d''apprendre auprès de développeurs plus expérimentés.

Auparavant, j''ai travaillé 10 ans à titre de facteur chez Postes Canada, où j''ai développé d''excellentes compétences en gestion du temps, service client et résolution de problèmes. Ces compétences se sont avérées précieuses dans ma transition vers la programmation.

## Objectifs professionnels

Je suis à la recherche d''opportunités pour débuter ma carrière en tant que développeur junior. Mon objectif est de rejoindre une équipe où je pourrai continuer à apprendre, contribuer à des projets concrets, et progresser dans mes compétences techniques. Je suis particulièrement intéressé par les postes axés sur le développement backend, le développement web et la programmation embarquée.'
WHERE content_type = 'bio'::content_type AND title = 'bio';

UPDATE content
SET markdown =
'## Une recette pour ne pas l''oublier parce que j''oublie d''utiliser mes favoris
[Zucchini Bread Recipe - Sally''s Baking Addiction](https://sallysbakingaddiction.com/zucchini-bread/)

**Prep Time**: 20 minutes
**Cook Time**: 45-55 minutes
**Total Time**: 1 hour, 15 minutes
**Yield**: 1 loaf (8 slices)

## Ingredients

- 1 and 1/2 cups (190g) all-purpose flour (spooned & leveled)
- 1/2 teaspoon baking powder
- 1/2 teaspoon baking soda
- 1/2 teaspoon ground cinnamon
- 1/2 teaspoon ground nutmeg
- 1/4 teaspoon salt
- 1/2 cup (115g) unsalted butter, melted
- 3/4 cup (150g) packed light or dark brown sugar
- 1 large egg, at room temperature
- 1/2 cup (120g) plain yogurt, at room temperature
- 2 teaspoons pure vanilla extract
- 1 and 1/2 cups (230g) grated zucchini (about 1 medium zucchini)
- Optional: 3/4 cup (130g) semi-sweet chocolate chips, chopped nuts, or raisins

## Instructions

1. **Preheat Oven**: Preheat the oven to 350°F (177°C) and grease a 9×5 inch (or 8×4 inch) loaf pan.

2. **Dry Ingredients**: In a large bowl, whisk the flour, baking powder, baking soda, cinnamon, nutmeg, and salt together.

3. **Wet Ingredients**: In a medium bowl, whisk the melted butter, brown sugar, egg, yogurt, and vanilla extract together until combined.

4. **Combine Wet & Dry**: Pour the wet ingredients into the dry ingredients and gently fold them together. Do not over-mix.

5. **Add Zucchini**: Gently fold in the grated zucchini. If adding chocolate chips, nuts, or raisins, fold them in as well.

6. **Pour Batter**: Pour the batter into the prepared loaf pan.

7. **Bake**: Bake for 45-55 minutes or until a toothpick inserted into the center of the loaf comes out clean.

8. **Cool**: Allow the bread to cool in the pan on a wire rack for 15 minutes. Then remove the bread from the pan and place it on a wire rack to cool completely before slicing.

## Notes

- **Make-Ahead & Freezing Instructions**: Zucchini bread can be frozen for up to 3 months. Thaw overnight in the refrigerator and bring to room temperature before serving.
- **Optional Add-Ins**: You can add 3/4 cup of chocolate chips, chopped nuts, or raisins to the batter.

Enjoy your zucchini bread!'
WHERE content_type = 'blog'::content_type AND title = 'recette';

UPDATE content
SET markdown =
'## Introduction à NixOS (Partie 1) WIP

## Caractéristiques principales

### Gestion déclarative

Avec NixOS, l''installation de paquets se fait de manière déclarative dans le fichier `configuration.nix`. Voici un exemple avec les paquets `wget`, `neovim` et `git`, qui seront installés lors de la construction du système. Il suffit de mettre à jour la configuration et de reconstruire le système et pour les enlever, ils suffit de les supprimer et de reconstruire le système :

#### Exemple NixOS

```nix
{ config, pkgs, ... }:

{
  environment.systemPackages = with pkgs; [
    wget
    neovim
    git
  ];
}
```
```
{ config, pkgs, ... }:

{
  environment.systemPackages = with pkgs; [
  ];
}

```
```bash
sudo nixos-rebuild switch
```

#### Exemple Debian

Avec APT, l''installation de paquets est impérative et se fait via des commandes shell. Voici un exemple typique :
```bash
sudo apt-get update
sudo apt-get install neovim
sudo apt-get remove neovim
```

## Isolation des paquets dans NixOS

NixOS utilise un système unique de gestion des paquets qui permet une isolation complète des logiciels. Cela évite les conflits de dépendances courants dans d''autres distributions Linux. Voici quelques points clés pour illustrer cette différence :

### 1. Installation dans le Nix Store

Contrairement à la plupart des distributions Linux, où les paquets sont installés dans des chemins communs (comme `/usr/bin`), NixOS installe tous les paquets dans un répertoire spécifique appelé le **Nix Store**. Ce répertoire se trouve généralement à `/nix/store`.

Chaque paquet a son propre chemin unique, ce qui signifie que deux versions différentes du même logiciel peuvent coexister sans conflit.

### 2. Dépendances isolées

Chaque paquet dans NixOS inclut ses propres dépendances, ce qui permet d''éviter les problèmes de compatibilité. Cela signifie que les mises à jour d''un paquet n''affectent pas les autres.

### 3. Comparaison avec Debian

Dans une distribution comme Debian, les paquets partagent des chemins communs. Par exemple, l''installation de `neovim` pourrait écraser une version précédente ou entraîner des conflits de dépendances.

#### Exemple NixOS

Pour voir où `neovim` est installé dans NixOS, vous pouvez utiliser la commande suivante :

```bash
which neovim
```

Cela affichera un chemin comme :

```
/nix/store/xxxxxxxx-neovim-0.10.1/bin/neovim
```

#### Exemple Debian

En revanche, dans Debian, la commande `which` pour `neovim` donnerait un chemin commun comme :

```bash
which neovim
```

Cela pourrait afficher :

```
/usr/bin/neovim
```

## Conclusion

NixOS est une distribution puissante pour ceux qui cherchent à gérer leur système de manière déclarative et reproductible. Que vous soyez développeur ou administrateur système, NixOS offre des outils robustes pour répondre à vos besoins.

## Références

- [Documentation officielle de NixOS](https://nixos.org/manual)
- [GitHub de NixOS](https://github.com/NixOS/nixpkgs)
- [Communauté NixOS sur Reddit](https://www.reddit.com/r/NixOS)'
WHERE content_type = 'blog'::content_type AND title = 'nixos_intro';

UPDATE content
SET markdown =
'## Mon Site Personnel : Un Laboratoire Web Innovant


## 🎯 Objectif du Projet
- Mise en pratique de compétences acquises : CI/CD, bases de données, backend, frontend
- Création d''une plateforme dynamique avec gestion de contenu en temps réel
- Implémentation d''une authentification sécurisée via Cloudflare Access
- Exploration et maîtrise de technologies émergentes (HTMX, Tailwind CSS, Templ)

## 🛠 Stack Technologique

### Backend et Frontend : Golang
Choisi pour sa performance exceptionnelle et sa simplicité élégante
#### Outils de Développement
[Air](https://github.com/air-verse/air) [Templ](https://github.com/a-h/templ) [Sqlc](https://sqlc.dev/) [Goose](https://github.com/pressly/goose):  Un ensemble d''outils puissants pour le développement Go, du rechargement en temps réel à la gestion de base de données.

### Frontend : HTMX
Interactions backend-frontend fluides pour des pages ultra-réactives

### Authentification : Cloudflare Access
Génération de tokens JWT pour une sécurité sans faille

### Environnement : Nix et NixOS
Garantie de reproductibilité et portabilité de l''environnement de développement

## 🏋️ Défis Techniques Relevés
- **Backend** : Gestion de la concurrence et mise en cache du contenu pour des performances maximales.
- **Frontend** : Combinaison de Go Templ, HTMX, et TailwindCSS pour une expérience utilisateur fluide et réactive.
- **DevOps** : Environnement de développement reproductible, conteneurisation efficace, mise en place d''un CI/CD avec déploiement simplifié via NixOS.
- **Sécurité** : Utilisation de Cloudflare Access pour l''authentification.

## 📊 Performances
**[GTmetrix](https://gtmetrix.com/reports/dev.lpdufour.xyz/cqFuzynR/#)** **[PageSpeed](https://pagespeed.web.dev/analysis/https-dev-lpdufour-xyz/0r8n2a6bra?form_factor=desktop)** **[WebPageTest](https://www.webpagetest.org/result/240926_AiDc67_AB3/)**

## 💼 Compétences Démontrées
- Maîtrise de Golang pour le développement full-stack
- Conception et implémentation d''APIs RESTful robustes
- Adoption et intégration de technologies de pointe (HTMX, Nix)
- Gestion de bases de données et migrations
- Déploiement et maintenance d''un site web

## 👁️ Explorer le Projet
🔗 [Code Source sur GitHub](https://github.com/L-PDufour/homepage)

##### Exemple de golang pour le backend

```go
// routes.go
// Configuration de la route pour récupérer le contenu
mux.HandleFunc("/content", s.Handler.GetContent())

// handlers.go
// Handler pour récupérer le contenu par ID
func (h *Handler) GetContent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "", http.StatusNotFound)
			return
		}
		props, err := h.Service.GetContentById(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		h.renderContentList(w, r, props)
	}
}

// utils.go
// Fonction pour convertir le contenu de markdown à html avec mise en cache
func GetHTMLContent(markdownContent string) (string, error) {
	cacheKey := hash(markdownContent)

	cacheMutex.RLock()
	cached, exists := htmlCache[cacheKey]
	if exists && time.Since(cached.Timestamp) < cacheExpiration {
		htmlCache[cacheKey] = cached
		cacheMutex.RUnlock()
		return cached.HTML, nil
	}
	cacheMutex.RUnlock()

	htmlContent, err, _ := singleflightGroup.Do(cacheKey, func() (interface{}, error) {
		return convertAndSanitize(markdownContent)
	})

	if err != nil {
		return "", err
	}

	cacheMutex.Lock()
	htmlCache[cacheKey] = models.CachedHTML{
		HTML:       htmlContent.(string),
		Timestamp:  time.Now(),
		LastAccess: time.Now(),
	}
	cacheMutex.Unlock()

	return htmlContent.(string), nil
}
```
##### Exemple de go templ, HTMX et tailwindcss

```go
// Exemple d''un bouton avec go templ et HTMX
templ ReadMoreButton(contentID int32) {
	<button
		hx-get={ "/content?id=" + strconv.Itoa(int(contentID)) }
		hx-target="#main-body"
		hx-swap="innerHTML settle:1s"
		hx-push-url={ "/content?id=" + strconv.Itoa(int(contentID)) }
		class="text-text underline px-4 py-2 hover:text-green transition duration-200 ease-in-out"
	>
		Agrandir
	</button>
}

// Template de base avec intégration de TailwindCSS et HTMX
templ Base(child templ.Component) {
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8"/>
			<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
			<meta name="description" content="Page personnelle de Léon-Pierre Dufour"/>
			<link href="/assets/css/output.css" rel="stylesheet"/>
			<link href="/assets/favicon.ico" rel="icon"/>
			<title>Page personnelle</title>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/pdf.js/4.6.82/pdf.min.mjs" type="module"></script>
			<script src="https://cdnjs.cloudflare.com/ajax/libs/pdf.js/4.6.82/pdf.worker.min.mjs" type="module"></script>
			<script src="/assets/js/htmx.min.js" defer></script>
		</head>
		<body id="main-body" class="flex flex-col min-h-screen bg-surface1">
			@header()
			<main class="flex-grow py-4">
				if child != nil {
					@child
				} else {
					{ children... }
				}
			</main>
			@footer()
		</body>
	</html>
}

```

##### Exemple de Nix et la Magie de la Déclarativité

Avec l''aide de [direnv](https://direnv.net/), chaque fois que j''accède à mon projet, mon environnement se crée automatiquement avec tous les outils nécessaires. Ainsi, chaque projet dispose de son propre environnement unique, sans que rien ne soit installé globalement.

Nix est également un langage de programmation qui s''accompagne d''outils puissants. Dans l''exemple ci-dessous, lorsque je lance la commande `nix run #container`, cela crée une image Docker contenant mon fichier binaire ainsi que ses dépendances et son environnement.

```nix
// flake.nix
// Configuration de l''environnement de développement
devShells.default = pkgs.mkShell {
  buildInputs = with pkgs; [
    go
    go-tools
    air
    gomod2nix.packages.${system}.default
    templ.packages.${system}.templ
    tailwindcss
    postgresql
  ];
};

// Configuration pour la création d''une image Docker
packages.container = pkgs.dockerTools.buildLayeredImage {
  name = "ldufour/goserver";
  tag = "latest";
  contents = [
    packages.default
    pkgs.postgresql
    pkgs.cacert
  ];
  config = {
    Cmd = [ "${packages.default}/bin/api" ];
    WorkingDir = "/";
    Volumes = {
      "/data" = { };
    };
    Env = [
      "IN_CONTAINER=true"
      "DB_HOST="
      "DB_PORT="
      "DB_NAME="
      "DB_USER="
      "DB_PASSWORD="
    ];
  };
};
```
Sur mon serveur, je peux déclarer l''image pour héberger mon site web comme suit :

```nix
///configuration.nix
virtualisation.oci-containers = {
  backend = "docker";
  containers = {
    homepage = {
      image = "docker.io/ldufour/goserver:latest";
      autoStart = true;
      ports = [ "80:80" ];
      environment = {
        "PORT" = "80";
      };
      extraOptions = [
        "--pull=always"
        "--add-host=host.docker.internal:host-gateway"
      ];
      volumes = [
        "/var/lib/goserver/data:/app/data"
        "/var/lib/goserver/config:/app/config"
      ];
    };
  };
};
```'
WHERE content_type = 'blog'::content_type AND title = 'website';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
