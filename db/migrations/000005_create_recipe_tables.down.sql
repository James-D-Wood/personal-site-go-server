DROP TABLE recipe_ingredients (
    recipe text NOT NULL,
    ingredient text NOT NULL,
    measurement real,
    units text
);

DROP TABLE recipes (
    name text NOT NULL,
    serves integer,
    url text
);

DROP TABLE ingredient_categories (
    name text NOT NULL
);

DROP TABLE ingredients (
    name text NOT NULL,
    category text
);


DROP TABLE units (
    name text NOT NULL
);
