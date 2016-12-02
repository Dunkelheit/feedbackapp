<template>
    <div class="container">
        <form class="form-signin">
            <h2 class="form-signin-heading">Please sign in</h2>
            <label for="inputUsername" class="sr-only">IceMobile username</label>
            <input v-model="username" type="text" id="inputUsername" class="form-control" placeholder="IceMobile username" required autofocus>
            <label for="inputPassword" class="sr-only">Password</label>
            <input v-model="password" type="password" id="inputPassword" class="form-control" placeholder="Password" required>
            <button class="btn btn-lg btn-primary btn-block" type="submit" v-on:click="login">Sign in</button>
        </form>
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
    methods: {
        login() {
            axios.post('/api/login', {
                username: this.username,
                password: this.password
            }).then(response => {
                console.log(response);
            })
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
body {
    padding-top: 40px;
    padding-bottom: 40px;
    background-color: #eee;
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
