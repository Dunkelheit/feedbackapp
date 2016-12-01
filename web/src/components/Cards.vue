<template>
    <div>
        <h1>Cards</h1>
        <div class="loading" v-if="loading">
            Loading...
        </div>
        <div v-if="error" class="error">
        </div>
        <div v-if="cards.length > 0" class="content">
            <dl>
                <dt v-for="card in cards">
                    <dd>
                        <ul>
                            <li>Title: {{card.title}}</li>
                            <li>Category: {{card.category}}</li>
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
    name: 'cards',
    data () {
        return {
            loading: false,
            cards: [],
            error: null
        };
    },
    created() {
        this.fetchCards();
    },
    watch: {
        '$route': 'fetchCards'
    },
    methods: {
        fetchCards() {
            this.error = null;
            this.cards = [];
            this.loading = true;
            axios.get('/api/cards').then(response => {
                this.loading = false;
                this.cards = response.data;
                console.log(this.cards);
            });
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
