import Layout from './views/layout.svelte';

import Index from './views/index.svelte';
import Challenge from './views/challenge.svelte';
import Leaderboard from './views/leaderboard.svelte';
import RemovalLeaderboard from './views/removal_leaderboard.svelte';
import NewChallenge from './views/new_challenge.svelte';
import Admin from './views/admin.svelte';

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
    }, {
        name: '/admin/leaderboards/:id',
        component: RemovalLeaderboard,
        layout: Layout
    }, {
        name: '/dashboard',
        component: Admin,
        layout: Layout
    } , {
        name: '/admin/challenges/new',
        component: NewChallenge,
        layout: Layout
    }
];

export { routes };
