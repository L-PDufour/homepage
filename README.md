# Page Personnelle de Léon-Pierre Dufour

## Description
Ce projet est mon site web personnel, conçu pour présenter mon profil professionnel et mes compétences en développement. Il est construit avec une stack technologique moderne et innovante, mettant en avant mes compétences en Go, développement web, et configuration système.

## Technologies utilisées
- Go templ : Pour la génération de templates HTML
- HTMX : Pour des interactions dynamiques côté client
- Nix : Pour la gestion de l'environnement de développement et des dépendances

## Structure du projet
- `flake.nix` : Déclaration de l'environnement de développement et gestion des dépendances

## Fonctionnalités
- Présentation de mon profil professionnel
- Mise en avant de mes compétences techniques
- (Autres fonctionnalités à venir)

## Installation et exécution locale(WIP)
1. Assurez-vous d'avoir Nix installé sur votre système et d'avoir une connexion avec postgress
2. Clonez ce dépôt
3. Naviguez dans le dossier du projet
4. Créer un fichier .env
```
PORT=
DB_HOST=
DB_PORT=
DB_NAME=
DB_USER=
DB_PASSWORD=
CONN_STRING=
GO_ENV="development"
ADMIN_EMAIL=
```
5. `nix run .#watch`
6. Se diriger vers la page avec votre navigateur au port spécifié


## Développement
Ce projet est en cours de développement. Les prochaines étapes incluent :
- Un outils pour planifier la semaine et les repas

