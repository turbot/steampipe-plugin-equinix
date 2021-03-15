# Table: metal_project

Projects the authenticated user has access to.

## Examples

### List all projects

```sql
select
  *
from
  metal_project
```

### Project with Organization information

```sql
select
  p.name as project_name,
  o.name as org_name
from
  metal_project as p,
  metal_organization as o
where
  p.organization_id = o.id
```
