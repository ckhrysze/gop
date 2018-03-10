/* Package main assumes a 1password json object for the op cli tool's
get command is fed in via stdin
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var SectionName string = "Credentials"
var AccessField string = "access"
var SecretField string = "secret"
var RegionField string = "region"

// The minimal structure needed to get the AWS credentials
type OnePassItem struct {
	Details struct {
		Sections []struct {
			Title  string `json:title`
			Fields []struct {
				T string `json:t`
				V string `json:v`
			} `json:fields`
		} `json:sections`
	} `json:details`
}

// main takes
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
