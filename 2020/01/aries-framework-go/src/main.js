import Vue from 'vue'
import App from './App.vue'
import Eagle from 'eagle.js'
import 'animate.css'
import 'eagle.js/dist/themes/gourmet/gourmet.css'

Vue.use(Eagle)

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
