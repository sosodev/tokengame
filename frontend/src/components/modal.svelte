<script>
    export let active = false;
</script>

<style>
    .modal-background {
        background-color: rgba(10, 10, 10, 0.5);
    }
</style>

<div class="modal" class:is-active={active}>
    <div class="modal-background" on:click={() => active = false}></div>
    <div class="modal-content">
        <slot/>
    </div>
    <button class="modal-close is-large" aria-label="close" on:click={() => active = false}></button>
</div>
