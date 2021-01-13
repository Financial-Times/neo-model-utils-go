# neo-model-utils-go

[![CircleCI](https://circleci.com/gh/Financial-Times/neo-model-utils-go/tree/master.svg)](https://circleci.com/gh/Financial-Times/neo-model-utils-go/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Financial-Times/neo-model-utils-go)](https://goreportcard.com/report/github.com/Financial-Times/neo-model-utils-go)
[![Coverage Status](https://coveralls.io/repos/github/Financial-Times/neo-model-utils-go/badge.svg)](https://coveralls.io/github/Financial-Times/neo-model-utils-go)

Provides various packages for managing the space between Neo4j and our Models.

## Packages

### mapper

Has the following features:

1. APIURL - Establishing the ApiURL given a whether the Label is a Person, Organisation or Company (Public or Private)
1. IDURL - Adds the appropriate prefix e.g `http://api.ft.com/things/`
1. TypeURIs - Builds up the type URI based on type e.g `http://www.ft.com/ontology/Person`
