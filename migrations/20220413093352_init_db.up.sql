do
$$
    begin
        execute 'ALTER DATABASE ' || current_database() || ' SET timezone = ''+06''';
    end;
$$;

create table users
(
                     id serial primary key,
                     first_name text not null default '',
                     last_name text not null default ''
);
create index users_first_name_idx
    on users (first_name);
create index users_last_name_idx
    on users (last_name);