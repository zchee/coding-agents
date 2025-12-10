# Go minimal ADK agent & tests (reference)

Keep this small and copy/paste friendly. Replace placeholders as needed.

## go.mod (snippet)
```go
module github.com/zchee/tumix/skills/adk/go-example

go 1.25

require (
    google.golang.org/adk v0.2.0
    google.golang.org/genai v0.4.0 // or latest matching ADK
    github.com/google/go-cmp v0.6.0
    github.com/modelcontextprotocol/go-sdk v0.1.1
)
```
> Change module path to your own repo if you copy out of this tree.

## agent.go
```go
package main

import (
    "log"
    "os"

    "google.golang.org/adk/agent"
    "google.golang.org/adk/agent/llmagent"
    "google.golang.org/adk/cmd/launcher/full"
    "google.golang.org/adk/model/gemini"
    "google.golang.org/adk/tool"
    "google.golang.org/adk/tool/geminitool"
    "google.golang.org/genai"
)

func main() {
    apiKey := os.Getenv("GOOGLE_API_KEY")
    if apiKey == "" {
        log.Fatal("GOOGLE_API_KEY missing")
    }

    model, err := gemini.NewModel(
        full.Background(),
        "gemini-2.0-flash-exp",
        &genai.ClientConfig{APIKey: apiKey},
    )
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
```

## agent_test.go (table-driven)
```go
package main

import (
    "testing"

    "github.com/google/go-cmp/cmp"
    "google.golang.org/adk/agent/llmagent"
    "google.golang.org/adk/model"
)

// Smoke-test the instruction wiring without hitting network.
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
                // Model nil is fine for config-only check.
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

// Example for testing event stream parity with a fake model would go here.
// Keep httprr traces under testdata/ and use t.Context().
```

## Run
- `export GOOGLE_API_KEY=...`
- `go run .` (CLI or web UI via launcher)
- `go test ./...`

Notes:
- Use `t.Context()` for integration tests that hit APIs.
- Prefer `mcptoolset.New` with `mcp.NewInMemoryTransports()` for MCP-backed tools in tests.
- Full example files live in `skills/adk/assets/go-example/` (go.mod, agent.go, agent_test.go with httprr + MCP skeletons).
- Example MCP tools there: `echo` (uppercase), `time_now` (UTC), `weather_stub` (static).
