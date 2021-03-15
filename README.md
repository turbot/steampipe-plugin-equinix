
<p align="center">
    <h1 align="center">Equinix Metal Plugin for Steampipe</h1>
</p>
<p align="center">
  <a aria-label="Steampipe logo" href="https://steampipe.io">
    <img src="https://steampipe.io/images/steampipe_logo_wordmark_padding.svg" height="28">
  </a>
  <a aria-label="License" href="LICENSE">
    <img alt="" src="https://img.shields.io/static/v1?label=license&message=MPL-2.0&style=for-the-badge&labelColor=777777&color=F3F1F0">
  </a>
</p>

## DRAFT - Under active development

This plugin is under active development and not yet available on the Steampipe Hub. It can be built by:

```
git clone https://github.com/turbot/steampipe-plugin-equinix-metal.git
cd steampipe-plugin-equinix-metal
make
cp config/*.spc ~/.steampipe/config
# Set your token inside ~/.steampipe/config/equinix_metal.spc
vi ~/.steampipe/config/equinix_metal.spc
steampipe query
steampipe> select * from metal_project
```

## Query Equinix Metal with SQL

Use SQL to Query infrastructure including servers, networks, facilities and more from Equinix Metal. For example:

```sql
select
  hostname,
  ip_addresses
from
  metal_device
```

Learn about [Steampipe](https://steampipe.io/).

## Get started

**[Table documentation and examples &rarr;](https://hub.steampipe.io/plugins/turbot/equinix-metal)**

Install the plugin:

```shell
steampipe plugin install equinix-metal
```

## Get involved

### Community

The Steampipe community can be found on [GitHub Discussions](https://github.com/turbot/steampipe/discussions), where you can ask questions, voice ideas, and share your projects.

Our [Code of Conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md) applies to all Steampipe community channels.

### Contributing

Please see [CONTRIBUTING.md](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md).
