<script lang="ts">
	interface Message {
		role: 'user' | 'assistant';
		content: string;
	}

	interface ChatResponse {
		response: string;
		department: string;
		model: string;
	}

	let messages: Message[] = $state([
		{
			role: 'assistant',
			content: "Hi! I'm Opus, your store assistant. Ask me about inventory, schedules, or any department questions.",
		},
	]);
	let input = $state('');
	let loading = $state(false);
	let aiStatus = $state<'online' | 'offline' | 'error'>('online');
	let currentModel = $state('Llama 3');
	let messagesContainer: HTMLDivElement;

	async function sendMessage() {
		if (!input.trim() || loading) return;

		const userMessage = input.trim();
		input = '';
		messages = [...messages, { role: 'user', content: userMessage }];
		loading = true;

		try {
			// Build conversation history (exclude system greeting)
			const history = messages.slice(1).map(m => ({
				role: m.role,
				content: m.content
			}));

			const response = await fetch('/api/v1/chat', {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({
					message: userMessage,
					history: history.slice(-10) // Keep last 10 messages for context
				})
			});

			if (!response.ok) {
				const error = await response.json().catch(() => ({ error: 'Unknown error' }));
				throw new Error(error.error || 'Failed to get response');
			}

			const data: ChatResponse = await response.json();
			messages = [...messages, { role: 'assistant', content: data.response }];
			currentModel = data.model || 'Llama 3';
			aiStatus = 'online';
		} catch (error) {
			aiStatus = 'error';
			const errorMessage = error instanceof Error ? error.message : 'Unknown error';
			messages = [
				...messages,
				{
					role: 'assistant',
					content: `Sorry, I encountered an error: ${errorMessage}\n\nMake sure Ollama is running with: ollama serve`
				},
			];
		} finally {
			loading = false;
		}
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
			<span class="status" class:status-error={aiStatus === 'error'}>
				{aiStatus === 'error' ? 'Connection error' : 'Online'} · {currentModel}
			</span>
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

	.status-error {
		color: var(--color-danger);
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
