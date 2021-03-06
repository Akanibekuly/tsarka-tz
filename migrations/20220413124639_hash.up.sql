create table hash
(
    id text not null primary key,
    status text not null default 'PENDING',
    result int not null default 0
);
create index hash_id_idx
    on hash (id);
create index hash_status_idx
    on hash (status);
create index hash_result_idx
    on hash (result);