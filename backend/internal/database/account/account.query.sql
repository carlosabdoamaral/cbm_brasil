-- Get public infos -> list
SELECT 
    id,
    name,
    email,
    created_at
FROM account_tb;





-- Get private infos -> single
SELECT
    account_tb.id,
    account_tb.name,
    account_tb.email,
    account_location_tb.latitude,
    account_location_tb.longitude,
    token_tb.code,
    token_tb.created_at,
    account_tb.password,
    account_tb.is_deleted,
    account_tb.created_at,
    account_tb.updated_at
FROM account_tb
INNER JOIN account_location_tb on account_location_tb.id_account = account_tb.id
INNER JOIN token_tb on token_tb.id_account = account_tb.id
WHERE
    account_tb.id = 1;





-- Get public infos -> single
SELECT
    id,
    name,
    email,
    created_at
FROM
    account_tb
WHERE
    account_tb.id = 1;





-- Create account
INSERT INTO account_tb(name, email, password) VALUES ('Carlos Amaral', 'carlos@gmail.com', 'Senha_admin123');
INSERT INTO account_location_tb(id_account, latitude, longitude) VALUES (1, 0000, 0000);





-- Create token
INSERT INTO token_tb(code, id_account) VALUES ('xxxxxx', 1);





-- Update account
UPDATE account_tb
SET name = '', email = '', password = ''
WHERE account_tb.id = 1;

UPDATE account_location_tb
SET longitude = '', latitude = ''
WHERE account_location_tb.id_account = 1;

UPDATE token_tb
SET code = 'xxxxxx', updated_at = NOW()
WHERE id_account = 1;





-- Delete account
UPDATE occurrence_tb
SET id_account = null
WHERE id_account = 1;

DELETE FROM token_tb WHERE id_account = 1;
DELETE FROM account_location_tb WHERE id_account = 1;
DELETE FROM account_tb WHERE id = 1;




-- SoftDelete account
UPDATE account_tb
SET is_deleted = true
WHERE id = 1;