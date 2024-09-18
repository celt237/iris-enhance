package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeConvert(t *testing.T) {
	tests := []struct {
		name        string
		strValue    string
		dataType    string
		ptr         bool
		expected    interface{}
		expectedErr bool
	}{
		{"String", "test", "string", false, "test", false},
		{"String Ptr", "test", "string", true, stringPtr("test"), false},
		{"Int", "123", "int", false, 123, false},
		{"Int Ptr", "123", "int", true, intPtr(123), false},
		{"Int8", "123", "int8", false, int8(123), false},
		{"Int8 Ptr", "123", "int8", true, int8Ptr(123), false},
		{"Int16", "123", "int16", false, int16(123), false},
		{"Int16 Ptr", "123", "int16", true, int16Ptr(123), false},
		{"Int32", "123", "int32", false, int32(123), false},
		{"Int32 Ptr", "123", "int32", true, int32Ptr(123), false},
		{"Int64", "123", "int64", false, int64(123), false},
		{"Int64 Ptr", "123", "int64", true, int64Ptr(123), false},
		{"Uint", "123", "uint", false, uint(123), false},
		{"Uint Ptr", "123", "uint", true, uintPtr(123), false},
		{"Uint8", "123", "uint8", false, uint8(123), false},
		{"Uint8 Ptr", "123", "uint8", true, uint8Ptr(123), false},
		{"Uint16", "123", "uint16", false, uint16(123), false},
		{"Uint16 Ptr", "123", "uint16", true, uint16Ptr(123), false},
		{"Uint32", "123", "uint32", false, uint32(123), false},
		{"Uint32 Ptr", "123", "uint32", true, uint32Ptr(123), false},
		{"Uint64", "123", "uint64", false, uint64(123), false},
		{"Uint64 Ptr", "123", "uint64", true, uint64Ptr(123), false},
		{"Float32", "123.45", "float32", false, float32(123.45), false},
		{"Float32 Ptr", "123.45", "float32", true, float32Ptr(123.45), false},
		{"Float64", "123.45", "float64", false, 123.45, false},
		{"Float64 Ptr", "123.45", "float64", true, float64Ptr(123.45), false},
		{"Bool", "true", "bool", false, true, false},
		{"Bool Ptr", "true", "bool", true, boolPtr(true), false},
		{"Invalid Type", "test", "invalid", false, nil, true},
		{"Invalid Int", "abc", "int", false, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TypeConvert(tt.strValue, tt.dataType, tt.ptr)
			if tt.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

// Helper functions to create pointers
func stringPtr(s string) *string    { return &s }
func intPtr(i int) *int             { return &i }
func int8Ptr(i int8) *int8          { return &i }
func int16Ptr(i int16) *int16       { return &i }
func int32Ptr(i int32) *int32       { return &i }
func int64Ptr(i int64) *int64       { return &i }
func uintPtr(i uint) *uint          { return &i }
func uint8Ptr(i uint8) *uint8       { return &i }
func uint16Ptr(i uint16) *uint16    { return &i }
func uint32Ptr(i uint32) *uint32    { return &i }
func uint64Ptr(i uint64) *uint64    { return &i }
func float32Ptr(f float32) *float32 { return &f }
func float64Ptr(f float64) *float64 { return &f }
func boolPtr(b bool) *bool          { return &b }
