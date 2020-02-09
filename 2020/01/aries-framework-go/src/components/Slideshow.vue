<template lang='pug'>
#Slideshow.eg-theme-agrume
  .eg-slideshow
    slide
      h1 Aries Framework Go
      h4 Hello World!
    <DemoSetup :routerURL=routerAgentURL :agentURL=humanAgentURL role="" />
    <DemoSetup :routerURL=routerAgentURL :agentURL=botAgentURL role=" (bot)" v-if="this.showBotSetup != 'false'" />
    <DemoBasicMessage :botAgentURL=botAgentURL :humanAgentURL=humanAgentURL :humanAgentWebhookURL=humanAgentWebhookURL />
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
    showBotSetup: { default: process.env.VUE_APP_SHOW_BOT_SETUP },
  },
  mixins: [ Slideshow ],
  components: {
    DemoSetup,
    DemoBasicMessage
  }
}
</script>
