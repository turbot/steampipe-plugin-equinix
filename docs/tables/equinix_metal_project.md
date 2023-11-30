---
title: "Steampipe Table: equinix_metal_project - Query Equinix Metal Projects using SQL"
description: "Allows users to query Equinix Metal Projects, providing detailed information about each project's properties and configuration."
---

# Table: equinix_metal_project - Query Equinix Metal Projects using SQL 

Equinix Metal Projects are organizational entities in the Equinix Metal platform that allow you to manage and group your resources. They provide a way to control user access, permissions, and billing. Projects are especially useful when you need to isolate resources for different environments, departments, or billing purposes.

## Table Usage Guide

The `equinix_metal_project` table provides insights into Projects within the Equinix Metal platform. As a system administrator, you can explore project-specific details through this table, including project IDs, names, and associated metadata. Utilize it to manage and monitor your Equinix Metal resources effectively, ensuring proper access control and resource allocation.

## Examples

### List all projects
Explore all active projects within your organization to gain a comprehensive overview of ongoing work and resources. This can assist in effective resource allocation and project management.

```sql
select
  *
from
  equinix_metal_project
```

### Project with Organization information
Explore the association between different projects and their corresponding organizations. This can be useful to understand the organizational structure and distribution of projects within your company.

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