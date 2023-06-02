CREATE DATABASE exercise_sql;
USE exercise_sql;


CREATE TABLE role_actors (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    role_name VARCHAR(100),

    PRIMARY KEY(id)
)

CREATE TABLE actors (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    username VARCHAR(100),
    password VARCHAR(500),
    role_id BIGINT UNSIGNED,
    is_verified BOOLEAN,
    is_active BOOLEAN,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    FOREIGN KEY(role_id) REFERENCES role_actors(id)
)

CREATE TABLE customers (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    email VARCHAR(100),
    avatar VARCHAR(100),
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    PRIMARY KEY(id)
)

CREATE TABLE admin_reg (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    admin_id BIGINT UNSIGNED,
    super_admin_id BIGINT UNSIGNED,
    status VARCHAR(100),

    PRIMARY KEY(id),
    FOREIGN KEY(admin_id) REFERENCES actors(id),
    FOREIGN KEY(super_admin_id) REFERENCES actors(id)
)

