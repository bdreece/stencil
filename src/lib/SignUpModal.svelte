<script lang="ts">
    import { writable } from 'svelte/store';
    import Modal from '../components/Modal.svelte';

    export let isActive = writable(false);
</script>

<Modal background {isActive}>
    <div slot="title">Sign Up for Free!</div>
    <div slot="body">Sign up form goes here</div>
</Modal>
