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
    ADD_SAVED_TRACK(state, newTrack) {
      state.savedTracks.push(newTrack);
    },
    REMOVE_SAVED_TRACK_AT_INDEX(state, index) {
      state.savedTracks.splice(index, 1);
    },
  },
  actions: {
    setSavedTracks(context, newTracks) {
      context.commit('SET_SAVED_TRACKS', newTracks);
    },
    pushSavedTrack(context, newTrack) {
      context.commit('ADD_SAVED_TRACK', newTrack);
    },
    removeSavedTrackAtIndex(context, index) {
      context.commit('REMOVE_SAVED_TRACK_AT_INDEX', index);
    },
  },
  getters: {
    getSavedTracks: (state) => {
      return state.savedTracks;
    },
  },
});
