<template>
  <v-app>
    <v-main>
      <Main/>
    </v-main>
  </v-app>
</template>

<script>
import Main from './components/Main';

export default {
  name: 'App',

  components: {
    Main,
  },

  data: () => ({
    //
  }),
};
</script>

<style scoped>
  .v-application {
    background-image: linear-gradient(to bottom right, #696969 70%, #4b4b4b 100%) !important;
  }
</style>>