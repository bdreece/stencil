<script lang="ts">
    import { onMount } from 'svelte';

    /* 
     * stencil - A boilerplate code generator for bootstrapping new projects.
     * Copyright (C) 2022 Brian Reece

     * This program is free software: you can redistribute it and/or modify
     * it under the terms of the GNU Affero General Public License as published
     * by the Free Software Foundation, either version 3 of the License, or
     * (at your option) any later version.

     * This program is distributed in the hope that it will be useful,
     * but WITHOUT ANY WARRANTY; without even the implied warranty of
     * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
     * GNU Affero General Public License for more details.

     * You should have received a copy of the GNU Affero General Public License
     * along with this program.  If not, see <https://www.gnu.org/licenses/>.
     */

    import { fade } from 'svelte/transition';
    let condition = false;
    let duration = 500;

    onMount(() => (condition = true));
</script>

{#if condition}
    <div class="container is-fluid">
        <h2
            class="is-size-2 has-text-centered my-4"
            transition:fade={{ delay: 250, duration }}
        >
            Welcome to Stencil!
        </h2>
        <p
            class="has-text-centered mb-4"
            transition:fade={{ delay: 500, duration }}
        >
            A boilerplate code generator for bootstrapping new projects.
        </p>
        <div
            class="buttons is-justify-content-center"
            transition:fade={{ delay: 750, duration }}
        >
            <button class="button is-link">Create New Project</button>
            <button class="button is-link">Open Existing Project</button>
        </div>
    </div>
{/if}
