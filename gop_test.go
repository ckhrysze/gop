package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEntryParsing(t *testing.T) {
	json := []byte(`{
		"uuid":"ignored",
		"vaultUuid":"ignored",
		"overview":{
			"title":"value for op get item"
		},
		"details":{
			"fields":[
			],
			"sections":[
				{
					"title":"Credentials",
					"name":"Section_ignored",
					"fields":[
						{"k":"string","n":"ignored","t":"access","v":"access_value"},
						{"k":"concealed","n":"ignored","t":"secret","v":"secret_value"},
						{"k":"string","n":"ignored","t":"region","v":"region_value"}
					]
				}
			]
		}
	}`)

	lines := parse1PasswordEntry(bytes.NewBuffer(json))
	expectedLines := []string{
		"export AWS_ACCESS_KEY_ID=access_value",
		"export AWS_SECRET_ACCESS_KEY=secret_value",
		"export AWS_DEFAULT_REGION=region_value",
	}

	if !reflect.DeepEqual(lines, expectedLines) {
		t.Errorf("Actual output %v does not equal expected %v", lines, expectedLines)
	}
}
