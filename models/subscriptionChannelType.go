// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Code generated by go generate; DO NOT EDIT.

// This file was generated by robots at
// 2021-03-03 10:55:45.160354 +0000 UTC
// SubscriptionChannelType is documented here http://hl7.org/fhir/ValueSet/subscription-channel-type
type SubscriptionChannelType int

const (
	SubscriptionChannelTypeRestHook SubscriptionChannelType = iota
	SubscriptionChannelTypeWebsocket
	SubscriptionChannelTypeEmail
	SubscriptionChannelTypeSms
	SubscriptionChannelTypeMessage
)

func (code SubscriptionChannelType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SubscriptionChannelType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "rest-hook":
		*code = SubscriptionChannelTypeRestHook
	case "websocket":
		*code = SubscriptionChannelTypeWebsocket
	case "email":
		*code = SubscriptionChannelTypeEmail
	case "sms":
		*code = SubscriptionChannelTypeSms
	case "message":
		*code = SubscriptionChannelTypeMessage
	default:
		return fmt.Errorf("unknown SubscriptionChannelType code `%s`", s)
	}
	return nil
}
func (code SubscriptionChannelType) String() string {
	return code.Code()
}
func (code SubscriptionChannelType) Code() string {
	switch code {
	case SubscriptionChannelTypeRestHook:
		return "rest-hook"
	case SubscriptionChannelTypeWebsocket:
		return "websocket"
	case SubscriptionChannelTypeEmail:
		return "email"
	case SubscriptionChannelTypeSms:
		return "sms"
	case SubscriptionChannelTypeMessage:
		return "message"
	}
	return "<unknown>"
}
func (code SubscriptionChannelType) Display() string {
	switch code {
	case SubscriptionChannelTypeRestHook:
		return "Rest Hook"
	case SubscriptionChannelTypeWebsocket:
		return "Websocket"
	case SubscriptionChannelTypeEmail:
		return "Email"
	case SubscriptionChannelTypeSms:
		return "SMS"
	case SubscriptionChannelTypeMessage:
		return "Message"
	}
	return "<unknown>"
}
func (code SubscriptionChannelType) Definition() string {
	switch code {
	case SubscriptionChannelTypeRestHook:
		return "The channel is executed by making a post to the URI. If a payload is included, the URL is interpreted as the service base, and an update (PUT) is made."
	case SubscriptionChannelTypeWebsocket:
		return "The channel is executed by sending a packet across a web socket connection maintained by the client. The URL identifies the websocket, and the client binds to this URL."
	case SubscriptionChannelTypeEmail:
		return "The channel is executed by sending an email to the email addressed in the URI (which must be a mailto:)."
	case SubscriptionChannelTypeSms:
		return "The channel is executed by sending an SMS message to the phone number identified in the URL (tel:)."
	case SubscriptionChannelTypeMessage:
		return "The channel is executed by sending a message (e.g. a Bundle with a MessageHeader resource etc.) to the application identified in the URI."
	}
	return "<unknown>"
}
