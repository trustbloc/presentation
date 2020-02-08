<template >
    <div class="eg-transition" :enter='enter' :leave='leave'>
    <div class="eg-slide" v-if="active">
    <div class="eg-slide-content">
        <h2>DIDComm Mediation</h2>
        <div>Get an Invitation from the mediator -
            <button class="button" v-on:click="getRouterInvitation">Retrieve</button>
            <eg-transition enter='fadeIn' leave='bounceOutLeft' style=padding-top:10px>
                <eg-code-block lang="json" v-if="invitation != null && step === 1" enter='flipInY'>{{ invitation }}</eg-code-block>
            </eg-transition>
            <p v-if="invitation != null && step > 1">Router endpoint: <b>{{ invitation.serviceEndpoint }}</b></p>
        </div>

        <div v-if="step === 2">Connect to the router -
        <button class="button" v-on:click="connectToRouter">Connect</button>
            <eg-transition enter='fadeIn' leave='bounceOutLeft'>
                <p v-if="connectionStatus != null && step === 2">{{ connectionStatus }}</p>
            </eg-transition>
        </div>

        <div v-if="step === 3">Check the connection status with the router  -
            <button class="button" v-on:click="routerConnStatus">Status</button>
            <button class="button" v-on:click="clearRouterConnStatus">Clear</button>
            <p>{{ routerConnnectionStatus }}</p>
        </div>

        <div v-if="step === 4">Register the router -
        <button class="button" v-on:click="registerRouter">Register</button>
            <p>{{ registerStatus }}</p>
        </div>
    </div>
    </div>
    </div>
</template>

<script>
    import axios from 'axios';
    import { Slide } from 'eagle.js'

    var routerUrl = process.env.VUE_APP_ROUTER_AGENT_URL
    var agentUrl = process.env.VUE_APP_HUMAN_AGENT_URL

    export default {
        name: 'DemoSetup',
        props: {
            steps: { default: 4 },
        },
        mixins: [Slide],
        data() {
            return {
                invitation: null,
                connectionStatus: null,
                connectionID: null,
                registerStatus: null,
                routerConnnectionStatus: null,
            };
        },
        metaInfo: {
            title: 'DIDComm Router Invitation'
        },
        methods: {
            getRouterInvitation: function () {
                axios
                    .post(routerUrl + '/connections/create-invitation', {})
                    .then(res => {
                        this.invitation = res.data.invitation;
                    })
            },
            connectToRouter: async function () {
                if (this.invitation === null) {
                    this.connectionStatus = "Retrieve the invitation before proceeding"
                } else {
                    let res = await axios.post(agentUrl + '/connections/receive-invitation', this.invitation)
                    this.connectionID = res.data.connection_id

                    await new Promise(r => setTimeout(r, 1000));

                    res = await axios.get( agentUrl + "/connections/"+ this.connectionID)
                    this.connectionStatus = res.data.result.State
                }
            },
            registerRouter: function () {
                if (this.routerConnnectionStatus !== "success") {
                    this.registerStatus = "Make sure connection with router is complete."
                } else {
                    var registerRouterUrl = agentUrl + "/route/register"
                    axios
                        .post(registerRouterUrl, {
                            "connectionID": this.connectionID
                        })
                        .then(res => {
                            this.registerStatus = "success"
                        }).catch(error => {
                        this.registerStatus = error
                    })
                }
            },
            routerConnStatus: async function () {
                let res = await axios.get( agentUrl + "/connections/"+ this.connectionID)

                if (res.data.result.State == 'completed') {
                    this.routerConnnectionStatus = "success"
                } else {
                    this.routerConnnectionStatus = "in-progress"
                }
            },
            clearRouterConnStatus : function () {
                this.routerConnnectionStatus = ""
            }
        }
    }
</script>

<style>
    .button {
        background-color: #a1f1a1; /* Green */
    }
</style>