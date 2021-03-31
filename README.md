## Note for users new to Steampipe
The Equinix Plugin for Steampipe can be managed automatically with the Steampipe CLI.
For more information on how to get started view the [documentation](https://hub.steampipe.io/plugins/turbot/equinix).

# The Equinix Plugin for Steampipe

Use SQL to query infrastructure including servers, networks, facilities and more from Equinix Metal. For example:

```sql
select hostname, ip_addresses from equinix_metal_device
```

- [Documentation](https://hub.steampipe.io/plugins/turbot/equinix)
- [Tables & schemas](https://hub.steampipe.io/plugins/turbot/equinix/tables)
- [Equinix plugin issues](https://github.com/turbot/steampipe-plugin-equinix/issues)
- [Steampipe issues](https://github.com/turbot/steampipe/issues)
- [Discussion forums](https://github.com/turbot/steampipe/discussions)

## Prerequisites

- Install [Steampipe](https://steampipe.io/downloads) v0.3.3 or greater
- Install [Golang](https://golang.org/doc/install) 1.15.x or greater

## Building the Plugin

Clone the repository:

```sh
$ mkdir -p $GOPATH/src/github.com/turbot; cd $GOPATH/src/github.com/turbot
$ git clone git@github.com:turbot/steampipe-plugin-equinix
```


Enter the plugin directory and build the plugin:

```sh
$ cd $GOPATH/src/github.com/turbot/steampipe-plugin-equinix
$ make
```


Copy configuration files from repo to local steampipe config:
```sh
$ cp $GOPATH/src/github.com/turbot/steampipe-plugin-equinix/config/*.spc ~/.steampipe/config
```


Read and update the configuration file setting(s) for your environment [see the plugin docs](https://hub.steampipe.io/plugins/turbot/equinix) for additional info:
```sh
$ pico ~/.steampipe/config/equinix.spc
```

## Using the Plugin

During the `make` process, the script will output the plugin to `~/.steampipe/plugins/hub.steampipe.io/plugins/turbot/equinix@latest/` which is the default location for steampipe plugins. Restart Steampipe if already running. Then try a test query:

```bash
$ steampipe query "select hostname from equinix_metal_device"
```

## Developing the Plugin

To add features to the Plugin, install [Go](http://www.golang.org) and configure your your [GOPATH](http://golang.org/doc/code.html#GOPATH)

Compile the Plugin by running `make`. The Plugin binary will output to your Steampipe plugin directory.

```sh
$ make
```

## Community

The Steampipe community can be found on [GitHub Discussions](https://github.com/turbot/steampipe/discussions), where you can ask questions, voice ideas, and share your projects. Our [Code of Conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md) applies to all Steampipe community channels.

## Contributing

Please see [CONTRIBUTING.md](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md).

`Help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Equinix Plugin](https://github.com/turbot/steampipe-plugin-equinix/labels/help%20wanted)

## License

By contributing to Steampipe and Steampipe plugins you agree that your contributions will be licensed as defined on the [LICENSE](LICENSE) file.