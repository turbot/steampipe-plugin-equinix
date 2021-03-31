# Table: equinix_metal_spot_price

Spot prices for each facility and plan.

## Examples

### List all spot prices

```sql
select
  *
from
  equinix_metal_spot_price
order by
  facility,
  plan
```

### Find the cheapest facility to run a c3.medium.x86

```sql
select
  *
from
  equinix_metal_spot_price
where
  plan = 'c3.medium.x86'
order by
  current_price
```

### 10 cheapest servers

```sql
select
  *
from
  equinix_metal_spot_price
order by
  current_price
limit
  10
```
