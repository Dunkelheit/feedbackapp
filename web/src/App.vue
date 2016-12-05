<template>
    <div id="app">
        <nav v-if="loggedIn" class="navbar navbar-inverse navbar-fixed-top">
            <div class="container">
                <div class="navbar-header">
                    <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
                        <span class="sr-only">Toggle navigation</span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                        <span class="icon-bar"></span>
                    </button>
                    <a class="navbar-brand" href="#">The Feedback App</a>
                </div>
                <div id="navbar" class="collapse navbar-collapse">
                    <ul class="nav navbar-nav">
                        <router-link tag="li" to="/cards">
                            <a>Cards</a>
                        </router-link>
                        <router-link tag="li" to="/users">
                            <a>Users</a>
                        </router-link>
                        <router-link tag="li" to="/reviews">
                            <a>Reviews</a>
                        </router-link>
                    </ul>
                    <ul class="nav navbar-nav navbar-right">
                        <li><a href="#" v-on:click="logout">Logout</a></li>
                    </ul>
                </div><!--/.nav-collapse -->
            </div>
        </nav>
        <div class="container">
            <router-view></router-view>
        </div>
    </div>
</template>

<script>
export default {
    name: 'app',
    methods: {
        logout() {
            this.$store.commit('logout');
            delete localStorage.feedbackAppToken;
        }
    },
    computed: {
        loggedIn() {
            return this.$store.state.loggedIn;
        }
    }
}
</script>

<style>
body {
    padding-top: 50px;
}
</style>
