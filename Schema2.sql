CREATE TABLE images (
  image_id serial PRIMARY KEY,
  url varchar NOT NULL,
  name varchar NOT NULL,
  description varchar NOT NULL
);

CREATE TABLE image_tags (
  image_id int not null references images(image_id) ON UPDATE CASCADE ON DELETE CASCADE,
  tags text[] not null default '{}'
);

unique index image_tags_id_image_id on (image_id);
index image_tags_id_tags using gin (tags);