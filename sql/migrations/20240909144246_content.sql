-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS content (
    id SERIAL PRIMARY KEY,
    type TEXT NOT NULL,
    title TEXT NOT NULL,
    markdown TEXT,         -- Field for storing Markdown content
    image_url TEXT,        -- Optional field for storing image URLs
    link TEXT,             -- Optional field for storing links
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_title_per_type UNIQUE (type, title)  -- Ensure uniqueness of title per type
);

INSERT INTO content (type, title, markdown, image_url, link, created_at, updated_at)
VALUES
('about', 'bio',
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
, NULL, NULL, NOW(), NOW());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS content;
-- +goose StatementEnd

