CREATE TABLE clients (
    id BIGSERIAL PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    full_name TEXT NOT NULL,
    passport TEXT NOT NULL,
    birthdate DATE NOT NULL,
    status TEXT NOT NULL DEFAULT 'INACTIVE' CHECK ( status IN ('ACTIVE', 'INACTIVE')),
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE cards (
    id BIGSERIAL PRIMARY KEY,
    number TEXT NOT NULL,
    balance BIGINT NOT NULL DEFAULT 0,
    issuer TEXT NOT NULL CHECK ( issuer IN ('VISA', 'MASTER', 'MIR')),
    owner_id BIGINT NOT NULL REFERENCES clients,
    status TEXT NOT NULL DEFAULT 'INACTIVE' CHECK ( status IN ('ACTIVE', 'INACTIVE')),
    created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    category TEXT NOT NULL
);

CREATE TABLE descriptions (
    id BIGSERIAL PRIMARY KEY,
    description TEXT NOT NULL
);

CREATE TABLE logos (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL
);

CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    card_id BIGINT NOT NULL REFERENCES cards,
    amount BIGINT NOT NULL DEFAULT 0,
    category_id BIGINT NOT NULL REFERENCES categories,
    description_id BIGINT NOT NULL REFERENCES descriptions,
    logo_id BIGINT NOT NULL REFERENCES logos
);