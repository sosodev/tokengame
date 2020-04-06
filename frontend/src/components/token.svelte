<script>
    export let cost;
</script>

<style>
    div {
        min-width: 150px;
        cursor: pointer;
    }
</style>

<div class="box has-text-centered hljs" on:click>
    <slot/> <span class="has-text-info"> ({cost}) </span>
</div>
