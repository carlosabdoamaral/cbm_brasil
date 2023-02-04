CREATE TABLE account_tb(
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    cpf VARCHAR(11) NOT NULL,
    birth_date TIMESTAMP,
    passwd VARCHAR(255) DEFAULT 'Senha_admin123',
    two_factor_code VARCHAR(6) DEFAULT '000000',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    soft_deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE account_location_tb(
    id SERIAL PRIMARY KEY,
    id_account INT REFERENCES account_tb(id),
    cep VARCHAR(8) NOT NULL,
    country VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    neighborhood VARCHAR(255) NOT NULL,
    street VARCHAR(255) NOT NULL,
    place_number INT NOT NULL,
    complement VARCHAR(255) NOT NULL
);

CREATE TABLE firefighter_account_tb(
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    cpf VARCHAR(11) NOT NULL,
    birth_date TIMESTAMP,
    passwd VARCHAR(255) DEFAULT 'Senha_admin123',
    two_factor_code VARCHAR(6) DEFAULT '000000',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    soft_deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE firefighter_account_location_tb(
    id SERIAL PRIMARY KEY,
    id_account INT REFERENCES firefighter_account_tb(id),
    cep VARCHAR(8) NOT NULL,
    country VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    neighborhood VARCHAR(255) NOT NULL,
    street VARCHAR(255) NOT NULL,
    place_number INT NOT NULL,
    complement VARCHAR(255) NOT NULL
);

CREATE TABLE occurrence_tb (
    id SERIAL PRIMARY KEY,
    id_author INT REFERENCES account_tb(id) NOT NULL,
    id_firefighter INT REFERENCES firefighter_account_tb(id),
    banner_x64 VARCHAR(500) NOT NULL,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    soft_deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    accepted_at TIMESTAMP DEFAULT NOW(),
    canceled_at TIMESTAMP DEFAULT NOW(),
    finished_at TIMESTAMP DEFAULT NOW(),
    is_accepted BOOLEAN DEFAULT FALSE,
    is_canceled BOOLEAN DEFAULT FALSE,
    is_finished BOOLEAN DEFAULT FALSE
);

CREATE TABLE occurrence_location_tb(
    id SERIAL PRIMARY KEY,
    id_occurrence INT REFERENCES occurrence_tb(id),
    cep VARCHAR(8) NOT NULL,
    country VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    neighborhood VARCHAR(255) NOT NULL,
    street VARCHAR(255) NOT NULL,
    place_number INT NOT NULL,
    complement VARCHAR(255) NOT NULL,
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL
);


CREATE TABLE occurrence_logs_tb (
    id SERIAL PRIMARY KEY,
    id_occurrence INT REFERENCES occurrence_tb(id),
    logged_at TIMESTAMP DEFAULT NOW(),
    msg VARCHAR(255) NOT NULL
);