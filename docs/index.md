---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/equinix.svg"
brand_color: "#ED1C24"
display_name: Equinix
name: equinix
description: Steampipe plugin for querying Equinix Metal servers, networks, facilities, and more.
og_description: Query Equinix with SQL! Open source CLI. No DB required. 
og_image: "/images/plugins/turbot/equinix-social-graphic.png"
---

# Equinix + Steampipe

[Equinix Metal](https://metal.equinix.com) delivers global, interconnected, integrated, and on-demand bare-metal infrastructure as a service. As the name implies, these bare metal servers do not require virtualization or multi-tenancy. The provisioned infrastructure is deployed at Equinix data center locations allowing enabling direct network interconnects to other cloud providers and a global network backbone.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select 
  name, 
  slug,
  distro,
  version
from 
  equinix_metal_operating_system
```

```
+-----------------+---------------+----------+---------+
| name            | slug          | distro   | version |
+-----------------+---------------+----------+---------+
| CentOS 7        | centos_7      | centos   | 7       |
| CentOS 6        | centos_6      | centos   | 6       |
| Alpine 3        | alpine_3      | alpine   | 3       |
| OpenSUSE 42.3   | opensuse_42_3 | opensuse | 42.3    |
| Custom / Manual | custom        | custom   | 1       |
+-----------------+---------------+----------+---------+
```

## Documentation

- **[Table definitions & examples →](https://hub.steampipe.io/plugins/turbot/equinix/tables)**

## Get started

### Install

Download and install the latest Equinix Metal plugin:

```bash
steampipe plugin install equinix
```

### Configuration

Installing the latest Equinix plugin will create a config file (`~/.steampipe/config/equinix.spc`) with a single connection named `equinix`:

```hcl
connection "equinix" {
  plugin = "equinix"

  # API Token for your Equinix Metal account
  # Reference: https://metal.equinix.com/developers/docs/accounts/users/
  # Env variables (in order of precedence): PACKET_AUTH_TOKEN
  # token = "YOUR_EQUINIX_METAL_API_TOKEN"
}
```

### Example Configurations

- Connect to a single account:

  ```hcl
  connection "equinix" {
    plugin = "equinix"
    token  = "C1eaP7ysSR25eQBJHzWHjg2Z4RbbKZQL"
  }
  ```

- Create connections to multiple accounts:

  ```hcl
  connection "equinix1" {
    plugin = "equinix"
    token  = "K1zaP7ysFR74eQBWBzWHjg2Z9RfbKZQL"
  }

  connection "equinix2" {
    plugin = "equinix"
    token  = "PuTxs7UsHxVVpGnjhJn56mdb63xdCN7q"
  }

  connection "equinix3" {
    plugin = "equinix"
    token  = "utyTsss96BY6oENAmeAZsiwffykqUppb"
  }
  ```

## Get Involved

* Open source: https://github.com/turbot/steampipe-plugin-equinix
* Community: [Slack Channel](https://steampipe.io/community/join)
