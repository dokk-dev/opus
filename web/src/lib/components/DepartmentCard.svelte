<script lang="ts">
	interface Department {
		id: string;
		name: string;
		status: 'good' | 'warning' | 'danger';
		alerts: number;
	}

	let { dept }: { dept: Department } = $props();

	const statusColors = {
		good: 'var(--color-success)',
		warning: 'var(--color-warning)',
		danger: 'var(--color-danger)',
	};
</script>

<a href="/departments/{dept.id}" class="card">
	<div class="card-header">
		<span class="dept-name">{dept.name}</span>
		<span class="status-dot" style="background: {statusColors[dept.status]}"></span>
	</div>
	<div class="card-body">
		{#if dept.alerts > 0}
			<span class="badge badge-{dept.status === 'danger' ? 'danger' : 'warning'}">
				{dept.alerts} alert{dept.alerts > 1 ? 's' : ''}
			</span>
		{:else}
			<span class="badge badge-success">All clear</span>
		{/if}
	</div>
</a>

<style>
	.card {
		display: block;
		text-decoration: none;
		color: inherit;
		transition: transform 0.2s, box-shadow 0.2s;
	}

	.card:hover {
		transform: translateY(-2px);
		box-shadow: 0 8px 16px -4px rgb(0 0 0 / 0.2);
	}

	.card-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 0.75rem;
	}

	.dept-name {
		font-weight: 600;
	}

	.status-dot {
		width: 10px;
		height: 10px;
		border-radius: 50%;
	}

	.card-body {
		display: flex;
		align-items: center;
	}
</style>
