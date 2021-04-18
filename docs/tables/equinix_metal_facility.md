# Table: equinix_metal_facility

Information about Equinix Metal facilities.

## Examples

### List all Equinix Metal facilities

```sql
select
  *
from
  equinix_metal_facility
```

### Get metro information for each facility

```sql
select
  f.name as facility_name,
  m.name as metro_name
from
  equinix_metal_facility as f,
  equinix_metal_metro as m
where
  f.metro_id = m.id
```
