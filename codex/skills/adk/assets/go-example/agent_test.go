package main

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/modelcontextprotocol/go-sdk/mcp"

	"google.golang.org/adk/agent/llmagent"
	"google.golang.org/adk/tool/mcptoolset"
)

func TestAgentConfig(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		name string
	}{
		"success: name set": {name: "hello_time_agent"},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			cfg := llmagent.Config{
				Name:        tt.name,
				Description: "d",
				Instruction: "i",
			}

			if diff := cmp.Diff(tt.name, cfg.Name); diff != "" {
				t.Fatalf("name mismatch (-want +got):\n%s", diff)
			}
			if cfg.Model != nil {
				t.Fatalf("expected nil model for config-only test")
			}
		})
	}
}

// Skeleton for replaying recorded HTTP interactions (httprr).
// Provide a trace at testdata/sample.httprr before enabling.
func TestHttprrReplaySkeleton(t *testing.T) {
	t.Parallel()

	trace := "testdata/sample.httprr"
	if _, err := os.Stat(trace); errors.Is(err, fs.ErrNotExist) {
		t.Skipf("add httprr trace at %s to enable replay", trace)
	}

	// TODO: wire a replay transport to genai.ClientConfig.HTTPClient and
	// run a lightweight agent invocation, asserting emitted events.
}

// Minimal MCP in-memory toolset smoke test.
func TestMCPToolsetBuilds(t *testing.T) {
	t.Parallel()

	clientTransport, serverTransport := mcp.NewInMemoryTransports()

	type echoInput struct {
		Text string `json:"text"`
	}
	type echoOutput struct {
		Upper string `json:"upper"`
	}

	echo := func(ctx context.Context, req *mcp.CallToolRequest, in echoInput) (*mcp.CallToolResult, echoOutput, error) {
		return nil, echoOutput{Upper: strings.ToUpper(in.Text)}, nil
	}

	type timeOutput struct {
		Now string `json:"now"`
	}
	now := func(ctx context.Context, req *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, timeOutput, error) {
		return nil, timeOutput{Now: "2025-01-01T00:00:00Z"}, nil
	}

	server := mcp.NewServer(&mcp.Implementation{Name: "example", Version: "v0.1.0"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: "echo", Description: "echo text"}, echo)
	mcp.AddTool(server, &mcp.Tool{Name: "time_now", Description: "stub time"}, now)

	if _, err := server.Connect(t.Context(), serverTransport, nil); err != nil {
		t.Fatalf("connect server: %v", err)
	}

	ts, err := mcptoolset.New(mcptoolset.Config{
		Transport: clientTransport,
	})
	if err != nil {
		t.Fatalf("create toolset: %v", err)
	}

	tools, err := ts.Tools(nil)
	if err != nil {
		t.Fatalf("list tools: %v", err)
	}
	if len(tools) < 2 {
		t.Fatalf("expected at least two tools, got %d", len(tools))
	}
}
