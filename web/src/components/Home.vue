<template>
    <div>
        <div v-if="loggedIn" class="page-header">
            <h1>Welecome to the Feedback App!</h1>
        </div>
        <div v-if="!loggedIn" class="container">
            <form v-on:submit="login" class="form-signin">
                <h2 class="form-signin-heading">The Feedback App</h2>
                <label for="inputUsername" class="sr-only">IceMobile username</label>
                <input v-model="username" type="text" id="inputUsername" class="form-control" placeholder="IceMobile username" required autofocus>
                <label for="inputPassword" class="sr-only">Password</label>
                <input v-model="password" type="password" id="inputPassword" class="form-control" placeholder="Password" required>
                <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
            </form>
            <div v-if="error" class="alert alert-danger" role="alert">
                <strong>Oh snap!</strong> Looks like your credentials are not correct.
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    name: 'home',
    data () {
        return {
            loading: false,
            error: null,
            usermame: null,
            password: null
        };
    },
    computed: {
        loggedIn() {
            return this.$store.state.loggedIn;
        },
        isAdmin() {
            return this.$store.state.role === 'admin';
        }
    },
    methods: {
        login() {
            this.error = null;
            axios.post('/api/login', {
                username: this.username,
                password: this.password
            }).then(response => {
                const token = response.headers['x-auth-token'];
                const role = response.data.role;
                this.$store.commit('login', {
                    token: token,
                    role: role
                });
                localStorage.feedbackAppToken = token;
                localStorage.feedbackAppRole = role;
                this.username = null;
                this.password = null;
            }).catch(error => {
                this.error = error;
            });
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.alert {
    max-width: 430px;
    margin: 0 auto;
}
.form-signin {
    max-width: 330px;
    padding: 15px;
    margin: 0 auto;
}
.form-signin .form-signin-heading,
.form-signin .checkbox {
    margin-bottom: 10px;
}
.form-signin .checkbox {
    font-weight: normal;
}
.form-signin .form-control {
    position: relative;
    height: auto;
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
    padding: 10px;
    font-size: 16px;
}
.form-signin .form-control:focus {
    z-index: 2;
}
.form-signin input[type="email"] {
    margin-bottom: -1px;
    border-bottom-right-radius: 0;
    border-bottom-left-radius: 0;
}
.form-signin input[type="password"] {
    margin-bottom: 10px;
    border-top-left-radius: 0;
    border-top-right-radius: 0;
}
</style>
