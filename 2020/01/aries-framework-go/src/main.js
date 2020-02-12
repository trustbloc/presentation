import Vue from 'vue'
import App from './App.vue'

// highlight.js
import hljs from 'highlight.js/lib/highlight';
import javascript from 'highlight.js/lib/languages/javascript';
import json from 'highlight.js/lib/languages/json';
import http from 'highlight.js/lib/languages/http';
import go from 'highlight.js/lib/languages/go';
hljs.registerLanguage('javascript', javascript);
hljs.registerLanguage('json', json);
hljs.registerLanguage('http', http);
hljs.registerLanguage('go', go);

// eagle.js
import Eagle, { Modal, CodeBlock, Options } from 'eagle.js'
import 'animate.css'
import 'eagle.js/dist/themes/agrume/agrume.css'
Options.hljs = hljs

import { Aries } from "@hyperledger/aries-framework-go"
Vue.prototype.$aries = new Aries({})

Eagle.use(Modal)
Eagle.use(CodeBlock)

Vue.use(Eagle)

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
