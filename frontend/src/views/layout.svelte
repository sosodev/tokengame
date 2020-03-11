<script>
    import { Route } from 'svelte-router-spa';

    import Navbar from '../components/navbar.svelte';

    export let currentRoute;
    const params = {};
</script>

<div class="app">
    <Navbar/>
    <Route {currentRoute} {params} />
</div>
