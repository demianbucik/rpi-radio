<template>
  <div id="app">
    <div id="nav">
      <router-link to="/">Home</router-link> |
      <router-link to="/about">About</router-link>
      <SidebarButton id="sidebarButton"/>
    </div>
    <router-view/>
  </div>
</template>

<script>
import SidebarButton from '@/components/SidebarButton.vue'

export default {
  name: 'App',
  components:{
    SidebarButton
  }
}
</script>>

<style>
html{
  height: 100%;
}

body{
  background-image: linear-gradient(to bottom right, #696969 70%, #4b4b4b 100%);
}

#sidebarButton{
  position: absolute;
  top: 30px;
  left: 30px;
}

#logo{
  max-width: 60%;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #b0cbe6;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}
</style>
