CREATE TABLE account_location_tb(
    id SERIAL PRIMARY KEY,
    cep VARCHAR(8) NOT NULL,
    country VARCHAR(255) NOT NULL,
    state VARCHAR(255) NOT NULL,
    city VARCHAR(255) NOT NULL,
    neighborhood VARCHAR(255) NOT NULL,
    street VARCHAR(255) NOT NULL,
    place_number INT NOT NULL,
    complement VARCHAR(255) NOT NULL
);

CREATE TABLE account_tb(
    id SERIAL PRIMARY KEY,
    id_main_location INT REFERENCES account_location_tb(id),
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