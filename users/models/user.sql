CREATE TABLE users (
    id CHAR(36) NOT NULL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255),
    address VARCHAR(255),
    phone_number CHAR(10) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    display_name VARCHAR(255),
    gender BOOLEAN DEFAULT TRUE,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME
);
