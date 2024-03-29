package mapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func allLabelsFor(label string) []string {
	allLabels := []string{}
	for t := label; t != ""; t = ParentType(t) {
		allLabels = append(allLabels, t)
	}
	s, _ := SortTypes(allLabels)
	return s
}

var (
	thingLabels                       = allLabelsFor("Thing")
	conceptLabels                     = allLabelsFor("Concept")
	boardRoleLabels                   = allLabelsFor("BoardRole")
	membershipRoleLabels              = allLabelsFor("MembershipRole")
	brandLabels                       = allLabelsFor("Brand")
	personLabels                      = allLabelsFor("Person")
	organisationLabels                = allLabelsFor("Organisation")
	membershipLabels                  = allLabelsFor("Membership")
	companyLabels                     = allLabelsFor("Company")
	publicCompanyLabels               = allLabelsFor("PublicCompany")
	privateCompanyLabels              = allLabelsFor("PrivateCompany")
	classificationLabels              = allLabelsFor("Classification")
	subjectLabels                     = allLabelsFor("Subject")
	sectionLabels                     = allLabelsFor("Section")
	genreLabels                       = allLabelsFor("Genre")
	topicLabels                       = allLabelsFor("Topic")
	locationLabels                    = allLabelsFor("Location")
	specialReportLabels               = allLabelsFor("SpecialReport")
	alphavilleSeriesLabels            = allLabelsFor("AlphavilleSeries")
	financialInstrumentsLabels        = allLabelsFor("FinancialInstrument")
	industryClassificationLabels      = allLabelsFor("IndustryClassification")
	naicsIndustryClassificationLabels = allLabelsFor("NAICSIndustryClassification")
	aniIndustryClassificationLabels   = allLabelsFor("FTAnIIndustryClassification")

	thingURI                       = "http://www.ft.com/ontology/core/Thing"
	conceptURI                     = "http://www.ft.com/ontology/concept/Concept"
	boardRoleURI                   = "http://www.ft.com/ontology/organisation/BoardRole"
	membershipRoleURI              = "http://www.ft.com/ontology/MembershipRole"
	classificationURI              = "http://www.ft.com/ontology/classification/Classification"
	personURI                      = "http://www.ft.com/ontology/person/Person"
	brandURI                       = "http://www.ft.com/ontology/product/Brand"
	organisationURI                = "http://www.ft.com/ontology/organisation/Organisation"
	membershipURI                  = "http://www.ft.com/ontology/organisation/Membership"
	companyURI                     = "http://www.ft.com/ontology/company/Company"
	publicCompanyURI               = "http://www.ft.com/ontology/company/PublicCompany"
	privateCompanyURI              = "http://www.ft.com/ontology/company/PrivateCompany"
	subjectURI                     = "http://www.ft.com/ontology/Subject"
	sectionURI                     = "http://www.ft.com/ontology/Section"
	genreURI                       = "http://www.ft.com/ontology/Genre"
	topicURI                       = "http://www.ft.com/ontology/Topic"
	locationURI                    = "http://www.ft.com/ontology/Location"
	specialReportURI               = "http://www.ft.com/ontology/SpecialReport"
	alphavilleSeriesURI            = "http://www.ft.com/ontology/AlphavilleSeries"
	financialInstrumentURI         = "http://www.ft.com/ontology/FinancialInstrument"
	industryClassificationURI      = "http://www.ft.com/ontology/IndustryClassification"
	naicsIndustryClassificationURI = "http://www.ft.com/ontology/NAICSIndustryClassification"
	aniIndustryClassificationURI   = "http://www.ft.com/ontology/FTAnIIndustryClassification"

	uuid = "92f4ec09-436d-4092-a88c-96f54e34007c"

	env = "prod"

	baseAPIURL         = "http://api.ft.com/"
	thingAPIURL        = baseAPIURL + "things/" + uuid
	personAPIURL       = baseAPIURL + "people/" + uuid
	organisationAPIURL = baseAPIURL + "organisations/" + uuid
	contentAPIURL      = baseAPIURL + "content/" + uuid
	brandAPIURL        = baseAPIURL + "brands/" + uuid

	brandHierarchy = []string{
		"http://www.ft.com/ontology/core/Thing",
		"http://www.ft.com/ontology/concept/Concept",
		"http://www.ft.com/ontology/classification/Classification",
		"http://www.ft.com/ontology/product/Brand",
	}

	publicCompanyHierarchy = []string{
		"http://www.ft.com/ontology/core/Thing",
		"http://www.ft.com/ontology/concept/Concept",
		"http://www.ft.com/ontology/organisation/Organisation",
		"http://www.ft.com/ontology/company/Company",
		"http://www.ft.com/ontology/company/PublicCompany",
	}

	boardRoleHierarchy = []string{
		"http://www.ft.com/ontology/core/Thing",
		"http://www.ft.com/ontology/concept/Concept",
		"http://www.ft.com/ontology/MembershipRole",
		"http://www.ft.com/ontology/organisation/BoardRole",
	}
)

func TestTypeURIsForThings(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI}, TypeURIs(thingLabels))
}

func TestTypeURIsForConcepts(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI}, TypeURIs(conceptLabels))
}

func TestTypeURIsForClassifications(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI}, TypeURIs(classificationLabels))
}

func TestTypeURIsForPeople(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, personURI}, TypeURIs(personLabels))
}

func TestTypeURIsForOrganisations(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, organisationURI}, TypeURIs(organisationLabels))
}

func TestTypeURIsForMemberships(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, membershipURI}, TypeURIs(membershipLabels))
}

func TestTypeURIsForMembershipRoles(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, membershipRoleURI}, TypeURIs(membershipRoleLabels))
}

func TestTypeURIsForBoardRoles(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, membershipRoleURI, boardRoleURI}, TypeURIs(boardRoleLabels))
}

func TestTypeURIsForCompany(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, organisationURI, companyURI}, TypeURIs(companyLabels))
}

func TestTypeURIsForPublicCompany(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, organisationURI, companyURI, publicCompanyURI}, TypeURIs(publicCompanyLabels))
}

func TestTypeURIsForPrivateCompany(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, organisationURI, companyURI, privateCompanyURI}, TypeURIs(privateCompanyLabels))
}

func TestTypeURIsForBrand(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, brandURI}, TypeURIs(brandLabels))
}

func TestTypeURIsForSubject(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, subjectURI}, TypeURIs(subjectLabels))
}

func TestTypeURIsForSection(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, sectionURI}, TypeURIs(sectionLabels))
}

func TestTypeURIsForGenre(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, genreURI}, TypeURIs(genreLabels))
}

func TestTypeURIsForTopic(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, topicURI}, TypeURIs(topicLabels))
}

func TestTypeURIsForLocation(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, locationURI}, TypeURIs(locationLabels))
}

func TestTypeURIsForSpecialReport(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, specialReportURI}, TypeURIs(specialReportLabels))
}

func TestTypeURIsForIndustryClassification(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, industryClassificationURI}, TypeURIs(industryClassificationLabels))
}

func TestTypeURIsForNAICSIndustryClassification(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, industryClassificationURI, naicsIndustryClassificationURI}, TypeURIs(naicsIndustryClassificationLabels))
}

func TestTypeURIsForFTAnIIndustryClassification(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, industryClassificationURI, aniIndustryClassificationURI}, TypeURIs(aniIndustryClassificationLabels))
}

func TestTypeURIsForAlphavilleSeries(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, classificationURI, alphavilleSeriesURI}, TypeURIs(alphavilleSeriesLabels))
}

func TestTypeURIForFinancialInstruments(t *testing.T) {
	assert.New(t).EqualValues([]string{thingURI, conceptURI, financialInstrumentURI}, TypeURIs(financialInstrumentsLabels))
}

func TestContentAPIURLs(t *testing.T) {
	neoLabels := []string{"Content"}
	assert.New(t).EqualValues(contentAPIURL, APIURL(uuid, neoLabels, env))
}

func TestContentPackageAPIURLs(t *testing.T) {
	neoLabels := []string{"Thing", "ContentPackage", "Content"}
	assert.New(t).EqualValues(contentAPIURL, APIURL(uuid, neoLabels, env))
}

func TestCompanyAPIURLs(t *testing.T) {
	neoLabels := []string{"PublicCompany", "Organisation", "Company"}
	assert.New(t).EqualValues(organisationAPIURL, APIURL(uuid, neoLabels, env))
}

func TestOrganisationAPIURLs(t *testing.T) {
	neoLabels := []string{"Organisation", "Concept", "Thing"}
	assert.New(t).EqualValues(organisationAPIURL, APIURL(uuid, neoLabels, env))
}

func TestPeopleAPIURLs(t *testing.T) {
	neoLabels := []string{"Person"}
	assert.New(t).EqualValues(personAPIURL, APIURL(uuid, neoLabels, env))

}

func TestBrandAPIURLs(t *testing.T) {
	neoLabels := []string{"Brand"}
	assert.New(t).EqualValues(brandAPIURL, APIURL(uuid, neoLabels, env))

}

func TestDefaultAPIURLs(t *testing.T) {
	neoLabels := []string{"thing", "relationship", "otherPrivateType"}
	assert.New(t).EqualValues(thingAPIURL, APIURL(uuid, neoLabels, env))

}

func TestMostSpecific(t *testing.T) {
	assert := assert.New(t)

	for _, t := range []struct {
		input    []string
		expected string
		err      error
	}{
		{
			[]string{"Organisation"},
			"Organisation",
			nil,
		}, {
			[]string{"Organisation", "PublicCompany", "Company"},
			"PublicCompany",
			nil,
		}, {
			[]string{"PublicCompany", "Organisation"},
			"PublicCompany",
			nil,
		}, {
			[]string{"Organisation", "PublicCompany", "PrivateCompany", "Company"},
			"",
			ErrNotHierarchy,
		}, {
			[]string{"zzzzzz", "yyyyyy"},
			"",
			ErrNotHierarchy,
		},
	} {
		ms, err := MostSpecificType(t.input)
		assert.Equal(t.expected, ms)
		assert.Equal(t.err, err)
	}
}

func TestTypeSorter(t *testing.T) {
	assert := assert.New(t)

	for _, t := range []struct {
		input    []string
		expected []string
		err      error
	}{
		{
			[]string{"Organisation"},
			[]string{"Organisation"},
			nil,
		}, {
			[]string{"Organisation", "PublicCompany", "Company"},
			[]string{"Organisation", "Company", "PublicCompany"},
			nil,
		}, {
			[]string{"PublicCompany", "Organisation"},
			[]string{"Organisation", "PublicCompany"},
			nil,
		}, {
			[]string{"Content", "Thing"},
			[]string{"Thing", "Content"},
			nil,
		}, {
			[]string{"Organisation", "PublicCompany", "PrivateCompany", "Company"},
			[]string{"Organisation", "PublicCompany", "PrivateCompany", "Company"},
			ErrNotHierarchy,
		}, {
			[]string{"zzzzzz", "yyyyyy"},
			[]string{"zzzzzz", "yyyyyy"},
			ErrNotHierarchy,
		},
	} {
		sorted, err := SortTypes(t.input)
		assert.Equal(t.expected, sorted)
		assert.Equal(t.err, err)
	}
}

func TestFullTypeHierarchy(t *testing.T) {
	assert := assert.New(t)
	for _, t := range []struct {
		startingType      string
		expectedHierarchy []string
	}{
		{
			"Brand",
			brandHierarchy,
		},
		{

			"http://www.ft.com/ontology/product/Brand",
			brandHierarchy,
		},
		{

			"PublicCompany",
			publicCompanyHierarchy,
		},
		{

			"http://www.ft.com/ontology/company/PublicCompany",
			publicCompanyHierarchy,
		},
		{

			"BoardRole",
			boardRoleHierarchy,
		},
		{

			"http://www.ft.com/ontology/organisation/BoardRole",
			boardRoleHierarchy,
		},
	} {
		convertedHierarchy := FullTypeHierarchy(t.startingType)
		assert.Equal(t.expectedHierarchy, convertedHierarchy)
	}

}
func TestTypesFromURIs(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		TypeURIs []string
		Types    []string
	}{
		"Things":                      {TypeURIs: []string{thingURI}, Types: thingLabels},
		"Concepts":                    {TypeURIs: []string{thingURI, conceptURI}, Types: conceptLabels},
		"Classifications":             {TypeURIs: []string{thingURI, conceptURI, classificationURI}, Types: classificationLabels},
		"People":                      {TypeURIs: []string{thingURI, conceptURI, personURI}, Types: personLabels},
		"Organisations":               {TypeURIs: []string{thingURI, conceptURI, organisationURI}, Types: organisationLabels},
		"Memberships":                 {TypeURIs: []string{thingURI, conceptURI, membershipURI}, Types: membershipLabels},
		"MembershipRoles":             {TypeURIs: []string{thingURI, conceptURI, membershipRoleURI}, Types: membershipRoleLabels},
		"BoardRoles":                  {TypeURIs: []string{thingURI, conceptURI, membershipRoleURI, boardRoleURI}, Types: boardRoleLabels},
		"Company":                     {TypeURIs: []string{thingURI, conceptURI, organisationURI, companyURI}, Types: companyLabels},
		"PublicCompany":               {TypeURIs: []string{thingURI, conceptURI, organisationURI, companyURI, publicCompanyURI}, Types: publicCompanyLabels},
		"PrivateCompany":              {TypeURIs: []string{thingURI, conceptURI, organisationURI, companyURI, privateCompanyURI}, Types: privateCompanyLabels},
		"Brand":                       {TypeURIs: []string{thingURI, conceptURI, classificationURI, brandURI}, Types: brandLabels},
		"Subject":                     {TypeURIs: []string{thingURI, conceptURI, classificationURI, subjectURI}, Types: subjectLabels},
		"Section":                     {TypeURIs: []string{thingURI, conceptURI, classificationURI, sectionURI}, Types: sectionLabels},
		"Genre":                       {TypeURIs: []string{thingURI, conceptURI, classificationURI, genreURI}, Types: genreLabels},
		"Topic":                       {TypeURIs: []string{thingURI, conceptURI, topicURI}, Types: topicLabels},
		"Location":                    {TypeURIs: []string{thingURI, conceptURI, locationURI}, Types: locationLabels},
		"SpecialReport":               {TypeURIs: []string{thingURI, conceptURI, classificationURI, specialReportURI}, Types: specialReportLabels},
		"IndustryClassification":      {TypeURIs: []string{thingURI, conceptURI, industryClassificationURI}, Types: industryClassificationLabels},
		"NAICSIndustryClassification": {TypeURIs: []string{thingURI, conceptURI, industryClassificationURI, naicsIndustryClassificationURI}, Types: naicsIndustryClassificationLabels},
		"FTAnIIndustryClassification": {TypeURIs: []string{thingURI, conceptURI, industryClassificationURI, aniIndustryClassificationURI}, Types: aniIndustryClassificationLabels},
		"AlphavilleSeries":            {TypeURIs: []string{thingURI, conceptURI, classificationURI, alphavilleSeriesURI}, Types: alphavilleSeriesLabels},
		"FinancialInstruments":        {TypeURIs: []string{thingURI, conceptURI, financialInstrumentURI}, Types: financialInstrumentsLabels},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result, err := TypesFromURIs(test.TypeURIs)
			assert.NoError(t, err)
			assert.EqualValues(t, test.Types, result)
		})
	}
}

func TestTypesFromURIsErrors(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		TypeURIs []string
	}{
		"invalid type uri": {
			TypeURIs: []string{
				"invalid-type-uri",
			},
		},
		"unsortable types": {
			TypeURIs: []string{
				organisationURI,
				publicCompanyURI,
				privateCompanyURI,
				companyURI,
			},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := TypesFromURIs(test.TypeURIs)
			assert.Error(t, err)
		})
	}
}
