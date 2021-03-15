# Table: metal_organization

Organizations the authenticated user has access to.

## Examples

### List all organizations

```sql
select
  *
from
  metal_organization
```

### List all projects for the organization

```sql
select
  o.id as org_id,
  o.name as org_name,
  p.id as project_id,
  p.name as project_name
from
  metal_organization as o,
  jsonb_array_elements_text(o.project_ids) as opid,
  metal_project as p
where
  p.id = opid
```
