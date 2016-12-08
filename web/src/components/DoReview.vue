<template>
    <div>
        <div class="page-header">
            <h1>Review somebody</h1>
        </div>
        <div class="loading" v-if="loading">
            Loading...
        </div>
        <div v-if="error" class="alert alert-danger" role="alert">
            <strong>Oh snap!</strong> Something went wrong.
        </div>

        <div class="content cards">

            <div class="card-container" v-for="card in cards">
                <div class="card">
                    <div class="front">
                        <p>{{card.title}}</p>
                    </div>
                    <div class="back">
                        <p>{{card.category === 0 ? 'This is a GOOD thing, congrats!' : 'This is a BAD thing. Get your sh*t together!'}}</p>
                    </div>
                </div>
            </div>

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
            review: null,
            cards: [],
            error: null
        };
    },
    created() {
        this.fetchReviewAndCards();
    },
    watch: {
        '$route': 'fetchReviewAndCards'
    },
    methods: {
        fetchReviewAndCards() {
            this.error = null;
            this.review = null;
            this.cards = [];
            this.loading = true;
            axios.get('/api/admin/cards', {
                headers: {
                    'x-auth-token': this.$store.state.loggedIn
                }
            }).then(response => {
                this.cards = response.data;
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

/* entire container, keeps perspective */
.card-container {
	perspective: 1000px;
    float: left;
    margin: 0 1em 1em 0;
}
	/* flip the pane when hovered */
	.card-container:hover .card, .flip-container.hover .card {
		transform: rotateY(180deg);
	}

.card-container, .front, .back {
	width: 160px;
	height: 240px;
}

/* flip speed goes here */
.card {
	transition: 0.6s;
	transform-style: preserve-3d;

	position: relative;
}

/* hide back of pane during swap */
.front, .back {
	backface-visibility: hidden;
	position: absolute;
	top: 0;
	left: 0;
    padding: 1em;
    text-align: center;
}

/* front pane, placed above back */
.front {
	z-index: 2;
	/* for firefox 31 */
	transform: rotateY(0deg);
    background-color: lightgreen;
}

/* back, initially hidden pane */
.back {
	transform: rotateY(180deg);
    background-color: lightblue;
}
</style>
