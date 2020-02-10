<template lang='pug'>
#Slideshow.eg-theme-agrume
  .eg-slideshow
    slide
      h1 Aries Framework Go
      h4 Enables trusted communication between agent apps and DIDComm services.
      h4 Decentralized identity standards including DIDs, DIDComm, DID resolution, and verifiable credentials.
      h4 Aries protocol and message implementations in Go.
    slide
      h1 Controller bindings
      h4 Supplies a higher layer API to control the framework's services, message handlers and protocols.
      h4 REST and WebAssembly (with JavaScript) variations. Available via docker image and npm.
    slide
      h1 Additional goals
      h4 A pluggable dependency framework, where implementors can customize primitives.
      h4 Batteries included default depdendencies. Only requires the Go compiler.
      h4 Portable across environments - runs in both servers and browser sandboxes.
    slide
      h1 Go project structure
      h4 pkg/framework/aries and pkg/client/* contains higher layer Go API
      h4 cmd/aries-agent-rest and pkg/controller/rest contains the OpenAPI endpoints and spec annotations.
      h4 cmd/aries-js-worker and pkg/controller/command enables usage via a JavaScript API.
    slide
      h1 Go project structure
      h4 pkg/didcomm contains DIDComm and Aries packages.
      h4 pkg/doc/did and pkg/doc/vc contains DID and VC encoding and decoding.
      h4 pkg/vdri contains registries and resolvers along with pkg/store abstracting storage.
    slide
      h1 Supports mediation and relays
      h4 Framework can run as a router (supporting Forward and Route Coordination).
      h4 Framework can run as a agent that is registered to a router.
      h4 HTTP and WebSocket transports are supported.
    slide
      h1 Demonstration
      h4 Send a basic message between edge agents using a mediator to facilitate transport.
    <DemoSetup :routerURL=routerAgentURL :agentURL=humanAgentURL role="" />
    <DemoSetup :routerURL=routerAgentURL :agentURL=botAgentURL role=" (bot)" v-if="this.showBotSetup" />
    <DemoBasicMessage :botAgentURL=botAgentURL :humanAgentURL=humanAgentURL :humanAgentWebhookURL=humanAgentWebhookURL :autoBotMsgRegister=autoBotMsgRegister />
</template>

<script>
import { Slideshow } from 'eagle.js'
import DemoSetup from './DemoSetup.vue'
import DemoBasicMessage from './DemoBasicMessage.vue'

export default {
  name: 'Slideshow',
  props: {
    routerAgentURL: { default: process.env.VUE_APP_ROUTER_AGENT_URL },
    humanAgentURL: { default: process.env.VUE_APP_HUMAN_AGENT_URL },
    humanAgentWebhookURL: { default: process.env.VUE_APP_HUMAN_AGENT_WEBHOOK_URL },
    botAgentURL: { default: process.env.VUE_APP_BOT_AGENT_URL },
    showBotSetup: { default: process.env.VUE_APP_SHOW_BOT_SETUP != 'false' },
    autoBotMsgRegister: { default: process.env.VUE_APP_AUTO_BOT_MSG_REGISTER != 'false' },
  },
  mixins: [ Slideshow ],
  components: {
    DemoSetup,
    DemoBasicMessage
  }
}
</script>
