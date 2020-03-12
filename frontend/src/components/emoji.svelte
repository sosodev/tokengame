<script>
    import { onDestroy, onMount } from 'svelte';

    function getRandomInt(max) {
        return Math.floor(Math.random() * Math.floor(max));
    }

    export let hidden = false;

    let x = 0;
    let y = 0;

    let winX;
    let winY;

    const interval = setInterval(() => {
        if (x > winX) {
            x = getRandomInt(winX);
        }

        if (y > winY) {
            y = 0;
        }

        x = x + 1;
        y = y + 1;
    }, 10);

    onMount(() => {
        x = getRandomInt(winX);
        y = getRandomInt(winY);
    });

    onDestroy(() => {
        clearInterval(interval);
    });
</script>

<style>
    div {
        position: absolute;
        font-size: 1.5rem;
    }
</style>

<svelte:window bind:innerWidth={winX} bind:innerHeight={winY} />

<div style="top: {y}px; left: {x}px;" class:is-hidden={!hidden}>
    <slot/>
</div>
