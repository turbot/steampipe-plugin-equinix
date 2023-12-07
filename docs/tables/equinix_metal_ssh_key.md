---
title: "Steampipe Table: equinix_metal_ssh_key - Query Equinix Metal SSH Keys using SQL"
description: "Allows users to query Equinix Metal SSH Keys, specifically their details and configurations, providing insights into the management and usage of SSH Keys within the Equinix Metal service."
---

# Table: equinix_metal_ssh_key - Query Equinix Metal SSH Keys using SQL

Equinix Metal SSH Keys are resources within the Equinix Metal service that enable secure shell (SSH) access to devices. They are used to authenticate users and ensure secure communication between devices within the network. SSH Keys are crucial for managing and controlling access to devices, and ensuring the security of communications.

## Table Usage Guide

The `equinix_metal_ssh_key` table provides insights into SSH Keys within the Equinix Metal service. As a network administrator, you can explore key-specific details through this table, including their labels, IDs, and created and updated timestamps. Utilize it to uncover information about keys, such as their ownership, usage, and the devices they are associated with, helping you manage access control and maintain the security of your network.

## Examples

### List all SSH keys
Explore the full range of SSH keys available in your Equinix Metal environment with this query. It's useful for verifying key access and ensuring appropriate security measures are in place.

```sql+postgres
select
  *
from
  equinix_metal_ssh_key;
```

```sql+sqlite
select
  *
from
  equinix_metal_ssh_key;
```

### Get SSH keys labelled Dwight
Discover the segments that use SSH keys labelled 'Dwight'. This can be beneficial in understanding the distribution and usage of specific SSH keys within your infrastructure, enhancing overall security management.

```sql+postgres
select
  *
from
  equinix_metal_ssh_key
where
  label = 'Dwight';
```

```sql+sqlite
select
  *
from
  equinix_metal_ssh_key
where
  label = 'Dwight';
```