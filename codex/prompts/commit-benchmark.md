---
description: git commit in multiple with Go benchmark results
---
Please commit current changes with a meaningful message. Commit related changes in multiple commits.
* If have `BenchmarkXXX`, take a benchmark using `sync; /usr/local/sbin/purge; go test -v -run='^$' -timeout (N) -count (N) -benchtime (N)x -benchmem -bench={TARGET_BENCHMARK_NAME} ./...`
  - Before execute above command, run benchmarks against existing code and measure the difference with the `benchstat` command
  - If performance degrades, consider a different approach or revert the changes if none exist
* If have `benchstat` result, including raw output to commit message
