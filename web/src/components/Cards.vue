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
        <div>
            <input v-model="card.title" placeholder="Title">
            <select v-model="card.category">
                <option value="0">Positive</option>
                <option value="1">Negative</option>
            </select>
            <button v-on:click="createCard">Create</button>
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
            error: null,
            card: {
                title: '',
                category: 0
            }
        };
    },
    created() {
        this.fetchCards();
    },
    watch: {
        '$route': 'fetchCards'
    },
    methods: {
        createCard(event) {
            axios.post('/api/cards', {
                title: this.card.title,
                category: parseInt(this.card.category, 10)
            }).then(response => {
                this.cards.push(response.data);
            });
        },
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
