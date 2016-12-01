import Vue from 'vue';
import VueRouter from 'vue-router';
import App from './App';
import Cards from './components/Cards';
import Users from './components/Users';

Vue.use(VueRouter);

const Foo = {
    template: '<div>foo</div>'
}
const Bar = {
    template: '<div>bar</div>'
}

const routes = [
  { path: '/cards', component: Cards },
  { path: '/users', component: Users }
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
