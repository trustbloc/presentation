/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/


package api

type Demo struct {
	ID string `json:"id,omitempty"`
}

type Invitations struct {
	DIDExchangeInvitations []Invitation `json:"invitations,omitempty"`
}


type Invitation struct {
	Type   			string 		`json:"@type,omitempty"`
	ServiceEndpoint string 		`json:"serviceEndpoint,omitempty"`
	RecipientKeys 	[]string	`json:"recipientKeys,omitempty"`
	ID 				string 		`json:"@id,omitempty"`
	Label 			string  	`json:"label,omitempty"`
	DID 			string 		`json:"did,omitempty"`
	RoutingKeys 	[]string	`json:"routingKeys,omitempty"`
}
