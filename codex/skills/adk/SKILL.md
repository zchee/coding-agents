---
name: adk
description: Build, debug, and ship agents with Google’s Agent Development Kit (ADK) — especially the Go SDK (google/adk-go). Trigger when users ask to create/modify ADK agents, tools, workflows, launcher configs, tests, evaluations, or MCP toolsets bridging ADK and external tools.
---

# How to use this skill
- Use for: standing up or editing ADK agents (Go focus), wiring tools (gemini, MCP), running launcher/web UI, testing/evaluation, debugging orchestration.
- Load external refs only as needed to save context.

## Quick references (load on demand)
- ADK docs: https://google.github.io/adk-docs/ (Go quickstart, workflow agents, tools, eval, deploy)
- Go SDK repo: https://github.com/google/adk-go
- Minimal Go agent + tests: skills/adk/references/go_minimal_agent.md
- Ready-to-copy files: skills/adk/assets/go-example/ (go.mod, agent.go, agent_test.go)

## Prereqs
- Go ≥ 1.24.4, ADK Go ≥ v0.2.0.
- API key (e.g., GOOGLE_API_KEY for Gemini) in env/.env.
- `go get google.golang.org/adk@latest` to pull SDK; run `go mod tidy`.

## Core workflows
- New agent: create module; implement `llmagent.New` or workflow agents; prefer toolsets over ad hoc HTTP calls.
- Tools: start with built-ins (`geminitool.GoogleSearch`, etc.); add custom tools via `functiontool.New` or MCP bridge (`mcptoolset.New` with `mcp.NewInMemoryTransports` during tests).
- Launcher: use `cmd/launcher/full` for web UI + CLI; set `AgentLoader` to your agent(s).
- Sessions & memory: use `session.InMemoryService()` for local dev; avoid globals.
- Instruction templates: keep placeholders `{var}`; resolve via `instructionutil.InjectSessionState`.
- Testing: table-driven `tests := map[string]struct{...}`; use `t.Context()`; cover tool packing and event stream diffs with go-cmp; include httprr traces for LLM/http.
- Evaluation: mirror ADK eval guides—assert both final response and step events; prefer deterministic prompts.

## Troubleshooting
- Version drift: pin ADK and genai versions together; update go.sum if proto mismatch.
- Tool errors: check `CallToolResult.IsError` and message content; surface actionable details.
- API key missing: ensure env before launcher/web UI.
- MCP bridging: if no `structuredContent`, pack text response under `output`.

## Suggested prompts to trigger this skill
- “Design a Go ADK agent that does ___; include tools and launcher setup.”
- “Add an MCP toolset to my ADK agent for ___ API; show code + config.”
- “Write tests for an ADK tool using httprr and go-cmp.”
- “Debug why my ADK agent’s GoogleSearch tool isn’t returning results.”
