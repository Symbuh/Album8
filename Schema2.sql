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

create unique index image_tags_id_image_id on image_tags (image_id);
create index image_tags_id_tags on image_tags using gin (tags);
Create index image_id_idx on images (image_id);