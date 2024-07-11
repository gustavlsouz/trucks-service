create table
  truck (
    id uuid primary key default gen_random_uuid(),
    plateNumber varchar(32) not null unique,
    createdAt timestamp without time zone not null default now()
  );