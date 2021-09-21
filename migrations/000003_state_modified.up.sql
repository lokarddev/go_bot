alter table state drop column previous;
alter table state drop column bot_user_id;
alter table state add column bot_user_id int not null constraint id references bot_user default 1;