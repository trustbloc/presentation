/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

'use strict'

import axios from 'axios';

/* @class AriesREST provides Aries agent functions. */
export const AriesREST = function(agentURL) {

    this.didexchange = {
        CreateInvitation: async function(data) {
            const endpoint = agentURL + "/connections/create-invitation"
            return axios.post(endpoint, data)
        },
        ReceiveInvitation: async function(data) {
            const endpoint = agentURL + "/connections/receive-invitation"
            return axios.post(endpoint, data)
        },
        AcceptInvitation: async function() {
            // TODO
            // POST /connections/{id}/accept-invitation
        },
        AcceptExchangeRequest: async function() {
            // TODO
            // POST /connections/{id}/accept-request
        },
        CreateImplicitInvitation: async function(data) {
            const endpoint = AriesREST._agentURL + "/connections/create-implicit-invitation"
            return axios.post(endpoint, data)
        },
        RemoveConnection: async function() {
            // TODO
            // POST /connections/{id}/remove
        },
        QueryConnectionByID: async function(connectionID) {
            const endpoint = agentURL + "/connections/" + connectionID
            return axios.get(endpoint)
        },
        QueryConnections: async function() {
            const endpoint = agentURL + "/connections"
            return axios.get(endpoint)
        }
    }

    this.messaging = {
        RegisteredServices: async function() {
            const endpoint = AriesREST._agentURL + "/message/services"
            return axios.get(endpoint)
        },
        RegisterMessageService: async function(data) {
            const endpoint = agentURL + "/message/register-service"
            return axios.post(endpoint, data)
        },
        RegisterHTTPMessageService: async function(data) {
            const endpoint = AriesREST._agentURL + "/http-over-didcomm/register"
            return axios.post(endpoint, data)
        },
        UnregisterMessageService: async function(data) {
            const endpoint = agentURL + "/message/unregister-service"
            return axios.post(endpoint, data)
        },
        SendNewMessage: async function(data) {
            const endpoint = agentURL + "/message/send"
            return axios.post(endpoint, data)
        },
        SendReplyMessage: async function(data) {
            const endpoint = AriesREST._agentURL + "/message/reply"
            return axios.post(endpoint, data)
        }
    }

    this.router = {
        Register: async function(data) {
            const endpoint = agentURL + "/route/register"
            return axios.post(endpoint, data)
        },
        Unregister: async function(data) {
            const endpoint = agentURL + "/route/unregister"
            return axios.delete(endpoint, data)
        },
    }
}

export const AriesWebHook = function(webhookURL) {
    this.topics = {
        Check: async function() {
            const endpoint = webhookURL + "/checktopics"
            return axios.get(endpoint);
        }
    }
}
