<script lang="ts">
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
    import { fly } from 'svelte/transition';
    import ComponentPanel from '../lib/ComponentPanel.svelte';
    import CodeView from './CodeView.svelte';
    import DiagramView from './DiagramView.svelte';

    enum PreviewTab {
        CODE,
        DIAGRAM,
    }

    let tab = PreviewTab.CODE;
</script>

<div class="columns" in:fly={{ delay: 400, x: 100 }}>
    <div class="column is-one-third">
        <ComponentPanel />
    </div>
    <div class="column">
        <div class="tabs is-centered is-toggle">
            <ul>
                <li class:is-active={tab == PreviewTab.CODE}>
                    <!-- svelte-ignore a11y-missing-attribute -->
                    <a on:click={() => (tab = PreviewTab.CODE)}>
                        <span class="icon is-small">
                            <i class="fa-solid fa-code" />
                        </span>
                        <span>Code</span>
                    </a>
                </li>
                <li class:is-active={tab == PreviewTab.DIAGRAM}>
                    <!-- svelte-ignore a11y-missing-attribute -->
                    <a on:click={() => (tab = PreviewTab.DIAGRAM)}>
                        <span class="icon is-small">
                            <i class="fa-solid fa-diagram-project" />
                        </span>
                        <span>Diagram</span>
                    </a>
                </li>
            </ul>
        </div>
        {#if tab == PreviewTab.CODE}
            <CodeView source={`print("hello, world!")`} />
        {:else}
            <DiagramView />
        {/if}
    </div>
</div>

<style>
    .columns {
        height: 100%;
    }
</style>
