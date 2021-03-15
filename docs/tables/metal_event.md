# Table: metal_event

Events across all devices.

## Examples

### List all events

```sql
select
  *
from
  metal_event
```

### All SSH key events

```sql
select
  *
from
  metal_event
where
  type like 'ssh_key.%'
```
