package translator

// A string stash for english language
var english = make(map[string]string)

// The app supported languages stash
// Currently has only english
// Can have same key-value maps for all languages
var translations = make(map[string]map[string]string)

func init() {
	english["child_addition_succeeded"] = "CHILD_ADDITION_SUCCEEDED"
	english["child_addition_failed"] = "CHILD_ADDITION_FAILED"
	english["person_not_found"] = "PERSON_NOT_FOUND"
	english["node"] = "NONE"

	translations["en"] = english
}

// Get returns string in given locale
func Get(stringType string, locale string) string {
	// We probably would call some translation interface here
	// But as we dont have it now
	// let's just return english strings
	locale = "en" // Re-assignment to english whatsoever
	return translations[locale][stringType]
}
