<script>
    import { Router } from 'svelte-router-spa';
    import { routes } from './routes.js';
</script>

<svelte:head>
    <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/bulma/0.8.0/css/bulma.css">
</svelte:head>

<Router {routes} />
