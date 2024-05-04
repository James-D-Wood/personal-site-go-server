CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    topic TEXT UNIQUE,
    takeaways TEXT[] NOT NULL,
    questions TEXT[] NOT NULL,
    exercises TEXT[] NOT NULL,
    dt_created TIMESTAMP NOT NULL,
    dt_updated TIMESTAMP
);

CREATE TABLE lesson_tags (
    lesson_id INTEGER,
    tag TEXT,
    PRIMARY KEY (lesson_id, tag),
    FOREIGN KEY (lesson_id) REFERENCES lessons(id) ON DELETE CASCADE
);

CREATE TABLE lesson_references (
    lesson_id INTEGER,
    title TEXT,
    author TEXT,
    url TEXT,
    PRIMARY KEY (lesson_id, title),
    FOREIGN KEY (lesson_id) REFERENCES lessons(id) ON DELETE CASCADE
);
