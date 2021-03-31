# Table: equinix_metal_capacity

Capacity information for each facility and plan.

## Examples

### List all capacities

```sql
select
  *
from
  equinix_metal_capacity
order by
  facility,
  plan
```

### Show all servers with limited capacity

```sql
select
  facility,
  plan
from
  equinix_metal_capacity
where
  level = 'limited'
order by
  facility,
  plan
```
