---
title: "Steampipe Table: equinix_metal_device - Query Equinix Metal Devices using SQL"
description: "Allows users to query Equinix Metal Devices, specifically providing details about the device's ID, hostname, state, and associated project ID."
---

# Table: equinix_metal_device - Query Equinix Metal Devices using SQL

Equinix Metal is a bare metal infrastructure service that provides users with high-performance, on-demand, and globally available bare metal servers. It offers a fully automated platform for deploying infrastructure at scale, with features such as API-driven provisioning, hourly and reserved billing, and a wide range of server configuration options. Equinix Metal Devices are the individual servers that users can provision and manage within this service.

## Table Usage Guide

The `equinix_metal_device` table provides insights into individual servers within the Equinix Metal service. As a system administrator or DevOps engineer, you can explore device-specific details through this table, including the device's ID, hostname, state, and associated project ID. Use this table to manage and monitor your infrastructure, identify unused or underutilized resources, and ensure compliance with your organization's infrastructure policies.

## Examples

### List all devices
Explore all devices in your Equinix Metal account to manage and monitor your resources more effectively. This query is useful in providing a comprehensive overview of your devices, aiding in resource allocation and troubleshooting.

```sql
select
  *
from
  equinix_metal_device
```

### Get IP addresses for a device
Determine the IP addresses associated with a specific device to better understand its network connections and interactions. This information can be useful for troubleshooting network issues or for security audits.

```sql
select
  jsonb_array_elements(ip_addresses)->>'address' as ip_address
from
  equinix_metal_device
where
  hostname = 'ny5-c3-medium-x86-01'
```

### Find devices tagged as production
Explore which devices are tagged as production. This is useful for identifying and managing devices specifically used in the production environment.

```sql
select
  hostname,
  tags
from
  equinix_metal_device
where
  tags->'production' is not null
```

### Group devices by facility
Determine the distribution of devices across different facilities to understand where resources are concentrated. This can aid in resource allocation and management.

```sql
select
  f.code,
  f.name,
  count(*) as num_devices
from
  equinix_metal_device as d,
  equinix_metal_facility as f
where
  d.facility_id = f.id
group by
  f.code,
  f.name
order by
  num_devices desc
```

### List devices with OS information
Explore devices and their corresponding operating systems to gain insights into system compatibility and version distribution across your network. This could be particularly useful for IT administrators looking to maintain system uniformity or troubleshoot software issues.

```sql
select
  d.hostname,
  os.name,
  os.version
from
  equinix_metal_device as d,
  equinix_metal_operating_system as os
where
  d.operating_system_slug = os.slug
```

### Group devices by OS distribution
Discover the segments that are grouped by their operating system distribution, which can help in identifying the most commonly used systems and managing resources effectively. This can assist in making informed decisions about resource allocation and system updates.

```sql
select
  os.distro,
  count(*)
from
  equinix_metal_device as d,
  equinix_metal_operating_system as os
where
  d.operating_system_slug = os.slug
group by
  os.distro
order by
  count desc
```