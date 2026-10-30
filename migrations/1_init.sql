create table if not exists embeddings_meta (
    id integer primary key,
    value text not null,
    offset int not null unique,
    created_at timestamp not null
);

create table if not exists free_offset (
    offset bigint primary key
);
