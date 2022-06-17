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
    import type { ProjectData } from 'src/scripts/project';
</script>

<script lang="ts">
    import Modal from '../components/Modal.svelte';

    export let isActive: Writable<boolean>;
    export let project: Writable<ProjectData>;

    async function loadProject() {
        if (file == null) {
            alert('Please upload a stencil.json file');
        }
        const text = await file.text();
        return JSON.parse(text);
    }

    let files: FileList;
    $: file = files != null && files.length > 0 ? files.item(0) : null;
</script>

<Modal background {isActive}>
    <svelte:fragment slot="title">
        <span class="icon">
            <i class="fa-solid fa-file-import" />
        </span>
        Open Existing Project
    </svelte:fragment>
    <svelte:fragment slot="body">
        <div class="field">
            <label class="label" for="projectFile">
                Upload Project File (stencil.json)
            </label>
            <div id="projectFile" class="control">
                <div class="file has-name">
                    <label class="file-label">
                        <input
                            accept="application/json"
                            class="file-input"
                            type="file"
                            bind:files
                        />
                        <span class="file-cta">
                            <span class="file-icon">
                                <i class="fas fa-upload" />
                            </span>
                            <span class="file-label"> Choose a file… </span>
                        </span>
                        <span class="file-name">
                            {#if file != null}
                                {file.name}
                            {/if}
                        </span>
                    </label>
                </div>
            </div>
        </div>
    </svelte:fragment>
    <div slot="footer" class="buttons">
        <button
            class="button is-primary"
            on:click={() => {
                loadProject().then((p) => ($project = p));
            }}
        >
            Save changes
        </button>
        <button class="button is-light">Cancel</button>
    </div>
</Modal>
