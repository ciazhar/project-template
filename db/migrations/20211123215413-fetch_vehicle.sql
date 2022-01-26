-- +migrate Up
-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION fetch_vehicle()
    RETURNS TABLE
            (
                id           uuid,
                brand        varchar,
                model        varchar,
                images       jsonb,
                lowest_price bigint
            )
    LANGUAGE sql
AS
$function$
select v.id,
       brand,
       model,
       coalesce(jsonb_agg(distinct jsonb_build_object(
               'id', vi.id,
               'dir', vi.dir,
               'file', vi.file
           )) filter ( where vi.id is not null ), '[]') as images,
       coalesce(min(vp.price), 0)                       as lowest_price
from vehicle v
         left join vehicle_image vi on v.id = vi.vehicle_id and vi.deleted_at is null
         left join vehicle_price vp on v.id = vp.vehicle_id and vp.deleted_at is null
where v.deleted_at is null
group by v.id;
$function$;
-- +migrate StatementEnd

-- +migrate Down
