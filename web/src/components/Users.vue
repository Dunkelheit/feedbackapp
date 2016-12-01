<template>
    <div>
        <div class="loading" v-if="loading">
            Loading...
        </div>
        <div v-if="error" class="error">
        </div>
        <div v-if="users.length > 0" class="content">
            <dl>
                <dt>Users</dt>
                <dd>
                    <ul v-for="user in users">
                        <li>Name: {{user.fullName}}</li>
                        <li>Email: {{user.email}}</li>
                        <li>Job Title: {{user.jobTitle}}</li>
                        <li>Department: {{user.department}}</li>
                        <li>Company: {{user.company}}</li>
                    </ul>
                </dd>
            </dl>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    name: 'users',
    data () {
        return {
            loading: false,
            users: [],
            error: null
        };
    },
    created() {
        this.fetchUsers();
    },
    watch: {
        '$route': 'fetchUsers'
    },
    methods: {
        fetchUsers() {
            this.error = null;
            this.users = [];
            this.loading = true;
            axios.get('/api/users').then(response => {
                this.loading = false;
                this.users = response.data;
                console.log(this.users);
            });
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
