<script>
    import { Route } from 'svelte-router-spa';

    import Navbar from '../components/navbar.svelte';

    export let currentRoute;
    const params = {};
</script>

<style>
    .app {
        min-height: 100%;
        position: relative;
    }
</style>

<div class="app">
    <Navbar/>
    <Route {currentRoute} {params} />
</div>
