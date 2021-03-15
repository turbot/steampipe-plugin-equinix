# Table: metal_operating_system

Information about operating systems offered by Equinix Metal.

## Examples

### List all

```sql
select
  *
from
  metal_operating_system
order by
  name
```

### Group by distribution

```sql
select
  distro,
  count(*)
from
  metal_operating_system
group by
  distro
order by
  count desc
```

### List all Ubuntu OS versions available

```sql
select
  *
from
  metal_operating_system
where
  distro = 'ubuntu'
order by
  name
```

### Which operating systems can be provisioned on a c3.medium.x86 device?

```sql
select
  *
from
  metal_operating_system
where
  provisionable_on ? 'c3.medium.x86'
```

### List all OS device combinations

```sql
select
  slug,
  device
from
  metal_operating_system,
  jsonb_array_elements_text(provisionable_on) as device
```
