<script lang="ts">
    import type { Writable } from 'svelte/store';
    import { writable } from 'svelte/store';
    import Modal from '../components/Modal.svelte';

    export let isActive = writable(false);

    let files: FileList;
    $: file = files != null && files.length > 0 ? files.item(0) : null;
</script>

<Modal background {isActive}>
    <svelte:fragment slot="title">Open Existing Project</svelte:fragment>
    <svelte:fragment slot="body">
        <div class="field">
            <label class="label" for="projectFile">
                Upload Project File (stencil.json)
            </label>
            <div id="projectFile" class="control">
                <div class="file has-name is-boxed">
                    <label class="file-label">
                        <input class="file-input" type="file" bind:files />
                        <span class="file-cta">
                            <span class="file-icon">
                                <i class="fas fa-upload" />
                            </span>
                            <span class="file-label"> Choose a file… </span>
                        </span>
                        <span class="file-name">
                            {#if file}
                                {file.name}
                            {/if}
                        </span>
                    </label>
                </div>
            </div>
        </div>
    </svelte:fragment>
    <div slot="footer" class="buttons">
        <button class="button is-primary">Save changes</button>
        <button class="button is-light">Cancel</button>
    </div>
</Modal>
