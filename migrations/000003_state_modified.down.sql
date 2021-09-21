alter table state drop column bot_user_id;
alter table state add column previous varchar(30);
alter table state add column bot_user_id int;
