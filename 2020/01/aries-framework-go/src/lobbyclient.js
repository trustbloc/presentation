/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/


'use strict'

import axios from 'axios';

/* @class LobbyClient provides the client API for the lobby server. */
export const LobbyClient = function(lobbyURL) {
    this.demo = {
        CreateDemo: async function(routerURL) {
            const endpoint = lobbyURL + "/demo"
            return axios.post(endpoint, {"routerUrl": routerURL})
        },
        GetDemo: async function(demoID) {
            const endpoint = lobbyURL + "/demo/" + demoID
            return axios.get(endpoint)
        }
    }

    this.invitation = {
        UploadInvitation: async function(invitation, demoID) {
            const endpoint = lobbyURL + "/demo/" + demoID + "/invitation"
            alert(endpoint)
            return axios.post(endpoint, invitation)
        },
        GetInvitations: async function(demoID) {
            const endpoint = lobbyURL + "/demo/" + demoID + "/invitations"
            return axios.get(endpoint)
        }
    }
}
