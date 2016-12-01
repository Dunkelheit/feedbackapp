<template>
    <div>
        <h1>Cards</h1>
        <dl>
            <dt v-for="card in cards">
                <dd>
                    <ul>
                        <li>Title: {{card.title}}</li>
                        <li>Category: {{card.category}}</li>
                    </ul>
                </dd>
            </dt>
        </ul>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    name: 'cards',
    data: function () {
        return {
            cards: [{ title: 'Empty' }]
        };
    },
    beforeMount() {
        axios.get('/api/cards').then(response => {
            this.cards = response.data;
            console.log(this.cards);
        });
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
