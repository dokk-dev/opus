<script lang="ts">
	import { page } from '$app/stores';

	const navItems = [
		{ href: '/', label: 'Dashboard', icon: 'home' },
		{ href: '/inventory', label: 'Inventory', icon: 'package' },
		{ href: '/schedule', label: 'Schedule', icon: 'calendar' },
		{ href: '/alerts', label: 'Alerts', icon: 'bell' },
		{ href: '/reports', label: 'Reports', icon: 'chart' },
		{ href: '/settings', label: 'Settings', icon: 'settings' },
	];
</script>

<nav class="sidebar">
	<div class="logo">
		<span class="logo-icon">O</span>
		<span class="logo-text">Opus</span>
	</div>

	<ul class="nav">
		{#each navItems as item}
			<li>
				<a
					href={item.href}
					class:active={$page.url.pathname === item.href}
				>
					<span class="icon">{item.icon.charAt(0).toUpperCase()}</span>
					{item.label}
				</a>
			</li>
		{/each}
	</ul>

	<div class="sidebar-footer">
		<div class="user">
			<div class="avatar">DM</div>
			<div class="user-info">
				<span class="user-name">Demo Manager</span>
				<span class="user-role">Store Manager</span>
			</div>
		</div>
	</div>
</nav>

<style>
	.sidebar {
		width: 240px;
		background: var(--color-bg-secondary);
		border-right: 1px solid var(--color-border);
		display: flex;
		flex-direction: column;
		padding: 1rem;
	}

	.logo {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		padding: 0.5rem;
		margin-bottom: 2rem;
	}

	.logo-icon {
		width: 36px;
		height: 36px;
		background: var(--color-primary);
		border-radius: 8px;
		display: flex;
		align-items: center;
		justify-content: center;
		font-weight: 700;
		font-size: 1.25rem;
	}

	.logo-text {
		font-size: 1.25rem;
		font-weight: 600;
	}

	.nav {
		list-style: none;
		flex: 1;
	}

	.nav a {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		padding: 0.75rem;
		border-radius: var(--radius);
		color: var(--color-text-muted);
		text-decoration: none;
		transition: all 0.2s;
	}

	.nav a:hover {
		background: var(--color-bg-tertiary);
		color: var(--color-text);
	}

	.nav a.active {
		background: var(--color-primary);
		color: white;
	}

	.icon {
		width: 20px;
		height: 20px;
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 0.75rem;
		font-weight: 600;
		background: var(--color-bg-tertiary);
		border-radius: 4px;
	}

	.active .icon {
		background: rgba(255, 255, 255, 0.2);
	}

	.sidebar-footer {
		border-top: 1px solid var(--color-border);
		padding-top: 1rem;
	}

	.user {
		display: flex;
		align-items: center;
		gap: 0.75rem;
	}

	.avatar {
		width: 36px;
		height: 36px;
		background: var(--color-bg-tertiary);
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 0.75rem;
		font-weight: 600;
	}

	.user-info {
		display: flex;
		flex-direction: column;
	}

	.user-name {
		font-size: 0.875rem;
		font-weight: 500;
	}

	.user-role {
		font-size: 0.75rem;
		color: var(--color-text-muted);
	}
</style>
