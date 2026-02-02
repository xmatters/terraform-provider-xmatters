# Changelog

## [0.3.0] - TBD

FEATURES:

- Added group resource: new `xmatters_group` resource for managing groups.
- Added group roster resource: new `xmatters_group_roster` resource for managing group membership.
- Added group data sources: new `xmatters_group` and `xmatters_groups` data sources for querying groups.
- Added support for dynamic groups with search criteria configuration.
- Added comprehensive filtering capabilities for groups data source including status, group type, license type, creation date ranges, and member license type filters.

ENHANCEMENTS:

- Added GroupTypeValidator to enforce that criteria can only be set on DYNAMIC groups.
- Added ExpandGroupCriteriaObject, ExpandGroupCriterionSet functions for converting Terraform objects to xMatters SearchCriteria API format.
- Added FlattenGroupCriteriaObject, FlattenGroupCriterionObject, FlattenGroupCriterionSet functions for converting xMatters API objects to Terraform types.
- Added comprehensive unit test coverage for group resources, data sources, and utility functions.
- Added documentation for group resource and data sources with usage examples.
- Enhanced test helper utilities with group-related functions.

CHANGES:

- Updated provider registration to include group and group roster resources and data sources.
- Expanded translators module with new object types: GroupCriteriaObjectType, GroupCriterionObjectType.
- Added unit tests for all Flatten and Expand functions in translators.

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
