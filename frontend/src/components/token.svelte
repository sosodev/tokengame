<script>
    export let cost;
</script>

<style>
    div {
        min-width: 50px;
        cursor: pointer;
        font-family: Consolas, monospace;
        font-size: 1.2rem;
    }
</style>

<div class="box has-text-centered hljs" on:click>
    <slot/> <span class="has-text-info"> ({cost}) </span>
</div>
