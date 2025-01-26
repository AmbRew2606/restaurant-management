CREATE TABLE menu_items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    nutrition JSONB,
    price NUMERIC(10, 2)
);
