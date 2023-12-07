---
title: "Steampipe Table: equinix_metal_operating_system - Query Equinix Metal Operating Systems using SQL"
description: "Allows users to query Operating Systems in Equinix Metal, specifically providing details about the name, version, and other relevant properties of the operating system."
---

# Table: equinix_metal_operating_system - Query Equinix Metal Operating Systems using SQL

Equinix Metal Operating System is a resource within the Equinix Metal infrastructure that defines the software on which a computer system runs. It manages the hardware resources of a computer and hosts applications that run on the computer. Equinix Metal Operating System is a key component in ensuring the efficient operation of your applications and services in the Equinix Metal environment.

## Table Usage Guide

The `equinix_metal_operating_system` table provides insights into the operating systems within Equinix Metal. As a systems administrator or DevOps engineer, you can explore details about each operating system through this table, including its name, version, and other relevant properties. Utilize it to manage and optimize the performance of your applications and services running on these operating systems.

## Examples

### List all
Explore all available operating systems in the Equinix Metal service, sorted by name. This can be useful in determining which operating systems are available for deployment.

```sql+postgres
select
  *
from
  equinix_metal_operating_system
order by
  name;
```

```sql+sqlite
select
  *
from
  equinix_metal_operating_system
order by
  name;
```

### Group by distribution
Analyze the distribution of different operating systems in use within your Equinix Metal infrastructure. This can help identify the most commonly used systems and inform decisions regarding system support and maintenance.

```sql+postgres
select
  distro,
  count(*)
from
  equinix_metal_operating_system
group by
  distro
order by
  count desc;
```

```sql+sqlite
select
  distro,
  count(*)
from
  equinix_metal_operating_system
group by
  distro
order by
  count(*) desc;
```

### List all Ubuntu OS versions available
Explore the various Ubuntu operating system versions available to understand the range of options for your system setup. This is beneficial for assessing compatibility and planning system upgrades.

```sql+postgres
select
  *
from
  equinix_metal_operating_system
where
  distro = 'ubuntu'
order by
  name;
```

```sql+sqlite
select
  *
from
  equinix_metal_operating_system
where
  distro = 'ubuntu'
order by
  name;
```

### Which operating systems can be provisioned on a c3.medium.x86 device?
Determine the variety of operating systems that can be installed on a specific device model. This is useful for understanding the flexibility and compatibility of different devices in your network.

```sql+postgres
select
  *
from
  equinix_metal_operating_system
where
  provisionable_on ? 'c3.medium.x86';
```

```sql+sqlite
Error: SQLite does not support the '?' operator used for JSON objects in PostgreSQL.
```

### List all OS device combinations
Explore various operating system and device combinations that can be provisioned on Equinix Metal. This can be useful for assessing compatibility and planning infrastructure deployment.

```sql+postgres
select
  slug,
  device
from
  equinix_metal_operating_system,
  jsonb_array_elements_text(provisionable_on) as device;
```

```sql+sqlite
select
  slug,
  device.value
from
  equinix_metal_operating_system,
  json_each(provisionable_on) as device;
```