package main

import "encoding/json"
//import "time"
//import "log"

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

	if p == "/var/log/gitlab/gitaly/current" {
		// This log is 100% json decode it
		json.Unmarshal([]byte(h["message"].(string)), &h)
		// Add a tag to indicate this operation occurred
		addTags(&h, []string{"DecodedJsonToGitalyFields"})
		// Create a field using a regex
		oniguruma("justtime", "time", &h, `(T[0-9:Z]+$)`)
	}

//	if p == "/var/log/syslog" {
		// Create multiple fields from one sequential pattern
//		if !grok(&h, `%{(?P<a>\S+):field1} %{(?P<b>\S+):field2} %{(?P<c>\S+):field3} %{(?P<d>\S+):field4} %{(?P<e>\S+):field5}`) {
			// If failure add a tag in this doc
//			addTags(&h, []string{"IfGrokErrorSyslog"})
//		}
//	}

	if p == "/var/log/mym.log" {
		// This log is 100% json decode it
		json.Unmarshal([]byte(h["message"].(string)), &h)
		// Add a tag to indicate this operation occurred
		addTags(&h, []string{"DecodedJsonToMailFields"})
		// Assign timestamp to actual receive time not log arrival
		h["@timestamp"] = h["Received"]
	}

	c <- &h // Although you write code here this line is required
}
