-- name: FetchVehicleCategory :many
select id, name
from vehicle_category
where deleted_at is null;

-- name: FetchVehicleType :many
select id, name
from vehicle_type
where deleted_at is null;

-- name: StoreVehicle :exec
insert into vehicle(type_id, model, brand, category_id, range, wheels, max_load, top_speed, waterproof, weight)
VALUES (@type_id::uuid, @model::varchar, @brand::varchar, @category_id::uuid, @range::int, @wheels::float,
        @max_load::int, @top_speed::int, @waterproof::int, @weight::int);

-- name: ValidateVehicle :one
select result::varchar
from validate_vehicle(@type_id::uuid, @category_id::uuid, @vehicle_id::uuid, @duration_type_id::uuid);

-- name: StoreVehicleImage :exec
insert into vehicle_image (file, dir, vehicle_id)
values (@file::varchar, @dir::varchar, @vehicle_id::uuid);

-- name: FetchVehicle :many
select id::uuid,
       model::varchar,
       brand::varchar,
       images::json,
       lowest_price::bigint
from fetch_vehicle()
limit $1 offset $2;

-- name: DetailVehicle :one
select id::uuid,
       type_id::uuid,
       type_name::varchar,
       model::varchar,
       brand::varchar,
       category_id::uuid,
       category_name::varchar,
       range::integer,
       wheels::float,
       max_load::integer,
       top_speed::integer,
       waterproof::integer,
       weight::integer,
       images::json,
       prices::json,
       created_at::timestamp,
       updated_at::timestamp
from detail_vehicle(@id::uuid);

-- name: UpsertVehiclePrice :exec
insert into vehicle_price(vehicle_id, duration, duration_type_id, price)
values (@vehicle_id::uuid, @duration::int, @duration_type_id::uuid, @price::bigint)
on conflict(vehicle_id,duration,duration_type_id) DO UPDATE
    SET duration   = @duration::int,
        price      = @price::bigint,
        updated_at = now(),
        deleted_at = null;

-- name: DeleteVehiclePrice :exec
update vehicle_price
set deleted_at = now()
where id = $1;

-- name: UpdateVehicle :exec
update vehicle
set type_id       = (case
                         when @type_id::uuid = '00000000-0000-0000-0000-000000000000' then type_id
                         else @type_id::uuid end),
    model         = (case when @model::varchar = '' then model else @model::varchar end),
    brand         = (case when @brand::varchar = '' then brand else @brand::varchar end),
    category_id   = (case
                         when @category_id::uuid = '00000000-0000-0000-0000-000000000000' then category_id
                         else @category_id::uuid end),
    range         = (case when @range::int = 0 then range else @range::int end),
    wheels        = (case when @wheels::float = 0 then wheels else @wheels::float end),
    max_load      = (case when @max_load::int = 0 then max_load else @max_load::int end),
    top_speed     = (case when @top_speed::int = 0 then top_speed else @top_speed::int end),
    waterproof    = (case when @waterproof::int = 0 then waterproof else @waterproof::int end),
    weight        = (case when @weight::int = 0 then weight else @weight::int end),
    plat_number   = (case when @plat_number::varchar = '' then plat_number else @plat_number::varchar end),
    frame_number  = (case when @frame_number::varchar = '' then frame_number else @frame_number::varchar end),
    engine_number = (case when @engine_number::varchar = '' then engine_number else @engine_number::varchar end),
    updated_at    = now()
where id = @id::uuid;