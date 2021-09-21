alter table bot_user
    drop column first_name,
    drop column last_name,
    drop column t_id;

alter table state
    drop column bot_user_id;