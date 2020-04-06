<svelte:head>
    <title>
        Play Fun Programming Challenges - Token Tournament
    </title>
</svelte:head>

<style>
    nav {
        max-width: 800px;
    }

    .left-button {
        margin-left: auto;
        margin-right: 0;
    }
</style>

<script>
    async function getChallenges() {
        return fetch('/api/challenges').then(resp => resp.json())
    }
</script>

<section class="hero">
    <div class="hero-body">
        <div class="container">
            <h1 class="title">
                Token Tournament
            </h1>
            <nav class="panel">
                <p class="panel-heading">
                    Challenges
                </p>
                {#await getChallenges()}
                    <a class="panel-block" href="/">
                        Loading challenges...
                    </a>
                {:then json}
                    {#each json.challenges as challenge}
                        <a class="panel-block" href="/challenges/{challenge.ID}">
                            #{challenge.ID} - {challenge.title}
                            <span class="left-button">
                                <a class="button is-info" href="/leaderboards/{challenge.ID}">
                                    Leaderboards
                                </a>
                            </span>
                        </a>
                    {/each}
                {:catch error}
                    <a class="panel-block" href="/">
                        Failed to fetch challenges... Try refreshing the page shortly
                    </a>
                {/await}
            </nav>
        </div>
    </div>
</section>
