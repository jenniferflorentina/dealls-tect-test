CREATE TABLE users
(
    id         uuid         PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4() ,
    username   VARCHAR(255) NOT NULL,
    password   VARCHAR(255) NOT NULL,
    name       VARCHAR(255) NOT NULL,
    phone      VARCHAR(15)  NOT NULL,
    email      VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE
);
