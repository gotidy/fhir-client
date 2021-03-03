#!/usr/bin/env bash

mkdir -p ./../../tmp
curl https://www.hl7.org/fhir/definitions.json.zip -o ./../../tmp/definitions.zip
unzip ./../../tmp/definitions.zip profiles-types.json valuesets.json -d ./../tmp
rm ./../../tmp/definitions.zip
curl http://hl7.org/fhir/bundle.profile.json  -o ./../../tmp/bundle.json
curl http://hl7.org/fhir/codesystem.profile.json -o ./../../tmp/codesystem.json
curl http://hl7.org/fhir/structuredefinition.profile.json -o ./../../tmp/structuredefinition.json
curl http://hl7.org/fhir/valueset.profile.json -o ./../../valueset.json 

# wget -O ./../../tmp/definitions.zip https://www.hl7.org/fhir/definitions.json.zip
# unzip ./../../tmp/definitions.zip profiles-types.json valuesets.json -d ./../../tmp
# rm ./../../tmp/definitions.zip
# wget -O ./../../tmp/bundle.json http://hl7.org/fhir/bundle.profile.json
# wget -O ./../../tmp/codesystem.json http://hl7.org/fhir/codesystem.profile.json
# wget -O ./../../tmp/structuredefinition.json http://hl7.org/fhir/structuredefinition.profile.json
# wget -O ./../../tmp/valueset.json http://hl7.org/fhir/valueset.profile.json

go generate ./fhir 

rm -rf ./../../tmp/*.json
