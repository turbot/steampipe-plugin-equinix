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

- **[Table definitions & examples →](equinix/tables)**

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

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-equinix
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)