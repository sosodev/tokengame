import Layout from './views/layout.svelte';

import Index from './views/index.svelte';
import Challenge from './views/challenge.svelte';
import Leaderboard from './views/leaderboard.svelte';

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
        name: '/leaderboards/:id',
        component: Leaderboard,
        layout: Layout
    }
];

export { routes };
