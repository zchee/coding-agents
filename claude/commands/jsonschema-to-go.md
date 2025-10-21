Could you generate the **all** Go types from https://raw.githubusercontent.com/a2aproject/A2A/refs/heads/main/specification/json/a2a.json JSON Schema into internal/types.go with ultrathink?
- **MUST** struct tag only support `json`.
- **MUST** use `omitzero` if make sence. DO NOT pointer type with `omitempty`.
- Go types order are **SHOULD** same as JSON Schema.
- struct fields order as well.\n- Implements `MarshalJSON` and `UnmarshalJSON` methods if needed.
