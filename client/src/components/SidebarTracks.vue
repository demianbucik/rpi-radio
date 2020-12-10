<template>
  <div>
    <v-container v-for="savedTrack in savedTracks" :key="savedTrack.id">
      <div>{{ savedTrack.name }}</div>
    </v-container>
  </div>
</template>

<script>
import { mapState, mapGetters } from 'vuex';

export default {
  name: 'SidebarTracks',
  data() {
    return {};
  },
  computed: {
    ...mapState({
      savedTracks: 'savedTracks',
    }),
    ...mapGetters(['getSavedTracks']),
  },
};
</script>

<style></style>
