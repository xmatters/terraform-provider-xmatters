package helpers

import (
	"github.com/google/uuid"
)

// Helper function to get a string pointer
func StringPointer(value string) *string {
	return &value
}

// Helper function to get an int pointer
func IntPointer(value int) *int {
	return &value
}

func Int32Pointer(value int32) *int32 {
	return &value
}

// Helper function to get a bool pointer
func BoolPointer(value bool) *bool {
	return &value
}

// isValidUUID checks if a string is a valid UUID and returns a bool value.
func IsValidUUID(u *string) bool {
	_, err := uuid.Parse(*u)
	return err == nil
}

func DerefNonNilStringSlice(ptrs []*string) []string {
	var strs []string
	for _, ptr := range ptrs {
		if ptr != nil {
			strs = append(strs, *ptr)
		}
	}
	return strs
}
