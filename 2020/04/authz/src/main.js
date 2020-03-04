import Vue from 'vue'
import Vuetify from "vuetify"
import "vuetify/dist/vuetify.min.css";
import App from './App.vue'
import Eagle from "eagle.js"
import "animate.css"

Vue.config.productionTip = false

Vue.use(Eagle)
Vue.use(Vuetify)

new Vue({
  render: h => h(App),
}).$mount('#app')
