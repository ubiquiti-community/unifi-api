package main

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestFieldInfoFromValidation(t *testing.T) {
	for i, c := range []struct {
		expectedType      string
		expectedComment   string
		expectedOmitEmpty bool
		validation        any
	}{
		{"string", "", true, ""},
		{"string", "default|custom", true, "default|custom"},
		{"string", ".{0,32}", true, ".{0,32}"},
		{"string", "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$", false, "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$"},

		{"int64", "^([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$", true, "^([1-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$|^$"},
		{"int64", "", true, "^[0-9]*$"},

		{"float64", "", true, "[-+]?[0-9]*\\.?[0-9]+"},
		// this one is really an error as the . is not escaped
		{"float64", "", true, "^([-]?[\\d]+[.]?[\\d]*)$"},
		{"float64", "", true, "^([\\d]+[.]?[\\d]*)$"},

		{"bool", "", false, "false|true"},
		{"bool", "", false, "true|false"},
	} {
		t.Run(fmt.Sprintf("%d %s %s", i, c.expectedType, c.validation), func(t *testing.T) {
			resource := &ResourceInfo{
				StructName:     "TestType",
				Types:          make(map[string]*FieldInfo),
				FieldProcessor: func(name string, f *FieldInfo) error { return nil },
			}

			fieldInfo, err := resource.fieldInfoFromValidation("fieldName", c.validation)
			// actualType, actualComment, actualOmitEmpty, err := fieldInfoFromValidation(c.validation)
			if err != nil {
				t.Fatal(err)
			}
			if fieldInfo.FieldType != c.expectedType {
				t.Fatalf("expected type %q got %q", c.expectedType, fieldInfo.FieldType)
			}
			if fieldInfo.FieldValidation != c.expectedComment {
				t.Fatalf("expected comment %q got %q", c.expectedComment, fieldInfo.FieldValidation)
			}
			if fieldInfo.OmitEmpty != c.expectedOmitEmpty {
				t.Fatalf("expected omitempty %t got %t", c.expectedOmitEmpty, fieldInfo.OmitEmpty)
			}
		})
	}
}

func TestResourceTypes(t *testing.T) {
	testData := `
{
  "note": ".{0,1024}",
  "date": "^$|^(20[0-9]{2}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])Z?$",
  "mac": "^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$",
  "number": "\\d+",
  "boolean": "true|false",
	"nested_type": {
    "nested_field": "^$"
  },
  "nested_type_array": [{
    "nested_field": "^$"
  }]
}
	`
	expectedFields := map[string]*FieldInfo{
		"Note":    NewFieldInfo("Note", "note", "string", ".{0,1024}", true, false, false, ""),
		"Date":    NewFieldInfo("Date", "date", "string", "^$|^(20[0-9]{2}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])T([01][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9])Z?$", false, false, false, ""),
		"MAC":     NewFieldInfo("MAC", "mac", "string", "^([0-9A-Fa-f]{2}:){5}([0-9A-Fa-f]{2})$", true, false, false, ""),
		"Number":  NewFieldInfo("Number", "number", "int64", "", true, false, true, "types.Number"),
		"Boolean": NewFieldInfo("Boolean", "boolean", "bool", "", false, false, false, ""),
		"NestedType": {
			FieldName:       "NestedType",
			JSONName:        "nested_type",
			FieldType:       "StructNestedType",
			FieldValidation: "",
			OmitEmpty:       true,
			IsPointer:       true,
			IsArray:         false,
			Fields: map[string]*FieldInfo{
				"NestedFieldModified": NewFieldInfo("NestedFieldModified", "nested_field", "string", "^$", false, false, false, ""),
			},
		},
		"NestedTypeArray": {
			FieldName:       "NestedTypeArray",
			JSONName:        "nested_type_array",
			FieldType:       "StructNestedTypeArray",
			FieldValidation: "",
			OmitEmpty:       true,
			IsPointer:       false,
			IsArray:         true,
			Fields: map[string]*FieldInfo{
				"NestedFieldModified": NewFieldInfo("NestedFieldModified", "nested_field", "string", "^$", false, false, false, ""),
			},
		},
	}

	expectedStruct := map[string]*FieldInfo{
		"Struct": {
			FieldName:       "Struct",
			JSONName:        "path",
			FieldType:       "struct",
			FieldValidation: "",
			OmitEmpty:       false,
			IsArray:         false,
			Fields: map[string]*FieldInfo{
				"   ID":      NewFieldInfo("ID", "_id", "string", "", true, false, false, ""),
				"   SiteID":  NewFieldInfo("SiteID", "site_id", "string", "", true, false, false, ""),
				"   _Spacer": nil,
				"  Hidden":   NewFieldInfo("Hidden", "attr_hidden", "bool", "", true, false, false, ""),
				"  HiddenID": NewFieldInfo("HiddenID", "attr_hidden_id", "string", "", true, false, false, ""),
				"  NoDelete": NewFieldInfo("NoDelete", "attr_no_delete", "bool", "", true, false, false, ""),
				"  NoEdit":   NewFieldInfo("NoEdit", "attr_no_edit", "bool", "", true, false, false, ""),
				"  _Spacer":  nil,
				" _Spacer":   nil,
			},
		},
	}

	for k, v := range expectedFields {
		expectedStruct["Struct"].Fields[k] = v
	}

	expectation := &ResourceInfo{
		StructName:   "Struct",
		ResourcePath: "path",

		Types: map[string]*FieldInfo{
			"Struct":                expectedStruct["Struct"],
			"StructNestedType":      expectedStruct["Struct"].Fields["NestedType"],
			"StructNestedTypeArray": expectedStruct["Struct"].Fields["NestedTypeArray"],
		},

		FieldProcessor: func(name string, f *FieldInfo) error {
			if name == "NestedField" {
				f.FieldName = "NestedFieldModified"
			}
			return nil
		},
	}

	t.Run("structural test", func(t *testing.T) {
		resource := NewResource("Struct", "path")
		resource.FieldProcessor = expectation.FieldProcessor

		err := resource.processJSON(([]byte)(testData))

		assert.NoError(t, err, "No error processing JSON")
		assert.Equal(t, expectation.StructName, resource.StructName)
		assert.Equal(t, expectation.ResourcePath, resource.ResourcePath)
		assert.Equal(t, expectation.Types, resource.Types)
	})
}

func TestNewResourcePaths(t *testing.T) {
	tests := []struct {
		structName       string
		resourcePath     string
		itemResourcePath string
		expectedPath     string
		expectedItemPath string
	}{
		{structName: "APGroup", resourcePath: "apgroups", expectedPath: "/v2/api/site/{site}/apgroups", expectedItemPath: "/v2/api/site/{site}/apgroups/{id}"},
		{
			structName:       "NetworkMembersGroup",
			resourcePath:     "network-members-groups",
			itemResourcePath: "network-members-group",
			expectedPath:     "/v2/api/site/{site}/network-members-groups",
			expectedItemPath: "/v2/api/site/{site}/network-members-group/{id}",
		},
		{structName: "PowerSupervisor", resourcePath: "power-supervisors", expectedPath: "/v2/api/site/{site}/power-supervisors", expectedItemPath: "/v2/api/site/{site}/power-supervisors/{id}"},
	}

	for _, tt := range tests {
		t.Run(tt.structName, func(t *testing.T) {
			resource := &ResourceInfo{
				StructName:       tt.structName,
				ResourcePath:     tt.resourcePath,
				ItemResourcePath: tt.itemResourcePath,
			}
			assert.Equal(t, tt.expectedPath, collectionPath(resource))
			assert.Equal(t, tt.expectedItemPath, itemPath(resource))
		})
	}
}
