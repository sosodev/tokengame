<svelte:head>
    <title>
        Leaderboards - Token Tournament
    </title>
</svelte:head>

<style>
    .container {
        padding-top: 1rem;
    }
</style>

<script>
    async function getHighscores() {
        return fetch('/api/highscores').then(resp => resp.json())
    }

    async function getChallenge(id) {
        return fetch('/api/challenges/' + id).then(resp => resp.json())
    }
</script>

<div class="container">
    <h1 class="title">
        Latest Highscores
    </h1>
    {#await getHighscores()}
        <h3 class="subtitle">Loading highscores...</h3>
    {:then json}
        <table class="table">
            <thead>
                <tr>
                    <th>Nickname</th>
                    <th>Score</th>
                    <th>Challenge</th>
                </tr>
            </thead>
            <tbody>
                {#each json.highscores as highscore}
                    <tr>
                        <td>{highscore.nickname}</td>
                        <td>{highscore.score}</td>
                        <td>
                            {#await getChallenge(highscore.challenge_id)}
                            <td>Loading challenge</td>
                            {:then json}
                            <td>
                            <span>
                                <a href="/challenges/{json.ID}">
                                    {json.title}
                                </a>
                            </span>
                            </td>
                            {/await}
                        </td>
                    </tr>
                {/each}
            </tbody>
        </table>
    {:catch error}
        <h3 class="subtitle">
            OOOPSIE WOOPSIE!!!
        </h3>
        <p>
            Uwu We made a fucky wucky!! A wittle fucko boingo! The code monkeys
            at our headquarters are working VEWY HAWD to fix this!
        </p>
    {/await}
</div>
