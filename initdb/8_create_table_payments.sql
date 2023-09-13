CREATE TABLE payments
(
    id            UUID PRIMARY KEY,
    order_id      UUID NOT NULL,
    user_id       UUID NOT NULL,
    payment_type  VARCHAR NOT NULL,
    created_at    TIMESTAMP  NOT NULL DEFAULT current_timestamp,
    updated_at    TIMESTAMP  NOT NULL DEFAULT current_timestamp,
    CONSTRAINT fk_order_id FOREIGN KEY (order_id) REFERENCES orders(id),
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);
