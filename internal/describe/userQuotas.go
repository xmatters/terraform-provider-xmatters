package describe

const (
	// --------------------------------------------------
	// UserQuotas Data Source
	// --------------------------------------------------
	UserQuotasDataSourceDescription = "Returns license quota and usage information for users in the xMatters instance.There are two user license types in xMatters: fullUsers and stakeholderUsers. Based on your pricing plan, your account has a specific number of Full Users licenses, and could be entitled to a number of Stakeholder User licenses. Stakeholder Users are users who don't need to take action during the response process. Users with a Stakeholder license type can be assigned the same roles and permissions as full users and they can access information in the web user interface and mobile app, but they cannot respond to any notifications they receive, act as an incident resolver, trigger forms that send messages, or initiate incidents or flows."
	// UserQuotas Data Source Schema
	UserQuotasStakeholderUsersEnabled = "Indicates whether stakeholder licensing is enabled for the instance. Available values are: true, false."
	UserQuotasStakeholderUsers        = "Quota and usage information for the stakeholder user license type."
	UserQuotasFullUsers               = "Quota and usage information for the full user license type."
	UserQuotasQuotaTotal              = "The total number of user licenses, both unused and active, available for your xMatters instance."
	UserQuotasQuotaActive             = "The number of active, or used, user licenses for your instance."
	UserQuotasQuotaUnused             = "The number of unused, or available, user licenses for your instance."
)
