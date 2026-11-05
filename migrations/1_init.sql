create table if not exists collections (
    id uuid primary key,
    title text not null,
    embedding_size integer not null,
    binary_filepath text not null,
    created_at timestamp not null
);

create table if not exists embeddings_meta (
    id uuid primary key,
    collection_id uuid not null references collections(id),
    value text not null,
    offset int not null unique,
    created_at timestamp not null
);

create table if not exists free_offsets (
    offset bigint primary key
);
