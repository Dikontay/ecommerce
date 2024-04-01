CREATE TABLE product_shelve (
                                ID INTEGER PRIMARY KEY,
                                product_id INTEGER NOT NULL,
                                shelve_id INTEGER NOT NULL,
                                isPrimary BOOLEAN NOT NULL DEFAULT 0, -- SQLite does not have a built-in BOOLEAN type, it uses INTEGER.
                                FOREIGN KEY (product_id) REFERENCES products (ID),
                                FOREIGN KEY (shelve_id) REFERENCES shelves (shelve_id)
);