<script lang="ts">
    import { writable } from 'svelte/store';
    import Modal from '../components/Modal.svelte';
    import type { ProjectObject } from '../scripts/project';
    import languages from '../assets/languages.json';

    let _project: ProjectObject = {
        name: '',
        language: '',
    };

    export let isActive = writable(false);
    export let project = writable(_project);
</script>

<Modal background {isActive}>
    <svelte:fragment slot="title">Create New Project</svelte:fragment>
    <svelte:fragment slot="body">
        <div class="field">
            <label class="label" for="name">Project Name</label>
            <div id="name" class="control">
                <input class="input" type="text" bind:value={_project.name} />
            </div>
        </div>
        <div class="field">
            <label class="label" for="brief">Project Brief</label>
            <div id="brief" class="control">
                <input class="input" type="text" bind:value={_project.about} />
            </div>
        </div>
        <div class="field">
            <label class="label" for="description">Project Description</label>
            <div id="description" class="control">
                <textarea class="textarea" bind:value={_project.longAbout} />
            </div>
        </div>
        <div class="field">
            <label class="label" for="language">Project Language</label>
            <div id="language" class="control">
                <div class="select">
                    <select bind:value={_project.language}>
                        {#each languages as lang}
                            <option value={lang.id}>{lang.name}</option>
                        {/each}
                    </select>
                </div>
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
