<template>
    <div>
        <div class="SidebarButton">
            <v-btn class="mx-2" small outlined color="white" @click="state = !state">
            <v-icon dark>mdi-format-list-bulleted-square</v-icon>
            </v-btn>
        </div>
        <v-navigation-drawer v-model="state" app class="Sidebar">
            
        </v-navigation-drawer>
    </div>  
</template>

<script>
export default {
    name: 'Sidebar',

    data() {
        return{
            state: false
        }
    }
}
</script>

<style>
    .SidebarButton {
        margin-top: 15px;
        margin-left: 10px;
    }

    .Sidebar {
        width: 200px !important; 
        background-image: linear-gradient(to bottom left, #6ce4b6 70%, #2e8562 100%) !important;
    }
</style>