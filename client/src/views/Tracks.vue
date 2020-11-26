<template>
  <div style="text-align: center">
    <v-container v-for="savedTrack in savedTracks" :key="savedTrack.id">
      <div id="SearchResultsTrack">
        <v-layout row wrap>
          <v-flex xs3 id="TrackImage">
            <img :src="savedTrack.thumbnail" width="100%" height="100%" id="BoxShadow" />
          </v-flex>
          <v-flex xs7 id="TrackTitle">
            <p>{{ savedTrack.name }}</p>
          </v-flex>
          <v-flex id="Buttons">
            <v-btn class="Buttons">play</v-btn>
            <v-btn class="Buttons">enqueue</v-btn>
            <v-btn class="Buttons" @click="removeTrack(savedTrack.id)">remove</v-btn>
          </v-flex>
        </v-layout>
      </div>
    </v-container>
  </div>
</template>

<script>
import axios from 'axios';
import { mapState, mapGetters } from 'vuex';

export default {
  data() {
    return {};
  },
  computed: {
    ...mapState({
      savedTracks: 'savedTracks',
    }),
    ...mapGetters(['getSavedTracks']),
  },
  created() {
    axios.get(`http://${process.env.SERVER_URL}/tracks`).then((res) => {
      this.$store.dispatch('setSavedTracks', res.data);
    });
  },
  methods: {
    removeTrack(id) {
      axios.delete(`http://${process.env.SERVER_URL}/tracks/${id}`).then(() => {
        var index = this.savedTracks.findIndex((x) => x.id === id);
        this.savedTracks.splice(index, 1);
      });
    },
  },
};
</script>

<style></style>
