/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/


package api

import (
	"encoding/json"
	"errors"
	"fmt"
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
	vars := mux.Vars(r)
	id := vars["uid"]

	demo, err  := api.getDemo(id)
	if err != nil {
		http.Error(w, "error = " +err.Error(), http.StatusInternalServerError)
		return
	}

	demoBytes, err := json.Marshal(demo)

	w.Write(demoBytes)
}

func (api *API)  PostInvitation (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["uid"]

	_, err  := api.getDemo(id)
	if err != nil {
		http.Error(w, "Invalid Demo ID : error = " +err.Error(), http.StatusInternalServerError)
		return
	}

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read payload", http.StatusInternalServerError)
		return
	}

	invitation := &Invitation{}
	err = json.Unmarshal(payload, invitation)
	if err != nil {
		http.Error(w, "payload marshal error", http.StatusInternalServerError)
		return
	}

	invitations, err  := api.getInvitations(id)
	if err != nil && !errors.Is(err, store.ErrDataNotFound){
		http.Error(w, "error = " +err.Error(), http.StatusInternalServerError)
		return
	}

	if errors.Is(err, store.ErrDataNotFound) {
		invitations = &Invitations{
			DIDExchangeInvitations: []Invitation{
			},
		}
	}

	newInvitations := &Invitations{
		DIDExchangeInvitations:append(invitations.DIDExchangeInvitations, *invitation),
	}

	invitationByte, err := json.Marshal(newInvitations)

	err  = api.store.Put(invitationDBKey(id), invitationByte)
	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}
}

func (api *API)  GetInvitations (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["uid"]

	_, err  := api.getDemo(id)
	if err != nil {
		http.Error(w, "Invalid Demo ID : error = " +err.Error(), http.StatusInternalServerError)
		return
	}

	invitations, err  := api.getInvitations(id)
	if err != nil {
		http.Error(w, "error = " +err.Error(), http.StatusInternalServerError)
		return
	}

	invitationsByte, err := json.Marshal(invitations)

	w.Write(invitationsByte)
}

func (api *API)  getInvitations (id string) (*Invitations, error) {
	data, err := api.store.Get(invitationDBKey(id))
	if err != nil {
		return nil, err
	}

	invitations := &Invitations{}
	err = json.Unmarshal(data, invitations)
	if err != nil {
		return nil, fmt.Errorf("db data unmarshal : %w", err)
	}

	return invitations, nil
}

func (api *API)  getDemo (id string) (*Demo, error) {
	data, err := api.store.Get(id)
	if err != nil {
		return nil, err
	}

	demo := &Demo{}
	err = json.Unmarshal(data, demo)
	if err != nil {
		return nil, fmt.Errorf("db data unmarshal : %w", err)
	}

	return demo, err
}

func invitationDBKey(id string) string {
	return "invitations-" + id
}
