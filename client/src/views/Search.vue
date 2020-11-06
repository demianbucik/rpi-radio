<template> 
  <div>
      <v-form id="SearchBarPos" onSubmit="return false;">
          <v-container>
              <v-text-field v-model="query" prepend-icon="mdi-magnify" color="white" clearable @keydown.enter="search"></v-text-field>
          </v-container>
      </v-form>
      <div style="text-align:center">
          <v-container v-for="video in videos" :key="video.id">              
            <div id="SearchResultsVideo">
              <v-layout row wrap>
                  <v-flex xs3 id="VideoImage">
                      <img :src="video.thumbnail" width="100%" height="100%" id="BoxShadow">
                  </v-flex>
                  <v-flex xs7 id="VideoTitle">
                      <p>{{video.title}}</p>
                  </v-flex>
                  <v-flex id="Buttons">
                      <v-btn class="Buttons">play</v-btn>
                      <v-btn class="Buttons">play next</v-btn>
                      <v-btn class="Buttons" @click="save">save</v-btn>
                  </v-flex>
              </v-layout>
            </div>
          </v-container>
      </div>
  </div>
</template>

<script src="https://apis.google.com/js/api.js"></script>
<script>
  import axios from 'axios'
  const yturl = 'https://www.googleapis.com'
  
  export default {
    name: "Search",
    data() {
        return {
            query: '',
            videos: []
        }
    },
    methods: {
        search() {
            axios.get(yturl + '/youtube/v3/search', {
                params: {
                    part: 'snippet',
                    maxResults: 15,
                    q: this.query,
                    key: process.env.API_KEY,
                    type: 'video'
                }
            }).then(res => {
                const videoIds = res.data.items.map(item => item.id.videoId)
                const videos = res.data.items.map(item => ({id: item.id.videoId, title: item.snippet.title, thumbnail: item.snippet.thumbnails.medium.url}))
                
                
                axios.get(yturl + '/youtube/v3/videos', {
                    params: {
                        part: 'contentDetails',
                        id: videoIds.join(','),
                        key: process.env.API_KEY,
                        type: 'video'
                    }
                }).then(vidRes => {
                    const duration = vidRes.data.items.map(item => ({id: item.id, duration: item.contentDetails.duration}))
                    
                    let tempObj = {}
                    for(const item of res.data.items){
                        tempObj[item.id.videoId] = {id: item.id.videoId, title: item.snippet.title, thumbnail: item.snippet.thumbnails.medium.url}
                    } 
                    for(const item of vidRes.data.items){
                        tempObj[item.id].duration = item.contentDetails.duration
                    }
                    this.videos = Object.values(tempObj)
                }).catch(err => {console.log(err)})
            }).catch(err => {console.log(err)})
        },

        save() {
            axios.post("http://localhost:8000/tracks", {
                name: "Matter/persons from porlock ~ MANA",
                url: "https://www.youtube.com/watch?v=i2NlVQi9XUE",
                thumbnail: "https://i.ytimg.com/vi/i2NlVQi9XUE/hq720.jpg?sqp=-oaymwEZCOgCEMoBSFXyq4qpAwsIARUAAIhCGAFwAQ==&rs=AOn4CLCztjOHnyi2HuJ7CvHQ9T6CVwcT8g"
            })
        }
    }
}
</script>

<style scoped>
    #SearchBarPos { 
        max-width: 600px;
        margin-top: 10px;
        margin-left: 15%;
        margin-right: 10px
    }    

    #VideoImage {
        height: 100%;
        margin: 15px 10px 10px 25px;
    }

    #BoxShadow {
        box-shadow: 5px 5px 15px rgb(145, 145, 145);
    }

    #VideoTitle {
        margin-top: 15px;
        color: aliceblue;
    }

    #SearchResultsVideo {
        margin-top: -5px;
        background-color: rgb(66, 66, 66);
        border-radius: 10px;
    }

    @media screen and (max-width: 540px){
        #SearchResultsVideo {
            margin-left: 7.5%;
            margin-right: 7.5%;
            max-width: 85%;
        }

        #Buttons {
            margin: -10px 0px 10px 0px              
        }

        .Buttons {
            margin: 5px 5px 5px 10px;
        }
    }

    @media screen and (min-width: 540px){
        #SearchResultsVideo {            
            margin-left: 10%;
            max-width: 50%;
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