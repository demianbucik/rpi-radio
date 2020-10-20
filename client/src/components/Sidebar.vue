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

            <v-expansion-panels
             popout
             multiple
            >
                <v-expansion-panel class="Panels">
                    <v-expansion-panel-header>
                        <v-btn class="Btns" @click="$router.push('/player')">Player</v-btn>
                    </v-expansion-panel-header>
                    <v-expansion-panel-content id="Player">
                        <SidebarPlayer/>
                    </v-expansion-panel-content>
                </v-expansion-panel>

               <v-expansion-panel class="Panels">
                    <v-expansion-panel-header>
                       <v-btn class="Btns" @click="$router.push('/player')">Playlists</v-btn>
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
                       <v-btn class="Btns" @click="$router.push('/player')">Songs</v-btn>
                    </v-expansion-panel-header>
                   <v-expansion-panel-content>
                       Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
                   </v-expansion-panel-content>
               </v-expansion-panel>
            
            </v-expansion-panels>            
            
            <v-btn class="BtnHome" block color="secondary" dark @click="$router.push('/')">Home</v-btn>

        </v-navigation-drawer>
    </div>  
</template>

<script>
  import SidebarPlayer from './SidebarPlayer';

  export default {
    name: 'Sidebar',    
    components: {
        SidebarPlayer,
    },

    data() {
        return{
            state: false,            
        }
    }
}
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

    #Subtitle{
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

    .BtnHome{
        position: fixed !important;
        bottom: 0;
    }

    .Btns {
        margin-left: -10px;
        margin-right: 10px;
        background-color: #6ce4b6 !important;
    }
</style>