CREATE TABLE users (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    address VARCHAR(255),
    email VARCHAR(255),
    born_date TIMESTAMP
);

-- Untuk migration (membuat tabel)
-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/restapi_gin" -path database/migrations up

-- Untuk migration (menghapus tabel)
-- migrate -database "mysql://root:@tcp(127.0.0.1:3306)/restapi_gin" -path database/migrations down

-- UNTUK MEMBUAT FILE MIGRATION
-- migrate create -ext sql -dir database/migrations -seq create_users_table                        