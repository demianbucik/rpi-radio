<template>
  <div>
    <v-container v-for="savedTrack in savedTracks" :key="savedTrack.id">
      <div>
        <v-container>
          <v-row id="TrackContainer">
            <v-col id="TrackImg">
              <v-img height="20" width="20" src="../assets/PlayLogo.png"></v-img>
            </v-col>
            <v-col id="TrackTitle">
              <marquee-text class="NoSelect">{{ savedTrack.name }}</marquee-text>
            </v-col>
          </v-row>
        </v-container>
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
#TrackImg {
  margin: -5px 10px -5px -30px !important;
}

#TrackTitle {
  margin-top: -42.25px;
  color:  rgb(50, 50, 50);
}

#TrackContainer {
  margin-top: -30px !important;
}

.NoSelect {
  -webkit-user-select: none;
  user-select: none;
}
</style>
