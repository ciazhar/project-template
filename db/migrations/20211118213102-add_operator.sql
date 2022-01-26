-- +migrate Up
create operator <-> (procedure = extension.geography_distance_knn, leftarg = extension.geography, rightarg = extension.geography, commutator = <->);
-- +migrate Down
