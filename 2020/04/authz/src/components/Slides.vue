<template lang="pug">
#Slideshow.eg-theme-agrume.eg-theme-custom
    .eg-slideshow
        slide
            h1 DIDAuthZ
        slide
            h2 Introduction
            p Introductory text.
        slide
            h2 Service Provider
            p Connect your wallet with the SP:
            <br/>
            <v-btn flat target="#">Connect</v-btn>
            <br/>
            p {show did-exchange invitation followed by resulting DIDs}
        slide
            h2 Relying Party
            p Connect your wallet with the RP:
            <br/>
            <v-btn flat target="#">Connect</v-btn>
            <br/>
            p {show did-exchange invitation followed by resulting DIDs}
            p TODO: should we demonstrate an ephemeral challenge instead rather than connecting to RP? (Out-Of-Band protocol)
        slide
            h2 Challenge for Credentials
            p The RP presents a challenge for credentials in the form of an auth request
            <br/>
            <v-btn flat target="#">Authenticate</v-btn>
            <br/>
            p {show challenge object}
            p {show transactional peer DID}
        slide
            h2 Select provider
            p Select the service provider
            <br/>
            p {display list of providers (1) here}
            p {show SP invitation + peer DID}
            p {show Consent Receipt object}
        slide
            h2 Confirmation
            p RP confirms your credentials
            <br/>
            p {notification: verification successful}
            p {show verifiable presentation of CR sent by RP}
            p {show resource received from SP}
        slide
            h2 Verify consents given
            p Inspect the consent that you just gave.
            <br/>
            p {display list of consents (1) here}
        slide
            h2 What just happened?
            <br/>
            p {explain the flow}
</template>

<script>
    import { Slideshow } from "eagle.js"
    export default {
        name: "Slides",
        mixins: [Slideshow]
    }
</script>