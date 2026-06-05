# apiDiff

A fast, lightweight CLI tool to monitor API structural changes.
No dashboards. No cloud. Just your terminal.

## Install

```bash
go install github.com/Ax-0m/apiDiff@latest
```

## Usage

### Compare two JSON files directly
```bash
apidiff old.json new.json
```

### Project based workflow
```bash
# Initialize a project
apidiff init users-service

# Edit apidiff.config.json with your base_url and endpoints

# Save snapshots
apidiff snapshot users-service

# Compare against saved snapshots
apidiff compare users-service
```

## What it detects

- ✚ Added fields
- ✖ Removed fields
- ~ Modified values  
- ⚠ Type changes (e.g. number → string)

## Example output
--- diff for /users ---
⚠  user.id: type changed  number → string
✚  user.city: added (Bangalore)
✖  user.address: removed

## Config format

```json
{
  "project": "users-service",
  "base_url": "https://api.example.com",
  "endpoints": [
    "/users",
    "/users/profile"
  ]
}
```

## Roadmap

### v3 — Smarter Diffing
- [ ] Array diffing — detect added, removed, reordered elements
- [ ] Auto endpoint discovery from OpenAPI/Swagger spec URL
- [ ] Auth header support in config (bearer tokens, API keys)
- [ ] `apidiff add-endpoint <project> <endpoint>` command

### v4 — Developer Workflow
- [ ] Scheduled snapshots — auto snapshot every X hours via cron
- [ ] `apidiff list` — list all projects and their last snapshot time
- [ ] Diff history — keep last N snapshots and compare across time
- [ ] GitHub Action — run apidiff compare in CI and fail on breaking changes
- [ ] Slack/webhook alerts on breaking changes
