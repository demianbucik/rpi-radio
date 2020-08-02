<template>
  <div class="home">
    <div>
      <SidebarButton/>
    </div>
    <img id="logo" alt="Vue logo" src="../assets/RadioV10Logo.png">
    <Main/>
  </div>
</template>

<script>
// @ is an alias to /src
import Main from '@/components/Main.vue'

export default {
  name: 'Home',
  components:{
    Main
  }
}
</script>
