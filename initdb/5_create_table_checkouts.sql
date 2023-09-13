CREATE TABLE orders
(
    id            UUID PRIMARY KEY,
    cart_id       UUID NOT NULL,
    user_id       UUID NOT NULL,
    created_at    TIMESTAMP  NOT NULL DEFAULT current_timestamp,
    updated_at    TIMESTAMP  NOT NULL DEFAULT current_timestamp,
    CONSTRAINT fk_cart_id FOREIGN KEY (cart_id) REFERENCES carts(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);
