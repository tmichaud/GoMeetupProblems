package domains

import (
	"net/url"
	"strings"
)

// given a slice of strings, return a comma separated string containing all of
// the unique domains assocated with either HTTP or HTTPS URLS.
//
// the domains in the result list should be squashed to lower case
//
// if the domain starts with www., w3., or web., that part of it should be
// stripped from the domain.
//
// For example, if the slice contains:
//
//  hi, i typically like to search <a href='https://www.google.com/'>http://google.com/!</a>
//  for information but sometimes i enjoy using http://www.duckduckgo.com/
//
//  sometimes it returns results from http://golang.org and https://reddit.com/
//
// the resulting slice(?) should contain:
//
//  duckduckgo.com,google.com,golang.org,reddit.com
//

func testSlice(s string, slice []string) bool {
	for _, x := range slice {
		if x == s {
			return true
		}
	}
	return false
}

func cleanUpPrefix(s string) string {
	if len(s) >= 3 && s[0:3] == "w3." {
		return s[3:]
	}
	if len(s) >= 4 && (s[0:4] == "www." || s[0:4] == "web.") {
		return s[4:]
	}
	return s
}

// 3nd attempt
// This attempt also walk backwards from end - but uses cleans up several parts
func ExtractDomains(str []string) string {
	found := []string{}
	s := strings.Join(str, "")
	j := len(s) // End of possible url
	i := 0      // Start of possible url

	// Find the last instance of http
	i = strings.LastIndex(s, "http")
	for i != -1 {
		// If found, parse.  (If not, we fell through)
		u, err := url.Parse(strings.Fields(s[i:j])[0])
		if err == nil && u != nil && len(u.Hostname()) > 0 {
			// We have a url.  Process it and add it to our saved list
			// Process it: means - cleanup the prefix, test to see if this url is unique, and if so, add it to our found slice
			tmp := cleanUpPrefix(u.Hostname())
			if !testSlice(tmp, found) {
				found = append(found, tmp)
			}
			// In any case - set our end (j) to our current i and find the last http again
			j = i
			i = strings.LastIndex(s[0:j], "http")
		} else {
			// We may have found http in the url itself (www.httprules.com).  Leave j where it is and look for an earlier one
			i = strings.LastIndex(s[0:i-1], "http")
		}
	}

	// Return the results, in order found
	ret := []string{}
	for y := len(found) - 1; y >= 0; y-- {
		ret = append(ret, found[y])
	}
	return strings.Join(ret, ",")
}
