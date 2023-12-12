---
title: "Steampipe Table: equinix_metal_organization - Query Equinix Metal Organizations using SQL"
description: "Allows users to query Equinix Metal Organizations, specifically providing insights into the organization's details, including ID, name, description, and created and updated timestamps."
---

# Table: equinix_metal_organization - Query Equinix Metal Organizations using SQL

Equinix Metal Organizations are a fundamental resource in the Equinix Metal cloud infrastructure platform. They represent a collection of users, with billing and permissions managed at the organization level. Organizations are the top level of the Equinix Metal resource hierarchy, and they contain projects which in turn contain the resources.

## Table Usage Guide

The `equinix_metal_organization` table provides insights into Equinix Metal Organizations within the Equinix Metal cloud infrastructure platform. As a cloud engineer or administrator, explore organization-specific details through this table, including ID, name, description, and created and updated timestamps. Utilize it to uncover information about organizations, such as their associated users, billing details, and permission settings.

## Examples

### List all organizations
Explore the various organizations within your Equinix Metal account to understand their structure and relationships. This is useful for managing and organizing resources within a large account.

```sql+postgres
select
  *
from
  equinix_metal_organization;
```

```sql+sqlite
select
  *
from
  equinix_metal_organization;
```

### List all projects for the organization
Explore which projects are associated with your organization. This is useful for gaining a comprehensive view of all ongoing projects, allowing for better management and coordination.

```sql+postgres
select
  o.id as org_id,
  o.name as org_name,
  p.id as project_id,
  p.name as project_name
from
  equinix_metal_organization as o,
  jsonb_array_elements_text(o.project_ids) as opid,
  equinix_metal_project as p
where
  p.id = opid;
```

```sql+sqlite
select
  o.id as org_id,
  o.name as org_name,
  p.id as project_id,
  p.name as project_name
from
  equinix_metal_organization as o,
  json_each(o.project_ids) as opid,
  equinix_metal_project as p
where
  p.id = opid.value;
```