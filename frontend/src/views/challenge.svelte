<script>
    export let currentRoute;

    import { onMount, onDestroy } from 'svelte';
    import Token from '../components/token.svelte';
    import Modal from '../components/modal.svelte';
    import Notification from '../components/notification.svelte';

    import beautify from 'js-beautify';
    import hljs from 'highlight.js/lib/highlight';
    import javascript from 'highlight.js/lib/languages/javascript';

    hljs.registerLanguage('javascript', javascript);

    let challengeData = {title: '', head: '', foot: '', tokens: [], testcases: []}
    fetch('/api/challenges/' + currentRoute.namedParams.id).then(resp => resp.json()).then(data => {
        let tokens = data.tokens;
        for (let i = tokens.length - 1; i > 0; i--) {
            const j = Math.floor(Math.random() * (i + 1));
            [tokens[i], tokens[j]] = [tokens[j], tokens[i]];
        }
        data.tokens = tokens;
        challengeData = data;
    });

    let seconds = 0;
    onMount(() => {
        let ticker = setInterval(() => {
            seconds += 1;
        }, 1000);

        let keyboardEvent = document.addEventListener('keydown', function(event) {
            if (event.keyCode === 13) {
                if (searched_tokens.length === 1) {
                    useToken(searched_tokens[0]);
                    search_input = "";
                }
            }
        });

        onDestroy(() => {
            document.removeEventListener(keyboardEvent);
            clearInterval(ticker);
        });
    });

    $: title = challengeData.title;
    $: head = challengeData.head;
    $: foot = challengeData.foot;
    $: tokens = challengeData.tokens;
    $: testcases = challengeData.testcases;
    $: code = hljs.highlight('javascript', beautify(head + user_code + foot)).value;

    let finished = false;
    let user_code = "";
    let token_page = 0;
    let initial_budget = 500;
    let budget = initial_budget;
    let score = 0;
    let move_stack = [];
    let search_input = "";
    let nickname = "";
    $: searched_tokens = tokens.filter(el => el.includes(search_input));

    // split the tokens into subsections
    let column = [];
    $: {
        column = [];
        for (let i = 0; i < searched_tokens.length; i++) {
            column.push(searched_tokens[i]);
        }
        column = column;
    }

    function forward() {
        if (token_page * 5 < tokens.length - 1) {
            token_page++;
        }
    }

    function backward() {
        if (token_page > 0)
            token_page--;
    }

    function calculateCost(token) {
        if (token.length > 3)
            return 3;
        else
            return token.length;
    }

    function useToken(token) {
        move_stack.push({
            code: user_code,
            token: token,
        })

        user_code += token;
        budget -= calculateCost(token);
    }

    function undoToken() {
        let last_move = move_stack.pop()
        if (last_move) {
            user_code = last_move.code;
            budget += calculateCost(last_move.token);
        }
    }

    let failed_test = "";
    $: failed_test_message = "Sorry, but it looks like your code failed this test: " + failed_test;
    // Run the testcases against the user's code
    function runTests() {
        let full_code = head + user_code + foot;
        eval.call(window, full_code);

        let passed = true;

        // run the test cases and fail if the return is false
        testcases.forEach(test => {
            if (!eval.call(window, test)) {
                passed = false;
                failed_test = test;
                return;
            }
        });

        if (passed) {
            score = 10000 * (budget / initial_budget) - seconds * 4;
        }

        finished = passed;
    }

    function submitHighscore() {
        fetch('/api/highscores/new', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            cache: 'no-cache',
            body: JSON.stringify({
                nickname: nickname,
                score: score,
                challenge_id: Number(currentRoute.namedParams.id),
            }),
        }).then(resp => resp.json()).then(data => {
            window.location.href = '/leaderboards/' + data.challenge_id;
        });
    }
</script>

<svelte:head>
    <title>
        {title} Challenge - Token Tournament
    </title>

    <link rel="stylesheet"
          href="http://cdnjs.cloudflare.com/ajax/libs/highlight.js/9.18.1/styles/school-book.min.css">
</svelte:head>

<style>
    h1 {
        margin-top: 1rem;
        font-size: 1.7rem;
    }

    pre {
        font-size: 1.2rem;
    }

    input {
        font-family: Consolas, monospace;
        font-size: 1.2rem;
        margin-bottom: 1rem;
    }

    code {
        font-family: Consolas, monospace;
        font-size: 1.2rem;
    }

    .token-box {
        padding: 1rem;
        margin-bottom: 1rem;
    }
</style>

<div class="container" style="padding: 1rem;">
    {#if failed_test !== ""}
    <Notification text={failed_test_message}></Notification>
    {/if}
    <h1 class="title">
        {title}
    </h1>
    <h3 class="subtitle">
        Budget: {budget}
    </h3>
    <hr/>
<pre><code class="hljs javascript">{@html code}</code></pre>
    <div class="token-box has-background-light">
    <input class="input" type="text" placeholder="Search for a token" bind:value={search_input}>
    <div class="columns is-centered is-mobile is-multiline">
        {#each column as token}
            <div class="column is-narrow">
                <Token cost="{calculateCost(token)}" on:click={() => useToken(token)}>
                    {token}
                </Token>
            </div>
        {/each}
    </div>
    </div>
    <button class="button is-warning is-large" on:click={runTests}>
        Submit
    </button>
    <button class="button is-info is-large" on:click={undoToken}>
        Undo
    </button>
</div>

<Modal bind:active="{finished}">
    <div class="card">
        <div class="card-content">
            <h2 style="font-size: 1.5rem; font-weight: bold; white-space: nowrap;">
                You finished the {title} challenge! 🎉
            </h2>
            <h3 class="subtitle">
                Final Score: {score}
            </h3>
            <div class="field-body">
                <div class="field">
                    <div class="control is-expanded">
                        <input class="input" type="text" placeholder="your nickname" bind:value={nickname}>
                    </div>
                </div>
            </div>
            <button class="button is-warning" on:click={submitHighscore}>
                Post Highscore
            </button>
        </div>
    </div>
</Modal>

