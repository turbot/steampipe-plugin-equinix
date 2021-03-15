# Table: metal_ssh_key

List all SSH keys defined for the authenticated user.

## Examples

### List all SSH keys

```sql
select
  *
from
  metal_ssh_key
```

### Get SSH keys labelled Dwight

```sql
select
  *
from
  metal_ssh_key
where
  label = 'Dwight'
```
