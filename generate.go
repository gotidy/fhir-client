////go:generate go run ./gen/types/. -i ./tmp -o ./gen/types/fhir
////go:generate go run ./gen/types/. -i ./tmp -o ./models
//go:generate gofmt -s -w .
//go:generate goimports -w .
package fhir
