-- +migrate Up
CREATE TABLE vehicle_price
(
    id               uuid      DEFAULT extension.uuid_generate_v4() NOT NULL,
    vehicle_id       uuid                                           NOT NULL,
    duration         int                                            NOT NULL,
    duration_type_id uuid                                           NOT NULL,
    price            int8                                           NOT NULL,
    created_at       timestamp DEFAULT now()                        NOT NULL,
    updated_at       timestamp DEFAULT now()                        NOT NULL,
    deleted_at       timestamp,
    PRIMARY KEY (id),
    CONSTRAINT vehicle_price_vehicle_duration_type
        UNIQUE (vehicle_id, duration_type_id, duration)
);
CREATE INDEX vehicle_price_vehicle_id
    ON vehicle_price (vehicle_id);
CREATE INDEX vehicle_price_duration_type_id
    ON vehicle_price (duration_type_id);
ALTER TABLE vehicle_price
    ADD CONSTRAINT FKvehicle_pr484952 FOREIGN KEY (vehicle_id) REFERENCES vehicle (id);
ALTER TABLE vehicle_price
    ADD CONSTRAINT FKvehicle_pr259855 FOREIGN KEY (duration_type_id) REFERENCES duration_type (id);

-- +migrate Down
ALTER TABLE vehicle_price
    DROP CONSTRAINT FKvehicle_pr484952;
ALTER TABLE vehicle_price
    DROP CONSTRAINT FKvehicle_pr259855;
DROP TABLE IF EXISTS vehicle_price CASCADE;
