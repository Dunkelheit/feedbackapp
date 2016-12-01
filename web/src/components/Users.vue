<template>
    <div>
        <h1>Users</h1>
        <dl>
            <dt v-for="user in users">
                <dd>
                    <ul>
                        <li>Name: {{user.fullName}}</li>
                        <li>Job Title: {{user.jobTitle}}</li>
                    </ul>
                </dd>
            </dt>
        </ul>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    name: 'users',
    data: function () {
        return {
            users: [{ title: 'Empty' }]
        };
    },
    beforeMount() {
        axios.get('/api/users').then(response => {
            this.users = response.data;
            console.log(this.users);
        });
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
