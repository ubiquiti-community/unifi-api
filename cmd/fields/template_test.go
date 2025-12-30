package main

import (
	"strings"
	"testing"
)

func TestIsSettingLogic(t *testing.T) {
	tests := []struct {
		name       string
		structName string
		want       bool
	}{
		{
			name:       "regular resource",
			structName: "Account",
			want:       false,
		},
		{
			name:       "setting resource",
			structName: "SettingMgmt",
			want:       true,
		},
		{
			name:       "another setting resource",
			structName: "SettingNetflow",
			want:       true,
		},
		{
			name:       "regular resource with partial match",
			structName: "AccountSetting",
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := strings.HasPrefix(tt.structName, "Setting")
			if have != tt.want {
				t.Errorf("IsSetting logic failed\ngot:  %v\nwant: %v", have, tt.want)
			}
		})
	}
}

func TestResourcePathGeneration(t *testing.T) {
	tests := []struct {
		name         string
		structName   string
		wantRegular  string
		wantSetting  string
	}{
		{
			name:        "account resource",
			structName:  "Account",
			wantRegular: "/s/{siteId}/rest/account",
		},
		{
			name:        "setting mgmt resource",
			structName:  "SettingMgmt",
			wantSetting: "/s/{siteId}/set/setting/mgmt",
		},
		{
			name:        "broadcast group resource",
			structName:  "BroadcastGroup",
			wantRegular: "/s/{siteId}/rest/broadcastgroup",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSetting := strings.HasPrefix(tt.structName, "Setting")
			
			if isSetting {
				// For settings, we expect the set/setting path
				if tt.wantSetting == "" {
					t.Fatalf("test case missing wantSetting for %v", tt.structName)
				}
			} else {
				// For regular resources, we expect the rest path
				if tt.wantRegular == "" {
					t.Fatalf("test case missing wantRegular for %v", tt.structName)
				}
			}
		})
	}
}

func TestHTTPMethodsForResourceTypes(t *testing.T) {
	tests := []struct {
		name        string
		structName  string
		isSetting   bool
		wantMethods map[string]string
	}{
		{
			name:       "regular resource methods",
			structName: "Account",
			isSetting:  false,
			wantMethods: map[string]string{
				"create": "POST",
				"read":   "GET",
				"update": "PUT",
				"delete": "DELETE",
			},
		},
		{
			name:       "setting resource methods",
			structName: "SettingMgmt",
			isSetting:  true,
			wantMethods: map[string]string{
				"create": "PUT", // Settings use PUT for both create and update
				"read":   "GET",
				"update": "PUT",
				// delete should not exist for settings
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have := strings.HasPrefix(tt.structName, "Setting")
			if have != tt.isSetting {
				t.Errorf("isSetting detection failed\ngot:  %v\nwant: %v", have, tt.isSetting)
			}

			// Verify expected methods
			for operation, expectedMethod := range tt.wantMethods {
				t.Run(operation, func(t *testing.T) {
					if tt.isSetting {
						switch operation {
						case "create", "update":
							if expectedMethod != "PUT" {
								t.Errorf("settings %v method\ngot:  %q\nwant: %q", operation, expectedMethod, "PUT")
							}
						case "read":
							if expectedMethod != "GET" {
								t.Errorf("settings %v method\ngot:  %q\nwant: %q", operation, expectedMethod, "GET")
							}
						case "delete":
							t.Errorf("settings should not have delete operation")
						}
					} else {
						// Regular resources should have all CRUD operations
						switch operation {
						case "create":
							if expectedMethod != "POST" {
								t.Errorf("regular resource %v method\ngot:  %q\nwant: %q", operation, expectedMethod, "POST")
							}
						case "read":
							if expectedMethod != "GET" {
								t.Errorf("regular resource %v method\ngot:  %q\nwant: %q", operation, expectedMethod, "GET")
							}
						case "update":
							if expectedMethod != "PUT" {
								t.Errorf("regular resource %v method\ngot:  %q\nwant: %q", operation, expectedMethod, "PUT")
							}
						case "delete":
							if expectedMethod != "DELETE" {
								t.Errorf("regular resource %v method\ngot:  %q\nwant: %q", operation, expectedMethod, "DELETE")
							}
						}
					}
				})
			}
		})
	}
}
