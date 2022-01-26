-- +migrate Up
CREATE TABLE duration_type
(
    id         uuid      DEFAULT extension.uuid_generate_v4() NOT NULL,
    name       varchar                                        NOT NULL,
    created_at timestamp DEFAULT now()                        NOT NULL,
    updated_at timestamp DEFAULT now()                        NOT NULL,
    deleted_at timestamp,
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS duration_type CASCADE;
