# Opus

AI-powered grocery retail operations platform - unified intelligence for store management.

## Overview

Opus brings together your store's disparate systems (POS, inventory, scheduling) into a unified interface that department managers can access through chat, web dashboard, or messaging platforms.

## Features

- **Unified Dashboard** - Real-time view of all departments
- **AI Assistant** - Natural language queries for inventory, schedules, and operations
- **Multi-Channel** - Access via web, Slack, Teams, SMS (coming soon)
- **Department Agents** - Specialized AI for each department's needs
- **Real-time Alerts** - Proactive notifications for issues
- **Local-first AI** - Powered by Llama 3 with Claude fallback

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                         CLOUD                               │
│  ┌───────────────────────────────────────────────────────┐  │
│  │  Go Backend (API + WebSocket Gateway)                 │  │
│  │  - REST API for data operations                       │  │
│  │  - WebSocket for real-time updates                    │  │
│  │  - AI routing to department agents                    │  │
│  └───────────────────────────────────────────────────────┘  │
│  ┌───────────────────────────────────────────────────────┐  │
│  │  SvelteKit Frontend (Dashboard)                       │  │
│  └───────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
                              │
                    Secure Tunnel (Future)
                              │
┌─────────────────────────────────────────────────────────────┐
│                      ON-PREMISE                             │
│  ┌─────────┐  ┌──────────┐  ┌─────────┐  ┌───────────────┐  │
│  │   POS   │  │ Periscope│  │   UKG   │  │ Store Servers │  │
│  └─────────┘  └──────────┘  └─────────┘  └───────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

## Tech Stack

| Component | Technology |
|-----------|------------|
| Backend | Go 1.26 |
| Frontend | SvelteKit 2 + TypeScript |
| AI (Local) | Ollama (Llama 3) |
| AI (Fallback) | Claude API |
| Database | PostgreSQL (planned) |
| Real-time | WebSockets |

## Getting Started

### Prerequisites

- Go 1.26+
- Node.js 22+
- Ollama with Llama 3 model

### Development

1. Clone the repository:
```bash
git clone https://github.com/dokk-dev/opus.git
cd opus
```

2. Start Ollama:
```bash
ollama serve
```

3. Start the backend:
```bash
cp .env.example .env
go run cmd/server/main.go
```

4. Start the frontend:
```bash
cd web
npm install
npm run dev
```

5. Open http://localhost:5173

## Project Structure

```
opus/
├── cmd/
│   └── server/          # Main application entry
├── internal/
│   ├── api/             # HTTP handlers
│   ├── gateway/         # WebSocket server
│   ├── ai/              # AI integration (Ollama, Claude)
│   ├── connectors/      # External system integrations
│   ├── models/          # Data models
│   └── config/          # Configuration
├── pkg/
│   └── utils/           # Shared utilities
└── web/                 # SvelteKit frontend
    └── src/
        ├── routes/      # Pages
        └── lib/         # Components and utilities
```

## Roadmap

### Phase 1: Demo (Current)
- [x] Project structure
- [x] Basic dashboard UI
- [x] Local AI integration (Ollama)
- [ ] Demo data for all departments
- [ ] Chat functionality

### Phase 2: Core Features
- [ ] Database integration
- [ ] User authentication
- [ ] Real inventory/schedule APIs
- [ ] Alert system

### Phase 3: Integrations
- [ ] Slack channel adapter
- [ ] Teams channel adapter
- [ ] On-premise bridge for POS/Periscope

### Phase 4: Production
- [ ] Audit logging
- [ ] Role-based access control
- [ ] Multi-store support

## License

Private - All rights reserved
