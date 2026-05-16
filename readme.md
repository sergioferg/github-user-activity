# github-activity

A lightweight, zero-dependency command-line interface (CLI) tool built in Go that fetches, aggregates, and summarizes a user's recent GitHub activity using the official GitHub REST API.

## Features

- **Activity Aggregation:** Groups and counts duplicate event types (such as Pushes, Issue Comments, and Pull Requests) by repository for a clean terminal overview.
- **Robust CLI Validation:** Actively guards against missing or excessive command-line arguments to prevent runtime panics.
- **Defensive HTTP Client:** Utilizes a custom `http.Client` configured with a strict 10-second timeout to handle network hangs gracefully.
- **Standard Stream Routing:** Idiomatically routes successful activity metrics to Standard Output (`stdout`) and error logs to Standard Error (`stderr`).

## Installation

Because this tool is compiled down to a single static binary, you can install it globally on your machine with a single command without needing to manually clone the source code:

```bash
go install [github.com/sergioferg/github-user-activity@latest](https://github.com/sergioferg/github-user-activity@latest)
```
