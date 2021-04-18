# Table: equinix_metal_plan

Information about Equinix Metal plans.

## Examples

### List all plans

```sql
select
  *
from
  equinix_metal_plan
```

### Get CPU details for each plan

```sql
select
  name,
  cpu ->> 'count' as cpu_count,
  cpu ->> 'type' as cpu_count
from
  equinix_metal_plan,
  jsonb_array_elements(specs -> 'cpus') as cpu
```

### Plans by price

```sql
select
  name,
  pricing ->> 'hour' as hourly_cost
from
  equinix_metal_plan
order by
  hourly_cost
```

### Plans available by metro

```sql
select
  m.name as metro_name,
  p.name as plan_name
from
  equinix_metal_plan as p,
  jsonb_array_elements_text(p.available_in_metros) as mid,
  equinix_metal_metro as m
where
  m.id = mid
order by
  m.name,
  p.name
```
