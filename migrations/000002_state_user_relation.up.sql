alter table bot_user
    add column first_name varchar(50),
    add column last_name  varchar(50),
    add column t_id       int;

alter table state
    add column bot_user_id int;