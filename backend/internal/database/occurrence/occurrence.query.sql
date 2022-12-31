-- 
SELECT
    occurrence_tb.id as occurrence_id,
    occurrence_tb.title as occurrence_title,
    occurrence_tb.description as occurrence_description,
    occurrence_tb.image64 as occurrence_image,
    occurrence_location_tb.latitude as location_latitude,
    occurrence_location_tb.longitude as location_longitude,
    account_tb.name as account_name,
    account_tb.email as account_email
FROM occurrence_tb
INNER JOIN account_tb on account_tb.id = occurrence_tb.id_account
INNER JOIN occurrence_location_tb on occurrence_tb.id = occurrence_location_tb.id_occurrence;





--
SELECT
    occurrence_tb.id as occurrence_id,
    occurrence_tb.title as occurrence_title,
    occurrence_tb.description as occurrence_description,
    occurrence_tb.image64 as occurrence_image,
    occurrence_location_tb.latitude as location_latitude,
    occurrence_location_tb.longitude as location_longitude,
    account_tb.name as account_name,
    account_tb.email as account_email
FROM occurrence_tb
INNER JOIN account_tb on account_tb.id = occurrence_tb.id_account
INNER JOIN occurrence_location_tb on occurrence_tb.id = occurrence_location_tb.id_occurrence
WHERE occurrence_tb.id = 1;





--
SELECT
    occurrence_tb.id as occurrence_id,
    occurrence_tb.title as occurrence_title,
    occurrence_tb.description as occurrence_description,
    occurrence_tb.image64 as occurrence_image,
    occurrence_location_tb.latitude as location_latitude,
    occurrence_location_tb.longitude as location_longitude,
    account_tb.name as account_name,
    account_tb.email as account_email
FROM occurrence_tb
INNER JOIN account_tb on account_tb.id = occurrence_tb.id_account
INNER JOIN occurrence_location_tb on occurrence_tb.id = occurrence_location_tb.id_occurrence
WHERE
    occurrence_location_tb.longitude BETWEEN 100 and 200
    AND
    occurrence_location_tb.latitude BETWEEN 200 and 300;





--
INSERT INTO occurrence_tb(id_account, title, description, image64) VALUES ('');
INSERT INTO occurrence_location_tb(id_occurrence, latitude, longitude) VALUES ('');
INSERT INTO occurrence_tag_tb(id_occurrence, tags) VALUES ('');





--
UPDATE occurrence_tb
SET title = '', description = '', image64 = ''
WHERE id = 1;

UPDATE occurrence_location_tb
SET latitude = 1, longitude = 1
WHERE id_occurrence = 1;

UPDATE occurrence_tag_tb
SET tags = ''
WHERE id_occurrence = 1;




--
DELETE FROM occurrence_tag_tb WHERE id = 1;
DELETE FROM occurrence_location_tb WHERE id = 1;
DELETE FROM occurrence_tb WHERE id = 1;




--
UPDATE occurrence_tb SET is_deleted = true
WHERE id = 1;