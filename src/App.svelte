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

    import { writable } from 'svelte/store';

    import Navbar from './lib/Navbar.svelte';
    import LandingView, { LandingViewProps } from './views/LandingView.svelte';
    import ProjectView from './views/ProjectView.svelte';
    import View from './scripts/view';

    const activeView = writable(View.LANDING);
    const project = writable({
        name: '',
        language: '',
    });

    const landingView: LandingViewProps = {
        activeView,
        project,
    };
</script>

<div id="app">
    <Navbar />
    {#if $activeView == View.LANDING}
        <LandingView {...landingView} />
    {:else if $activeView == View.PROJECT}
        <ProjectView />
    {/if}
</div>

<style>
    #app {
        position: absolute;
        width: 100%;
        height: 100%;
    }
</style>
