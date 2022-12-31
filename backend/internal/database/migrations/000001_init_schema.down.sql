UPDATE schema_migrations
SET dirty = false;

DROP TABLE IF EXISTS account_location_tb CASCADE;
DROP TABLE IF EXISTS token_tb CASCADE;
DROP TABLE IF EXISTS account_tb CASCADE;
DROP TABLE IF EXISTS occurrence_location_tb CASCADE;
DROP TABLE IF EXISTS occurrence_tag_tb CASCADE;
DROP TABLE IF EXISTS occurrence_tb CASCADE;
DROP TABLE IF EXISTS tutorial_tb CASCADE;
DROP TABLE IF EXISTS tutorial_update_tb CASCADE;
DROP TABLE IF EXISTS tutorial_viewed_by_tb CASCADE;
DROP TABLE IF EXISTS tutorial_step_tb CASCADE;