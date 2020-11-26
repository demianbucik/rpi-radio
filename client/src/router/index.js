import Vue from 'vue';
import VueRouter from 'vue-router';
/*import { component } from 'vue/types/umd'*/
import Home from '../views/Home.vue';
import Player from '../views/Player.vue';
import Search from '../views/Search.vue';
import Tracks from '../views/Tracks.vue';

Vue.use(VueRouter);

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/player',
    name: 'Player',
    component: Player,
  },
  {
    path: '/search',
    name: 'Search',
    component: Search,
  },
  {
    path: '/tracks',
    name: 'Tracks',
    component: Tracks,
  },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes,
});

export default router;
