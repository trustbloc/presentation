<template lang='pug'>
#Slideshow
  .eg-slideshow
    slide
      h1 Aries Framework Go
      h4 Hello World!
</template>

<script>
import { Slideshow } from 'eagle.js'
export default {
  name: 'Slideshow',
  mixins: [ Slideshow ],
}
</script>
