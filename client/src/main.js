import Vue from 'vue';
import App from './App.vue';
import './registerServiceWorker';
import router from './router';
import vuetify from './plugins/vuetify';
import MarqueeText from 'vue-marquee-text-component';
import store from './store';

Vue.config.productionTip = false;
Vue.component('marquee-text', MarqueeText);

new Vue({
  router,
  vuetify,
  store,
  render: (h) => h(App),
}).$mount('#app');
