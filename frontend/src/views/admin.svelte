<script>
    let username;
    let password;

    async function getChallenges() {
        return fetch('/api/challenges').then(resp => resp.json())
    }

    function deleteChallenge(id) {
        fetch('/admin/challenges/' + id, {
            method: 'DELETE',
            cache: 'no-cache'
        }).then(resp => {
            challenges = getChallenges();
        });
    }

    let challenges = getChallenges();
</script>

<style>
    .card {
        margin: 1rem auto;
    }
</style>

<div class="container">
    <div class="card">
        <div class="card-content">
            <h1 class="title has-text-centered">  
                The Back Panel
            </h1>
           {#await challenges}
                <p>Fetching challenges...</p>
            {:then json}
                {#each json.challenges as challenge}
                    <div class="level">
                        <div class="level-left">
                            <p class="level-item is-bold">
                                {challenge.title}
                            </p>
                        </div>
                        <div class="level-right">
                            <div class="level-item">
                                <button class="button is-danger" on:click={() => deleteChallenge(challenge.ID)}>
                                    Delete
                                </button>
                            </div>
                            <a class="level-item button is-danger" href="/admin/leaderboards/{challenge.ID}">
                                Remove Highscore
                            </a>
                        </div>
                    </div>
                {/each}
            {:catch error}
                <p>Failed to fetch challenges</p>
            {/await}
            <a class="level-item button is-success" href="/admin/challenges/new">
                New Challenge
            </a>
        </div>
    </div>
</div>
