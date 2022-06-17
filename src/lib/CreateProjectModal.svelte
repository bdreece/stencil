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

    import type { ProjectData } from '../scripts/project';
</script>

<script lang="ts">
    import { writable } from 'svelte/store';
    import Modal from '../components/Modal.svelte';
    import languages from '../assets/languages.json';

    let _project: ProjectData = {
        name: '',
        language: '',
    };

    export let isActive = writable(false);
    export let project = writable(_project);
</script>

<Modal background {isActive}>
    <svelte:fragment slot="title">
        <span class="icon">
            <i class="fa-solid fa-file-circle-plus" />
        </span>
        Create New Project
    </svelte:fragment>
    <svelte:fragment slot="body">
        <div class="field">
            <label class="label" for="name">Name</label>
            <div id="name" class="control">
                <input
                    class="input is-rounded"
                    type="text"
                    bind:value={_project.name}
                />
            </div>
        </div>
        <div class="field">
            <label class="label" for="language">Language</label>
            <div id="language" class="control">
                <div class="select is-rounded">
                    <select bind:value={_project.language}>
                        {#each languages as lang}
                            <option value={lang.id}>{lang.name}</option>
                        {/each}
                    </select>
                </div>
            </div>
        </div>
        <div class="field">
            <label class="label" for="brief">Brief</label>
            <div id="brief" class="control">
                <input
                    class="input is-rounded"
                    type="text"
                    bind:value={_project.about}
                />
            </div>
        </div>
        <div class="field">
            <label class="label" for="description">Description</label>
            <div id="description" class="control">
                <textarea
                    class="textarea is-rounded"
                    bind:value={_project.longAbout}
                />
            </div>
        </div>
    </svelte:fragment>
    <div slot="footer" class="buttons">
        <button
            class="button is-primary"
            on:click={() => {
                $project = _project;
                $isActive = false;
            }}
        >
            Save changes
        </button>
        <button class="button is-light" on:click={() => ($isActive = false)}>
            Cancel
        </button>
    </div>
</Modal>
