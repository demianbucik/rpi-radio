<template>
  <v-app>
    <v-main>
      <Sidebar />
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script>
import Sidebar from './components/Sidebar';
import axios from 'axios';

export default {
  name: 'App',

  components: {
    Sidebar,
  },

  data: () => ({
    //
  }),
  created() {
    axios.get(`http://${process.env.SERVER_URL}/tracks`).then((res) => {
      this.$store.dispatch('setSavedTracks', res.data);
    });
  },
};
</script>

<style>
.v-application {
  background-image: linear-gradient(to bottom right, #696969 70%, #4b4b4b 100%) !important;
  color: whitesmoke !important;
}

::selection {
  background: rgba(43, 43, 43, 0.459);
}
</style>
