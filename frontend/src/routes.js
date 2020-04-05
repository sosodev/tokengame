import Layout from './views/layout.svelte';

import Index from './views/index.svelte';
import Challenge from './views/challenge.svelte';
import leaderboards from './views/leaderboards.svelte';

const routes = [
    {
        name: '/',
        component: Index,
        layout: Layout
    }, {
        name: '/challenges/:id',
        component: Challenge,
        layout: Layout
    }, {
        name: '/leaderboards',
        component: leaderboards,
        layout: Layout
    }
];

export { routes };
