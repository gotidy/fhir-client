.PHONY: gen test clean

SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')
TMP = ./tmp

mod:
	go mod tidy

## Format the Code.
fmt:
	gofmt -s -l -w $(SRCS)

download: 
	mkdir -p $(TMP)
	curl https://www.hl7.org/fhir/definitions.json.zip -o $(TMP)/definitions.zip
	unzip $(TMP)/definitions.zip profiles-types.json valuesets.json -d $(TMP)/tmp -o
	curl http://hl7.org/fhir/bundle.profile.json  -o $(TMP)/bundle.json
	curl http://hl7.org/fhir/codesystem.profile.json -o $(TMP)/codesystem.json
	curl http://hl7.org/fhir/structuredefinition.profile.json -o $(TMP)/structuredefinition.json
	curl http://hl7.org/fhir/valueset.profile.json -o $(TMP)/valueset.json 

gen: 
	# mkdir -p tmp
	# # wget -O ./tmp/definitions.zip https://www.hl7.org/fhir/definitions.json.zip
	# curl https://www.hl7.org/fhir/definitions.json.zip --output $(TMP)/definitions.zip
	# unzip $(TMP)/definitions.zip profiles-resources.json profiles-types.json valuesets.json -d $(TMP)/
	# rm -rf $(TMP)/definitions.zip
	go generate ./...
	# gofmt -s -l -w $(SRCS)

test: mod gen
	go test ./...	

clean: 
	rm -rf tmp	
	rm -rf .tmp	

