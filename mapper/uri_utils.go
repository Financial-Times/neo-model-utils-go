package mapper

import (
	"fmt"
	"log"
)

var apiPaths = map[string]string{
	"Organisation": "organisations",
	"Person":       "people",
	"Brand":        "brands",
	"Thing":        "things",
	"Content":      "content",
}

var typeURIs = map[string]string{
	"Thing":                       "http://www.ft.com/ontology/core/Thing",
	"Concept":                     "http://www.ft.com/ontology/concept/Concept",
	"Role":                        "http://www.ft.com/ontology/organisation/Role",
	"BoardRole":                   "http://www.ft.com/ontology/organisation/BoardRole",
	"MembershipRole":              "http://www.ft.com/ontology/MembershipRole",
	"Classification":              "http://www.ft.com/ontology/classification/Classification",
	"Person":                      "http://www.ft.com/ontology/person/Person",
	"Organisation":                "http://www.ft.com/ontology/organisation/Organisation",
	"Membership":                  "http://www.ft.com/ontology/organisation/Membership",
	"Company":                     "http://www.ft.com/ontology/company/Company",
	"PublicCompany":               "http://www.ft.com/ontology/company/PublicCompany",
	"PrivateCompany":              "http://www.ft.com/ontology/company/PrivateCompany",
	"Brand":                       "http://www.ft.com/ontology/product/Brand",
	"Subject":                     "http://www.ft.com/ontology/Subject",
	"Section":                     "http://www.ft.com/ontology/Section",
	"Topic":                       "http://www.ft.com/ontology/Topic",
	"Location":                    "http://www.ft.com/ontology/Location",
	"Genre":                       "http://www.ft.com/ontology/Genre",
	"SpecialReport":               "http://www.ft.com/ontology/SpecialReport",
	"AlphavilleSeries":            "http://www.ft.com/ontology/AlphavilleSeries",
	"FinancialInstrument":         "http://www.ft.com/ontology/FinancialInstrument",
	"IndustryClassification":      "http://www.ft.com/ontology/IndustryClassification",
	"NAICSIndustryClassification": "http://www.ft.com/ontology/NAICSIndustryClassification",
}

// APIURL - Establishes the ApiURL given a whether the Label is a Person, Organisation or Company (Public or Private)
func APIURL(uuid string, labels []string, env string) string {
	base := "http://api.ft.com/"
	if env == "test" {
		base = "http://test.api.ft.com/"
	}

	path := ""
	mostSpecific, err := MostSpecificType(labels)
	if err == nil {
		for t := mostSpecific; t != "" && path == ""; t = ParentType(t) {
			path = apiPaths[t]
		}
	}
	if path == "" {
		// TODO: I don't thing we should default to this, but I kept it
		// for compatibility and because this function can't return an error
		path = "things"
	}
	return base + path + "/" + uuid
}

// IDURL - Adds the appropriate prefix e.g http://api.ft.com/things/
func IDURL(uuid string) string {
	return "http://api.ft.com/things/" + uuid
}

// TypeURIs - Builds up the type URI based on type e.g http://www.ft.com/ontology/Person
func TypeURIs(labels []string) []string {
	var results []string
	sorted, err := SortTypes(labels)
	if err != nil {
		log.Printf("ERROR - %v", err)
		return []string{}
	}
	for _, label := range sorted {
		uri := typeURIs[label]
		if uri != "" {
			results = append(results, uri)
		}
	}
	return results
}

// TypesFromURIs converts list of typeURIs to types based on the "typeURIs" mapping.
// The returned concept type list is sorted based on the "parentTypes" hierarchy.
// If a typeURI is not part of the "typeURIs" mapping, or if the type list is un-sortable it returns an error.
func TypesFromURIs(types []string) ([]string, error) {
	var result []string
	for _, t := range types {
		var found bool
		for key, val := range typeURIs {
			if val != t {
				continue
			}
			result = append(result, key)
			found = true
			break
		}
		if !found {
			return nil, fmt.Errorf("type uri '%s' not part of the type heirarcy", t)
		}
	}
	sorted, err := SortTypes(result)
	if err != nil {
		return nil, err
	}
	return sorted, nil
}
