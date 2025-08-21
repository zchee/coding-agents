# IMPORTANT

**First, MUST DON'T HOLD BACK. GIVE IT YOUR ALL.**

**Second, MUST after receiving your generated code or tool results, carefully reflect on their quality and determine optimal next steps before proceeding. Use your thinking to plan and iterate based on this new information, and then take the best next action.**

**Third, for maximum efficiency, whenever you need to perform multiple independent operations, **MUST** invoke all relevant tools simultaneously whenever possible, rather than sequentially.**

**Fourth, Actively use the Todo List (TodoList) tool with ultrathink. **MUST** always maintain a step-by-step, meaningful list of at least 25 items.**

# How to use MCP servers

* **MUST ACTIVELY USE `sequential-thinking` and `memory` MCP server with ultrathink.**
* **MUST ACTIVELY USE `abcoder` MCP server with ultrathink.**
    - @include ./instructions/abcoder.md
* **MUST ACTIVELY USE `pinecone` vector databases for the your memory with ultrathink.**
    - Tool Usage for Code Generation
        - When not found index, create index with current repository name.
        - When generating code related to Pinecone, always use the `pinecone` MCP and the `search_docs` tool.
        - Perform at least two distinct searches per request using different, relevant questions to ensure comprehensive context is gathered before writing code.
    - Error Handling
        - If an error occurs while executing Pinecone-related code, immediately invoke the `pinecone` MCP and the `search_docs` tool.
        - Search for guidance on the specific error encountered and incorporate any relevant findings into your resolution strategy.
    - Syntax and Version Accuracy
        - Before writing any code, verify and use the correct syntax for the latest stable version of the Pinecone SDK.
        - Prefer official code snippets and examples from documentation over generated or assumed field values.
        - Do not fabricate field names, parameter values, or request formats.
* **MUST ACTIVELY USE `gemini-google-search` MCP server that Google Gemini web search. MUST BE ALWAYS use this for web search instead of the builtin `Web_Search` tool. AGAIN, DONT USE builtin `Web_Search` tool.**
* **If you need more detailed information about the library or API details, MUST ACTIVELY USE the following MCP servers with deep thinking.**
    - `context7`
    - `deepwiki`

# Go programming language

@include ./instructions/Go.md
