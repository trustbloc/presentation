<template >
    <Layout>
        <h2>DIDComm Mediation</h2>
        <ul>
            <li>Get an Invitation from the mediator -
                <button class="button" v-on:click="getRouterInvitation">Retrieve</button>
                <p>{{ invitation }}</p>
            </li>

            <li>Connect to the router</li>

            <li>Set the router</li>
        </ul>
    </Layout>
</template>

<script>
    import axios from 'axios';

    export default {
        name: 'DemoSetup',
        data() {
            return {
                invitation: null,
            };
        },
        metaInfo: {
            title: 'DIDComm Router Invitation'
        },
        methods: {
            getRouterInvitation: function () {
                axios
                    .post('https://didcomm.troyronda.com/router/admin/connections/create-invitation', {
                        alias: 'rolson'
                    })
                    .then(res => {
                        this.invitation = res.data;
                    })
            }
        }
    }
</script>

<style>
    .button {
        background-color: #a1f1a1; /* Green */
    }
</style>