
-- +migrate Up
INSERT INTO emobi_service.vehicle_category (id, name, created_at, updated_at, deleted_at) VALUES ('6bb611f3-eeb0-438a-9565-db40656adccf', 'Seated', now(), now(), null);
INSERT INTO emobi_service.vehicle_category (id, name, created_at, updated_at, deleted_at) VALUES ('31498010-6615-4246-a9f2-77fc4dd6b4df', 'Regular', now(), now(), null);
INSERT INTO emobi_service.vehicle_category (id, name, created_at, updated_at, deleted_at) VALUES ('1bad54ec-eb4a-4a7a-b7fb-61b68c7766b8', 'Premium', now(), now(), null);

-- +migrate Down
