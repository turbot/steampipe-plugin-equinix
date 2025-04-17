## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#62](https://github.com/turbot/steampipe-plugin-equinix/pull/62))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#62](https://github.com/turbot/steampipe-plugin-equinix/pull/62))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#60](https://github.com/turbot/steampipe-plugin-equinix/pull/60))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#60](https://github.com/turbot/steampipe-plugin-equinix/pull/60))

## v0.7.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#55](https://github.com/turbot/steampipe-plugin-equinix/pull/55))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#55](https://github.com/turbot/steampipe-plugin-equinix/pull/55))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-equinix/blob/main/docs/LICENSE). ([#55](https://github.com/turbot/steampipe-plugin-equinix/pull/55))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#54](https://github.com/turbot/steampipe-plugin-equinix/pull/54))

## v0.6.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#46](https://github.com/turbot/steampipe-plugin-equinix/pull/46))

## v0.6.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#44](https://github.com/turbot/steampipe-plugin-equinix/pull/44))
- Recompiled plugin with Go version `1.21`. ([#44](https://github.com/turbot/steampipe-plugin-equinix/pull/44))

## v0.5.0 [2023-07-17]

_Enhancements_

- Updated the `docs/index.md` file to include multi-account configuration examples. ([#36](https://github.com/turbot/steampipe-plugin-equinix/pull/36))

## v0.4.0 [2023-04-06]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#34](https://github.com/turbot/steampipe-plugin-equinix/pull/34))

## v0.3.0 [2022-09-26]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#31](https://github.com/turbot/steampipe-plugin-equinix/pull/31))
- Recompiled plugin with Go version `1.19`. ([#31](https://github.com/turbot/steampipe-plugin-equinix/pull/31))

## v0.2.0 [2022-04-27]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#26](https://github.com/turbot/steampipe-plugin-equinix/pull/26))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#27](https://github.com/turbot/steampipe-plugin-equinix/pull/27))

## v0.1.0 [2021-12-15]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk-v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) ([#23](https://github.com/turbot/steampipe-plugin-equinix/pull/23))
- Recompiled plugin with Go version 1.17 ([#23](https://github.com/turbot/steampipe-plugin-equinix/pull/23))

## v0.0.7 [2021-09-22]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v161--2021-09-21) ([#19](https://github.com/turbot/steampipe-plugin-equinix/pull/19))
- Changed plugin license to Apache 2.0 per [turbot/steampipe](https://github.com/turbot/steampipe/issues/488) ([#17](https://github.com/turbot/steampipe-plugin-equinix/pull/17))

## v0.0.6 [2021-04-18]

_What's new?_

- New tables
  - [equinix_metal_metro](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_metro)

- New columns
  - [equinix_metal_device.metro_id](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_device)
  - [equinix_metal_facility.metro_id](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_facility)
  - [equinix_metal_plan.available_in_metros](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_plan)
  - [equinix_metal_plan.href](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_plan)

_Bug fixes_

- Fixed: Return an error if the token is not set in credentials or env var ([#15](https://github.com/turbot/steampipe-plugin-equinix/issues/15))


## v0.0.5 [2021-04-02]

_Bug fixes_

- Fixed: `Table definitions & examples` link now points to the correct location ([#12](https://github.com/turbot/steampipe-plugin-equinix/pull/12))


## v0.0.4 [2021-03-31]

_Bug fixes_

- Fix HTML line break in docs home page


## v0.0.3 [2021-03-31]

_What's new?_

- Updated README and docs home page


## v0.0.2 [2021-03-31]

_What's new?_

- New tables added
  - [equinix_metal_capacity](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_capacity)
  - [equinix_metal_spot_price](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_spot_price)


## v0.0.1 [2021-03-19]

_What's new?_

- New tables added
  - [equinix_metal_device](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_device)
  - [equinix_metal_event](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_event)
  - [equinix_metal_facility](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_facility)
  - [equinix_metal_operating_system](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_operating_system)
  - [equinix_metal_organization](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_organization)
  - [equinix_metal_plan](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_plan)
  - [equinix_metal_project](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_project)
  - [equinix_metal_ssh_key](https://hub.steampipe.io/plugins/turbot/equinix/tables/equinix_metal_ssh_key)
