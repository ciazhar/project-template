-- +migrate Up
-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION detail_vehicle(_id uuid)
    RETURNS TABLE
            (
                id            uuid,
                type_id       uuid,
                type_name     varchar,
                model         varchar,
                brand         varchar,
                category_id   uuid,
                category_name varchar,
                range         integer,
                wheels        float,
                max_load      integer,
                top_speed     integer,
                waterproof    integer,
                weight        integer,
                images        jsonb,
                prices        jsonb,
                created_at    timestamp,
                updated_at    timestamp
            )
    LANGUAGE sql
AS
$function$
select v.id,
       vt.id                                            as type_id,
       coalesce(vt.name, '')                            as type_name,
       model,
       brand,
       vc.id                                            as category_id,
       coalesce(vc.name, '')                            as category_name,
       coalesce(range, 0)                               as range,
       coalesce(wheels, 0)                              as wheels,
       coalesce(max_load, 0)                            as max_load,
       coalesce(top_speed, 0)                           as top_speed,
       coalesce(waterproof, 0)                          as waterproof,
       coalesce(weight, 0)                              as weight,
       coalesce(jsonb_agg(DISTINCT jsonb_build_object(
               'id', vi.id,
               'dir', vi.dir,
               'file', vi.file
           )) filter ( where vi.id is not null ), '[]') as images,
       coalesce(jsonb_agg(DISTINCT jsonb_build_object(
               'id', vp.id,
               'duration', vp.duration,
               'duration_type', dt.name,
               'proce', vp.price
           )) filter ( where vp.id is not null ), '[]') as prices,
       v.created_at,
       v.updated_at
from vehicle v
         left join vehicle_type vt on v.type_id = vt.id and vt.deleted_at is null
         left join vehicle_category vc on v.category_id = vc.id and vc.deleted_at is null
         left join vehicle_image vi on v.id = vi.vehicle_id and vi.deleted_at is null
         left join vehicle_price vp on v.id = vp.vehicle_id and vp.deleted_at is null
         left join duration_type dt on vp.duration_type_id = dt.id and dt.deleted_at is null
where v.deleted_at is null
  and v.id = _id
group by v.id, vt.id, vc.id
$function$;
-- +migrate StatementEnd

-- +migrate Down
