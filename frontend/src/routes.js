import Layout from './views/layout.svelte';

import Index from './views/index.svelte';

const routes = [
    {
        name: '/',
        component: Index,
        layout: Layout
    }
];

export { routes };
