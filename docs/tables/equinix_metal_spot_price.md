---
title: "Steampipe Table: equinix_metal_spot_price - Query Equinix Metal Spot Prices using SQL"
description: "Allows users to query Equinix Metal Spot Prices, specifically the current and historical prices for different Equinix Metal instances."
---

# Table: equinix_metal_spot_price - Query Equinix Metal Spot Prices using SQL

Equinix Metal Spot Price is a feature of the Equinix Metal cloud service that allows users to bid on spare Equinix Metal computing capacity. This feature offers an opportunity for users to run their applications at a lower cost. The prices for these instances fluctuate based on supply and demand for instances.

## Table Usage Guide

The `equinix_metal_spot_price` table provides insights into the current and historical prices for different Equinix Metal instances. If you are a cost-conscious cloud architect or a DevOps engineer, you can use this table to track price trends and make cost-effective decisions about resource allocation. By analyzing the data, you can identify the most affordable instances and times to run your applications on Equinix Metal.

## Examples

### List all spot prices
Explore the varying spot prices across different facilities and plans. This is useful for comparing and making informed decisions on the most cost-effective options.

```sql
select
  *
from
  equinix_metal_spot_price
order by
  facility,
  plan
```

### Find the cheapest facility to run a c3.medium.x86
Discover the most cost-effective facility to operate a 'c3.medium.x86' by sorting through available options based on their current price. This is beneficial for budget optimization and cost-effective resource allocation.

```sql
select
  *
from
  equinix_metal_spot_price
where
  plan = 'c3.medium.x86'
order by
  current_price
```

### 10 cheapest servers
Discover the segments that offer the most cost-effective server options. This query is useful for budget-conscious businesses looking to minimize their server expenses.

```sql
select
  *
from
  equinix_metal_spot_price
order by
  current_price
limit
  10
```