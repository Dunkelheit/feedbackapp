<template>
    <div>
        <div class="page-header">
            <button type="button" class="btn btn-primary pull-right" v-on:click="openModal">Create</button>
            <h1>Cards</h1>
        </div>
        <div class="loading" v-if="loading">
            Loading...
        </div>
        <div v-if="error">
        </div>
        <div v-if="cards.length > 0">
            <table class="table table-striped">
                <thead>
                    <tr>
                        <th>Title</th>
                        <th>Category</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="card in cards">
                        <td><a href="#" v-on:click="loadCard(card.id, $event)">{{card.title}}</a></td>
                        <td>{{card.category === 0 ? 'Positive' : 'Negative'}}</td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div id="createModal" class="modal fade" tabindex="-1" role="dialog">
            <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
                        <h4 class="modal-title">{{ card.id ? 'Update' : 'Create' }} card</h4>
                    </div>
                    <div class="modal-body">
                        <form>
                            <input v-model="card.id" type="hidden" />
                            <div class="form-group">
                                <label for="inputTitle">Title</label>
                                <input type="text" v-model="card.title" class="form-control" id="inputTitle" placeholder="Type the name of the card here">
                            </div>
                            <div class="form-group">
                                <label for="inputCategory">Category</label>
                                <select id="inputCategory" class="form-control" v-model="card.category">
                                    <option value="0">Positive</option>
                                    <option value="1">Negative</option>
                                </select>
                            </div>
                        </form>
                    </div>
                    <div class="modal-footer">
                        <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
                        <button type="button" class="btn btn-primary" v-on:click="persistCard">{{ card.id ? 'Update' : 'Create' }}</button>
                    </div>
                </div>
            </div>
        </div>
    <div>
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
        openModal(event) {
            if (event) {
                event.preventDefault();
            }
            this.card.id = null;
            this.card.title = '';
            $('#createModal').modal('show');
        },
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
                this.fetchCards();
                $('#createModal').modal('hide');
            });
        },
        fetchCards() {
            this.error = null;
            this.cards = [];
            this.loading = true;
            axios.get('/api/cards').then(response => {
                this.loading = false;
                this.cards = response.data;
            });
        },
        loadCard(cardId, event) {
            if (event) {
                event.preventDefault();
            }
            this.cards.forEach(card => {
                if (card.id === cardId) {
                    this.card.id = card.id;
                    this.card.title = card.title;
                    this.card.category = card.category;
                    $('#createModal').modal('show');
                    return;
                }
            });
        }
    }
}
</script>

