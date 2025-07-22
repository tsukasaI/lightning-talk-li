# Playwright for E2E Testing with Go - Slidev Presentation

This is a Slidev presentation about Playwright for E2E testing with Go, demonstrating browser automation testing patterns.

## Getting Started

To start the slide show:

- `pnpm install`
- `pnpm dev`
- visit <http://localhost:3030>

Edit the [slides.md](./slides.md) to see the changes.

## Demo Application

The presentation includes a live demo with a Go web server and Playwright tests:

```bash
# Run the demo server
cd sample
go run server.go

# Run the tests (in another terminal)
go test -v
```

Learn more about Slidev at the [documentation](https://sli.dev/).
