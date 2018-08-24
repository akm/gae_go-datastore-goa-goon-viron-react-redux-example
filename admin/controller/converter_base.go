package controller

import (
	"fmt"
	"strconv"

	"google.golang.org/appengine/datastore"
)

var NoModelGiven = fmt.Errorf("No model given")
var NoPayloadGiven = fmt.Errorf("No payload given")
var NoMediaTypeGiven = fmt.Errorf("No media type given")

func BoolPointerToBool(v *bool) bool {
	return BoolPointerToBoolWith(v, false)
}

func BoolPointerToBoolWith(v *bool, d bool) bool {
	if v == nil {
		return d
	}
	return *v
}

func IntPointerToInt(v *int) int {
	return IntPointerToIntWith(v, 0)
}

func IntPointerToIntWith(v *int, d int) int {
	if v == nil {
		return d
	}
	return *v
}

func IntToInt64(v int) int64 {
	return int64(v)
}

func IntToInt64Pointer(v int) *int64 {
	r := IntToInt64(v)
	return &r
}

func StringPointerToString(v *string) string {
	return StringPointerToStringWith(v, "")
}

func StringPointerToStringWith(v *string, d string) string {
	if v == nil {
		return d
	}
	return *v
}

func StringToInt64(v string) (int64, error) {
	return strconv.ParseInt(v, 10, 64)
}

func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func Int64ToStringPointer(v int64) *string {
	s := Int64ToString(v)
	return &s
}

func StringToDatastoreKeyPointer(v string) (*datastore.Key, error) {
	return datastore.DecodeKey(v)
}

func StringPointerToDatastoreKeyPointer(v *string) (*datastore.Key, error) {
	if v == nil {
		return nil, nil
	}
	return StringToDatastoreKeyPointer(*v)
}

func DatastoreKeyPointerToString(key *datastore.Key) string {
	if key == nil {
		return ""
	}
	return key.Encode()
}

func DatastoreKeyPointerToStringPointer(key *datastore.Key) *string {
	if key == nil {
		return nil
	}
	s := DatastoreKeyPointerToString(key)
	return &s
}
