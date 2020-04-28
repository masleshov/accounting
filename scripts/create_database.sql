SET SEARCH_PATH TO accounting;

create sequence plan_plan_id_seq
;

create sequence user_ref_user_id_seq
;

create table plan
(
	plan_id serial not null
		constraint plan_pkey
			primary key,
	user_id integer not null,
	date_beg date not null,
	date_end date not null,
	plan_no varchar not null
)
;

create index plan_user_id_index
	on plan (user_id)
;

create table user_ref
(
	user_id serial not null
		constraint user_ref_pkey
			primary key,
	pwd_hash uuid not null,
	name varchar(30) not null,
	second_name varchar(30) not null,
	surname varchar(30) not null
)
;

create unique index user_ref_user_id_uindex
	on user_ref (user_id)
;

alter table plan
	add constraint plan_user_ref_user_id_fk
		foreign key (user_id) references user_ref
;

