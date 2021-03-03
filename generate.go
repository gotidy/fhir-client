//go:generate go run ./gen/types/*.go -i ./tmp -o ./gen/types/fhir
////go:generate go run ./gen/types/*.go -i ./tmp -o ./models
//go:generate gofmt -s -l -w *.go
package fhir
