<script lang="ts">
	import { onMount } from 'svelte';

	interface Message {
		role: 'user' | 'assistant';
		content: string;
	}

	let messages: Message[] = $state([
		{
			role: 'assistant',
			content: "Hi! I'm Opus, your store assistant. Ask me about inventory, schedules, or any department questions.",
		},
	]);
	let input = $state('');
	let loading = $state(false);
	let messagesContainer: HTMLDivElement;

	async function sendMessage() {
		if (!input.trim() || loading) return;

		const userMessage = input.trim();
		input = '';
		messages = [...messages, { role: 'user', content: userMessage }];
		loading = true;

		try {
			// For demo, simulate AI response
			// In production, this calls the backend API
			await new Promise((resolve) => setTimeout(resolve, 1000));

			const response = getDemoResponse(userMessage);
			messages = [...messages, { role: 'assistant', content: response }];
		} catch (error) {
			messages = [
				...messages,
				{ role: 'assistant', content: 'Sorry, I encountered an error. Please try again.' },
			];
		} finally {
			loading = false;
		}
	}

	function getDemoResponse(query: string): string {
		const q = query.toLowerCase();

		if (q.includes('milk') || q.includes('dairy')) {
			return "Current 2% milk inventory: 48 units. We have 24 units expiring in 2 days. Recommended action: Consider a markdown or feature display to move expiring product.";
		}

		if (q.includes('schedule') || q.includes('who')) {
			return "Today's closing shift:\n- Dairy: Sarah M.\n- Produce: Mike T.\n- Front End: Jennifer L.\n\nWould you like me to show tomorrow's schedule?";
		}

		if (q.includes('register') || q.includes('pos')) {
			return "Register 3 has been offline for 15 minutes. Error code: NET_TIMEOUT. This typically indicates a network connectivity issue. IT has been notified. Registers 1, 2, and 4 are operational.";
		}

		if (q.includes('order') || q.includes('reorder')) {
			return "I can help with ordering. Which department or product are you looking to reorder? You can say something like 'reorder organic bananas' or 'show dairy reorder suggestions'.";
		}

		return "I can help you with:\n- Inventory levels and reorders\n- Staff schedules\n- System status and alerts\n- Department performance\n\nWhat would you like to know?";
	}

	function handleKeydown(event: KeyboardEvent) {
		if (event.key === 'Enter' && !event.shiftKey) {
			event.preventDefault();
			sendMessage();
		}
	}

	$effect(() => {
		if (messagesContainer) {
			messagesContainer.scrollTop = messagesContainer.scrollHeight;
		}
	});
</script>

<div class="card chat">
	<div class="chat-header">
		<div class="chat-avatar">O</div>
		<div>
			<h3>Opus Assistant</h3>
			<span class="status">Online · Using Llama 3</span>
		</div>
	</div>

	<div class="messages" bind:this={messagesContainer}>
		{#each messages as message}
			<div class="message message-{message.role}">
				<div class="message-content">
					{message.content}
				</div>
			</div>
		{/each}
		{#if loading}
			<div class="message message-assistant">
				<div class="message-content loading">
					<span class="dot"></span>
					<span class="dot"></span>
					<span class="dot"></span>
				</div>
			</div>
		{/if}
	</div>

	<div class="chat-input">
		<input
			type="text"
			placeholder="Ask about inventory, schedules..."
			bind:value={input}
			onkeydown={handleKeydown}
			disabled={loading}
		/>
		<button onclick={sendMessage} disabled={loading || !input.trim()}>Send</button>
	</div>
</div>

<style>
	.chat {
		display: flex;
		flex-direction: column;
		height: 400px;
	}

	.chat-header {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		padding-bottom: 1rem;
		border-bottom: 1px solid var(--color-border);
		margin-bottom: 1rem;
	}

	.chat-avatar {
		width: 40px;
		height: 40px;
		background: var(--color-primary);
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		font-weight: 700;
	}

	.chat-header h3 {
		font-size: 0.875rem;
		font-weight: 600;
	}

	.status {
		font-size: 0.75rem;
		color: var(--color-success);
	}

	.messages {
		flex: 1;
		overflow-y: auto;
		display: flex;
		flex-direction: column;
		gap: 0.75rem;
	}

	.message {
		max-width: 85%;
	}

	.message-user {
		align-self: flex-end;
	}

	.message-assistant {
		align-self: flex-start;
	}

	.message-content {
		padding: 0.75rem;
		border-radius: var(--radius);
		font-size: 0.875rem;
		line-height: 1.4;
		white-space: pre-wrap;
	}

	.message-user .message-content {
		background: var(--color-primary);
	}

	.message-assistant .message-content {
		background: var(--color-bg-tertiary);
	}

	.loading {
		display: flex;
		gap: 4px;
		padding: 0.75rem 1rem;
	}

	.dot {
		width: 8px;
		height: 8px;
		background: var(--color-text-muted);
		border-radius: 50%;
		animation: bounce 1.4s infinite ease-in-out;
	}

	.dot:nth-child(1) { animation-delay: -0.32s; }
	.dot:nth-child(2) { animation-delay: -0.16s; }

	@keyframes bounce {
		0%, 80%, 100% { transform: scale(0.8); opacity: 0.5; }
		40% { transform: scale(1); opacity: 1; }
	}

	.chat-input {
		display: flex;
		gap: 0.5rem;
		margin-top: 1rem;
		padding-top: 1rem;
		border-top: 1px solid var(--color-border);
	}

	.chat-input input {
		flex: 1;
	}

	.chat-input button {
		flex-shrink: 0;
	}
</style>
