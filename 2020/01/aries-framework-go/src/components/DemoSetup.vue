<template >
    <Layout>
        <h2>DIDComm Mediation</h2>
        <ul>
            <li>Get an Invitation from the mediator -
                <button class="button" v-on:click="getRouterInvitation">Retrieve</button>
                <p>{{ invitation }}</p>
            </li>

            <li>Connect to the router -
            <button class="button" v-on:click="connectToRouter">Connect</button>
                <p>{{ connectionStatus }}</p>
            </li>

            <li>Check the connection status with the router  -
                <button class="button" v-on:click="routerConnStatus">Status</button>
                <button class="button" v-on:click="clearRouterConnStatus">Clear</button>
                <p>{{ routerConnnectionStatus }}</p>
            </li>

            <li>Register the router -
            <button class="button" v-on:click="registerRouter">Register</button>
                <p>{{ registerStatus }}</p>
            </li>

        </ul>
    </Layout>
</template>

<script>
    import axios from 'axios';

    var routerUrl = "https://didcomm.troyronda.com/router/admin"
    var agentUrl = "http://localhost:10081"

    export default {
        name: 'Connection',
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
                let res = await axios.post(agentUrl + '/connections/receive-invitation', this.invitation)
                this.connectionID = res.data.connection_id

                await new Promise(r => setTimeout(r, 1000));

                res = await axios.get( agentUrl + "/connections/"+ this.connectionID)
                this.connectionStatus = res.data.result.State
            },
            registerRouter: function () {
                var registerRouterUrl = agentUrl + "/route/register"
                axios
                    .post(registerRouterUrl, {
                        "connectionID" : this.connectionID
                    })
                    .then(res => {
                        console.log(res)
                        this.registerStatus = "success"
                    })
            },
            routerConnStatus: async function () {
                await axios.get( agentUrl + "/connections/"+ this.connectionID)
                this.routerConnnectionStatus = "success"
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