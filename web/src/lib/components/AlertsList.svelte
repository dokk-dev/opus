<script lang="ts">
	// Demo alerts data
	const alerts = [
		{
			id: 1,
			type: 'danger',
			title: 'Register 3 offline',
			department: 'Front End',
			time: '5 min ago',
		},
		{
			id: 2,
			type: 'warning',
			title: 'Low stock: Organic Bananas',
			department: 'Produce',
			time: '12 min ago',
		},
		{
			id: 3,
			type: 'warning',
			title: 'Expiring: 2% Milk (24 units)',
			department: 'Dairy',
			time: '1 hr ago',
		},
		{
			id: 4,
			type: 'info',
			title: 'Schedule change request',
			department: 'Bakery',
			time: '2 hr ago',
		},
	];
</script>

<div class="card alerts">
	<div class="alerts-header">
		<h3>Recent Alerts</h3>
		<a href="/alerts">View all</a>
	</div>

	<ul class="alerts-list">
		{#each alerts as alert}
			<li class="alert-item alert-{alert.type}">
				<div class="alert-indicator"></div>
				<div class="alert-content">
					<span class="alert-title">{alert.title}</span>
					<span class="alert-meta">{alert.department} · {alert.time}</span>
				</div>
			</li>
		{/each}
	</ul>
</div>

<style>
	.alerts-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 1rem;
	}

	.alerts-header h3 {
		font-size: 1rem;
		font-weight: 600;
	}

	.alerts-header a {
		font-size: 0.75rem;
		color: var(--color-primary);
		text-decoration: none;
	}

	.alerts-list {
		list-style: none;
	}

	.alert-item {
		display: flex;
		gap: 0.75rem;
		padding: 0.75rem 0;
		border-bottom: 1px solid var(--color-border);
	}

	.alert-item:last-child {
		border-bottom: none;
		padding-bottom: 0;
	}

	.alert-indicator {
		width: 8px;
		height: 8px;
		border-radius: 50%;
		margin-top: 0.25rem;
		flex-shrink: 0;
	}

	.alert-danger .alert-indicator {
		background: var(--color-danger);
	}

	.alert-warning .alert-indicator {
		background: var(--color-warning);
	}

	.alert-info .alert-indicator {
		background: var(--color-primary);
	}

	.alert-content {
		display: flex;
		flex-direction: column;
	}

	.alert-title {
		font-size: 0.875rem;
		font-weight: 500;
	}

	.alert-meta {
		font-size: 0.75rem;
		color: var(--color-text-muted);
		margin-top: 0.125rem;
	}
</style>
