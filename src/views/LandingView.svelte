<script lang="ts" context="module">
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

    import type { Writable } from 'svelte/store';
    import type { ProjectObject } from '../scripts/project';
    import View from '../scripts/view';

    export interface LandingViewProps {
        activeView: Writable<View>;
        project: Writable<ProjectObject>;
    }
</script>

<script lang="ts">
    import { onMount } from 'svelte';
    import { writable } from 'svelte/store';
    import { fade } from 'svelte/transition';

    import CreateProjectModal from '../lib/CreateProjectModal.svelte';
    import OpenProjectModal from '../lib/OpenProjectModal.svelte';

    export let activeView: Writable<View>;
    export let project: Writable<ProjectObject>;

    project.subscribe((p) => {
        if (!(p.name == '' || p.language == '')) {
            $activeView = View.PROJECT;
        }
    });

    let condition = false;
    onMount(() => (condition = true));

    const duration = 500;
    const createProject = writable(false);
    const openProject = writable(false);
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
            <button
                class="button is-link"
                on:click={() => ($createProject = true)}
            >
                Create New Project
            </button>
            <button
                class="button is-link"
                on:click={() => ($openProject = true)}
            >
                Open Existing Project
            </button>
        </div>
        <CreateProjectModal {project} isActive={createProject} />
        <OpenProjectModal isActive={openProject} />
    </div>
{/if}
