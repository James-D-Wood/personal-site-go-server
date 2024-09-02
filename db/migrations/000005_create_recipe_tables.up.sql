CREATE TABLE ingredient_categories (
    name text NOT NULL
);

CREATE TABLE ingredients (
    name text NOT NULL,
    category text
);

CREATE TABLE recipe_ingredients (
    recipe text NOT NULL,
    ingredient text NOT NULL,
    measurement real,
    units text
);

CREATE TABLE recipes (
    name text NOT NULL,
    serves integer,
    url text
);

CREATE TABLE units (
    name text NOT NULL
);

ALTER TABLE ONLY ingredient_categories
    ADD CONSTRAINT ingredient_categories_pkey PRIMARY KEY (name);

ALTER TABLE ONLY ingredients
    ADD CONSTRAINT ingredients_pkey PRIMARY KEY (name);

ALTER TABLE ONLY recipe_ingredients
    ADD CONSTRAINT recipe_ingredients_pkey PRIMARY KEY (recipe, ingredient);

ALTER TABLE ONLY recipes
    ADD CONSTRAINT recipes_pkey PRIMARY KEY (name);

ALTER TABLE ONLY units
    ADD CONSTRAINT units_pkey PRIMARY KEY (name);

ALTER TABLE ONLY ingredients
    ADD CONSTRAINT ingredients_category_fkey FOREIGN KEY (category) REFERENCES ingredient_categories(name);

ALTER TABLE ONLY recipe_ingredients
    ADD CONSTRAINT recipe_ingredients_ingredient_fkey FOREIGN KEY (ingredient) REFERENCES ingredients(name);

ALTER TABLE ONLY recipe_ingredients
    ADD CONSTRAINT recipe_ingredients_recipe_fkey FOREIGN KEY (recipe) REFERENCES recipes(name);

ALTER TABLE ONLY recipe_ingredients
    ADD CONSTRAINT recipe_ingredients_units_fkey FOREIGN KEY (units) REFERENCES units(name);
