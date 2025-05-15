package describe

const (
	// --------------------------------------------------
	// Service(s) Data Source
	// --------------------------------------------------
	ServiceDataSourceDescription  = "Returns a service in your xMatters instance by its unique identifier. Services let you define the business, technical, and external services performed by, within, or available to your enterprise – and the teams supporting them. Each service has a group assigned as its owner — a group can own multiple services, but each service can only have one owner."
	ServicesDataSourceDescription = "Returns a list of services in your xMatters instance that match the provided criteria. Services let you define the business, technical, and external services performed by, within, or available to your enterprise – and the teams supporting them. Each service has a group assigned as its owner — a group can own multiple services, but each service can only have one owner."
	// Services Data Source Schema
	ServicesOwner  = "Unique identifier (UUID) or target name of the group that owns the xMatters service. Partial matches are not permitted."
	ServicesSearch = "Optional search parameters to filter the list of services returned."
	ServicesList   = "List of available xMatters services."
	// Services Data Source Search Schema
	ServicesSearchTerms   = "List of of search terms, separated by spaces. Used with the operand and fields parameters to expand or limit search results. When two or more search terms are present, the result includes services that match either search term."
	ServicesSearchOperand = "Operands expand or limit the search query parameter. OR is the default operand and returns services that have any of the search terms in the name or description. AND returns services that have all search terms in the name or description."
	ServicesSearchFields  = "Set the field to search when a search term is specified. Available options are: NAME, DESCRIPTION."
	// Service Search Fields Schema
	ServiceIDSearch = "Unique ID (UUID) or target name of the xMatters service."
	// Service Data Source Schema
	ServiceID          = "Unique identifier (UUID) of the xMatters service."
	ServiceName        = "Name of the service."
	ServiceDescription = "Description of the service."
	ServiceType        = "xMatters service type."
	ServiceTier        = "xMatters service tier."
	ServiceLinks       = "List of links associated to the service and list of user-specified links for the service."
	// Service Links Schema
	ServiceLinkLabel = "Link text displayed to the user."
	ServiceLinkURL   = "Direct URL for the link."

	// --------------------------------------------------
	// Service Resource
	// --------------------------------------------------
	ServiceResourceDescription = "Create or update a service in your xMatters instance.\n\nServices are the business, technical, and external services performed by, within, or available to your enterprise – and the teams supporting them. Each service has a group assigned as its owner — a group can own multiple services, but each service can only have one owner."
	// Service Resource Schema
	ServiceResourceType  = "xMatters service type. Available options are: TECHNICAL, APPLICATION."
	ServiceResourceTier  = "xMatters service tier. Available options are: PLATINUM, GOLD, SILVER, BRONZE."
	ServiceResourceLinks = "List of user-specified links for the service."
	ServiceResourceOwner = "Unique identifier (UUID) of the group that owns the xMatters service. A service can only have one owner, but a group can own multiple services."

	// --------------------------------------------------
	// Service Dependency Resource
	// --------------------------------------------------
	ServiceDependencyResourceDescription = "Create a dependency between two existing xMatters services.\n\nService dependencies describe the relationship between a service and its direct dependent in your xMatters system. If a service has multiple dependencies or depends on multiple services, those relationships cannot currently be displayed in the request response via the xMatters REST API. To view multiple dependencies, use the service map in the xMatters web user interface. "
	// Service Dependency Resource Schema
	ServiceDependencyID                 = "Unique identifier (UUID) of the service dependency."
	ServiceDependencyServiceID          = "Unique identifier (UUID) of a service. This service is relied upon by the dependent service."
	ServiceDependencyDependentServiceID = "Unique identifier (UUID) of the dependent service."
)
