<template>
    <div>
        <div class="page-header">
            <h1>My Reviews</h1>
        </div>
        <div class="loading" v-if="loading">
            Loading...
        </div>
        <div v-if="error" class="alert alert-danger" role="alert">
            <strong>Oh snap!</strong> Something went wrong.
        </div>
        <div v-if="reviews.length > 0" class="content">
            <table class="table table-striped">
                <thead>
                    <tr>
                        <th>Your reviewee</th>
                        <th>Completed</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="review in reviews">
                        <td><router-link :to="{ name: 'doReview', params: { id: review.id }}">{{review.reviewee.fullName}}</router-link></a></td>
                        <td>{{review.completed ? 'Yes' : 'No'}}</td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    name: 'reviews',
    data () {
        return {
            loading: false,
            reviews: [],
            error: null
        };
    },
    created() {
        this.fetchMyReviews();
    },
    watch: {
        '$route': 'fetchMyReviews'
    },
    methods: {
        fetchMyReviews() {
            this.error = null;
            this.reviews = [];
            this.loading = true;
            axios.get('/api/my/reviews', {
                headers: {
                    'x-auth-token': this.$store.state.loggedIn
                }
            }).then(response => {
                this.reviews = response.data;
                this.loading = false;
            }).catch(error => {
                this.error = error;
            });
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
