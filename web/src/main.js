import Vue from 'vue';
import VueRouter from 'vue-router';
import Vuex from 'vuex';

import App from './App';
import Home from './components/Home';
import MyReviews from './components/MyReviews';
import DoReview from './components/DoReview';
import AdminCards from './components/admin/Cards';
import AdminUsers from './components/admin/Users';
import AdminReviews from './components/admin/Reviews';

Vue.use(VueRouter);
Vue.use(Vuex);

const routes = [{
    path: '/',
    component: Home,
    name: 'home'
}, {
    path: '/my/reviews',
    component: MyReviews,
    name: 'myReviews'
}, {
    path: '/my/reviews/:id',
    component: DoReview,
    name: 'doReview'
}, {
    path: '/admin/cards',
    component: AdminCards,
    name: 'adminCards'
}, {
    path: '/admin/users',
    component: AdminUsers,
    name: 'adminUsers'
}, {
    path: '/admin/reviews',
    component: AdminReviews,
    name: 'adminReviews'
}];

const store = new Vuex.Store({
    state: {
        loggedIn: localStorage.feedbackAppToken,
        role: localStorage.feedbackAppRole
    },
    mutations: {
        login(state, data) {
            state.loggedIn = data.token;
            state.role = data.role;
        },
        logout(state) {
            state.loggedIn = false;
            state.role = 'user';
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
