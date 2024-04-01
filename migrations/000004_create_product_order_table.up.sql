CREATE TABLE product_order (
                               Product_order_id INTEGER PRIMARY KEY,
                               product_id INTEGER NOT NULL,
                               order_id INTEGER NOT NULL,
                               quantity INTEGER NOT NULL,
                               discount REAL,
                               FOREIGN KEY (product_id) REFERENCES products (ID),
                               FOREIGN KEY (order_id) REFERENCES orders (ID)
);