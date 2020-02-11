<template >
    <div class="eg-transition" :enter='enter' :leave='leave'>
        <div class="eg-slide" v-if="active">
            <div class="eg-slide-content">
                <h2>Basic Message</h2>

                <div v-if="step === 1">
                    Create lobby -
                    <button class="button" v-on:click="createLobby">Create</button>
                    <p/>
                    Upload bot invitation -
                    <button class="button" v-on:click="uploadBotInvitation">Upload</button>
                    <p/>
                    Get invitations from the server -
                    <button class="button" v-on:click="getInvitations">Retrieve</button>
                    <eg-transition enter='fadeIn' leave='bounceOutLeft' style=padding-top:10px>
                        <eg-code-block lang="json" v-if="invitations != null && step === 1" enter='flipInY'>{{ invitations }}</eg-code-block>
                    </eg-transition>
                </div>

                <div v-if="step < 5">
                    <p v-if="invitations != null && step > 1">{{ invitations.label }} endpoint: <b>{{ invitations[0].serviceEndpoint }}</b></p>
                </div>

                <div v-if="step === 2">Connect to the bot -
                    <button class="button" v-on:click="connectToAgents">Connect</button>
                    <eg-transition enter='fadeIn' leave='fadeOut'>
                        <p v-if="connectionStatus != null && step === 2">{{ connectionStatus }}</p>
                    </eg-transition>
                </div>

                <div v-if="step === 3">Check the connection status with the bot  -
                    <button class="button" v-on:click="botConnStatus">Status</button>
                    <button class="button" v-on:click="clearBotConnStatus">Clear</button>
                    <eg-transition enter='fadeIn'>
                        <p v-if="botConnnectionStatus != null && step === 3">{{ botConnnectionStatus }}</p>
                    </eg-transition>
                </div>

                <div v-if="step === 4">Register for messages -
                    <button class="button" v-on:click="humanRegistration">Register</button>
                    <p v-if="registerStatus != null && step === 4">{{ registerStatus }}</p>
                </div>

                <div v-if="step === 5">
                    Chat with the bot
                    <eg-triggered-message trigger="botMessageContent != null" enter="bounceInRight" position="top right" leave="bounceOutRight">
                        <p>{{botMessageContent}}</p>
                    </eg-triggered-message>
<!--                    <eg-transition enter='fadeIn' leave='bounceOutLeft' style=padding-top:10px>-->
<!--                        <eg-code-block lang="json" v-if="botMessageResponse != null && (step === 5 || step === 6)" enter='flipInY'>{{ botMessageResponse }}</eg-code-block>-->
<!--                    </eg-transition>-->
                    <eg-transition enter='fadeIn' leave='bounceOutLeft' style=padding-top:10px>
                        <eg-code-block v-if="botMessageContent != null">{{ botMessageContent }}</eg-code-block>
                    </eg-transition>
                    <div class="center">
                        <div>What kind of pie do you like?</div>
                        <input v-model='pie' v-on:keyup.enter="sendMessage"/>
                        <button class="button" v-on:click="sendMessage">Send</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import { AriesREST, AriesWebHook } from '../ariesrest.js'
    import { Slide } from 'eagle.js'
    import {LobbyClient} from "../lobbyclient";
    const uuidv4 = require('uuid/v4');

    export default {
        name: 'DemoMultiUser',
        props: {
            steps: { default: 5 },
            lobbyURL: String,
            botAgentURL: String,
            humanAgentURL: String,
            routerURL: String,
            humanAgentWebhookURL: String,
            autoBotMsgRegister: Boolean
        },
        mixins: [Slide],
        data() {
            return {
                pie: "strawberry",
                invitations: [],
                connectionStatus: null,
                connectionIDs: [],
                registerStatus: null,
                botConnnectionStatus: null,
                botMessageResponse: null,
                justReceived: false,
                botMessageContent: null,
                demoID: null,
                ariesBotAgentClient: new AriesREST(this.botAgentURL),
                ariesHumanAgentClient: new AriesREST(this.humanAgentURL),
                ariesHumanWebHookClient: new AriesWebHook(this.humanAgentWebhookURL),
                lobbyClient: new LobbyClient(this.lobbyURL),
            };
        },
        metaInfo: {
            title: 'Basic Message via DIDComm'
        },
        methods: {
            createLobby: function() {
                this.lobbyClient.demo.CreateDemo(this.routerURL).then(res => {
                    this.demoID = res.data.uid
                })
            },
            uploadInvitation: function() {
                this.ariesHumanAgentClient.didexchange.CreateInvitation({}).then(res => {
                    this.lobbyClient.invitation.UploadInvitation(res.data.invitation, this.demoID)
                })
            },
            uploadBotInvitation: function() {
                this.ariesBotAgentClient.didexchange.CreateInvitation({}).then(res => {
                    this.lobbyClient.invitation.UploadInvitation(res.data.invitation, this.demoID)
                })
            },
            getInvitations: function() {
                this.lobbyClient.invitation.GetInvitations(this.demoID).then(res => {
                    this.invitations = res.data.invitations
                })
            },
            connectToAgents: async function() {
                if (this.invitations.length === 0) {
                    this.connectionStatus = "Retrieve the invitation before proceeding"
                } else {
                    for (var i = 0; i < this.invitations.length; i++) {
                        this.connectionStatus = "connecting to " + i

                        var connectionID

                        try {
                            let res = await this.ariesHumanAgentClient.didexchange.ReceiveInvitation(this.invitations[i])
                            connectionID = res.data.connection_id
                        } catch (err) {
                            this.connectionStatus = err
                            return
                        }
                        const attempts = 40

                        for (let i =0; i < attempts; i++) {
                            await new Promise(r => setTimeout(r, 250));
                            let res = await this.ariesHumanAgentClient.didexchange.QueryConnectionByID(connectionID)
                            if (res.data.result.State == 'completed') {

                                this.connectionStatus = res.data.result.State
                                this.connectionIDs.push(connectionID)
                                break;
                            }

                            if (i == attempts - 1) {
                                this.connectionStatus = "timed out at state: " + res.data.result.State
                            }
                        }
                    }
                }
            },
            botRegistration: async function() {
                try {
                    let res = await this.ariesBotAgentClient.messaging.RegisterMessageService({
                        "name" : "basic-message",
                        "type" : "https://didcomm.org/basicmessage/1.0/message"
                    })

                    console.log(res)
                } catch(err) {
                    console.log(err)
                }
            },
            humanRegistration: async function() {
                await this.unregisterMessage(this.ariesHumanAgentClient)
                await this.registerMessage(this.ariesHumanAgentClient)
            },
            registerMessage: async function(agentClient) {
                try {
                    let res = await agentClient.messaging.RegisterMessageService({
                        "name" : "basic-message",
                        "type" : "https://didcomm.org/basicmessage/1.0/message"
                    })

                    console.log(res)
                    this.registerStatus = "success"
                } catch(err) {
                    this.registerStatus = err
                }
            },
            unregisterMessage: async function(agentClient) {
                try {
                    let res = await agentClient.messaging.UnregisterMessageService({
                        "name" : "basic-message",
                    })
                    console.log(res)
                } catch(err) {
                    console.log(err)
                }
            },
            botConnStatus: async function() {
                for (const connectionID of this.connectionIDs) {
                    try {
                        let res = await this.ariesHumanAgentClient.didexchange.QueryConnectionByID(connectionID)

                        if (res.data.result.State == 'completed') {
                            this.botConnnectionStatus = "success"
                        } else {
                            this.botConnnectionStatus = "in-progress"
                        }
                    } catch(err) {
                        this.botConnnectionStatus = err
                    }
                }
            },
            clearBotConnStatus : function() {
                this.botConnnectionStatus = ""
            },
            sendMessage: async function() {
                for (const connectionID of this.connectionIDs) {
                    try {
                        await this.ariesHumanAgentClient.messaging.SendNewMessage({
                            "connection_ID": connectionID,
                            "message_body": {
                                "@id" : uuidv4(),
                                "@type" : "https://didcomm.org/basicmessage/1.0/message",
                                "content": "I love " + this.pie + " pie",
                                "sent_time": new Date().toISOString(),
                                "~l10n": {"locale":"en"}
                            },
                        })
                    } catch(err) {
                        this.botMessageResponse = err
                    }
                }

                this.nextStep()
            },
            refreshTopicsData: async function() {
                try {
                    var topicsResponse = await this.ariesHumanWebHookClient.topics.Check();
                    if (topicsResponse.data.message && topicsResponse.data.message['@type'] === "https://didcomm.org/basicmessage/1.0/message") {
                        this.botMessageResponse = topicsResponse.data;
                        this.botMessageContent = topicsResponse.data.message.content;
                        this.justReceived = true;
                    }
                } catch(err) {
                    this.botMessageResponse = err;
                }
            },
        },
        mounted: function () {
            if (this.autoBotMsgRegister) {
                this.botRegistration();
            }
            this.refreshTopicsData();

            setInterval(function () {
                this.refreshTopicsData();
            }.bind(this), 1000);
        }
    }

</script>

<style>
    .button {
        background-color: #a1f1a1; /* Green */
    }
</style>