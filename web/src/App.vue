<template>
  <div id="app">
    <p>
      <!-- use router-link component for navigation. -->
      <!-- specify the link by passing the `to` prop. -->
      <!-- <router-link> will be rendered as an `<a>` tag by default -->
      <router-link to="/cards">Cards</router-link>
      <router-link to="/users">Users</router-link>
      <router-link to="/reviews">Reviews</router-link>
    </p>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
    name: 'app',
    components: {
    }
}
</script>

<style>
</style>
