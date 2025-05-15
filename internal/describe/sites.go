package describe

const (
	// --------------------------------------------------
	// Site(s) Data Source
	// --------------------------------------------------
	// Site Data Source Schema
	SiteDataSourceDescription  = "Returns a site in your xMatters instance that matches the provided criteria. Sites in xMatters group users by their physical location. A user's default settings for location-specific settings such as language and time zones are determined from their site. Sites are also used to notify users based on their geographic location."
	SitesDataSourceDescription = "Returns a site in your xMatters instance by its unique identifier (UUID) or target name. Sites in xMatters group users by their physical location. A user's default settings for location-specific settings such as language and time zones are determined from their site."
	SiteIDSearch               = "Unique identifier (UUID) or target name of the xMatters site."
	// Sites Data Source Schema
	SitesList    = "List of available xMatters sites."
	SitesSearch  = "Search terms to filter the list of sites."
	SitesFilters = "Filter the list of sites by country, geocoded status, and site status."
	// Sites Data Source Search Schema
	SitesSearchTerms   = "List of search terms, separated by spaces. Used with the operand and fields parameters to expand or limit search results. When two or more search terms are present, the result includes services that match either search term."
	SitesSearchOperand = "Operands expand or limit the search query parameter. OR is the default operand and returns services that have any of the search terms in the name or description. AND returns services that have all search terms in the name or description."
	SitesSearchFields  = "Set the field to search when a search term is specified. Available options are: NAME, ADDRESS."
	// Sites Data Source Filter Schema
	SiteFilterCountry  = "Full name, two-letter, or three-letter code for a country. For example: CA, CAN, Canada. This field is not case-sensitive. See the <a href=\"https://en.wikipedia.org/wiki/ISO_3166-2\" target=\"_blank\">ISO 3166-2</a> standard for country codes."
	SiteFilterGeocoded = "If 'true', all sites with full latitude and longitude coordinates are returned. If 'false', all sites with partial coordinates or no latitude or longitude coordinates are returned."
	SiteFilterStatus   = "Returns sites with the specified status. Available options are: ACTIVE, INACTIVE."
	// --------------------------------------------------
	// Site Resource
	// --------------------------------------------------
	SiteResourceDescription = "Create or update a site in xMatters.\n\nSites in xMatters group users by their physical location. A user's default settings for location-specific settings such as language and time zones are determined from their site. Sites are also used to notify users based on their geographic location."
	// Site Resource Schema
	SiteAddress1   = "First line of the site address."
	SiteAddress2   = "Second line of the site address."
	SiteCity       = "City where the site is located."
	SiteCountry    = "Country where the site is located. Use either the full country name, or the two-letter or three-letter country code as specified by the ISO 3166-2 standard."
	SiteLanguage   = "Default language used by this site. Specify the language using the two-letter <a href=\"https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes\" target=\"_blank\">ISO 639</a> language code."
	SiteLatitude   = "Latitude of the site's physical location."
	SiteLongitude  = "Longitude of the site's physical location."
	SitePostalCode = "ZIP or postal code for the site."
	SiteState      = "Region, state or province where the site is located."
	SiteStatus     = "Status of the site in xMatters.  Available options are: ACTIVE, INACTIVE. Default is ACTIVE."
	SiteTimezone   = "Default time zone of the site. Use the two-letter country code followed by the city name. Three-letter country codes are not supported. For example: US/Los Angeles."
	SiteID         = "Unique identifier (UUID) of the xMatters site."
	SiteName       = "Name used to identify this site."
)
