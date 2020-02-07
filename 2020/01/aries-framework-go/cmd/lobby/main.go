/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/


package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/trustbloc/presentation/2020/01/aries-framework-go/cmd/lobby/api"
	"github.com/trustbloc/presentation/2020/01/aries-framework-go/cmd/lobby/store/mem"
)

func main() {
	// get configs
	addr := ":9071"

	api := api.NewAPI(mem.New())

	// REST Handlers
	r := mux.NewRouter()
	r.HandleFunc("/demo", api.CreateDemo)
	r.HandleFunc("/demo/{uid}", api.GetDemo)
	r.HandleFunc("/postInvitation", api.PostInvitation)

	// start the server
	log.Fatal(http.ListenAndServe(addr, r))
}


