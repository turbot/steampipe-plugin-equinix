# Table: metal_plan

Information about Equinix Metal plans.

## Examples

### List all plans

```sql
select
  *
from
  metal_plan
```

### Get CPU details for each plan

```sql
select
  name,
  cpu ->> 'count' as cpu_count,
  cpu ->> 'type' as cpu_count
from
  metal_plan,
  jsonb_array_elements(specs -> 'cpus') as cpu
```

### Plans by price

```sql
select
  name,
  pricing ->> 'hour' as hourly_cost
from
  metal_plan
order by
  hourly_cost
```
