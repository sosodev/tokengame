<svelte:head>
    <title>
        Leaderboards - Token Tournament
    </title>
</svelte:head>

<style>
    .container {
        padding-top: 3rem;
    }

    .table {
        width: 40%;
    }
</style>

<script>
    export let currentRoute;

    async function getChallenge(id) {
        return fetch('/api/challenges/' + id).then(resp => resp.json())
    }
</script>

<div class="container">
    {#await getChallenge(currentRoute.namedParams.id)}
        <h1 class="title">Loading highscores...</h1>
    {:then json}
        <h1 class="title">
            {json.title} Leaderboards
        </h1>
        <table class="table">
            <thead>
                <tr>
                    <th>Nickname</th>
                    <th>Score</th>
                </tr>
            </thead>
            <tbody>
                {#each json.highscores as highscore}
                    <tr>
                        <td>{highscore.nickname}</td>
                        <td>{highscore.score}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    {:catch error}
        <h1 class="title">
            OOOPSIE WOOPSIE!!!
        </h1>
        <h3 class="subtitle">
            Uwu We made a fucky wucky!! A wittle fucko boingo! The code monkeys
            at our headquarters are working VEWY HAWD to fix this!
        </h3>
    {/await}
</div>
