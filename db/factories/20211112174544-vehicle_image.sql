-- +migrate Up
CREATE TABLE vehicle_image
(
    id         uuid      DEFAULT extension.uuid_generate_v4() NOT NULL,
    vehicle_id uuid                                           NOT NULL,
    dir        varchar                                        NOT NULL,
    "file"     varchar                                        NOT NULL,
    created_at timestamp DEFAULT now()                        NOT NULL,
    updated_at timestamp DEFAULT now()                        NOT NULL,
    deleted_at timestamp,
    PRIMARY KEY (id)
);
CREATE INDEX vehicle_image_vehicle_id
    ON vehicle_image (vehicle_id);
ALTER TABLE vehicle_image
    ADD CONSTRAINT FKvehicle_im863779 FOREIGN KEY (vehicle_id) REFERENCES vehicle (id);

-- +migrate Down
ALTER TABLE vehicle_image
    DROP CONSTRAINT FKvehicle_im863779;
DROP TABLE IF EXISTS vehicle_image CASCADE;
