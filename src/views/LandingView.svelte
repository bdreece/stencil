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
    import type { ProjectData } from '../scripts/project';
    import View from '../scripts/view';

    export interface LandingViewProps {
        activeView: Writable<View>;
        project: Writable<ProjectData>;
    }
</script>

<script lang="ts">
    import { onMount } from 'svelte';
    import { writable } from 'svelte/store';
    import { fade, fly } from 'svelte/transition';

    import CreateProjectModal from '../lib/CreateProjectModal.svelte';
    import OpenProjectModal from '../lib/OpenProjectModal.svelte';

    export let activeView: Writable<View>;
    export let project: Writable<ProjectData>;

    project.subscribe((p) => {
        if (!(p.name == '' || p.language == '')) {
            $activeView = View.PROJECT;
        }
    });

    let condition = false;
    onMount(() => (condition = true));

    const duration = 500;
    const createProject = {
        isActive: writable(false),
        project,
    };
    const openProject = {
        isActive: writable(false),
        project,
    };
</script>

{#if condition}
    <div class="container is-fluid" out:fly={{ x: -100 }}>
        <div class="box">
            <section class="hero">
                <div class="hero-body">
                    <p class="title" in:fade={{ duration }}>
                        Welcome to Stencil!
                    </p>
                    <p class="subtitle" in:fade={{ delay: 200, duration }}>
                        A boilerplate code generator for bootstrapping new
                        projects.
                    </p>
                </div>
            </section>
            <div
                class="buttons is-justify-content-center"
                in:fade={{ delay: 400, duration }}
            >
                <button
                    class="button is-child is-link"
                    on:click={() => createProject.isActive.set(true)}
                >
                    <span class="icon">
                        <i class="fa-solid fa-file-circle-plus" />
                    </span>
                    <p>Create New Project</p>
                </button>
                <button
                    class="button is-child is-link is-light"
                    on:click={() => openProject.isActive.set(true)}
                >
                    <span class="icon">
                        <i class="fa-solid fa-file-import" />
                    </span>
                    <p>Open Existing Project</p>
                </button>
            </div>
        </div>
        <CreateProjectModal {...createProject} />
        <OpenProjectModal {...openProject} />
    </div>
{/if}
