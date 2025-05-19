DO $$ 
DECLARE 
    r RECORD;
    e RECORD;
BEGIN
    -- Удаляем все таблицы
    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') 
    LOOP
        EXECUTE 'DROP TABLE IF EXISTS public.' || quote_ident(r.tablename) || ' CASCADE';
    END LOOP;

    -- Удаляем все ENUM-типы
    FOR e IN (SELECT typname FROM pg_type WHERE typtype = 'e' AND typnamespace = 'public'::regnamespace) 
    LOOP
        EXECUTE 'DROP TYPE IF EXISTS public.' || quote_ident(e.typname) || ' CASCADE';
    END LOOP;
END $$;
