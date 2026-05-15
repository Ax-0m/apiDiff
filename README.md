# apiDiff

A fast, lightweight CLI tool to compare two JSON API responses and detect structural changes.

## Install

```bash
go install github.com/Ax-0m/apiDiff@latest
```

## Usage

```bash
apidiff old.json new.json
```

## What it detects

- ✚ Added fields
- ✖ Removed fields
- ~ Modified values
- ⚠ Type changes (e.g. number → string)

## Example

old.json

```json
{
  "user": {
    "id": 1,
    "name": "Prakhar"
  }
}
```

new.json

```json
{
  "user": {
    "id": "usr_001",
    "name": "Prakhar",
    "city": "Bangalore"
  }
}
```

Output:
⚠ user.id: type changed number → string
✚ user.city: added (Bangalore)

## Roadmap

- [ ] Array diffing
- [ ] Snapshot system — run `apidiff snapshot save` to store a response, then `apidiff snapshot diff` to compare later
- [ ] OpenAPI support

## Built with

Go
