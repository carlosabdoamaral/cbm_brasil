If is dirty, run:

- task migration-down

UPDATE public.schema_migrations
SET dirty = false
WHERE version = 1;
Then go to sql and run
