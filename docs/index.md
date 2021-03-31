---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/equinix.svg"
brand_color: "#ED1C24"
display_name: Equinix
name: equinix
description: Steampipe plugin for querying Equinix Metal servers, networks, facilities and more.
social_about: Use SQL to query infrastructure resources (e.g. servers, networks) from Equinix Metal. Open source CLI. No DB required. 
social_preview: "/images/plugins/turbot/equinix-social-graphic.png"
---

# Steampipe

The Steampipe CLI is open source software that allows you to perform real-time queries against cloud APIs using SQL; all without having to extract, transform and load data into a local DB. If you are just getting started checkout [steampipe.io](https://steampipe.io).

# Equinix

[Equinix Metal](https://metal.equinix.com) delivers global, interconnected, integrated, and on-demand bare-metal infrastructure as a service. As the name implies, these bare metal servers do not require virtualization or multi-tenancy. The provisioned infrastructure is deployed at Equinix data center locations allowing enabling direct network interconnects to other cloud providers and a global network backbone.

# Steampipe + Equinix

Steampipe allows you to query [Equinix Metal APIs](https://metal.equinix.com/developers/api/) with SQL; retrieving metadata about your servers, networks, facilities and more:

```sql
select 
  hostname, 
  ip_addresses,
  tags
from 
  equinix_metal_device;
```

```sql
select
  o.name as org_name,
  p.name as project_name,
  p.id as project_id
from
  equinix_metal_organization as o,
  jsonb_array_elements_text(o.project_ids) as opid,
  equinix_metal_project as p
where
  p.id = opid
```
```
+----------+--------------+--------------------------------------+
| org_name | project_name | project_id                           |
+----------+--------------+--------------------------------------+
| DMI      | WUPHF        | d2fe44f2-aa42-4424-be44-44f2b442faff |
| DMI      | DM Infinity  | 22da22ce-2b2d-44bf-224c-cc24ef244a24 |
+----------+--------------+--------------------------------------+
```

Browse all [available tables and their schemas](equinix/tables).

## Installation

Download and install the latest Equinix Metal plugin:

```bash
steampipe plugin install equinix
```

## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest equinix plugin will create a connection file (`~/.steampipe/config/equinix.spc`) with a single connection named `equinix`. You must modify this connection to include your personal credentials.

An [API Token](https://metal.equinix.com/developers/api/) is the recommended way to set credentials. Read scope is required (write is not):

```hcl
connection "equinix" {
  plugin  = "equinix"
  token   = "XV1JE4QXVHdCYoWT2wbr2NRdPYrZRx3N"
}
```

Credentials are resolved in this order:
1. `token` in Steampipe config.
2. `PACKET_AUTH_TOKEN` environment variable. (A legacy name.)

## Scope

A Equinix Metal connection is scoped to a single Equinix Metal account, with a single set of credentials.
