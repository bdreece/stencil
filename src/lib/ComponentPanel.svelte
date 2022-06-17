<script lang="ts">
    import { validate_each_argument } from 'svelte/internal';

    let components = [];
</script>

<nav class="panel">
    <p class="panel-heading">
        <span class="icon">
            <i class="fa-solid fa-cubes" />
        </span>
        Components
    </p>
    <div class="panel-block">
        <p class="control has-icons-left">
            <input class="input" type="text" placeholder="Search" />
            <span class="icon is-left">
                <i class="fas fa-search" aria-hidden="true" />
            </span>
        </p>
    </div>
    {#each components as component}
        <button class="panel-block is-active">
            <span class="panel-icon">
                <i class="fa-solid fa-cube" />
            </span>
            {component.value.name}
        </button>
    {/each}
</nav>

<style>
    .panel {
        height: 100%;
    }
</style>
