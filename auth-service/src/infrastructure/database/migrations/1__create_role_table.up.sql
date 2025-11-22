CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS roles
(
   id UUID PRIMARY KEY NOT NULL,
   created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,

   name VARCHAR(255) NOT NULL
);
CREATE INDEX idx_roles_name ON roles(name);

INSERT INTO roles (id, name)
VALUES
    (gen_random_uuid(), 'USER'),
    (gen_random_uuid(), 'ADMIN');