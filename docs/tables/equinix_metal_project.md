# Table: equinix_metal_project

Projects the authenticated user has access to.

## Examples

### List all projects

```sql
select
  *
from
  equinix_metal_project
```

### Project with Organization information

```sql
select
  p.name as project_name,
  o.name as org_name
from
  equinix_metal_project as p,
  equinix_metal_organization as o
where
  p.organization_id = o.id
```
