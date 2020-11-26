<template>
  <div>
    <div id="SidebarButton">
      <v-btn class="mx-2" small outlined color="white" @click="state = !state">
        <v-icon dark>mdi-format-list-bulleted-square</v-icon>
      </v-btn>
    </div>
    <v-navigation-drawer v-model="state" app id="Sidebar">
      <v-list-item>
        <v-list-item-content id="Menu">
          <v-list-item-title id="Title">Radio V10</v-list-item-title>
          <v-list-item-subtitle id="Subtitle">Menu</v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>

      <v-divider></v-divider>

      <v-expansion-panels popout multiple>
        <v-expansion-panel class="Panels">
          <v-expansion-panel-header>
            <v-btn
              class="Btns"
              @click="
                $router.push('/player');
                state = !state;
              "
            >
              Player
            </v-btn>
          </v-expansion-panel-header>
          <v-expansion-panel-content id="Player">
            <SidebarPlayer />
          </v-expansion-panel-content>
        </v-expansion-panel>

        <v-expansion-panel class="Panels">
          <v-expansion-panel-header>
            <v-btn
              class="Btns"
              @click="
                $router.push('/player');
                state = !state;
              "
            >
              Playlists
            </v-btn>
          </v-expansion-panel-header>
          <v-expansion-panel-content>
            <v-btn class="BtnPlaylist">Playlist 1</v-btn>
            <v-btn class="BtnPlaylist">Playlist 2</v-btn>
            <v-btn class="BtnPlaylist">Playlist 3</v-btn>
            <v-btn class="BtnPlaylist">Playlist 4</v-btn>
            <v-btn class="BtnPlaylist">Playlist 5</v-btn>
          </v-expansion-panel-content>
        </v-expansion-panel>

        <v-expansion-panel class="Panels">
          <v-expansion-panel-header>
            <v-btn
              class="Btns"
              @click="
                $router.push('/tracks');
                state = !state;
              "
            >
              Tracks
            </v-btn>
          </v-expansion-panel-header>
          <v-expansion-panel-content>
            <SidebarTracks />
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>

      <v-btn
        class="BtnHome"
        block
        color="secondary"
        dark
        @click="
          $router.push('/');
          state = !state;
        "
      >
        Home
      </v-btn>
    </v-navigation-drawer>
  </div>
</template>

<script>
import SidebarPlayer from './SidebarPlayer';
import SidebarTracks from './SidebarTracks';

export default {
  name: 'Sidebar',
  components: {
    SidebarPlayer,
    SidebarTracks,
  },

  data() {
    return {
      state: false,
    };
  },
};
</script>

<style>
#SidebarButton {
  margin: 15px 0px 0px 10px;
}

#Sidebar {
  width: 210px !important;
  background-image: linear-gradient(to bottom left, #6ce4b6 65%, #2e8562 100%) !important;
}

#Menu {
  margin: 15px 0px 15px 0px;
}

#Title {
  font-size: 30px !important;
}

#Subtitle {
  font-size: 20px !important;
}

#Player {
  margin: 0px -35px 0px -35px;
}

.Panels {
  margin-top: 10px;
  background-color: #48aa83 !important;
}

.BtnPlaylist {
  margin-top: 5px;
}

.BtnHome {
  position: fixed !important;
  bottom: 0;
}

.Btns {
  margin-left: -10px;
  margin-right: 10px;
  background-color: #6ce4b6 !important;
}
</style>
