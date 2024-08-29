-- +goose Up
-- +goose StatementBegin
INSERT INTO posts (title, content, author_id) VALUES
('Happy New Year!',
 'New Year is a widely celebrated occasion in the United Kingdom, marking the end of one year and the beginning of another.

## Top New Year Activities in the UK

* Attending a **Hogmanay** celebration in Scotland
* Taking part in a local *First-Foot* tradition in Scotland and Northern England
* Setting personal resolutions and goals for the upcoming year
* Going for a New Year''s Day walk to enjoy the fresh start
* Visiting a local pub for a celebratory toast and some cheer',
 1),
('May Day',
 'May Day is an ancient spring festival celebrated on the first of May in the United Kingdom, embracing the arrival of warmer weather and the renewal of life.

## Top May Day Activities in the UK

* Dancing around the **Maypole**, a traditional folk activity
* Attending local village fetes and fairs
* Watching or participating in *Morris dancing* performances
* Enjoying the public holiday known as *Early May Bank Holiday*',
 2),
('Bonfire Night',
 'Bonfire Night, also known as Guy Fawkes Night, is an annual commemoration observed on 5 November in the United Kingdom.

## Popular Bonfire Night Traditions

* Lighting bonfires and setting off fireworks
* Burning effigies known as "Guys"
* Enjoying traditional foods like toffee apples and baked potatoes
* Attending organized firework displays in local communities
* Reciting the rhyme "Remember, remember the fifth of November"',
 3);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM posts
WHERE title IN (
  'Happy New Year!',
  'May Day',
  'Bonfire Night'
);
-- +goose StatementEnd

