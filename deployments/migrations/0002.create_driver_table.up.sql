create table
  driver (
    id uuid primary key,
    document varchar(11) not null unique,
    createdAt timestamp without time zone not null default now ()
  );