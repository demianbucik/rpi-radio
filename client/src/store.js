import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    savedTracks: [],
  },
  mutations: {
    SET_SAVED_TRACKS(state, newTracks) {
      state.savedTracks = newTracks;
    },
  },
  actions: {
    setSavedTracks(context, newTracks) {
      context.commit('SET_SAVED_TRACKS', newTracks);
    },
  },
  getters: {
    getSavedTracks: (state) => {
      return state.savedTracks;
    },
  },
});
