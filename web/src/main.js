import Vue from 'vue';
import VueRouter from 'vue-router';
import Vuex from 'vuex'

import App from './App';
import Home from './components/Home';
import Cards from './components/Cards';
import Users from './components/Users';
import Reviews from './components/Reviews';

Vue.use(VueRouter);
Vue.use(Vuex);

const routes = [{
    path: '/',
    component: Home
}, {
    path: '/cards',
    component: Cards
}, {
    path: '/users',
    component: Users
}, {
    path: '/reviews',
    component: Reviews
}];

const store = new Vuex.Store({
    state: {
        loggedIn: localStorage.feedbackAppToken
    },
    mutations: {
        login(state, token) {
            state.loggedIn = token;
        },
        logout(state) {
            state.loggedIn = false;
        }
    }
});

const router = new VueRouter({
    routes, // short for routes: routes
    linkActiveClass: 'active'
});

/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    store,
    template: '<App/>',
    components: { App }
});
