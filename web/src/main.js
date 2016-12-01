import Vue from 'vue';
import VueRouter from 'vue-router';
import App from './App';
import Cards from './components/Cards';
import Users from './components/Users';
import Reviews from './components/Reviews';

Vue.use(VueRouter);

const routes = [
  { path: '/cards', component: Cards },
  { path: '/cards/:id', component: Cards, name: 'cardById' },
  { path: '/users', component: Users },
  { path: '/reviews', component: Reviews }
]

const router = new VueRouter({
  routes // short for routes: routes
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  template: '<App/>',
  components: { App },
  router: router
});
