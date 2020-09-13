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
            
            <v-list class="BtnMenu">  
                <v-list-group 
                 v-for="item in items"
                 :key="item.title"
                 v-model="item.active"
                 no-action
                 >
                    <template v-slot:activator>
                        <v-icon>mdi-format-list-bulleted-square</v-icon>
                        <v-list-item-title class="BtnTitle" v-text="item.title"></v-list-item-title>
                    </template>
                </v-list-group>
            </v-list>

            <v-btn class="BtnHome" block color="secondary" dark>Home</v-btn>

        </v-navigation-drawer>
    </div>  
</template>

<script>
//https://vuetifyjs.com/en/components/lists/
  export default {
    name: 'Sidebar',    
    
    data() {
        return{
            state: false,
            items: [
                {
                    action: '',
                    title: 'Player'
                },
                {
                    action: '',
                    title: 'Play lists'
                },
                {
                    action: '',
                    title: 'Songs'
                }
            ]
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

    .BtnMenu {
        margin-top: 20px !important;
    }

    .BtnTitle {
        margin-left: 10px;
        font-size: 20px !important;
    }

    .BtnHome{
        position: fixed !important;
        bottom: 0;
    }
</style>