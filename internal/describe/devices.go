package describe

const (
	// --------------------------------------------------
	// Device(s) Data Source
	// --------------------------------------------------
	// Device Data Source Schema
	DeviceIDSearch           = "Unique identifier (UUID) or target name of the xMatters device. The target name consists of owner’s target name, followed by the | (pipe) character, followed by the device name. For example, ‘mmcbride|Work Phone’."
	DeviceTargetName         = "Device ID in xMatters consisting of owner’s target name, followed by the | (pipe) character, followed by the device name. For example, ‘mmcbride|Work Phone’."
	DeviceCountry            = "Country code of the fax device."
	DeviceDefaultDevice      = "If 'true', the device can receive notifications when the user has no other active devices."
	DeviceEmailAddress       = "The email address of the device."
	DeviceExternallyOwned    = "Whether the device is managed by an external system."
	DeviceName               = "The name of the device. Device names are configured uniquely for each company."
	DevicePhoneNumber        = "Phone number associated with the device."
	DeviceProvider           = "The name of the provider to use when sending notifications to this device."
	DeviceSequence           = "The order in which devices are contacted, where 1 represents the first device."
	DeviceTimeframes         = "A list of timeframes when xMatters may contact this device."
	DeviceTimeframeDays      = "List of the days of the week this timeframe is active."
	DeviceTimeframeStartTime = "The time of day that the timeframe begins."
	DeviceTimeframeTimezone  = "The time zone of the start_time value."

	// Devices Data Source Schema
	DevicesFilters      = "Optional search parameters to filter the list of devices returned."
	DevicesList         = "List of available xMatters devices."
	DeviceFiltersStatus = "Returns all devices with the specified status. Available options are: ACTIVE, INACTIVE."
	DeviceFiltersType   = "Returns all devices with the specified type. Available options: ANDROID_PUSH, APPLE_PUSH, EMAIL, FAX, GENERIC, TEXT_PAGER, TEXT_PHONE, VOICE, VOICE_IVR."
	DeviceFiltersNames  = "A list of device names. Returns a list of devices with the specified names."

	// --------------------------------------------------
	// Device Resource
	// --------------------------------------------------
	DeviceID                         = "Unique identifier (UUID) of the xMatters device."
	DeviceResourceDefaultDevice      = "If 'true', the device can receive notifications when the user has no other active devices. Available values are: true, false."
	DeviceDelay                      = "The number of minutes to wait for a response on this device before contacting the next device."
	DeviceDeviceType                 = "The type of device. Available options are: EMAIL, TEXT_PAGER, TEXT_PHONE, VOICE, VOICE_IVR"
	DeviceResourceEmailAddress       = "The email address of the device. Required when the 'device_Type' is EMAIL. For example. someone@example.com"
	DeviceExternalKey                = "Unique identifier of a resource in an external system."
	DeviceResourceExternallyOwned    = "Whether the object is managed by an external system. Available options are: true, false. Default value is true."
	DeviceResourceName               = "The name of the device. Device names are configured uniquely for each company. For example, Work Phone, Work Email."
	DeviceOwner                      = "Unique identifier (UUID) of the user that owns the device."
	DeviceResourcePhoneNumber        = "Phone number associated with the device. Required when the 'device_Type' is TEXT_PAGER, TEXT_PHONE, VOICE, VOICE_IVR. Voice/SMS phone numbers must be in in E.164 format, and include the + sign and country code. For example: +441632960577"
	DevicePIN                        = "The PIN code of the pager. Required when the 'device_Type' is TEXT_PAGER."
	DevicePriorityThreshold          = "The minimum priority of an alert for it to be delivered to this device. Available options are: LOW, MEDIUM, HIGH."
	DeviceResourceProvider           = "The name of the provider to use when sending notifications to this device. If there is only one provider configured for this type of device, a value does not need to be included. "
	DeviceResourceSequence           = "The order in which devices are contacted, where 1 represents the first device. If the provided sequence number is higher than the number of existing devices, the device is added to the end of the device order."
	DeviceStatus                     = "Status of the device in xMatters. Available options are: ACTIVE, INACTIVE. Default is ACTIVE."
	DeviceTestStatus                 = "A code indicating whether the device has been tested or if testing is in progress. Available options are: TESTED, UNTESTED, PENDING, INVALID."
	DeviceResourceTimeframes         = "A list of timeframes when xMatters may contact this device. If the ‘default_device’ parameter is set to ‘true’, the device may be contacted outside of the specified timeframes in certain situations."
	DeviceResourceTimeframeDays      = "List of the days of the week this timeframe is active. Available values are: SU, MO, TU, WE, TH, FR, SA."
	DeviceTimeframeDuration          = "The length of the timeframe in minutes."
	DeviceTimeframeHolidays          = "If ‘true’, the timeframe is excluded from site holidays; i.e., the device is not active on holidays."
	DeviceTimeframeName              = "The name of the timeframe."
	DeviceResourceTimeframeStartTime = "The time of day that the timeframe begins. For example, 08:00, 15:30."
	DeviceResourceTimeframeTimezone  = "The time zone of the start_time value. For example, US/Pacific."
	DeviceTwoWayDevice               = "Required when the 'device_Type' is TEXT_PAGER. If 'true', the pager is able to send and receive messages. If 'false', the pager is only able to receive messages."
)
