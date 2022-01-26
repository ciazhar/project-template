-- +migrate Up
-- +migrate StatementBegin
CREATE OR REPLACE function validate_vehicle(
    _type_id uuid default '00000000-0000-0000-0000-000000000000',
    _category_id uuid default '00000000-0000-0000-0000-000000000000',
    _vehicle_id uuid default '00000000-0000-0000-0000-000000000000',
    _duration_type_id uuid default '00000000-0000-0000-0000-000000000000'
)
    RETURNS TABLE
            (
                result varchar
            )
    LANGUAGE sql
as
$function$
with cte_category_id as (
    select case
               when _category_id = '00000000-0000-0000-0000-000000000000' then true
               else count(*) = 1 end as correct
    from vehicle_category
    where id = _category_id
),
     cte_type_id as (
         select case
                    when _type_id = '00000000-0000-0000-0000-000000000000' then true
                    else count(*) = 1 end as correct
         from vehicle_type
         where id = _type_id
     ),
     cte_vehicle_id as (
         select case
                    when _vehicle_id = '00000000-0000-0000-0000-000000000000' then true
                    else count(*) = 1 end as correct
         from vehicle
         where id = _vehicle_id
     ),
     cte_duration_type_id as (
         select case
                    when _duration_type_id = '00000000-0000-0000-0000-000000000000' then true
                    else count(*) = 1 end as correct
         from duration_type
         where id = _duration_type_id
     )
select case
           when c.correct = false then 'Category ID Not Found'
           when t.correct = false then 'Type ID Not Found'
           when v.correct = false then 'Vehicle ID Not Found'
           when dt.correct = false then 'Duration Type ID Not Found'
           else 'Validated'
           end
           as result
from cte_type_id t
         cross join cte_category_id c
         cross join cte_vehicle_id v
         cross join cte_duration_type_id dt
$function$;
-- +migrate StatementEnd

-- +migrate Down
