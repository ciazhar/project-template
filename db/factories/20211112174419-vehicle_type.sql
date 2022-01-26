-- +migrate Up
CREATE TABLE vehicle_type
(
    id         uuid      DEFAULT extension.uuid_generate_v4() NOT NULL,
    name       varchar                                        NOT NULL,
    detail     varchar,
    created_at timestamp DEFAULT now()                        NOT NULL,
    updated_at timestamp DEFAULT now()                        NOT NULL,
    deleted_at timestamp,
    PRIMARY KEY (id)
);
COMMENT ON TABLE vehicle_type IS '1. Escooter
2. Motorcycle';

-- +migrate Down
DROP TABLE IF EXISTS vehicle_type CASCADE;
