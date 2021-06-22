DROP DATABASE IF EXISTS qanda;

CREATE DATABASE qanda;

\c qanda;


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
  tag_id int REFRENCES tags(tag_id) ON UPDATE CASCADE ON DELETE CASCADE,
  CONSTRAINT images_tags_pkey PRIMARY KEY (images_id, tags_id)
)

-- COPY questions FROM '/home/n/Desktop/questions.csv' DELIMITER ',' CSV HEADER;
-- COPY answers FROM '/home/n/Desktop/answers.csv' DELIMITER ',' CSV HEADER;
-- COPY answerphotos FROM '/home/n/Desktop/answers_photos.csv' DELIMITER ',' CSV HEADER;

Create index image_id_idx on images_tags (image_id);
create index tag_id_idx on imags_tags (tag_id);