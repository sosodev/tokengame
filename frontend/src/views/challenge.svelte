<script>
    export let currentRoute;

    import Token from '../components/token.svelte';
    import Modal from '../components/modal.svelte';

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

    $: title = challengeData.title;
    $: head = challengeData.head;
    $: foot = challengeData.foot;
    $: tokens = challengeData.tokens;
    $: testcases = challengeData.testcases;
    $: code = hljs.highlight('javascript', beautify(head + user_code + foot)).value;

    let finished = false;
    let user_code = "";
    let token_page = 0;
    let budget = 200;
    let move_stack = [];

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

    // Run the testcases against the user's code
    function runTests() {
        let full_code = head + user_code + foot;
        eval.call(window, full_code);

        let passed = true;

        // run the test cases and fail if the return is false
        testcases.forEach(test => {
            if (!eval.call(window, test)) {
                passed = false;
                return;
            }
        });

        finished = passed;
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

    nav {
        padding-top: 1%;
        padding-bottom: 1%;
    }

    input {
        margin-bottom: 1rem;
    }

    code {
        font-size: .9rem;
    }
</style>

<div class="container" style="padding: 1rem;">
    <h1 class="title">
        Challenge #{currentRoute.namedParams.id} - {title}
    </h1>
    <h3 class="subtitle">
        Budget: {budget}
    </h3>
    <hr/>
<pre><code class="hljs javascript">{@html code}</code></pre>
    <nav class="level has-background-grey-lighter">
        <div class="level-item">
            <div class="button" on:click={backward}>
                &lt;-
            </div>
        </div>
        {#each tokens.slice(token_page * 5, (token_page + 1) * 5) as token}
            <div class="level-item">
                <Token cost="{calculateCost(token)}" on:click={() => useToken(token)}>
                    {token}
                </Token>
            </div>
        {/each}
        <div class="level-item">
            <div class="button is-warning" on:click={forward}>
                -&gt;
            </div>
        </div>
    </nav>
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
            <h2 style="font-size: 1.75rem; font-weight: bold; white-space: nowrap;">
                You finished the {title} challenge! 🎉
            </h2>
            <h3 class="subtitle">
                Final Score: 3962
            </h3>
            <div class="field-body">
                <div class="field">
                    <div class="control is-expanded">
                        <input class="input" type="text" placeholder="your nickname">
                    </div>
                </div>
            </div>
            <button class="button is-warning">
                Post Highscore
            </button>
        </div>
    </div>
</Modal>

