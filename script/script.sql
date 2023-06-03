CREATE DATABASE exercise_sql;
USE exercise_sql;


CREATE TABLE role_actors (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    role_name VARCHAR(100),

    PRIMARY KEY(id)
);

CREATE TABLE actors (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    username VARCHAR(100) UNIQUE,
    password VARCHAR(500),
    role_id BIGINT UNSIGNED,
    is_verified BOOLEAN,
    is_active BOOLEAN,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    FOREIGN KEY(role_id) REFERENCES role_actors(id) ON DELETE CASCADE
);

CREATE TABLE customers (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    email VARCHAR(100),
    avatar VARCHAR(100),
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    PRIMARY KEY(id)
);

CREATE TABLE admin_reg (
    id BIGINT UNSIGNED AUTO_INCREMENT,
    admin_id BIGINT UNSIGNED,
    super_admin_id BIGINT UNSIGNED,
    status VARCHAR(100),

    PRIMARY KEY(id),
    FOREIGN KEY(admin_id) REFERENCES actors(id) ON DELETE CASCADE,
    FOREIGN KEY(super_admin_id) REFERENCES actors(id) ON DELETE CASCADE
);

CREATE TABLE authentications (
	token TEXT NOT NULL
);

-- insert role, 1 -> admin, 2 -> super_admin
INSERT INTO role_actors(id, role_name) 
VALUES (1, 'admin'),
(2, 'super_admin');

-- insert super admin, 
-- username: super_admin, 
-- password: password
INSERT INTO actors(id, username, password, role_id, is_verified, is_active) 
VALUES (1, 'super_admin', '$2a$04$e1it1T0mKhWvyvpIvbhMJuACG9qPS8DtV4laZnEpo6FPMTSk/CH1m', 2, 1, 1);

INSERT INTO customers(first_name, last_name, email, avatar)
VALUES (
        'George',
        'Bluth',
        'george.bluth@reqres.in',
        'https://reqres.in/img/faces/1-image.jpg'
    ),
    (
        'Janet',
        'Weaver',
        'janet.weaver@reqres.in',
        'https://reqres.in/img/faces/2-image.jpg'
    ),
    (
        'Emma',
        'Wong',
        'emma.wong@reqres.in',
        'https://reqres.in/img/faces/3-image.jpg'
    ),
    (
        'Eve',
        'Holt',
        'eve.holt@reqres.in',
        'https://reqres.in/img/faces/4-image.jpg'
    ),
    (
        'Charles',
        'Morris',
        'charles.morris@reqres.in',
        'https://reqres.in/img/faces/5-image.jpg'
    ),
    (
        'Tracey',
        'Ramos',
        'tracey.ramos@reqres.in',
        'https://reqres.in/img/faces/6-image.jpg'
    ),
    (
        'Michael',
        'Lawson',
        'michael.lawson@reqres.in',
        'https://reqres.in/img/faces/7-image.jpg'
    ),
    (
        'Lindsay',
        'Ferguson',
        'lindsay.ferguson@reqres.in',
        'https://reqres.in/img/faces/8-image.jpg'
    ),
    (
        'Tobias',
        'Funke',
        'tobias.funke@reqres.in',
        'https://reqres.in/img/faces/9-image.jpg'
    ),
    (
        'Byron',
        'Fields',
        'byron.fields@reqres.in',
        'https://reqres.in/img/faces/10-image.jpg'
    ),
    (
        'George',
        'Edwards',
        'george.edwards@reqres.in',
        'https://reqres.in/img/faces/11-image.jpg'
    ),
    (
        'Rachel',
        'Howell',
        'rachel.howell@reqres.in',
        'https://reqres.in/img/faces/12-image.jpg'
    );

