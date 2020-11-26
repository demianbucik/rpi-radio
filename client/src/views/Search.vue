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
              <p>{{ track.title }}</p>
            </v-flex>
            <v-flex id="Buttons">
              <v-btn class="Buttons">play</v-btn>
              <v-btn class="Buttons">enqueue</v-btn>
              <v-btn class="Buttons" @click="savetrack(track)">save</v-btn>
            </v-flex>
          </v-layout>
        </div>
      </v-container>
    </div>
  </div>
</template>

<script src="https://apis.google.com/js/api.js"></script>
<script>
import axios from 'axios';
const yturl = 'https://www.googleapis.com';

export default {
  name: 'Search',
  data() {
    return {
      query: '',
      tracks: [],
    };
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
          const tracks = res.data.items.map((item) => ({
            id: item.id.videoId,
            title: item.snippet.title,
            thumbnail: item.snippet.thumbnails.medium.url,
          }));

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
              const duration = vidRes.data.items.map((item) => ({
                id: item.id,
                duration: item.contentDetails.duration,
              }));

              let tempObj = {};
              for (const item of res.data.items) {
                tempObj[item.id.videoId] = {
                  id: item.id.videoId,
                  title: item.snippet.title,
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

    savetrack(track) {
      axios.post(`http://${process.env.SERVER_URL}/tracks`, {
        name: track.title,
        url: `https://www.youtube.com/watch?v=${track.id}`,
        thumbnail: track.thumbnail,
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
