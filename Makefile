gen: 
	# wget -O ./tmp/definitions.zip https://www.hl7.org/fhir/definitions.json.zip
	curl https://www.hl7.org/fhir/definitions.json.zip --output ./tmp/definitions.zip
	unzip ./tmp/definitions.zip profiles-resources.json profiles-types.json valuesets.json -d ./tmp/
	rm ./tmp/definitions.zip
	go generate ./...
