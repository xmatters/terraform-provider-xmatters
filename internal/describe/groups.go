package describe

const (
	// --------------------------------------------------
	// Group(s) Data Source
	// --------------------------------------------------
	// Group Search Fields Schema
	GroupIDSearch = "Unique identifier (UUID) or target name of the xMatters group."
	// Group Data Source Schema
	GroupID              = "Unique identifier (UUID) of the xMatters group."
	GroupTargetName      = "Name used to identify this group."
	GroupDescription     = "A description of the group, can be a maximum of 1024 characters."
	GroupStatus          = "Status of the group in xMatters.  Available options are: ACTIVE, INACTIVE."
	GroupExternalKey     = "Unique identifier of a group in an external system."
	GroupExternallyOwned = "Whether the group is managed by an external system."
	GroupAllowDuplicates = "When set to true group members can be added to an escalation timeline multiple times and recipients can receive multiple notifications for the same event targeting the group."
	GroupTimezone        = "Default time zone of the group represented by the the two-letter country code followed by the city name."
	GroupSite            = "Unique identifier (UUID) of the site where the group is assigned."
	GroupObservedByAll   = "True if groups can locate and send notifications to the group regardless of their role. If this value is false, only groups who have the selected roles can observe the group."
	GroupObservers       = "Unique identifier (UUID) or target name of the role or roles set as observers for the group. Adding observer roles via the xMatters REST API overwrites any existing observers for the group. You can only add specific observers to a group if observedByAll is set to false."
	GroupSupervisors     = "A list of the supervisors of the group."
	GroupServices        = "List of available xMatters services."
	GroupType            = "The type of group in xMatters. Available options are BROADCAST, DYNAMIC, and ON_CALL."
	// Groups Data Source Schema
	GroupsSearch  = "Basic Search parameters to find a list of xMatters groups."
	GroupsFilters = "Search Filters to improve the search results."
	GroupsOptions = "Sorting Options for search results."
	GroupsList    = "List of available xMatters groups."
	// Groups Data Source Search Schema
	GroupsSearchTerms   = "List of search terms, separated by spaces. Used with the operand and fields parameters to expand or limit search results. When two or more search terms are present, the result includes groups that match either search term."
	GroupsSearchOperand = "Operands expand or limit the search query parameter. OR is the default operand and returns groups that have any of the search terms in the name or description. AND returns groups that have all search terms in the name or description."
	GroupsSearchFields  = "Set the fields to search when a search term is specified. Available options are: NAME, DESCRIPTION, SERVICE_NAME."
	// Groups Data Source Filters Schema
	GroupsFilterGroupType          = "Returns a list of groups of the specified type: Available options are: ON_CALL, BROADCAST, DYNAMIC."
	GroupsFilterMemberExists       = "Returns a list of groups that have shifts created, but no members added to the shifts. Available options are: ALL_SHIFTS: Returns a list of groups that have no members added to any shifts. ANY_SHIFTS: Returns a list of groups that have at least one shift with no members."
	GroupsFilterMembers            = "Unique identifier (UUID) or Base64 encoded target name of a user, device, or group. Returns a list of groups that contains at least one of the specified members."
	GroupsFilterMembersLicenseType = "Returns a list of groups that contain at least one member (or a device that belongs to a user) who has the specified license type. The member does not have to be part of any shifts for the group to be included in the response. Available values are: FULL_USER, STAKEHOLDER_USER."
	GroupsFilterSites              = "Returns a list of groups for the specified sites. You can specify the site using its unique identifier (UUID) or target name (case-insensitive), or both. When two or more sites are sent in the request, the response includes groups for which either site is assigned."
	GroupsFilterStatus             = "Returns all groups with the specified status. Available options are: ACTIVE, INACTIVE."
	GroupsFilterSupervisors        = "Target names or unique identifiers (UUID) of group supervisors. Returns a list of groups assigned to the specified supervisors. Values can be combined in a comma-separated list. When multiple supervisors are specified, the response returns users who are assigned to at least one of the supervisors."
	GroupsFilterCreatedAfter       = "Timestamp in ISO format. Returns a list of groups created after the provided value. Can be used on its own or in conjunction with the ‘createdBefore’ and ‘createdTo’ parameters."
	GroupsFilterCreatedBefore      = "Timestamp in ISO format. Returns a list of groups created before the provided value. Can be used on its own or in conjunction with the ‘createdAfter’ and ‘createdFrom’ parameters."
	GroupsFilterCreatedFrom        = "Timestamp in ISO format. Returns a list of groups created at or after the provided value. Can be used on its own or in conjunction with the ‘createdTo’ and ‘createdBefore’ parameters. "
	GroupsFilterCreatedTo          = "Timestamp in ISO format. Returns a list of groups created up to and until the provided value. Can be used on its own or in conjunction with the ‘createdFrom’ and ‘createdAfter’ parameters."
	// Groups Data Source Options Schema
	GroupsOptionsSortBy    = "Criteria for sorting the returned results. Available options are: NAME, GROUPTYPE, STATUS."
	GroupsOptionsSortOrder = "Order in which returned groups are sorted. Available options are: ASCENDING, DESCENDING."

	// --------------------------------------------------
	// Group Resource
	// --------------------------------------------------
	// Group Resource Schema
	GroupResourceID                = "Unique identifier (UUID) of the xMatters group."
	GroupResourceTargetName        = "Name of the group, up to a maximum of 100 characters."
	GroupResourceStatus            = "Status of the group in xMatters. Available options are: ACTIVE, INACTIVE. Default is ACTIVE."
	GroupResourceDescription       = "Description of the group, to a maximum of 1024 characters."
	GroupResourceType              = "Type of group to create in xMatters. Available options are: ON_CALL, DYNAMIC, BROADCAST"
	GroupResourceAllowDuplicates   = "If set to 'true', group members can be added to an escalation timeline multiple times, and recipients can receive more than one notification for the same event targeting the group."
	GroupResourceSite              = "Unique identifier (UUID) of the site that the group uses for holidays. If this value is not provided, the group is set to not use site holidays."
	GroupResourceObservedByAll     = "If set to 'true' all roles can view and target this group. If set to 'false', use the 'observers' parameter to set which roles can target the group."
	GroupResourceObservers         = "Role or roles set as observers for group when the 'observed_by_all' parameter is set to 'false'. Updates overwrite any previously saved observer roles for the group. "
	GroupResourceUseDefaultDevices = "If set to ‘true’ the group can notify members on their failsafe (default) devices if none of the member's other devices are available."
	GroupResourceSupervisors       = "Comma-separated list of unique identifiers (UUID) that represent the group’s supervisors. Supervisors must have the role-based permissions required to supervise the group and the authenticating account must have permission to assign supervisors to the group. If empty or null, the group will have no supervisors. If the property is omitted from the request, the authenticating user who makes the request is set as the supervisor."
	GroupResourceExternalKey       = "Unique identifier of a resource in an external system."
	GroupResourceExternallyOwned   = "Whether the object is managed by an external system. Available options are: true, false. Default value is false."

	// --------------------------------------------------
	// Group Roster Resource
	// --------------------------------------------------
	// Group Roster Resource Schema
	GroupRosterResourceID         = "Unique identifier (UUID) of the group in xMatters to manage membership."
	GroupRosterResourceGroup      = "A reference to the group in which membership is being managed."
	GroupRosterResourceMembers    = "A list of members to add to the group. The members can be people, devices, or other groups. If the member is a group, the group must be a dynamic team."
	GroupRosterResourceMemberID   = "The unique identifier of the member to add to the group. The member can be a person, device, or group."
	GroupRosterResourceMemberType = "The type of the member to add to the group. Available options are: PERSON, DEVICE, GROUP."
)
