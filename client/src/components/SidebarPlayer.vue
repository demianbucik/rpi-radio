<template>
  <v-container>
    <v-row class="Rows">
      <v-col align="center"> Title-tracks </v-col>
    </v-row>
    <v-row id="BtnsPlayer" class="Rows">
      <v-col class="BtnPlayerPos">
        <v-btn class="BtnPlayer">
          <v-img height="32" width="32" src="../assets/PreviousLogo.png"></v-img>
        </v-btn>
      </v-col>
      <v-col class="BtnPlayerPos">
        <v-btn class="BtnPlayer">
          <v-img height="32" width="32" src="../assets/PlayLogo.png"></v-img>
        </v-btn>
      </v-col>
      <v-col class="BtnPlayerPos">
        <v-btn class="BtnPlayer">
          <v-img height="32" width="32" src="../assets/NextLogo.png"></v-img>
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'SidebarPlayer',
};
</script>

<style>
.Rows {
  margin-top: -25px;
}

.BtnPlayerPos {
  max-width: 70px !important;
  margin-top: 10px;
}

.BtnPlayer {
  background-color: #6ce4b6 !important;
}

#BtnsPlayer {
  margin-left: -10px;
  margin-bottom: -30px;
}
</style>
