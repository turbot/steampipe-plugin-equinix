---
title: "Steampipe Table: equinix_metal_metro - Query Equinix Metal Metros using SQL"
description: "Allows users to query Equinix Metal Metros, specifically the metro areas where Equinix Metal has facilities, providing insights into the geographical distribution and availability of these facilities."
---

# Table: equinix_metal_metro - Query Equinix Metal Metros using SQL

Equinix Metal Metros are geographical areas where Equinix Metal has facilities. These metro areas are significant for users as they provide the physical locations where users can deploy and manage their infrastructure. The metro areas cover multiple regions worldwide, offering users a wide range of options for infrastructure deployment.

## Table Usage Guide

The `equinix_metal_metro` table provides insights into the metro areas where Equinix Metal has facilities. As a DevOps engineer, explore metro-specific details through this table, including the geographical distribution and availability of these facilities. Utilize it to uncover information about metros, such as their code, name, and country, aiding in the selection of suitable locations for infrastructure deployment.

## Examples

### List all Equinix Metal metros
Explore all the metropolitan locations where Equinix Metal services are available. This allows you to identify where you might deploy your infrastructure according to your business needs.

```sql+postgres
select
  *
from
  equinix_metal_metro;
```

```sql+sqlite
select
  *
from
  equinix_metal_metro;
```