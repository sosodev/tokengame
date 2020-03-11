import Layout from './views/layout.svelte';

import Index from './views/index.svelte';
import BubbleSort from './views/bubble-sort.svelte';

const routes = [
    {
        name: '/',
        component: Index,
        layout: Layout
    }, {
        name: '/challenges/bubble_sort',
        component: BubbleSort,
        layout: Layout
    }
];

export { routes };
