---
title: "Steampipe Table: equinix_metal_plan - Query Equinix Metal Plans using SQL"
description: "Allows users to query Equinix Metal Plans, specifically providing information about the various plans available on Equinix Metal, including details such as the plan's ID, slug, and description."
---

# Table: equinix_metal_plan - Query Equinix Metal Plans using SQL

Equinix Metal Plans represent the various hardware and configuration options available to users on the Equinix Metal platform. These plans detail the specific configurations of the server, including the CPU, memory, storage, and network capabilities. It allows users to choose a plan that best fits their needs, based on their requirements and the workload they intend to run.

## Table Usage Guide

The `equinix_metal_plan` table provides insights into the various plans available within the Equinix Metal platform. As a system administrator or DevOps engineer, explore plan-specific details through this table, including the plan's ID, slug, and description. Utilize it to uncover information about the various hardware and configuration options available, allowing you to choose a plan that best fits your needs and workload.

## Examples

### List all plans
Explore all available plans within your Equinix Metal environment to better understand the resources and services available to you. This can help in strategic planning and optimizing resource allocation.

```sql+postgres
select
  *
from
  equinix_metal_plan;
```

```sql+sqlite
select
  *
from
  equinix_metal_plan;
```

### Get CPU details for each plan
Explore the specific CPU details for each plan to better understand the resources allocated. This can assist in identifying the most suitable plan based on your CPU requirements.

```sql+postgres
select
  name,
  cpu ->> 'count' as cpu_count,
  cpu ->> 'type' as cpu_count
from
  equinix_metal_plan,
  jsonb_array_elements(specs -> 'cpus') as cpu;
```

```sql+sqlite
select
  name,
  json_extract(cpu.value, '$.count') as cpu_count,
  json_extract(cpu.value, '$.type') as cpu_type
from
  equinix_metal_plan,
  json_each(specs, '$.cpus') as cpu;
```

### Plans by price
Analyze the pricing structure of various plans to understand the hourly cost for each, helping you make informed decisions about which plan suits your budget best.

```sql+postgres
select
  name,
  pricing ->> 'hour' as hourly_cost
from
  equinix_metal_plan
order by
  hourly_cost;
```

```sql+sqlite
select
  name,
  json_extract(pricing, '$.hour') as hourly_cost
from
  equinix_metal_plan
order by
  hourly_cost;
```

### Plans available by metro
Analyze the settings to understand the availability of different plans across various metropolitan areas. This is useful for assessing the distribution and accessibility of services in different urban regions.

```sql+postgres
select
  m.name as metro_name,
  p.name as plan_name
from
  equinix_metal_plan as p,
  jsonb_array_elements_text(p.available_in_metros) as mid,
  equinix_metal_metro as m
where
  m.id = mid
order by
  m.name,
  p.name;
```

```sql+sqlite
select
  m.name as metro_name,
  p.name as plan_name
from
  equinix_metal_plan as p,
  json_each(p.available_in_metros) as mid,
  equinix_metal_metro as m
where
  m.id = mid.value
order by
  m.name,
  p.name;
```