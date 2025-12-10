#!/usr/bin/env -S uv run -s -U
# /// script
# requires-python = ">=3.14"
# dependencies = [
# "orjson",
# "rich",
# "loguru",
# ]
# ///

"""
Deduplicate Codex-style history.jsonl by `text`,
keeping the entry with the newest `ts`.

Usage:
    python dedup_history.py input.jsonl > output.jsonl
    cat input.jsonl | python dedup_history.py > output.jsonl
"""

import sys
from dataclasses import dataclass
from rich import print
from loguru import logger

import orjson


@dataclass
class Record:
    ts: int
    index: int  # order of first appearance of this text
    line: bytes  # original JSON line (without trailing newline)


def main() -> None:
    # Input source: file path or stdin
    if len(sys.argv) > 1 and sys.argv[1] != "-":
        infile = open(sys.argv[1], "rb")
        close_infile = True
    else:
        infile = sys.stdin.buffer
        close_infile = False

    records: dict[str, Record] = {}
    order = 0

    logger.info("Parse JSON")
    try:
        for raw in infile:
            # Remove line-ending only
            line = raw.rstrip(b"\r\n")
            # Skip empty / whitespace-only lines
            if not line.strip():
                continue

            # Parse JSON
            try:
                data = orjson.loads(line)
            except Exception as e:
                print(f"failed to parse JSON: {e}\n")
                sys.exit(1)

            # Extract required fields
            try:
                text = data["text"]
                ts = data["ts"]
            except KeyError as e:
                print(f"missing required field {e} in line: {line!r}\n")
                sys.exit(1)

            if not isinstance(text, str):
                print(f"`text` must be string, got {type(text).__name__}: {line!r}\n")
                sys.exit(1)

            if not isinstance(ts, int):
                print(f"`ts` must be integer, got {type(ts).__name__}: {line!r}\n")
                sys.exit(1)

            rec = records.get(text)
            if rec is None:
                # First time this text appears
                records[text] = Record(ts=ts, index=order, line=line)
                order += 1
            else:
                # Newer timestamp wins; if equal, last wins
                if ts > rec.ts:
                    rec.ts = ts
                    rec.line = line
                elif ts == rec.ts:
                    # Same ts: keep the last one
                    rec.line = line

    finally:
        if close_infile:
            infile.close()

    # Deterministic output: sort by first appearance index
    sorted_records = sorted(records.values(), key=lambda r: r.index)

    out = sys.stdout.buffer
    for rec in sorted_records:
        out.write(rec.line)
        out.write(b"\n")


if __name__ == "__main__":
    main()
