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
