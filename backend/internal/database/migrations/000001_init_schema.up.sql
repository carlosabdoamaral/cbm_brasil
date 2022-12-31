


-- -- -- -- -- -- -- -- -- -- -- -- -- -- --
CREATE TABLE account_tb (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE account_location_tb (
    id SERIAL PRIMARY KEY ,
    id_account INT REFERENCES account_tb(id),
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL
);

CREATE TABLE token_tb (
    id   SERIAL PRIMARY KEY,
    id_account INT REFERENCES account_tb(id),
    code VARCHAR(6) DEFAULT 'xxxxxx',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);





-- -- -- -- -- -- -- -- -- -- -- -- -- -- --
CREATE TABLE occurrence_tb (
    id SERIAL PRIMARY KEY,
    id_account INT REFERENCES account_tb(id),
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    image64 VARCHAR(500) NOT NULL,
    is_deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE occurrence_location_tb (
    id SERIAL PRIMARY KEY,
    id_occurrence INT REFERENCES occurrence_tb(id),
    latitude FLOAT NOT NULL,
    longitude FLOAT NOT NULL
);

CREATE TABLE occurrence_tag_tb(
    id SERIAL PRIMARY KEY,
    id_occurrence INT REFERENCES occurrence_tb(id),
    tags VARCHAR(255) NOT NULL
);





-- -- -- -- -- -- -- -- -- -- -- -- -- -- --
CREATE TABLE tutorial_tb (
    id SERIAL PRIMARY KEY,
    id_author INT REFERENCES account_tb(id),
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE tutorial_update_tb (
    id SERIAL PRIMARY KEY,
    id_author INT REFERENCES account_tb(id),
    id_tutorial INT REFERENCES tutorial_tb(id),
    description VARCHAR(255) NOT NULL,
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE tutorial_viewed_by_tb (
    id SERIAL PRIMARY KEY,
    id_tutorial INT REFERENCES tutorial_tb(id),
    account_id INT REFERENCES account_tb(id)
);

CREATE TABLE tutorial_step_tb (
    id SERIAL PRIMARY KEY,
    id_tutorial INT REFERENCES tutorial_tb(id),
    step_number INT NOT NULL
);