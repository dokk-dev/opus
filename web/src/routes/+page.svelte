<script lang="ts">
	import DepartmentCard from '$lib/components/DepartmentCard.svelte';
	import AlertsList from '$lib/components/AlertsList.svelte';
	import Chat from '$lib/components/Chat.svelte';

	const departments = [
		{ id: 'dairy', name: 'Dairy', status: 'good', alerts: 0 },
		{ id: 'produce', name: 'Produce', status: 'warning', alerts: 2 },
		{ id: 'meat', name: 'Meat', status: 'good', alerts: 0 },
		{ id: 'bakery', name: 'Bakery', status: 'good', alerts: 1 },
		{ id: 'deli', name: 'Deli', status: 'good', alerts: 0 },
		{ id: 'grocery', name: 'Grocery', status: 'good', alerts: 0 },
		{ id: 'frontend', name: 'Front End', status: 'danger', alerts: 1 },
	];
</script>

<div class="dashboard">
	<header class="header">
		<div>
			<h1>Good morning, Manager</h1>
			<p class="subtitle">Here's your store overview for today</p>
		</div>
		<div class="header-stats">
			<div class="stat">
				<span class="stat-value">4</span>
				<span class="stat-label">Active Alerts</span>
			</div>
			<div class="stat">
				<span class="stat-value">12</span>
				<span class="stat-label">Staff On Duty</span>
			</div>
		</div>
	</header>

	<div class="grid">
		<section class="departments">
			<h2>Departments</h2>
			<div class="dept-grid">
				{#each departments as dept}
					<DepartmentCard {dept} />
				{/each}
			</div>
		</section>

		<aside class="sidebar-right">
			<AlertsList />
			<Chat />
		</aside>
	</div>
</div>

<style>
	.dashboard {
		max-width: 1400px;
		margin: 0 auto;
	}

	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 2rem;
	}

	h1 {
		font-size: 1.75rem;
		font-weight: 600;
	}

	.subtitle {
		color: var(--color-text-muted);
		margin-top: 0.25rem;
	}

	.header-stats {
		display: flex;
		gap: 2rem;
	}

	.stat {
		text-align: center;
	}

	.stat-value {
		display: block;
		font-size: 1.5rem;
		font-weight: 600;
	}

	.stat-label {
		font-size: 0.75rem;
		color: var(--color-text-muted);
	}

	.grid {
		display: grid;
		grid-template-columns: 1fr 380px;
		gap: 1.5rem;
	}

	.departments h2 {
		margin-bottom: 1rem;
		font-size: 1.125rem;
	}

	.dept-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(200px, 1fr));
		gap: 1rem;
	}

	.sidebar-right {
		display: flex;
		flex-direction: column;
		gap: 1.5rem;
	}

	@media (max-width: 1024px) {
		.grid {
			grid-template-columns: 1fr;
		}
	}
</style>
