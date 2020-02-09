/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/hyperledger/aries-framework-go/pkg/common/log"
	"github.com/hyperledger/aries-framework-go/pkg/controller/command/didexchange"
	"github.com/hyperledger/aries-framework-go/pkg/controller/command/messaging"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/common/service"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/messaging/service/basic"
)

var logger = log.New("aries-framework/webhook")

const (
	addressPattern        = ":%s"
	basicMsgPath          = "/basic-message"
	connectionsPath       = "/connections"
	msgServiceOperationID = "/message"
	sendNewMsg            = msgServiceOperationID + "/send"
)

var (
	apiURL   string
	replyMsg string
)

type webHookMsg struct {
	Message  service.DIDCommMsgMap `json:"message"`
	MyDID    string                `json:"mydid"`
	TheirDID string                `json:"theirdid"`
}

func connections(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	connMsg := didexchange.ConnectionMsg{}

	err = json.Unmarshal(msg, &connMsg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	logger.Infof("received DID exchange state transition event :: connID=%s state=%s", connMsg.ConnectionID, connMsg.State)
}

func messages(w http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	logger.Infof("received generic msg event")

	incomingMsg := webHookMsg{}
	json.Unmarshal(msg, &incomingMsg)

	logger.Infof("received basic message :: MyDID=%s TheirDID=%s msg %v", incomingMsg.MyDID, incomingMsg.TheirDID, incomingMsg)

	err = processMessage(&incomingMsg)
	if err != nil {
		logger.Infof("unable to process msg")
	}

	w.WriteHeader(http.StatusOK)
}

func processMessage(hookMsg *webHookMsg) error {
	basicFromMsg := basic.Message{}
	hookMsg.Message.Decode(&basicFromMsg)

	basicToMsg := basic.Message{
		ID:      uuid.New().String(),
		Type:    basic.MessageRequestType,
		Content: fmt.Sprintf(replyMsg, basicFromMsg.Content),
		I10n: struct {
			Locale string `json:"locale"`
		}{
			Locale: "en",
		},
		SentTime: time.Now(),
	}

	rawBytes, err := json.Marshal(&basicToMsg)
	if err != nil {
		return fmt.Errorf("failed to get raw message bytes:  %w", err)
	}

	request := &messaging.SendNewMessageArgs{
		TheirDID:    hookMsg.TheirDID,
		MessageBody: rawBytes,
	}

	// call controller to send message
	err = sendHTTP(http.MethodPost, apiURL+sendNewMsg, request, nil)
	if err != nil {
		return fmt.Errorf("failed to send message : %w", err)
	}

	return nil
}

func sendHTTP(method, destination string, param, result interface{}) error {
	message, err := json.Marshal(param)
	if err != nil {
		return fmt.Errorf("failed to prepare params : %w", err)
	}

	// create request
	req, err := http.NewRequest(method, destination, bytes.NewBuffer(message))
	if err != nil {
		return fmt.Errorf("failed to create new http '%s' request for '%s', cause: %s", method, destination, err)
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")

	// send http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to get response from '%s', cause :%s", destination, err)
	}

	defer closeResponse(resp.Body)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to read response from '%s', cause :%s", destination, err)
	}

	logger.Debugf("Got response from '%s' [method: %s], response payload: %s", destination, method, string(data))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get successful response from '%s', unexpected status code [%d], "+
			"and message [%s]", destination, resp.StatusCode, string(data))
	}

	if result == nil {
		return nil
	}

	return json.Unmarshal(data, result)
}

func closeResponse(c io.Closer) {
	err := c.Close()
	if err != nil {
		logger.Errorf("Failed to close response body : %s", err)
	}
}

func main() {
	port := os.Getenv("WEBHOOK_PORT")
	if port == "" {
		panic("port to be passed as WEBHOOK_PORT env variable")
	}

	apiURL = os.Getenv("WEBHOOK_API_URL")
	if apiURL == "" {
		panic("controller to be passed as WEBHOOK_API_URL env variable")
	}

	replyMsg = os.Getenv("WEBHOOK_REPLY_MSG")
	if replyMsg == "" {
		panic("replyMsg to be passed as WEBHOOK_REPLY_MSG env variable")
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(connectionsPath, connections).Methods(http.MethodPost)
	router.HandleFunc(basicMsgPath, messages).Methods(http.MethodPost)

	handler := cors.Default().Handler(router)

	logger.Fatalf("webhook server start error %s", http.ListenAndServe(fmt.Sprintf(addressPattern, port), handler))
}
