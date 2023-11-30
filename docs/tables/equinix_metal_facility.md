---
title: "Steampipe Table: equinix_metal_facility - Query Equinix Metal Facilities using SQL"
description: "Allows users to query Equinix Metal Facilities, providing detailed information about the physical locations where Equinix Metal services are available."
---

# Table: equinix_metal_facility - Query Equinix Metal Facilities using SQL

Equinix Metal is a bare metal cloud service that allows you to deploy physical infrastructure at global scale. Equinix Metal Facilities refer to the physical locations around the world where Equinix Metal services are available. These facilities provide the necessary infrastructure for the deployment of physical servers, storage, and networking resources.

## Table Usage Guide

The `equinix_metal_facility` table provides insights into Equinix Metal Facilities. As a system administrator or infrastructure manager, you can explore facility-specific details through this table, including location, features, and available capacity. Use it to plan and manage your Equinix Metal deployments, ensuring optimal location selection and resource allocation.

## Examples

### List all Equinix Metal facilities
Discover the segments that comprise all the Equinix Metal facilities. This would be beneficial for understanding the overall infrastructure and distribution of resources within your network.

```sql
select
  *
from
  equinix_metal_facility
```

### Get metro information for each facility
Explore the geographical relationship between facilities and metros. This query is useful for understanding where each facility is located in relation to metropolitan areas, which can help in planning logistics or assessing service coverage areas.

```sql
select
  f.name as facility_name,
  m.name as metro_name
from
  equinix_metal_facility as f,
  equinix_metal_metro as m
where
  f.metro_id = m.id
```