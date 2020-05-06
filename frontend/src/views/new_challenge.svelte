<script>
    let title = "";
    let head = "";
    let foot = "";

    let new_token = "";
    let tokens = [];

    let new_testcase = "";
    let testcases = [];

    function add_token() {
        tokens.push(new_token);
        new_token = "";
        tokens = tokens;
    }

    function add_testcase() {
        testcases.push(new_testcase);
        new_testcase = "";
        testcases = testcases;
    }

    function submit_challenge() {
        fetch('/admin/challenges', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            cache: 'no-cache',
            body: JSON.stringify({
                title: title,
                head: head,
                foot: foot,
                tokens: tokens,
                testcases: testcases
            })
        }).then(resp => resp.json()).then(data => {
            window.location.href = '/challenges/' + data.ID;
        });
    }

    function remove_last_entry(array) {
        if (array.length > 0) {
            array.pop();
            testcases = testcases;
            tokens = tokens;
        }
    }
</script>

<style>
    .card {
        margin: 1rem auto;
        max-width: 40%;
    }

    @media (max-width: 1024px) {
        .card {
            margin: 0;
            max-width: 100%;
        }
    }
</style>

<div class="container">
    <div class="card">
        <div class="card-content">
            <h1 class="title">
                Create a Challenge
            </h1>
            <div class="field">
                <label class="label">
                    Title
                </label>
                <div class="control">
                    <input class="input" type="text" bind:value={title}>
                </div>
            </div>
            <div class="field">
                <label class="label">
                    Header Code
                </label>
                <div class="control">
                    <textarea class="textarea" bind:value={head}></textarea>
                </div>
            </div>
            <div class="field">
                <label class="label">
                    Footer Code
                </label>
                <div class="control">
                    <textarea class="textarea" bind:value={foot}></textarea>
                </div>
            </div>
            <nav class="panel">
                <p class="panel-heading">
                    Tokens
                </p>
                {#each tokens as token}
                <div class="panel-block">
                {token}
                </div>
                {/each}
                {#if tokens.length > 0}
                    <div class="panel-block">
                        <button class="button is-danger" on:click={() => remove_last_entry(tokens)}>
                            Remove Last
                        </button>
                    </div>
                {/if}
               <div class="panel-block">
                    <div class="field has-addons" style="width: 100%;">
                        <div class="control is-expanded">
                            <input class="input" type="text" placeholder="new + token" bind:value={new_token}>
                        </div>
                        <div class="control">
                            <button class="button is-info" on:click={add_token}>
                                Add Token
                            </button>
                        </div>
                    </div>
                </div>
            </nav>
            <nav class="panel">
                <p class="panel-heading">
                    Test Cases
                </p>
                {#each testcases as testcase}
                <div class="panel-block">
                {testcase}
                </div>
                {/each}
                {#if testcases.length > 0}
                <div class="panel-block">
                    <button class="button is-danger" on:click={() => remove_last_entry(testcases)}>
                        Remove Last
                    </button>
                </div>
                {/if}
                <div class="panel-block">
                    <div class="field has-addons" style="width: 100%;">
                        <div class="control is-expanded">
                            <input class="input" type="text" placeholder="function test1() ...; test1();" bind:value={new_testcase}>
                        </div>
                        <div class="control">
                            <button class="button is-info" on:click={add_testcase}>
                                Add Testcase
                            </button>
                        </div>
                    </div>
                </div>
            </nav>
            <button class="button is-success" on:click={submit_challenge}>
                Post It!
            </button>
        </div>
    </div>
</div>
