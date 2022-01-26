-- +migrate Up
CREATE TABLE vehicle
(
    id            uuid      DEFAULT extension.uuid_generate_v4() NOT NULL,
    type_id       uuid                                           NOT NULL,
    model         varchar                                        NOT NULL,
    brand         varchar                                        NOT NULL,
    category_id   uuid                                           NOT NULL,
    range         int,
    wheels        float,
    max_load      int,
    top_speed     int,
    waterproof    int,
    weight        int,
    plat_number   varchar,
    frame_number  varchar,
    engine_number varchar,
    created_at    timestamp DEFAULT now()                        NOT NULL,
    updated_at    timestamp DEFAULT now()                        NOT NULL,
    deleted_at    timestamp,
    PRIMARY KEY (id)
);
CREATE INDEX vehicle_type_id
    ON vehicle (type_id);
CREATE INDEX vehicle_category_id
    ON vehicle (category_id);
ALTER TABLE vehicle
    ADD CONSTRAINT FKvehicle191822 FOREIGN KEY (type_id) REFERENCES vehicle_type (id);
ALTER TABLE vehicle
    ADD CONSTRAINT FKvehicle692909 FOREIGN KEY (category_id) REFERENCES vehicle_category (id);

-- +migrate Down
ALTER TABLE vehicle
    DROP CONSTRAINT FKvehicle191822;
ALTER TABLE vehicle
    DROP CONSTRAINT FKvehicle692909;
DROP TABLE IF EXISTS vehicle CASCADE;
