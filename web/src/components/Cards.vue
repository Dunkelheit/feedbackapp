<template>
    <div>
        <div class="loading" v-if="loading">
            Loading...
        </div>
        <div v-if="error" class="error">
        </div>
        <div v-if="cards.length > 0" class="content">
            <dl>
                <dt>Cards</dt>
                <dd>
                    <ul v-for="card in cards">
                        <li>Title: {{card.title}}</li>
                        <li>Category: {{card.category}}</li>
                        <li><router-link :to="{name: 'cardById', params: { id: card.id }}">Edit</router-link></li>
                    </ul>
                </dd>
            </dl>
        </div>
        <div>
            <input v-model="card.id" type="hidden" />
            <input v-model="card.title" placeholder="Title" />
            <select v-model="card.category">
                <option value="0">Positive</option>
                <option value="1">Negative</option>
            </select>
            <button v-on:click="persistCard">{{ card.id ? 'Update' : 'Create' }}</button>
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
                id: null,
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
        persistCard(event) {
            var action;
            var url;
            if (this.card.id !== null) {
                action = 'put';
                url = '/api/cards/' + this.card.id;
            } else {
                action = 'post';
                url = '/api/cards';
            }
            axios[action](url, {
                title: this.card.title,
                category: parseInt(this.card.category, 10)
            }).then(response => {
                this.card.id = null;
                this.card.title = '';
                if (action === 'put') {
                    this.$router.push('/cards');
                } else {
                    this.fetchCards();
                }
            });
        },
        fetchCards() {
            this.error = null;
            this.cards = [];
            this.loading = true;
            axios.get('/api/cards').then(response => {
                this.loading = false;
                this.cards = response.data;
                if (this.$route.params.id) {
                    this.loadCard(parseInt(this.$route.params.id, 10));
                }
            });
        },
        loadCard(cardId) {
            this.cards.forEach(card => {
                if (card.id === cardId) {
                    this.card.id = card.id;
                    this.card.title = card.title;
                    this.card.category = card.category;
                    return;
                }
            });
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
