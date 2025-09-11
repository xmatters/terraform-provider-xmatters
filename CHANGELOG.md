# Changelog

## [0.3.0] - 2025-09-11

ENHANCEMENTS:

- Added group support:
  - `xmatters_group` resource.
  - `xmatters_group` and `xmatters_groups` data sources.
  - `xmatters_group_roster` resource for managing group membership.
- Added documentation and example Terraform configs for group, group roster, and groups data source.

CHANGES:

- Improved example values and documentation for all group-related resources and data sources.

## [0.2.0] - 2025-09-09

ENHANCEMENTS:

- Added device support: new `xmatters_device` resource and `xmatters_device`/`xmatters_devices` data sources.
- Added documentation for device resource and data sources.
- Added example Terraform configs for device resource and data sources.
- Added unit tests for device resource and data sources.
- Upgraded dependency: `github.com/xmatters/xmatters-go` to v0.2.0.

CHANGES:

- Updated person resource documentation for externally_owned default value.
- Improved device-related test helpers and expanders.
- Removed unused custom sequence type and related tests.

## [0.1.1] - 2025-05-28

- Fix some documentation issues.

## [0.1.0] - 2025-05-22

- Initial release.
