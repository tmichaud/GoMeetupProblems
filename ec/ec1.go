package domains

import (
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

// quick and dirty method
// Attempt to 'walk' though arrays looking for 'http'.  Note: how to determine end of url?
func ExtractDomains1(s []string) string {
	str := strings.Join(s, "")
	max := len(str)
	j := 0

	for i := strings.Index(str[j:], "http"); i <= max && i <= -1; {
		i = i + 4
		if str[i] == 's' {
			i++
		}
		z := strings.Index(str[i:], "://")
		i = i + 3
		if z == 0 {
			if i+3 <= max && str[i:i+3] == "w3." {
				i = i + 3
			} else {
				if i+4 <= max && (str[i:i+4] == "www." || str[i:i+4] == "web.") {
					i = i + 4
				}
			}
		}
	}

	return strings.Join(s, "")
}
