CREATE TABLE IF NOT EXISTS users (
  id BIGSERIAL primary key,
  first_name TEXT not null,
  last_name TEXT,
  created_at TIMESTAMP default now()
);