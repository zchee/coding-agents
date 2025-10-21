#!/usr/bin/env python3.13

import asyncio
import json
from pathlib import Path
import sys


async def run_command(command, *args) -> str | None:
    process = await asyncio.create_subprocess_exec(
        command,
        *args,
        stdout=asyncio.subprocess.PIPE,
    )

    stdout, _ = await process.communicate()
    if stdout:
        return stdout.decode().strip()


""""Example JSON Input Structure:
{
  "session_id": "8635843f-4277-4e0e-aade-f84c434e0a2d",
  "transcript_path": "/Users/zchee/.claude/projects/-Users-zchee-src-github-com-gaudiy-vector-p/8635843f-4277-4e0e-aade-f84c434e0a2d.jsonl",
  "cwd": "/Users/zchee/src/github.com/gaudiy/vector-p",
  "model": {
    "id": "claude-sonnet-4-5-20250929[1m]",
    "display_name": "Sonnet 4.5 (with 1M token context)"
  },
  "workspace": {
    "current_dir": "/Users/zchee/src/github.com/gaudiy/vector-p",
    "project_dir": "/Users/zchee/src/github.com/gaudiy/vector-p"
  },
  "version": "2.0.14",
  "output_style": {
    "name": "default"
  },
  "cost": {
    "total_cost_usd": 0,
    "total_duration_ms": 1287,
    "total_api_duration_ms": 0,
    "total_lines_added": 0,
    "total_lines_removed": 0
  },
  "exceeds_200k_tokens": false
}
"""


async def main() -> None:
    # Read JSON from stdin
    data = json.load(sys.stdin)

    # Extract values
    model = data["model"]["id"]
    current_dir = Path(data["workspace"]["current_dir"]).name
    git_branch = await run_command("git", "branch", "--show-current")
    exceeds_200k_tokens = data["exceeds_200k_tokens"]
    print(f"  {model} |   {current_dir} |  {git_branch} | exceeds_200k_tokens: {exceeds_200k_tokens}")


# Check for git branch
# if os.path.exists('.git'):
#     try:
#         with open('.git/HEAD', 'r') as f:
#             ref = f.read().strip()
#             if ref.startswith('ref: refs/heads/'):
#                 git_branch = f" | 🌿 {ref.replace('ref: refs/heads/', '')}"
#     except:
#         pass

# print(f"{model} | 📁 {current_dir} |  {git_branch} | 💲 {data}")
asyncio.run(main())
