create table books
(
    id           int auto_increment
        primary key,
    name         varchar(100) null,
    description  varchar(100) null,
    medium_price int          null,
    author       varchar(100) null,
    image_url    varchar(100) null,
    created_at   datetime     null,
    updated_at   datetime     null,
    deleted_at   datetime     null
);