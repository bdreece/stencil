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

    export interface ModalProps {
        isActive: Writable<boolean>;
        background: boolean;
    }
</script>

<script lang="ts">
    import { slide } from 'svelte/transition';
    import { writable } from 'svelte/store';

    function closeModal() {
        $isActive = false;
    }

    export let isActive = writable(false);
    export let background = true;
</script>

<div class="modal" class:is-active={$isActive}>
    {#if background}
        <div class="modal-background" on:click={closeModal} />
    {/if}
    {#if $isActive}
        <div class="modal-card" transition:slide>
            <header class="modal-card-head">
                <p class="modal-card-title">
                    <slot name="title">Modal Title</slot>
                </p>
                <button
                    class="delete"
                    aria-label="close"
                    on:click={closeModal}
                />
            </header>
            <section class="modal-card-body">
                <slot name="body">
                    Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed
                    do eiusmod tempor incididunt ut labore et dolore magna
                    aliqua. Ut enim ad minim veniam, quis nostrud exercitation
                    ullamco laboris nisi ut aliquip ex ea commodo consequat.
                    Duis aute irure dolor in reprehenderit in voluptate velit
                    esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
                    occaecat cupidatat non proident, sunt in culpa qui officia
                    deserunt mollit anim id est laborum.
                </slot>
            </section>
            <footer class="modal-card-foot">
                <slot name="footer">
                    <button class="button is-success" on:click={closeModal}>
                        Save changes
                    </button>
                    <button class="button" on:click={closeModal}>
                        Cancel
                    </button>
                </slot>
            </footer>
        </div>
    {/if}
</div>

<svelte:window
    on:keydown={(event) => {
        if (event.code == '27') {
            closeModal();
        }
    }}
/>
