/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/


package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/trustbloc/presentation/2020/01/aries-framework-go/cmd/lobby/store"
)


type API struct {
	store store.Store
}

func NewAPI(store store.Store) *API{
	return &API{
		store : store,
	}
}

func (api *API) CreateDemo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "HTTP Method not allowed", http.StatusMethodNotAllowed)
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read payload", http.StatusInternalServerError)
		return
	}

	demo := &Demo{}
	err = json.Unmarshal(payload, demo)
	if err != nil {
		http.Error(w, "payload marshal error", http.StatusInternalServerError)
		return
	}

	// TODO validate fields
	id := uuid.New().String()

	err  = api.store.Put(id, payload)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	resp := map[string]string{
		"uid": id,
	}

	json.NewEncoder(w).Encode(resp)
}

func (api *API) GetDemo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "HTTP Method not allowed", http.StatusMethodNotAllowed)
	}


	vars := mux.Vars(r)
	id := vars["uid"]

	data, err := api.store.Get(id)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (api *API)  PostInvitation (w http.ResponseWriter, r *http.Request) {
	// TODO Implement
}