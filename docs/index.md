---
organization: Turbot
category: ["networking"]
icon_url: "/images/plugins/turbot/equinix_metal.svg"
brand_color: "#ED1C24"
display_name: Equinix Metal
name: equinix_metal
description: Steampipe plugin for querying Equinix Metal servers, networks, facilities and more.
---

# Equinix Metal

Query your Equinix Metal infrastructure including servers, networks, facilities and more.

## Installation

Download and install the latest Equinix Metal plugin:

```bash
steampipe plugin install equinix_metal
```

## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest equinix_metal plugin will create a connection file (`~/.steampipe/config/equinix_metal.spc`) with a single connection named `equinix_metal`. You must modify this connection to include your personal credentials.

An [API Token](https://metal.equinix.com/developers/api/) is the recommended way to set credentials. Read scope is required (write is not):

```hcl
connection "equinix_metal" {
  plugin  = "equinix_metal"
  token   = "XV1JE8QXVHdCYoWT8wbr7NRdPYrZRx3N"
}
```

Credentials are resolved in this order:
1. `token` in Steampipe config.
2. `PACKET_AUTH_TOKEN` environment variable.

## Scope

A Equinix Metal connection is scoped to a single Equinix Metal account, with a single set of credentials.
