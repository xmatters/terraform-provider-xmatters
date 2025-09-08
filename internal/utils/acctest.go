package utils

import (
	"math/rand/v2"

	"github.com/google/uuid"
)

const (
	// CharSetAlphaNum is the alphanumeric character set for use with
	// RandStringFromCharSet
	CharSetAlphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ012346789"

	// CharSetAlpha is the alphabetical character set for use with
	// RandStringFromCharSet
	CharSetAlpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// CharSetNum is the numeric character set for use with
	// RandStringFromCharSet
	CharSetNum = "012346789"
)

// RandUUID generates a random UUID
func RandUUID() string {
	return uuid.New().String()
}

// RandUUIDPointer generates a random UUID and returns a pointer
func RandUUIDPointer() *string {
	uuid := RandUUID()
	return &uuid
}

// RandInt32 generates a random 32bit integer
func RandInt32() int32 {
	return rand.Int32()
}

// RandInt32Pointer generates a random 32bit integer and returns a pointer
func RandInt32Pointer() *int32 {
	i := RandInt32()
	return &i
}

// RandInt64 generates a random 64bit integer
func RandInt64() int64 {
	return rand.Int64()
}

// RandInt64Pointer generates a random 64bit integer and returns a pointer
func RandInt64Pointer() *int64 {
	i := RandInt64()
	return &i
}

// RandFloat64 generates a random 64bit float
func RandFloat64() float64 {
	return rand.Float64()
}

// RandFloat64Pointer generates a random 64bit float and returns a pointer
func RandFloat64Pointer() *float64 {
	f := RandFloat64()
	return &f
}

// RandomLatitude generates a random latitude between -90.0 and 90.0.
func RandomLatitude() float64 {
	return -90.0 + rand.Float64()*(90.0-(-90.0))
}

// RandomLatitudePointer generates a random latitude between -90.0 and 90.0 and returns a pointer
func RandomLatitudePointer() *float64 {
	lat := RandomLatitude()
	return &lat
}

// RandomLongitude generates a random longitude between -180.0 and 180.0.
func RandomLongitude() float64 {
	return -180.0 + rand.Float64()*(180.0-(-180.0))
}

// RandomLongitudePointer generates a random longitude between -180.0 and 180.0 and returns a pointer
func RandomLongitudePointer() *float64 {
	lon := RandomLongitude()
	return &lon
}

// RandBool generates a random boolean
func RandBool() bool {
	return rand.IntN(2) == 0 //rand.IntN(2) == 0
}

// RandBoolPointer generates a random boolean and returns a pointer
func RandBoolPointer() *bool {
	b := RandBool()
	return &b
}

// RandNumString generates a random number between 0 and the max specified
func RandNumString(max int) string {
	return RandStringFromCharSet(max, CharSetNum)
}

// RandNumStringPointer generates a random number between 0 and the max specified and returns a pointer
func RandNumStringPointer(max int) *string {
	n := RandNumString(max)
	return &n
}

func RandNumInt(max int) int32 {
	return rand.Int32N(int32(max))
}

func RandNumIntPointer(max int) *int32 {
	i := RandNumInt(max)
	return &i
}

// RandString generates a random alphanumeric string of the length specified
func RandString(strlen int) string {
	return RandStringFromCharSet(strlen, CharSetAlphaNum)
}

// RandStringPointer generates a random alphanumeric string of the length specified and returns a pointer
func RandStringPointer(strlen int) *string {
	s := RandString(strlen)
	return &s
}

func RandStringList(count, strlen int) []string {
	var list []string
	for i := 0; i < count; i++ {
		list = append(list, RandString(strlen))
	}
	return list
}

func RandStringPointerList(count, strlen int) []*string {
	var list []*string
	for i := 0; i < count; i++ {
		s := RandString(strlen)
		list = append(list, &s)
	}
	return list
}

// RandStringFromCharSet generates a random string by selecting characters from the charset provided
func RandStringFromCharSet(strlen int, charSet string) string {
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = charSet[rand.IntN(len(charSet))]
	}
	return string(result)
}

func RandPhoneNum() string {
	return "+1" + RandStringFromCharSet(10, CharSetNum)
}

func RandStatusAttr() string {
	status := []string{
		"ACTIVE",
		"INACTIVE",
	}
	return status[rand.IntN(len(status))]
}

func RandTimezoneAttr() string {
	timezones := []string{
		"US/Pacific",
		"US/Mountain",
		"US/Central",
		"US/Eastern",
	}
	return timezones[rand.IntN(len(timezones))]
}

func RandGroupTypeAttr() string {
	groupTypes := []string{
		"ON_CALL",
		"BROADCAST",
	}
	return groupTypes[rand.IntN(len(groupTypes))]
}

func RandServiceTypeAttr() string {
	serviceTypes := []string{
		"TECHNICAL",
		"APPLICATION",
	}
	return serviceTypes[rand.IntN(len(serviceTypes))]
}

func RandServiceTierAttr() string {
	serviceTiers := []string{
		"PLATINUM",
		"GOLD",
		"SILVER",
		"BRONZE",
	}
	return serviceTiers[rand.IntN(len(serviceTiers))]
}

func RandDeviceTypeAttr() string {
	deviceTypes := []string{
		"EMAIL",
		"VOICE",
		"VOICE_IVR",
		"TEXT_PHONE",
		"TEXT_PAGER",
	}
	return deviceTypes[rand.IntN(len(deviceTypes))]
}

func RandDevicePriorityAttr() string {
	priorities := []string{
		"LOW",
		"MEDIUM",
		"HIGH",
	}
	return priorities[rand.IntN(len(priorities))]
}

func RandDeviceTestStatusAttr() string {
	testStatus := []string{
		"TESTED",
		"UNTESTED",
		"PENDING",
		"INVALID",
	}
	return testStatus[rand.IntN(len(testStatus))]
}

func RandDeviceNameAttr(dType string) string {
	voiceNames := []string{"Work Phone", "Mobile Phone", "Home Phone"}
	broadcastNames := []string{"Public Broadcast", "Company Broadcast"}
	emailNames := []string{"Work Email", "Home Email"}
	smsNames := "SMS Phone"
	pagerNames := []string{"Work Pager", "Home Pager"}

	switch dType {
	case "VOICE":
		return voiceNames[rand.IntN(len(voiceNames))]
	case "VOICE_IVR":
		return broadcastNames[rand.IntN(len(broadcastNames))]
	case "EMAIL":
		return emailNames[rand.IntN(len(emailNames))]
	case "TEXT_PHONE":
		return smsNames
	case "TEXT_PAGER":
		return pagerNames[rand.IntN(len(pagerNames))]
	default:
		return RandString(10)
	}
}
