# CLAUDE.md

@./instructions/chatmodes/Ultimate-Transparent-Thinking-Beast-Mode.chatmode.md

> Ultrathink carefully and implement the most concise solution that changes as little code as possible.

## IMPORTANT

**First, MUST DON'T HOLD BACK. GIVE IT YOUR ALL.**

**Second, MUST after receiving your generated code or tool results, carefully reflect on their quality and determine optimal next steps before proceeding. Use your thinking to plan and iterate based on this new information, and then take the best next action.**

**Third, MUST ACTIVELY USE the `TodoWrite` tool and `sequential-thinking` MCP servers with ultrathink. MUST always maintain a step-by-step, meaningful list of at least 25 items.**

**Fourth, for maximum efficiency, whenever you need to perform multiple independent operations, MUST invoke all relevant tools simultaneously whenever possible, rather than sequentially.**

## ABSOLUTE RULES

<!-- - NO PARTIAL IMPLEMENTATION -->
- NO SIMPLIFICATION
    * No `// This is simplified stuff for now, complete implementation would blablabla`
- NO CODE DUPLICATION
    * Check existing codebase to reuse functions and constants Read files before writing new functions. Use common sense function name to find them easily.
- NO DEAD CODE
    * Either use or delete from codebase completely
- IMPLEMENT TEST FOR EVERY FUNCTIONS
- NO CHEATER TESTS
    * Test must be accurate, reflect real usage and be designed to reveal flaws. No useless tests! Design tests to be verbose so we can use them for debuging.
- NO INCONSISTENT NAMING
    * Read existing codebase naming patterns.
- NO OVER-ENGINEERING
    * Don't add unnecessary abstractions, factory patterns, or middleware when simple functions would work. Don't think "enterprise" when you need "working"
- NO MIXED CONCERNS
    * Don't put validation logic inside API handlers, database queries inside UI components, etc. instead of proper separation
- NO RESOURCE LEAKS
    * Don't forget to close database connections, clear timeouts, remove event listeners, or clean up file handles

## Tone and Behavior

- Criticism is welcome. Please tell me when I am wrong or mistaken, or even when you think I might be wrong or mistaken.
- Please tell me if there is a better approach than the one I am taking.
- Please tell me if there is a relevant standard or convention that I appear to be unaware of.
- Be skeptical.
- Be concise.
- Short summaries are OK, but don't give an extended breakdown unless we are working through the details of a plan.
- Do not flatter, and do not give compliments unless I am specifically asking for your judgement.
- Occasional pleasantries are fine.
- Feel free to ask many questions. If you are in doubt of my intent, don't guess. Ask.

## Philosophy

### Error Handling

- **Fail fast** for critical configuration (missing text model)
- **Log and continue** for optional features (extraction model)
- **Graceful degradation** when external services unavailable
- **User-friendly messages** through resilience layer

### Testing

- Always use the test-runner agent to execute tests.
- Do not use mock services for anything ever.
- Do not move on to the next test until the current test is complete.
- If the test fails, consider checking if the test is structured correctly before deciding we need to refactor the codebase.
- Tests to be verbose so we can use them for debugging.

## Tools

* **MUST ACTIVELY USE `memory` tool ultrathink**

## MCP servers

* **MUST ACTIVELY USE `sequential-thinking` MCP server ultrathink**
* **MUST ACTIVELY USE `claude-context` MCP server ultrathink**
* **MUST ACTIVELY USE `gemini-google-search` MCP server for Google Gemini web search. MUST ALWAYS use this for web search instead of the builtin `WebSearch` tool. AGAIN, DON'T USE the builtin `WebSearch` tool**
* **MUST ACTIVELY USE the `context7` MCP servers with deep thinking if you need more detailed information about the library or API details.**

<!-- ## Go programming language -->
<!---->
<!-- @~/.claude/instructions/Go.md -->
<!---->

<!-- ## Terraform programming language -->
<!---->
<!-- **MUST ACTIVELY USE `terraform` MCP server -->
<!---->
<!-- ## Zsh programming language -->
<!---->
<!-- @~/.claude/instructions/Zsh.md -->

