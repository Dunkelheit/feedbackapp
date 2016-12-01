<template>
    <div>
        <h1>Reviews</h1>
        <div class="loading" v-if="loading">
            Loading...
        </div>
        <div v-if="error" class="error">
        </div>
        <div v-if="reviews.length > 0" class="content">
            <dl>
                <dt v-for="review in reviews">
                    <dd>
                        <ul>
                            <li>{{review.uuid}}</li>
                            <li>{{review.reviewer.fullName}} reviews {{review.reviewee.fullName}}</li>
                            <li>Completed: {{review.completed}} </li>
                        </ul>
                    </dd>
                </dt>
            </dl>
        </div>
        <div>
            <select v-model="review.reviewerId">
                <option v-for="user in users" v-bind:value="user.id">{{user.fullName}}</option>
            </select>
            <select v-model="review.revieweeId">
                <option v-for="user in users" v-bind:value="user.id">{{user.fullName}}</option>
            </select>
            <button v-on:click="createReview">Create</button>
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
            users: [],
            error: null,
            review: {
                reviewerId: null,
                revieweeId: null
            }
        };
    },
    created() {
        this.fetchReviewsAndUsers();
    },
    watch: {
        '$route': 'fetchReviewsAndUsers'
    },
    methods: {
        createReview() {
            axios.post('/api/reviews', {
                reviewerId: this.review.reviewerId,
                revieweeId: this.review.revieweeId
            }).then(response => {
                this.reviews.push(response.data);
            });
        },
        fetchReviewsAndUsers() {
            this.error = null;
            this.reviews = [];
            this.users = [];
            this.loading = true;
            axios.all([
                axios.get('/api/reviews'),
                axios.get('/api/users')
            ]).then(axios.spread((reviews, users) => {
                this.reviews = reviews.data;
                this.users = users.data;
                this.loading = false;
            }));
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
