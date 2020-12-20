<template>
  <div>
    <v-container v-for="savedTrack in savedTracks" :key="savedTrack.id">
      <div>
        <marquee-text @click="none" class="NoSelect">{{ savedTrack.name }}</marquee-text>
      </div>
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

<style scoped>
#TrackButton {
  max-width: 150px !important;
}

.NoSelect {
  -webkit-user-select: none;
  user-select: none;
}
</style>
