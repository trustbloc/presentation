/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/


package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hyperledger/aries-framework-go/pkg/common/log"
	"github.com/rs/cors"

	"github.com/trustbloc/presentation/2020/01/aries-framework-go/cmd/lobby/api"
	"github.com/trustbloc/presentation/2020/01/aries-framework-go/cmd/lobby/store/mem"
)

const (
	addressPattern    = ":%s"
	demoPath   = "/demo"
	getDemoPath   = demoPath + "/{uid}"
	postInvitationsPath   = demoPath + "/{uid}/invitation"
	getInvitationsPath   = demoPath + "/{uid}/invitations"
)

var logger = log.New("aries-framework/lobby-server")


func main() {
	// get configs
	port := os.Getenv("LOBBY_SERVER_PORT")
	if port == "" {
		panic("port to be passed as ENV variable")
	}

	api := api.NewAPI(mem.New())

	// REST Handlers
	r := mux.NewRouter()
	r.HandleFunc(demoPath, api.CreateDemo).Methods(http.MethodPost)
	r.HandleFunc(getDemoPath, api.GetDemo).Methods(http.MethodGet)
	r.HandleFunc(postInvitationsPath, api.PostInvitation).Methods(http.MethodPost)
	r.HandleFunc(getInvitationsPath, api.GetInvitations).Methods(http.MethodGet)

	handler := cors.Default().Handler(r)

	// start the server
	logger.Fatalf("lobby server start error %s", http.ListenAndServe(fmt.Sprintf(addressPattern, port), handler))
}