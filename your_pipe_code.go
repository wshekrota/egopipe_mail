package main

import "encoding/json"

//import "time"
//import "log"
import "regexp"

/*

   stage 2
   Where your filter code runs. The doc object is the h map

*/

// yourPipeCode objective is for a user to code the pipe transform stage in golang.
// That code would exist here and a ref to that completed map goes back in channel.
//
func yourPipeCode(h map[string]interface{}, c chan *map[string]interface{}) {

	// h is the hash representing your docs (which are a collection of fields)
	// keys are fieldnames
	// value is interface{} and must be asserted

	// Access log.file.path fields value
	//
	p := dotField(h, "log.file.path").(string)

	if p == "/var/log/zz.log" {

		// This log is 100% json decode it
		if json.Unmarshal([]byte(h["message"].(string)), &h) == nil {

			// Add a tag to indicate this operation occurred
			addTags(&h, []string{"DecodedJsonToMailFields"})

			// Assign timestamp to actual receive time not log arrival
			h["@timestamp"] = h["Received"]

			fld := h["Text"].(string)
			r1 := regexp.MustCompile(`(?i)devops`)
			if r1.Match([]byte(fld)) {
				h["type_job"] = "Devops"
			}
			r2 := regexp.MustCompile(`(?i)developer|programmer`)
			if r2.Match([]byte(fld)) {
				h["type_job"] = "Developer"
			}
			r3 := regexp.MustCompile(`(?i)analytics|elasticsearch`)
			if r3.Match([]byte(fld)) {
				h["type_job"] = "Analytics"
			}
		}
	}

	c <- &h // Although you write code here this line is required
}
