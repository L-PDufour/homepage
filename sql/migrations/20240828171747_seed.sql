-- +goose Up
-- +goose StatementBegin
INSERT INTO posts (title, content, author_id) VALUES
('Happy New Year!',
 'New Year is a widely celebrated occasion in the United Kingdom, marking the end of one year and the beginning of another.\n\n## Top New Year Activities in the UK\n\n* Attending a **Hogmanay** celebration in Scotland\n* Taking part in a local *First-Foot* tradition in Scotland and Northern England\n* Setting personal resolutions and goals for the upcoming year\n* Going for a New Year''s Day walk to enjoy the fresh start\n* Visiting a local pub for a celebratory toast and some cheer',
 1),

('May Day',
 'May Day is an ancient spring festival celebrated on the first of May in the United Kingdom, embracing the arrival of warmer weather and the renewal of life.\n\n## Top May Day Activities in the UK\n\n* Dancing around the **Maypole**, a traditional folk activity\n* Attending local village fetes and fairs\n* Watching or participating in *Morris dancing* performances\n* Enjoying the public holiday known as *Early May Bank Holiday*',
 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DELETE FROM posts
WHERE title IN (
  'Happy New Year!',
  'May Day'
);
-- +goose StatementEnd
