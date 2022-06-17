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
</script>

<script lang="ts">
    import View from '../scripts/view';

    export let activeView: Writable<View>;
</script>

<div>
    <nav class="navbar is-dark" aria-label="main navigation">
        <div class="navbar-brand">
            <a class="navbar-item is-size-1 has-text-weight-bold px-4" href="/">
                <slot name="brand">Stencil</slot>
            </a>

            <!-- svelte-ignore a11y-missing-attribute -->
            <a
                role="button"
                class="navbar-burger"
                aria-label="menu"
                aria-expanded="false"
                data-target="navbarBasicExample"
            >
                <span aria-hidden="true" />
                <span aria-hidden="true" />
                <span aria-hidden="true" />
            </a>
        </div>

        <div class="navbar-menu">
            <div class="navbar-end">
                <div class="navbar-item">
                    <div class="buttons">
                        <button
                            class="button is-link is-light"
                            on:click={() => ($activeView = View.ABOUT)}
                        >
                            About
                        </button>
                        <a
                            class="button is-link"
                            href="https://github.com/bdreece/stencil"
                        >
                            GitHub
                        </a>
                    </div>
                </div>
            </div>
        </div>
    </nav>
</div>
