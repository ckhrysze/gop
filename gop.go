// Package main assumes a 1password json object for the op cli tool's
// get command is fed in via stdin
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// variabled can be set with -ldflags "-X main.SectionName=<whatever>"

// SectionName is the 1password section containing AWS credentials
var SectionName = "Credentials"

// AccessField is the 1password field with the AWS access key
var AccessField = "access"

// SecretField is the 1password field with the AWS secret key
var SecretField = "secret"

// RegionField is the 1password field with the AWS default region
var RegionField = "region"

// OnePassItem sets up the minimal structure needed to get AWS credentials
type OnePassItem struct {
	Details struct {
		Sections []struct {
			Title  string `json:"title"`
			Fields []struct {
				T string `json:"t"`
				V string `json:"v"`
			} `json:"fields"`
		} `json:"sections"`
	} `json:"details"`
}

// parse1PasswordEntry takes json from the given reader and
// returns 'source'able export statements
func parse1PasswordEntry(input io.Reader) []string {
	dec := json.NewDecoder(input)

	var item OnePassItem
	err := dec.Decode(&item)
	if err != nil {
		log.Fatal(err)
	}

	output := make([]string, 0)
	for _, section := range item.Details.Sections {
		if section.Title == SectionName {
			for _, field := range section.Fields {
				switch field.T {
				case AccessField:
					output = append(output, fmt.Sprintf("export AWS_ACCESS_KEY_ID=%v", field.V))
				case SecretField:
					output = append(output, fmt.Sprintf("export AWS_SECRET_ACCESS_KEY=%v", field.V))
				case RegionField:
					output = append(output, fmt.Sprintf("export AWS_DEFAULT_REGION=%v", field.V))
				}
			}
		}
	}
	return output
}

// main sends stdin to parse1PasswordEntry and formats the returned statements to stdout
func main() {
	lines := parse1PasswordEntry(os.Stdin)
	for _, line := range lines {
		fmt.Println(line)
	}
}
