<script>
    let expanded = false;
</script>

<nav class="navbar is-info" role="navigation">
    <div class="container">
        <div class="navbar-brand">
            <a class="navbar-item" href="/">
                Token Tournament
            </a>

            <a role="button" class="navbar-burger burger" aria-label="menu" aria-expanded="{expanded}" on:click={() => expanded = !expanded}>
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
            </a>
        </div>
        <div class="navbar-menu" class:is-active={expanded}>
            <div class="navbar-start">
                <a class="navbar-item">
                    Home
                </a>
                <a class="navbar-item">
                    Leaderboards
                </a>
            </div>
        </div>
    </div>
</nav>