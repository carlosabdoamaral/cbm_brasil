If is dirty, run:

- task migration-down

Then go to sql and run
UPDATE public.schema_migrations
SET dirty = false
WHERE version = 1;
