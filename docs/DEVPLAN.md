# Opus Development Plan

> **Last Updated**: 2026-02-22
> **Current Phase**: Demo Build
> **Status**: In Progress

---

## Quick Resume

**To pick up where we left off, tell Claude:**
> "Let's continue working on Opus. Read the DEVPLAN.md and pick up from the current state."

---

## Project Overview

**Opus** is an AI-powered grocery retail operations platform that unifies store systems (POS, inventory, scheduling) into a single interface accessible via chat, web dashboard, or messaging platforms.

### Tech Stack
| Component | Technology |
|-----------|------------|
| Backend | Go 1.26 |
| Frontend | SvelteKit 5 + TypeScript |
| AI (Local) | Ollama (Llama 3) |
| AI (Fallback) | Claude API (planned) |
| Database | PostgreSQL (planned) |

### Target Users
- Store Managers
- Department Managers
- Assistant Department Managers

### Departments
Dairy, Produce, Meat, Bakery, Deli, Grocery, Front End

---

## What Has Been Done

### Phase 0: Project Setup ✅
- [x] Created GitHub repository (dokk-dev/opus)
- [x] Set up Go backend structure
- [x] Set up SvelteKit frontend structure
- [x] Created architecture documentation
- [x] Configured development environment (Go, Node, Ollama)

### Phase 1: Core AI Integration ✅
- [x] Implemented Ollama client for local AI
- [x] Created department-specific AI agents with specialized prompts
- [x] Built AI router for query routing based on keywords
- [x] Connected frontend chat to backend API
- [x] Added error handling and status display in chat UI

---

## Current State

### Working Features
1. **Dashboard UI** - Department cards, alerts list, chat interface
2. **Chat with AI** - Real queries to Llama 3 via Ollama
3. **Department Routing** - Queries routed to specialized agents
4. **WebSocket Gateway** - Infrastructure ready for real-time updates

### Known Limitations
- Demo data only (no real inventory/schedule data)
- No authentication
- No persistent conversation history
- No external channel integrations (Slack, Teams, SMS)

### How to Run
```bash
# Terminal 1: Ollama
ollama serve

# Terminal 2: Backend
cd ~/opus && go run cmd/server/main.go

# Terminal 3: Frontend
cd ~/opus/web && npm install && npm run dev

# Browser
open http://localhost:5173
```

---

## Next 3 Steps

### Step 1: Add Demo Data System
**Goal**: Create realistic mock data for inventory, schedules, and alerts that the AI can reference.

Tasks:
- [ ] Create demo data store (in-memory for now)
- [ ] Add inventory data per department (items, quantities, expiration dates)
- [ ] Add schedule data (shifts, employees, coverage)
- [ ] Add active alerts (system issues, low stock, expiring items)
- [ ] Update AI agents to query and reference this data

### Step 2: Build Department Detail Pages
**Goal**: Create individual pages for each department with inventory and schedule views.

Tasks:
- [ ] Create `/departments/[dept]` route
- [ ] Display department inventory table
- [ ] Display department schedule
- [ ] Show department-specific alerts
- [ ] Add quick actions (mark item low, request coverage)

### Step 3: Implement Alert System
**Goal**: Create a working alert system with real-time notifications.

Tasks:
- [ ] Define alert types and severity levels
- [ ] Create alert generation logic (low stock, expiring items, etc.)
- [ ] Push alerts via WebSocket to connected clients
- [ ] Add alert management UI (acknowledge, dismiss, escalate)
- [ ] Add alert history/log

---

## Future Roadmap

### Phase 2: Data & Auth
- [ ] PostgreSQL database integration
- [ ] User authentication (JWT)
- [ ] Role-based access control
- [ ] Conversation history persistence

### Phase 3: External Integrations
- [ ] Slack channel adapter
- [ ] Teams channel adapter
- [ ] SMS via Twilio
- [ ] Email notifications

### Phase 4: Production Connectors
- [ ] On-premise bridge architecture
- [ ] Periscope inventory connector
- [ ] UKG scheduling connector
- [ ] POS status connector

### Phase 5: Production Hardening
- [ ] Rate limiting
- [ ] Audit logging
- [ ] Error monitoring
- [ ] Performance optimization
- [ ] Multi-store support

---

## Session Log

| Date | Work Done |
|------|-----------|
| 2026-02-22 | Project setup, Go backend, SvelteKit frontend, AI integration with Ollama |

---

## Notes

- Using Llama 3 locally to avoid API costs during development
- Claude API fallback planned for complex queries
- Demo targets single store; multi-store is Phase 5
- Corporate/IT approval needed for production integrations (Periscope, UKG, POS)
