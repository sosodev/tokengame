<svelte:head>
    <title>
        Leaderboards Administration - Token Tournament
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

    async function getHighscores(id) {
        return fetch('/api/challenges/' + id).then(resp => resp.json())
    }

    function deleteHighscore(id) {
        fetch('/admin/highscores/' + id, {
            method: 'DELETE',
            cache: 'no-cache'
        }).then(resp => {
            highscores = getHighscores(currentRoute.namedParams.id);
        });
    }

    let highscores = getHighscores(currentRoute.namedParams.id);
</script>

<div class="container">
    {#await highscores}
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
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                {#each json.highscores as highscore}
                    <tr>
                        <td>{highscore.nickname}</td>
                        <td>{highscore.score}</td>
                        <td>
                            <button id="highscore-{highscore.ID}" class="button is-danger" on:click={() => deleteHighscore(highscore.ID)}>
                                Remove
                            </button>
                        </td>
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
