<template> 
  <div>
      <v-form id="SearchBarPos" onSubmit="return false;">
          <v-container>
              <v-text-field v-model="query" prepend-icon="mdi-magnify" color="white" clearable @keydown.enter="search"></v-text-field>
          </v-container>
      </v-form>
      <div style="text-align:center" id="SearchResults">
          <v-container v-for="video in videos" :key="video.id">
              <v-layout row>
                  <v-flex xs3>
                      <img :src="video.thumbnail" width="100%" height="100%">
                  </v-flex>
                  <v-flex>
                      <p>{{video.title}}</p>
                  </v-flex>
              </v-layout>
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
                    key: process.env.API_KEY
                }
            }).then(res => {
                const videoIds = res.data.items.map(item => item.id.videoId)
                const videos = res.data.items.map(item => ({id: item.id.videoId, title: item.snippet.title, thumbnail: item.snippet.thumbnails.medium.url}))
                
                axios.get(yturl + '/youtube/v3/videos', {
                    params: {
                        part: 'contentDetails',
                        id: videoIds.join(','),
                        key: process.env.API_KEY
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
        }
    }
}
</script>

<style>
    #SearchBarPos { 
        max-width: 600px;
        margin-top: 10px;
        margin-left: 15%;
        margin-right: 10px
    }

    #SearchResults {
        margin-left: 10%;
        margin-right: 10%;
        max-width: 50%;    
    }
</style>