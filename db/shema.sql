CREATE TABLE positions (
    positions_id serial PRIMARY KEY,
    add_price_id VARCHAR (16),
    close_price_id  VARCHAR (16),
    opened_at TIMESTAMP DEFAULT NOW()
);