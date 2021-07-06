DROP DATABASE IF EXISTS foundantImages;

CREATE DATABASE foundantImages;

\c foundantImages;


CREATE TABLE images (
  image_id serial PRIMARY KEY,
  url varchar NOT NULL,
  name varchar NOT NULL,
  description varchar NOT NULL
);


CREATE TABLE tags (
  tag_id serial PRIMARY KEY,
  name varchar NOT NULL
);

CREATE TABLE images_tags (
  image_id int REFERENCES images(image_id) ON UPDATE CASCADE ON DELETE CASCADE,
  tag_id int REFERENCES tags(tag_id) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT images_tags_pkey PRIMARY KEY (image_id, tag_id)
)

Create index image_id_idx on images (image_id);
create index tag_id_idx on images_tags (tag_id);