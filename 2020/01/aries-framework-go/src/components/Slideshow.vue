<template lang='pug'>
#Slideshow.eg-theme-agrume.eg-theme-custom
  .eg-slideshow
    slide
      h2 Aries Framework Go
      h4 Enables trusted communication between agent apps and DIDComm services.
      h4 Decentralized identity standards including DIDs, DIDComm, DID resolution, and verifiable credentials.
      h4 Aries protocol and message implementations in Go.
    slide
      h2 Controller bindings
      h4 Supplies a higher layer API to control the framework's services, message handlers and protocols.
      h4 REST and WebAssembly (with JavaScript) variations. Available via docker image and npm.
    slide
      h2 Additional goals
      h4 A pluggable dependency framework, where implementors can customize primitives.
      h4 Batteries included default depdendencies. Only requires the Go compiler.
      h4 Portable across environments - runs in both servers and browser sandboxes.
    slide
      h2 Go project structure
      h4 pkg/framework/aries and pkg/client/* contains higher layer Go API
      h4 cmd/aries-agent-rest and pkg/controller/rest contains the OpenAPI endpoints and spec annotations.
      h4 cmd/aries-js-worker and pkg/controller/command enables usage via a JavaScript API.
    slide
      h2 Go project structure
      h4 pkg/didcomm contains DIDComm and Aries packages.
      h4 pkg/doc/did and pkg/doc/vc contains DID and VC encoding and decoding.
      h4 pkg/vdri contains registries and resolvers along with pkg/store abstracting storage.
    slide
      h2 Go sample code
      pre
        eg-code-block(lang="go").
          framework, err := aries.New()
          if err != nil {...}

          ctx, err := framework.Context()
          if err != nil {...}

          didex, err := didexchange.New(ctx)
          if err != nil {...}

          invitation, err := didex.CreateInvitation("bob")
          if err != nil {...}
    slide
      h2 REST controller
      pre
        eg-code-block(lang="http").
          POST /connections/create-invitation HTTP/1.1
          Host: 0.0.0.0:9083
          Content-Type: application/json

          HTTP/1.1 200 OK
          {"invitation":{"serviceEndpoint":"routing:endpoint","recipientKeys":["2pHiF3XL6Y2K2zcpGCSrwE3odZsfB4DZvaQ422Jw1d3F"],"@id":"c115f8a0-3840-4be5-b5a8-bc10a345ca4c","label":"user-agent","@type":"https://didcomm.org/didexchange/1.0/invitation"},"alias":"","invitation_url":""}    slide
    slide
      h2 JavaScript controller
      pre
        eg-code-block(lang="javascript").
          import { Aries } from "@hyperledger/aries-framework-go"
          const aries = new Aries({})
          let res = aries.didexchange.createInvitation({})
    slide
      h2 Mediation and relays
      h4 Framework can run as a router (supporting Forward and Route Coordination).
      h4 Framework can run as a agent that is registered to a router.
      h4 HTTP and WebSocket transports are supported.
    slide
      h2 Demonstration
      h4 Send a basic message between edge agents using a mediator to facilitate transport.
    slide
      h3 Edge Agent Routing Topology
      p.center
        mermaid.
          graph TB;
          A1(Agent One) -->|HTTP| R(Mediator);
          R -->|WS| A(Agent Two);
    <DemoSetup :routerURL=routerAgentURL :agentURL=humanAgentURL role="" />
    <DemoSetup :routerURL=routerAgentURL :agentURL=botAgentURL role=" (bot)" v-if="this.showBotSetup" />
    <DemoBasicMessage :botAgentURL=botAgentURL :humanAgentURL=humanAgentURL :humanAgentWebhookURL=humanAgentWebhookURL :autoBotMsgRegister=autoBotMsgRegister />
</template>

<script>
import { Slideshow } from 'eagle.js'
import mermaid from './mermaid.vue'
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
    mermaid,
    DemoSetup,
    DemoBasicMessage
  }
}
</script>
