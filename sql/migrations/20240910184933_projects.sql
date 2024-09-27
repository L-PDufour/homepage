-- +goose Up
-- +goose StatementBegin
INSERT INTO content (type, title, markdown, image_url, link, created_at, updated_at)
VALUES
('project', 'website',
'## Mon premier projet complet

Ce site web personnel est mon premier projet complet, développé après ma première année d''apprentissage en programmation. Il illustre mon exploration des technologies web modernes pour créer une expérience utilisateur simple mais efficace. Le code source est disponible sur [GitHub](https://github.com/L-PDufour/homepage).

### Technologies explorées :
- **Backend :** Go (bibliothèque standard), me permettant de comprendre les bases d''un serveur web.
- **Frontend :**
  - Templ et Tailwind CSS pour découvrir le rendu côté serveur et le styling moderne.
  - HTMX pour expérimenter avec l''interactivité côté client de manière simple.
- **Déploiement :** Initiation à Docker pour la conteneurisation et à Nix pour la gestion de l''environnement.

Ce projet reflète mon enthousiasme à apprendre et à appliquer de nouvelles technologies.',
NULL, NULL, NOW(), NOW()),

('project', 'probono',
'Dans le cadre de mon apprentissage, j''ai eu l''opportunité de contribuer au développement d''une application web pour l''inscription et la gestion des bénévoles de la Guignolée du Centre de pédiatrie sociale de Québec. Ce projet collaboratif a été une excellente occasion d''appliquer mes connaissances dans un contexte réel.

### Technologies découvertes :
- **Backend :** Initiation à NestJS et Prisma, découvrant ainsi les frameworks backend modernes.
- **Frontend :** Premiers pas avec React pour comprendre le développement d''interfaces utilisateur.
- **Déploiement :** Découverte de Docker et de ses principes de base.

Ce projet, toujours en développement, marque mes débuts dans un environnement de développement collaboratif. Il m''a permis d''appliquer mes connaissances théoriques à un cas pratique, tout en apprenant de mes collègues plus expérimentés.',
NULL, NULL, NOW(), NOW()),

('project', 'github',
'Mon profil [GitHub](https://github.com/L-PDufour/) témoigne de mon parcours d''apprentissage. Il regroupe divers projets, exercices et tutoriels que j''ai réalisés pour développer mes compétences.

### Langages explorés et apprentissages :
- **C :** Découverte des bases de la programmation à travers de petits exercices algorithmiques.
- **Go :** Apprentissage de la création d''APIs simples et de serveurs web basiques.
- **Python :** Initiation à la programmation orientée objet via de petits projets guidés.',
NULL, NULL, NOW(), NOW());

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS content;
-- +goose StatementEnd
