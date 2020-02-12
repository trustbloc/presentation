<template >
    <div class="eg-transition" :enter='enter' :leave='leave'>
        <div class="eg-slide" v-if="active">
            <div class="eg-slide-content">
                <h2>DIDComm Mediation{{ role }}</h2>
                <div>Get an Invitation from the mediator -
                    <button class="button" v-on:click="getRouterInvitation">Retrieve</button>
                    <eg-transition enter='fadeIn' leave='bounceOutLeft' style=padding-top:10px>
                        <eg-code-block lang="json" v-if="invitation != null && step === 1" enter='flipInY'>{{ invitation }}</eg-code-block>
                    </eg-transition>
                    <eg-transition enter='fadeIn' leave='fadeOut'>
                        <p v-if="invitation != null && step > 1">{{ invitation.label }} endpoint: <b>{{ invitation.serviceEndpoint }}</b></p>
                    </eg-transition>
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
                    <eg-transition enter='fadeIn' leave='fadeOut'>
                        <p>{{ routerConnnectionStatus }}</p>
                    </eg-transition>
                </div>

                <div v-if="step === 4">Register the router -
                    <button class="button" v-on:click="registerRouter">Register</button>
                    <button class="button" v-on:click="unregisterRouter">Unregister</button>
                    <eg-transition enter='fadeIn' leave='fadeOut'>
                        <p>{{ registerStatus }}</p>
                    </eg-transition>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import { Slide } from 'eagle.js'
    import { AriesREST } from '../ariesrest.js'

    export default {
        name: 'DemoSetup',
        props: {
            steps: { default: 4 },
            routerURL: String,
            agentURL: String,
            role: String,
        },
        mixins: [Slide],
        data() {
            return {
                invitation: null,
                connectionStatus: null,
                connectionID: null,
                registerStatus: null,
                routerConnnectionStatus: null,
                ariesAgentClient: this.$aries,
                ariesRouterClient: new AriesREST({ agentURL: this.routerURL }),
            };
        },
        metaInfo: {
            title: 'DIDComm Router Invitation'
        },
        methods: {
            getRouterInvitation: function () {
                try {
                    this.ariesRouterClient.didexchange.CreateInvitation({}).then(res => {
                            this.invitation = res.data.invitation;
                        })
                } catch (err) {
                    this.invitation = err
                }
            },
            connectToRouter: async function () {
                if (this.invitation === null) {
                    this.connectionStatus = "Retrieve the invitation before proceeding"
                } else {
                    this.connectionStatus = "connecting"

                    try {
                        let res = await this.ariesAgentClient.didexchange.receiveInvitation(this.invitation)
                        this.connectionID = res.data.connection_id
                    } catch (err) {
                        this.connectionStatus = err
                        return
                    }

                    const attempts = 40
                    for (let i =0; i < attempts; i++) {
                        await new Promise(r => setTimeout(r, 250));
                        let res = await this.ariesAgentClient.didexchange.queryConnectionByID(this.connectionID)

                        if (res.data.result.State == 'completed') {
                            this.connectionStatus = res.data.result.State
                            break;
                        }

                        if (i == attempts - 1) {
                            this.connectionStatus = "timed out at state: " + res.data.result.State
                        }
                    }
                }
            },
            registerRouter: function () {
                if (this.routerConnnectionStatus !== "success") {
                    this.registerStatus = "Make sure connection with router is complete."
                } else {
                    this.ariesAgentClient.router.register({
                            "connectionID": this.connectionID
                        })
                        // eslint-disable-next-line no-unused-vars
                        .then(res => {
                            this.registerStatus = "success"
                        }).catch(error => {
                            this.registerStatus = error.response.data
                        })
                }
            },
            unregisterRouter: async function () {
                try {
                    await this.ariesAgentClient.router.unregister({})
                    this.registerStatus = "router unregistered"
                } catch (err) {
                    this.registerStatus = err
                }  
            },
            routerConnStatus: async function () {
                let res = await this.ariesAgentClient.didexchange.queryConnectionByID(this.connectionID)

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