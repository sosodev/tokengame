/*
	Some useful documentation links :)

	Svelte (the reactive framework): https://svelte.dev
	Svelte Router (enables the single-page-app): https://github.com/jorgegorka/svelte-router
	General JavaScript Stuff: https://developer.mozilla.org/en-US/docs/Web/JavaScript
*/

import App from './App.svelte';

const app = new App({
	target: document.body,
});

export default app;