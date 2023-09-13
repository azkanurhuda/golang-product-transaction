CREATE TABLE carts
(
    id            UUID PRIMARY KEY,
    user_id       UUID NOT NULL,
    product_id    UUID NOT NULL,
    created_at    TIMESTAMP  NOT NULL DEFAULT current_timestamp,
    updated_at    TIMESTAMP  NOT NULL DEFAULT current_timestamp,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_product_id FOREIGN KEY (product_id) REFERENCES products(id)
);
