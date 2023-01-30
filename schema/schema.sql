create table if not exists authors (
    id varchar(36) not null,
    name varchar(100) unique not null,
    primary key (id)
);

create table if not exists books (
    id varchar(36) not null,
    name varchar(100) unique not null,
    author_id varchar(36) not null,
    primary key (id),
    constraint fk_author foreign key(author_id) references authors(id)
);