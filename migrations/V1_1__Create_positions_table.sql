CREATE TABLE positions (
    position_id serial PRIMARY KEY,
    add_price double precision NOT NULL,
    close_price double precision DEFAULT NULL,
    symbol VARCHAR(6),
    opened_at TIMESTAMP DEFAULT NOW(),
    is_buy_type boolean NOT NULL,
    stop_loss double precision DEFAULT NULL,
    take_profit double precision DEFAULT NULL
);