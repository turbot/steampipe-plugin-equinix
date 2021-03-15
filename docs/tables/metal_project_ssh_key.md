# Table: metal_project_ssh_key

List all SSH keys for all projects.

## Examples

### List all SSH keys for all projects

```sql
select
  *
from
  metal_project_ssh_key
```

### Get SSH keys labelled Dwight

```sql
select
  *
from
  metal_project_ssh_key
where
  label = 'Dwight'
```

### List all SSH keys for a specific project

```sql
select
  *
from
  metal_project_ssh_key
where
  project_id = 'd8fe44f9-ac39-4482-be30-11f7b159faff'
```
