---
title: "Steampipe Table: equinix_metal_capacity - Query Equinix Metal Capacities using SQL"
description: "Allows users to query Capacities in Equinix Metal, specifically the available and utilized capacity data, providing insights into resource usage and availability."
---

# Table: equinix_metal_capacity - Query Equinix Metal Capacities using SQL

Equinix Metal is a bare metal cloud service that provides physical servers with the automation and flexibility of a virtual machine. The service offers a range of server types for different workloads and needs, from single-tenant servers for maximum privacy and performance, to multi-tenant servers for cost-effective scaling. It is a globally distributed platform with data centers in multiple locations around the world.

## Table Usage Guide

The `equinix_metal_capacity` table provides insights into available and utilized capacities within Equinix Metal. As a cloud architect or system administrator, explore capacity-specific details through this table, including available and utilized server types, plan slugs, and associated metadata. Utilize it to monitor resource usage, manage capacity planning, and optimize the allocation of resources.

## Examples

### List all capacities
Discover the segments that allow you to analyze available resources within different facilities and plans, helping you manage and allocate resources effectively.

```sql+postgres
select
  *
from
  equinix_metal_capacity
order by
  facility,
  plan;
```

```sql+sqlite
select
  *
from
  equinix_metal_capacity
order by
  facility,
  plan;
```

### Show all servers with limited capacity
Discover the segments that are operating at a limited capacity in your infrastructure. This is particularly useful for planning resources and preventing potential shortages.

```sql+postgres
select
  facility,
  plan
from
  equinix_metal_capacity
where
  level = 'limited'
order by
  facility,
  plan;
```

```sql+sqlite
select
  facility,
  plan
from
  equinix_metal_capacity
where
  level = 'limited'
order by
  facility,
  plan;
```