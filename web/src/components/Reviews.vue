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
                <dt v-for="card in reviews">
                    <dd>
                        <ul>
                            <li>Title: {{review.id}}</li>
                        </ul>
                    </dd>
                </dt>
            </dl>
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
        this.fetchReviews();
    },
    watch: {
        '$route': 'fetchReviews'
    },
    methods: {
        fetchReviews() {
            this.error = null;
            this.reviews = [];
            this.loading = true;
            axios.get('/api/reviews').then(response => {
                this.loading = false;
                this.reviews = response.data;
                console.log(this.reviews);
            });
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
