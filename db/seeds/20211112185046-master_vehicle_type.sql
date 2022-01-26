-- +migrate Up
insert into vehicle_type(name, detail)
values ('E-Scooter', 'Sekuter Elektrik');
insert into vehicle_type(name, detail)
values ('Motorcycle', 'Sepeda Motor');

-- +migrate Down
