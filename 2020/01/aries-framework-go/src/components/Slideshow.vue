<template lang='pug'>
#Slideshow
  .eg-slideshow
    slide
      h1 Aries Framework Go
      h4 Hello World!
    slide
       <DemoSetup />
</template>

<script>
import { Slideshow } from 'eagle.js'
import DemoSetup from './DemoSetup.vue'

export default {
  name: 'Slideshow',
  mixins: [ Slideshow ],
  components: {
    DemoSetup
  }
}
</script>
