<template>
  <v-app>
    <v-main>
      <Sidebar/>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
  import Sidebar from './components/Sidebar'

  export default {
    name: 'App',

    components: {
     Sidebar,
    },

    data: () => ({
      //
    }),
};
</script>

<style>
  .v-application {
    background-image: linear-gradient(to bottom right, #696969 70%, #4b4b4b 100%) !important;
    color: whitesmoke !important;
  }
  </style>