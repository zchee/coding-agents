package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/adk/agent"
	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/cmd/launcher/full"
	"google.golang.org/adk/model/gemini"
	"google.golang.org/adk/tool"
	"google.golang.org/adk/tool/geminitool"
	"google.golang.org/adk/tool/mcptoolset"
	"google.golang.org/genai"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_API_KEY missing")
	}

	ctx := context.Background()
	model, err := gemini.NewModel(ctx, "gemini-2.0-flash-exp", &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatalf("create model: %v", err)
	}

	timeAgent, err := llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       model,
		Description: "Tells current time in a city.",
		Instruction: "You are concise; answer with a one-line time summary.",
		Tools: []tool.Tool{
			geminitool.GoogleSearch{},
		},
		Toolsets: []tool.Toolset{
			newMCPToolset(ctx),
		},
	})
	if err != nil {
		log.Fatalf("build agent: %v", err)
	}

	launcher := full.NewLauncher()
	if err := launcher.Run(full.Config{
		AgentLoader: agent.NewSingleLoader(timeAgent),
	}); err != nil {
		log.Fatalf("launch: %v", err)
	}
}

func newMCPToolset(ctx context.Context) tool.Toolset {
	clientTransport, serverTransport := mcp.NewInMemoryTransports()

	server := mcp.NewServer(&mcp.Implementation{Name: "example", Version: "v0.1.0"}, nil)

	type echoInput struct {
		Message string `json:"message" jsonschema:"message to echo"`
	}
	type echoOutput struct {
		Upper string `json:"upper" jsonschema:"uppercased message"`
	}

	echo := func(_ context.Context, _ *mcp.CallToolRequest, in echoInput) (*mcp.CallToolResult, echoOutput, error) {
		return nil, echoOutput{Upper: strings.ToUpper(in.Message)}, nil
	}

	mcp.AddTool(server, &mcp.Tool{
		Name:        "echo",
		Description: "Echo text and uppercase it.",
	}, echo)

	type timeInput struct {
		City string `json:"city" jsonschema:"city name"`
	}
	type timeOutput struct {
		Now string `json:"now" jsonschema:"current time ISO8601"`
	}
	now := func(_ context.Context, _ *mcp.CallToolRequest, in timeInput) (*mcp.CallToolResult, timeOutput, error) {
		return nil, timeOutput{Now: time.Now().UTC().Format(time.RFC3339)}, nil
	}
	mcp.AddTool(server, &mcp.Tool{
		Name:        "time_now",
		Description: "Return current UTC time; city is accepted but unused.",
	}, now)

	type weatherInput struct {
		City string `json:"city" jsonschema:"city name"`
	}
	type weatherOutput struct {
		Summary string `json:"summary" jsonschema:"short weather summary"`
	}
	weather := func(_ context.Context, _ *mcp.CallToolRequest, in weatherInput) (*mcp.CallToolResult, weatherOutput, error) {
		return nil, weatherOutput{Summary: "Sunny in " + in.City}, nil
	}
	mcp.AddTool(server, &mcp.Tool{
		Name:        "weather_stub",
		Description: "Return a stub weather summary for a city.",
	}, weather)

	if _, err := server.Connect(ctx, serverTransport, nil); err != nil {
		log.Fatalf("connect mcp server: %v", err)
	}

	ts, err := mcptoolset.New(mcptoolset.Config{
		Transport: clientTransport,
	})
	if err != nil {
		log.Fatalf("build mcp toolset: %v", err)
	}
	return ts
}
