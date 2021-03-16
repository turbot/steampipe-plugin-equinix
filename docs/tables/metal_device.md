# Table: metal_device

Devices (servers) from all projects visible to the user.

## Examples

### List all devices

```sql
select
  *
from
  metal_device
```

### Get IP addresses for a device

```sql
select
  jsonb_array_elements(ip_addresses)->>'address' as ip_address
from
  metal_device
where
  hostname = 'ny5-c3-medium-x86-01'
```

### Find devices tagged as production

```sql
select
  hostname,
  tags
from
  metal_device
where
  tags->'production' is not null
```

### Group devices by facility

```sql
select
  f.code,
  f.name,
  count(*) as num_devices
from
  metal_device as d,
  metal_facility as f
where
  d.facility_id = f.id
group by
  f.code,
  f.name
order by
  num_devices desc
```

### List devices with OS information

```sql
select
  d.hostname,
  os.name,
  os.version
from
  metal_device as d,
  metal_operating_system as os
where
  d.operating_system_slug = os.slug
```

### Group devices by OS distribution

```sql
select
  os.distro,
  count(*)
from
  metal_device as d,
  metal_operating_system as os
where
  d.operating_system_slug = os.slug
group by
  os.distro
order by
  count desc
```
