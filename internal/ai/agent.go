package ai

import (
	"context"
	"fmt"
	"strings"
)

// Department represents a store department
type Department string

const (
	DeptDairy    Department = "dairy"
	DeptProduce  Department = "produce"
	DeptMeat     Department = "meat"
	DeptBakery   Department = "bakery"
	DeptDeli     Department = "deli"
	DeptGrocery  Department = "grocery"
	DeptFrontEnd Department = "frontend"
)

// Agent represents a department-specific AI agent
type Agent struct {
	department  Department
	ollama      *OllamaClient
	systemPrompt string
}

// NewAgent creates a new department agent
func NewAgent(dept Department, ollama *OllamaClient) *Agent {
	return &Agent{
		department:   dept,
		ollama:       ollama,
		systemPrompt: buildSystemPrompt(dept),
	}
}

func buildSystemPrompt(dept Department) string {
	base := `You are Opus, an AI assistant for grocery store operations. You help department managers and assistant managers with their daily tasks.

You have access to:
- Inventory data (stock levels, reorder points, expiration dates)
- Schedule information (shifts, coverage, time-off requests)
- Sales data and trends
- Alert and monitoring systems

Be concise, helpful, and action-oriented. When managers ask questions, provide direct answers and suggest next steps when appropriate.

Current context: You are assisting the %s department.

Department-specific knowledge:
%s

Always respond in a professional but friendly manner. If you don't have specific data, say so clearly and suggest how to get it.`

	deptKnowledge := getDepartmentKnowledge(dept)
	return fmt.Sprintf(base, dept, deptKnowledge)
}

func getDepartmentKnowledge(dept Department) string {
	knowledge := map[Department]string{
		DeptDairy: `- Monitor milk, cheese, yogurt, and egg inventory closely due to short shelf life
- Pay attention to expiration dates and FIFO rotation
- Temperature monitoring is critical (dairy coolers should be 35-38°F)
- Common shrink issues: expired products, temperature abuse
- Peak ordering: Sunday/Monday for weekend sales`,

		DeptProduce: `- Fresh produce requires daily quality checks
- Monitor ripeness levels for bananas, avocados, tomatoes
- Wet rack items need frequent misting and rotation
- Seasonal availability affects ordering
- High shrink department - track waste carefully
- Organic vs conventional inventory separation`,

		DeptMeat: `- Temperature critical (meat cases 28-32°F)
- Track sell-by dates carefully
- Monitor grinding logs and food safety compliance
- Marination and value-added prep scheduling
- Weekend and holiday demand spikes
- Special orders and custom cuts tracking`,

		DeptBakery: `- Production schedules based on sales patterns
- Fresh-baked timing for peak traffic
- Track ingredient inventory (flour, sugar, eggs)
- Special order management (cakes, platters)
- End-of-day markdown decisions
- Seasonal and holiday production planning`,

		DeptDeli: `- Hot bar and salad bar temperature monitoring
- Sliced meats and cheeses inventory
- Prepared foods production scheduling
- Catering and special orders
- Food safety compliance (time/temp logs)
- Peak lunch and dinner rush preparation`,

		DeptGrocery: `- Center store inventory management
- Shelf capacity and facing standards
- Promotional display execution
- Vendor deliveries and resets
- Overstock management
- Category planogram compliance`,

		DeptFrontEnd: `- Register operations and cash management
- Cashier scheduling for peak hours
- Customer service issues escalation
- Register/equipment status monitoring
- Cart availability and lot maintenance
- Checkout line management`,
	}

	if k, ok := knowledge[dept]; ok {
		return k
	}
	return "General store operations support."
}

// ProcessQuery handles a user query through the department agent
func (a *Agent) ProcessQuery(ctx context.Context, userQuery string, conversationHistory []Message) (string, error) {
	// Build conversation with system prompt
	messages := []Message{
		{Role: "system", Content: a.systemPrompt},
	}

	// Add conversation history
	messages = append(messages, conversationHistory...)

	// Add current user query
	messages = append(messages, Message{
		Role:    "user",
		Content: userQuery,
	})

	response, err := a.ollama.Chat(ctx, messages)
	if err != nil {
		return "", fmt.Errorf("agent query failed: %w", err)
	}

	return response, nil
}

// Router routes queries to the appropriate department agent
type Router struct {
	agents map[Department]*Agent
	ollama *OllamaClient
}

// NewRouter creates a new agent router
func NewRouter(ollama *OllamaClient) *Router {
	r := &Router{
		agents: make(map[Department]*Agent),
		ollama: ollama,
	}

	// Initialize all department agents
	departments := []Department{
		DeptDairy, DeptProduce, DeptMeat, DeptBakery,
		DeptDeli, DeptGrocery, DeptFrontEnd,
	}

	for _, dept := range departments {
		r.agents[dept] = NewAgent(dept, ollama)
	}

	return r
}

// Route determines which department should handle a query
func (r *Router) Route(query string) Department {
	query = strings.ToLower(query)

	// Simple keyword-based routing (can be enhanced with AI classification)
	routingRules := map[Department][]string{
		DeptDairy:    {"milk", "cheese", "yogurt", "egg", "butter", "cream", "dairy"},
		DeptProduce:  {"fruit", "vegetable", "produce", "banana", "apple", "lettuce", "organic"},
		DeptMeat:     {"meat", "beef", "chicken", "pork", "steak", "ground", "butcher"},
		DeptBakery:   {"bread", "cake", "pastry", "donut", "bakery", "fresh baked"},
		DeptDeli:     {"deli", "sandwich", "sliced", "hot bar", "salad bar", "catering"},
		DeptGrocery:  {"aisle", "shelf", "grocery", "canned", "cereal", "snack"},
		DeptFrontEnd: {"register", "cashier", "checkout", "front end", "cart", "customer service"},
	}

	for dept, keywords := range routingRules {
		for _, keyword := range keywords {
			if strings.Contains(query, keyword) {
				return dept
			}
		}
	}

	// Default to grocery for unmatched queries
	return DeptGrocery
}

// ProcessQuery routes and processes a query
func (r *Router) ProcessQuery(ctx context.Context, query string, history []Message) (string, Department, error) {
	dept := r.Route(query)
	agent := r.agents[dept]

	response, err := agent.ProcessQuery(ctx, query, history)
	if err != nil {
		return "", dept, err
	}

	return response, dept, nil
}
