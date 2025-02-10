package common

func StringFromPointer(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}
