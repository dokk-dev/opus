# Opus Architecture

## System Overview

Opus is a hybrid cloud/on-premise system designed to unify grocery store operations under a single AI-powered interface.

## Design Principles

1. **Local-first AI** - Use Ollama/Llama 3 for most queries to minimize costs and latency
2. **Security by default** - All connections encrypted, no secrets in client code
3. **Graceful degradation** - System works even when some integrations are unavailable
4. **Department isolation** - Each department has its own AI agent with specialized knowledge

## Components

### 1. Go Backend

The backend is a single Go binary that handles:

#### HTTP API (`internal/api/`)
- RESTful endpoints for CRUD operations
- JSON responses
- CORS handling
- Rate limiting (planned)

#### WebSocket Gateway (`internal/gateway/`)
- Real-time bidirectional communication
- Client session management
- Channel-based message routing
- Heartbeat/keepalive

#### AI Router (`internal/ai/`)
- Routes queries to appropriate department agent
- Manages conversation context
- Handles Ollama communication
- Falls back to Claude API when needed

### 2. SvelteKit Frontend

Single-page application providing:

- **Dashboard** - Store overview, department status, alerts
- **Chat Interface** - Natural language queries to Opus AI
- **Department Views** - Detailed inventory, schedules per department
- **Settings** - User preferences, notification config

### 3. AI System

```
┌──────────────────────────────────────────────┐
│                  AI Router                    │
│  - Analyzes query intent                      │
│  - Routes to department agent                 │
│  - Manages fallback logic                     │
└──────────────────┬───────────────────────────┘
                   │
    ┌──────────────┼──────────────┐
    ▼              ▼              ▼
┌────────┐   ┌──────────┐   ┌──────────┐
│ Dairy  │   │ Produce  │   │   ...    │
│ Agent  │   │  Agent   │   │  Agents  │
└────────┘   └──────────┘   └──────────┘
    │              │              │
    └──────────────┼──────────────┘
                   ▼
         ┌─────────────────┐
         │     Ollama      │
         │   (Llama 3)     │
         └────────┬────────┘
                  │ fallback
                  ▼
         ┌─────────────────┐
         │   Claude API    │
         │   (if needed)   │
         └─────────────────┘
```

### 4. Department Agents

Each department has a specialized agent with:

- **System prompt** - Department-specific context and knowledge
- **Tool access** - Can query inventory, schedules, alerts for its department
- **Conversation history** - Maintains context within a session

Departments:
- Dairy
- Produce
- Meat
- Bakery
- Deli
- Grocery
- Front End

### 5. Connectors (Planned)

Modular integrations for external systems:

```go
type Connector interface {
    Name() string
    Connect(ctx context.Context) error
    Disconnect() error
    Health() HealthStatus
}
```

Planned connectors:
- **POS Connector** - Register status, transaction data
- **Periscope Connector** - Inventory, production data
- **UKG Connector** - Employee schedules
- **Slack Adapter** - Message channel integration
- **Teams Adapter** - Message channel integration

## Data Flow

### Query Flow
```
User Input → Frontend → WebSocket → Gateway → AI Router
                                                  │
                                          Department Agent
                                                  │
                                              Ollama/Claude
                                                  │
                                          Response formatted
                                                  │
User Display ← Frontend ← WebSocket ← Gateway ←──┘
```

### Alert Flow (Planned)
```
System Monitor → Detects Issue → Creates Alert
                                       │
                                 Alert Engine
                                       │
              ┌────────────────────────┼────────────────────────┐
              ▼                        ▼                        ▼
        WebSocket Push           Slack Message            SMS (Twilio)
              │                        │                        │
         Dashboard                Slack App                  Phone
```

## Security Architecture

### Authentication (Planned)
- JWT tokens for API authentication
- Session management with secure cookies
- Role-based access control (Manager, Assistant Manager, Admin)

### Data Protection
- TLS 1.3 for all connections
- Secrets stored in environment variables
- No PII in logs
- Audit trail for sensitive operations

### Network Security
- On-premise bridge uses secure tunnel (Tailscale/WireGuard)
- API rate limiting
- Input validation at all boundaries

## Deployment

### Demo Mode
- Single machine deployment
- SQLite or in-memory data
- Mock integrations

### Production Mode
- Cloud backend (Fly.io recommended)
- PostgreSQL database
- On-premise bridge for local systems
- Multiple channel adapters

## Configuration

All configuration via environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| SERVER_ADDR | HTTP server address | :8080 |
| OLLAMA_URL | Ollama API endpoint | http://localhost:11434 |
| OLLAMA_MODEL | Default model | llama3 |
| CLAUDE_API_KEY | Claude API key (optional) | - |
| CLAUDE_FALLBACK | Enable Claude fallback | true |
| DATABASE_URL | PostgreSQL connection | - |
| JWT_SECRET | JWT signing key | - |
| CORS_ORIGINS | Allowed origins | http://localhost:5173 |
