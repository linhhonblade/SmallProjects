package common

func StringFromPointer(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}

func GetPointerString(s string) *string {
	return &s
}
