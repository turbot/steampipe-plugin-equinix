![image](https://hub.steampipe.io/images/plugins/turbot/equinix-social-graphic.png)

# Equinix Plugin for Steampipe

Use SQL to query infrastructure including servers, networks, facilities and more from Equinix Metal.

* **[Get started →](https://hub.steampipe.io/plugins/turbot/equinix)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/equinix/tables)
* Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-equinix/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):
```shell
steampipe plugin install equinix
```

Run a query:
```sql
select hostname, ip_addresses from equinix_metal_device
```

## Developing

Prerequisites:
- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-equinix.git
cd steampipe-plugin-equinix
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:
```
make
```

Configure the plugin:
```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/equinix.spc
```

Try it!
```
steampipe query
> .inspect equinix
```

Further reading:
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-equinix/blob/main/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Equinix Plugin](https://github.com/turbot/steampipe-plugin-equinix/labels/help%20wanted)
