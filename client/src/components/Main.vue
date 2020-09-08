<template>
  <v-container>
    <div class="text-center">
      
      <v-img id="logo"
        :src="require('../assets/RadioV10Logo.png')"
        class="my-3"
        contain
      />

      <h2 class="display-2 mb-3"> Welcome to</h2>
      <h1 class="display-3 font-weight-bold mb-3">Radio Vipavska 10</h1>
      <h3 class="display-1 mb-3"> No meme music!</h3>

      <p>
        For any and all information turn to
        <br><b>Demian</b> or <b>Andrej</b>.
      </p>
    </div>
    
  </v-container>
</template>

<script>
  export default {
    name: 'Main',

    data: () => ({
      
    }),
  }
</script>

<style>
  #logo {
    max-height: 250px;
    margin-top: 15px !important;
  }

  a {
    color: mediumspringgreen !important;
  }

  ::selection {
    background: rgba(43, 43, 43, 0.459) !important;
  }
</style>
