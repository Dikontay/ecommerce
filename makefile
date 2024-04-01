migrate :
	sqlite3 ecommerce.sqlite3 < migrations/000001_create_products_table.down.sql
	sqlite3 ecommerce.sqlite3 < migrations/000001_create_products_table.up.sql
	sqlite3 ecommerce.sqlite3 < migrations/000002_create_shelves_table.down.sql
	sqlite3 ecommerce.sqlite3 < migrations/000002_create_shelves_table.up.sql
	sqlite3 ecommerce.sqlite3 < migrations/000003_create_orders_table.down.sql
	sqlite3 ecommerce.sqlite3 < migrations/000003_create_orders_table.up.sql
	sqlite3 ecommerce.sqlite3 < migrations/000004_create_product_order_table.down.sql
	sqlite3 ecommerce.sqlite3 < migrations/000004_create_product_order_table.up.sql
	sqlite3 ecommerce.sqlite3 < migrations/000005_create_product_shelve_table.down.sql
	sqlite3 ecommerce.sqlite3 < migrations/000005_create_product_shelve_table.up.sql

