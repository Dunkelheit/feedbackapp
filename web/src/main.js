import Vue from 'vue';
import VueRouter from 'vue-router';
import Vuex from 'vuex';

import App from './App';
import Home from './components/Home';
import Cards from './components/admin/Cards';
import Users from './components/admin/Users';
import Reviews from './components/admin/Reviews';

Vue.use(VueRouter);
Vue.use(Vuex);

const routes = [{
    path: '/',
    component: Home,
    name: 'home'
}, {
    path: '/admin/cards',
    component: Cards,
    name: 'adminCards'
}, {
    path: '/admin/users',
    component: Users,
    name: 'adminUsers'
}, {
    path: '/admin/reviews',
    component: Reviews,
    name: 'adminReviews'
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
