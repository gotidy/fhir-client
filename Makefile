.PHONY: gen test download-models download-gen

SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')
TMP = ./tmp

mod:
	go mod tidy

## Format the Code.
fmt:
	gofmt -s -w $(SRCS)
	goimports -w .

download-gen: 
	# Definitions
	mkdir -p $(TMP)
	curl https://www.hl7.org/fhir/definitions.json.zip -o $(TMP)/definitions.zip
	unzip $(TMP)/definitions.zip profiles-types.json valuesets.json -d $(TMP) -o
	rm -rf $(TMP)/definitions.zip
	# Profiles
	curl http://hl7.org/fhir/bundle.profile.json  -o $(TMP)/bundle.json
	curl http://hl7.org/fhir/codesystem.profile.json -o $(TMP)/codesystem.json
	curl http://hl7.org/fhir/structuredefinition.profile.json -o $(TMP)/structuredefinition.json
	curl http://hl7.org/fhir/valueset.profile.json -o $(TMP)/valueset.json 

download-models: 
	# Definitions
	mkdir -p $(TMP)
	curl https://www.hl7.org/fhir/definitions.json.zip -o $(TMP)/definitions.zip
	unzip $(TMP)/definitions.zip profiles-resources.json profiles-types.json valuesets.json -d $(TMP) -o
	rm -rf $(TMP)/definitions.zip

gen-gen: 
	# go generate ./...
	rm -rf $(TMP)/*

	$(MAKE) download-gen
	go run ./gen/types/. -i ./tmp -o ./gen/types/fhir
	rm -rf $(TMP)/*

gen-models: 
	# go generate ./...
	rm -rf $(TMP)/*

	$(MAKE) download-models
	go run ./gen/types/. -i ./tmp -o ./models -d ./gen/client/definitions.go
	rm -rf $(TMP)/*

gen-client: 
	go run ./gen/client/. -o ./
	$(MAKE) fmt

gen: ge-gen ge-models gen-client fmt

test: mod gen
	go test ./...	

clean: 
	rm -rf tmp	
	rm -rf .tmp	

