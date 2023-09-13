CREATE TABLE products
(
    id                      UUID PRIMARY KEY,
    product_category_id     UUID NOT NULL,
    name                    VARCHAR(255) NOT NULL,
    description             TEXT NOT NULL,
    price                   BIGINT NOT NULL,
    stock                   INT NOT NULL,
    created_at    TIMESTAMP  NOT NULL DEFAULT current_timestamp,
    updated_at    TIMESTAMP  NOT NULL DEFAULT current_timestamp,
    CONSTRAINT fk_product_category_id FOREIGN KEY (product_category_id) REFERENCES product_categories(id)
);
