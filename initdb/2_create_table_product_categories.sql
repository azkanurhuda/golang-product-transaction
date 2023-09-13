CREATE TABLE product_categories
(
    id            UUID PRIMARY KEY,
    name          VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP  NOT NULL DEFAULT current_timestamp,
    updated_at    TIMESTAMP  NOT NULL DEFAULT current_timestamp
);
