---
title: "Steampipe Table: equinix_metal_event - Query Equinix Metal Events using SQL"
description: "Allows users to query Equinix Metal Events, specifically the event details, providing insights into event types, status, and associated metadata."
---

# Table: equinix_metal_event - Query Equinix Metal Events using SQL

Equinix Metal Events is a feature within the Equinix Metal platform that provides users with real-time information regarding the status and details of various events occurring within their infrastructure. This includes, but is not limited to, device state changes, operating system provisioning, and network maintenance events. It is a crucial tool for maintaining visibility and control over the infrastructure.

## Table Usage Guide

The `equinix_metal_event` table provides insights into events within the Equinix Metal platform. As a system administrator, you can explore event-specific details through this table, including event types, status, and associated metadata. Utilize it to monitor your infrastructure, track system changes, and respond to events in a timely manner.

## Examples

### List all events
Explore all events within your Equinix Metal environment to gain insights into past activities and changes. This can be useful for auditing, troubleshooting, and understanding the history of your infrastructure.

```sql
select
  *
from
  equinix_metal_event
```

### All SSH key events
Uncover the details of all events related to SSH keys. This is useful for tracking changes and maintaining security within your network.

```sql
select
  *
from
  equinix_metal_event
where
  type like 'ssh_key.%'
```