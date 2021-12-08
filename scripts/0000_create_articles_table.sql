CREATE TABLE articles (
    id SERIAL PRIMARY KEY,
    title TEXT UNIQUE NOT NULL,
    uri TEXT UNIQUE NOT NULL,
    summary TEXT NOT NULL,
    body_md TEXT NOT NULL
);