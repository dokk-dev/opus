<script lang="ts">
	import '../app.css';
	import Sidebar from '$lib/components/Sidebar.svelte';

	let { children } = $props();
</script>

<div class="app">
	<Sidebar />
	<main class="main">
		{@render children()}
	</main>
</div>

<style>
	.app {
		display: flex;
		height: 100vh;
	}

	.main {
		flex: 1;
		overflow-y: auto;
		padding: 1.5rem;
	}
</style>
