#!/usr/bin/env python3

import json
import os
from pathlib import Path

records = {}
codex_home = Path(os.getenv('CODEX_HOME', os.path.expanduser('~/.config/codex')))
with open(codex_home / 'history.jsonl', 'r') as f:
    for line in f:
        data = json.loads(line)
        # Use a unique identifier field (e.g., "session_id") as the key
        session_id = data.get('text')
        if session_id not in records:
            records[session_id] = data

for record in records.values():
    print(json.dumps(record))
