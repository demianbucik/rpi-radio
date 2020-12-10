<template>
  <div>
    <v-form id="SearchBarPos" onSubmit="return false;">
      <v-container>
        <v-text-field
          v-model="query"
          prepend-icon="mdi-magnify"
          color="white"
          clearable
          @keydown.enter="search"
        ></v-text-field>
      </v-container>
    </v-form>
    <div style="text-align: center">
      <v-container v-for="track in tracks" :key="track.id">
        <div id="SearchResultsTrack">
          <v-layout row wrap>
            <v-flex xs3 id="TrackImage">
              <img :src="track.thumbnail" width="100%" height="100%" id="BoxShadow" />
            </v-flex>
            <v-flex xs7 id="TrackTitle">
              <p>{{ track.name }}</p>
            </v-flex>
            <v-flex id="Buttons">
              <v-btn class="Buttons">play</v-btn>
              <v-btn class="Buttons">enqueue</v-btn>
              <v-btn class="Buttons" @click="saveTrack(track)">save</v-btn>
            </v-flex>
          </v-layout>
        </div>
      </v-container>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import { mapState, mapGetters } from 'vuex';
const yturl = 'https://www.googleapis.com';

export default {
  name: 'Search',
  data() {
    return {
      query: '',
      tracks: [],
    };
  },
  computed: {
    ...mapState({
      savedTracks: 'savedTracks',
    }),
    ...mapGetters(['getSavedTracks']),
  },
  methods: {
    search() {
      axios
        .get(yturl + '/youtube/v3/search', {
          params: {
            part: 'snippet',
            maxResults: 15,
            q: this.query,
            key: process.env.API_KEY,
            type: 'video',
          },
        })
        .then((res) => {
          const trackIds = res.data.items.map((item) => item.id.videoId);
          axios
            .get(yturl + '/youtube/v3/videos', {
              params: {
                part: 'contentDetails',
                id: trackIds.join(','),
                key: process.env.API_KEY,
                type: 'video',
              },
            })
            .then((vidRes) => {
              let tempObj = {};
              for (const item of res.data.items) {
                tempObj[item.id.videoId] = {
                  id: item.id.videoId,
                  name: item.snippet.title,
                  thumbnail: item.snippet.thumbnails.medium.url,
                };
              }
              for (const item of vidRes.data.items) {
                tempObj[item.id].duration = item.contentDetails.duration;
              }
              this.tracks = Object.values(tempObj);
            })
            .catch((err) => {
              console.log(err);
            });
        })
        .catch((err) => {
          console.log(err);
        });
    },

    saveTrack(track) {
      axios
        .post(`http://${process.env.SERVER_URL}/tracks`, {
          name: track.name,
          url: `https://www.youtube.com/watch?v=${track.id}`,
          thumbnail: track.thumbnail,
        })
        .then((res) => {
          this.$store.dispatch('pushSavedTrack', res.data);
        });
    },
  },
};
</script>

<style>
#SearchBarPos {
  max-width: 600px;
  margin-top: 10px;
  margin-left: 15%;
  margin-right: 10px;
}

#TrackImage {
  height: 100%;
  margin: 15px 10px 10px 25px;
}

#BoxShadow {
  box-shadow: 5px 5px 15px rgb(145, 145, 145);
}

#TrackTitle {
  margin-top: 15px;
  color: aliceblue;
}

#SearchResultsTrack {
  margin-top: -5px;
  background-color: rgb(66, 66, 66);
  border-radius: 10px;
}

@media screen and (max-width: 540px) {
  #SearchResultsTrack {
    margin-left: 7.5%;
    margin-right: 7.5%;
    max-width: 85%;
  }

  #Buttons {
    margin: -10px 0px 10px 0px;
  }

  .Buttons {
    margin: 5px 5px 5px 10px;
  }
}

@media screen and (min-width: 540px) {
  #SearchResultsTrack {
    margin-left: 10%;
    max-width: 50%;
    margin-right: 40%;
  }

  #Buttons {
    margin-left: 100px;
    margin-top: -15px;
  }

  .Buttons {
    margin: -72.5px 5px 0px 20px;
  }
}
</style>
