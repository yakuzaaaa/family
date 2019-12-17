package queryresolver

import (
	"strings"
)

// Represents types of query this resolver can handle
const (
	_addChild        = "ADD_CHILD"
	_getRelationship = "GET_RELATIONSHIP"
)

// Resolve takes a whole query string
// Parses it and call proper method in controller
func Resolve(queryString string) {
	queryParams := strings.Split(queryString, " ")
	// It's a valid query if and only if
	// It has more than two words in it
	if len(queryParams) > 1 {
		// The first word in a query represents the query type
		queryType := queryParams[0]
		switch queryType {
		case _addChild:
			break
		case _getRelationship:
			break
		}
	}
}
