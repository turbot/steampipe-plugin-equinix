---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/equinix.svg"
brand_color: "#ED1C24"
display_name: Equinix
name: equinix
description: Steampipe plugin for querying Equinix Metal servers, networks, facilities and more.
og_description: Query Equinix with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/equinix-social-graphic.png"
---

# Equinix + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[Equinix Metal](https://metal.equinix.com) delivers global, interconnected, integrated, and on-demand bare-metal infrastructure as a service. As the name implies, these bare metal servers do not require virtualization or multi-tenancy. The provisioned infrastructure is deployed at Equinix data center locations allowing enabling direct network interconnects to other cloud providers and a global network backbone.

For example:

```sql
select
  hostname,
  state,
  tags
from
  equinix_metal_device
```

```
+-----------------------+--------+---------------+
| hostname              | state  | tags          |
+-----------------------+--------+---------------+
| dc13-c3.medium.x86-01 | active | {"prod":true} |
+-----------------------+--------+---------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/equinix/tables)**

## Get started

### Install

Download and install the latest Equinix Metal plugin:

```bash
steampipe plugin install equinix
```

### Credentials

| Item | Description |
| - | - |
| Credentials | Use an [API Token](https://metal.equinix.com/developers/api/). |
| Permissions | `Read-only` |
| Radius | Each connection represents a single Equinix Metal account. |
| Resolution |  1. `token` in Steampipe config.<br />2. `PACKET_AUTH_TOKEN` environment variable. (A legacy name.) |

### Configuration

Installing the latest equinix plugin will create a config file (`~/.steampipe/config/equinix.spc`) with a single connection named `equinix`:

```hcl
connection "equinix" {
  plugin  = "equinix"
  token   = "XV1JE4QXVHdCYoWT2wbr2NRdPYrZRx3N"
}
```

## Multi-Account Connections

You may create multiple equinix connections:

```hcl
connection "equinix_dev" {
  plugin    = "equinix"
  token     = "XV1KL4QXVHcCYoWT2wbr2NRdRTrZRx3K"
}
connection "equinix_qa" {
  plugin    = "equinix"
  token     = "XV1KL4QXVHcCYoWT2wbr2NRdRTrZRx3L"
}
connection "equinix_prod" {
  plugin    = "equinix"
  token     = "XV1KL4QXVHcCYoWT2wbr2NRdRTrZRx3M"
}
```

Each connection is implemented as a distinct [Postgres schema](https://www.postgresql.org/docs/current/ddl-schemas.html). As such, you can use qualified table names to query a specific connection:

```sql
select * from equinix_qa.equinix_metal_organization
```

You can multi-account connections by using an [**aggregator** connection](https://steampipe.io/docs/using-steampipe/managing-connections#using-aggregators). Aggregators allow you to query data from multiple connections for a plugin as if they are a single connection.

```hcl
connection "equinix_all" {
  plugin      = "equinix"
  type        = "aggregator"
  connections = ["equinix_dev", "equinix_qa", "equinix_prod"]
}
```

Querying tables from this connection will return results from the `equinix_dev`, `equinix_qa`, and `equinix_prod` connections:

```sql
select * from equinix_all.equinix_metal_organization
```

Alternatively, can use an unqualified name and it will be resolved according to the [Search Path](https://steampipe.io/docs/guides/search-path). It's a good idea to name your aggregator first alphbetically, so that it is the first connection in the search path (i.e. `equinix_all` comes before `equinix_dev`):

```sql
select * from equinix_metal_organization
```

Steampipe supports the `*` wildcard in the connection names. For example, to aggregate all the equinix plugin connections whose names begin with `equinix_`:

```hcl
connection "equinix_all" {
  type        = "aggregator"
  plugin      = "equinix"
  connections = ["equinix_*"]
}
```

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-equinix
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
