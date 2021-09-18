create table bot_user
(
    id serial constraint user_pkey primary key,
    phone varchar(30) not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON bot_user
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


create table task (
    id serial constraint task_pkey primary key,
    name varchar(200) not null,
    description varchar(2000),
    status varchar(50),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON task
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();


create table dashboard (
    bot_user_id integer not null constraint bot_user_id references bot_user,
    task_id integer not null constraint task_id references task
);

create table state (
    id serial constraint state_key primary key,
    previous varchar(30),
    current varchar(30),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON state
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();
