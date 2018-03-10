// Package main assumes a 1password json object for the op cli tool's
// get command is fed in via stdin
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// variabled can be set with -ldflags "-X main.SectionName=<whatever>"

// SectionName is the 1password section containing AWS credentials
var SectionName string = "Credentials"

// AccessField is the 1password field with the AWS access key
var AccessField string = "access"

// SecretField is the 1password field with the AWS secret key
var SecretField string = "secret"

// RegionField is the 1password field with the AWS default region
var RegionField string = "region"

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

// main turns 1password json via std and output 'source'able env var output
func main() {
	dec := json.NewDecoder(os.Stdin)

	var item OnePassItem
	err := dec.Decode(&item)
	if err != nil {
		log.Fatal(err)
	}

	for _, section := range item.Details.Sections {
		if section.Title == SectionName {
			for _, field := range section.Fields {
				switch field.T {
				case AccessField:
					fmt.Printf("export AWS_ACCESS_KEY_ID=%v\n", field.V)
				case SecretField:
					fmt.Printf("export AWS_SECRET_ACCESS_KEY=%v\n", field.V)
				case RegionField:
					fmt.Printf("export AWS_DEFAULT_REGION=%v\n", field.V)
				}
			}
		}
	}
}
