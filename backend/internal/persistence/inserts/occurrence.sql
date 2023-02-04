INSERT INTO occurrence_tb(id_author, id_firefighter, banner_x64, title, description)
VALUES (1, 1, 'abcdex==', 'Title example', 'Description');

INSERT INTO occurrence_location_tb(id_occurrence, cep, country, state, city, neighborhood, street, place_number, complement, latitude, longitude)
VALUES (1, '00000000', 'BRAZIL', 'TEMPLATE', 'TEMPLATE', 'TEMPLATE', 'TEMPLATE', 999, 'TEMPLATE', 27.000, 12.0000);

INSERT INTO occurrence_logs_tb(id_occurrence, msg) VALUES (1, 'occurrence created');