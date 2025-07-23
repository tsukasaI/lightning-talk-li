# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Slidev presentation about Playwright for E2E testing with Go. The repository contains a slide deck built with Slidev (a presentation framework for developers) that demonstrates how to use Playwright-Go for browser automation testing.

## Architecture

- **Slidev Framework**: Uses Slidev for creating interactive presentations with Vue.js components
- **Presentation Content**: Main slides are in `slides.md` with Markdown + Vue syntax
- **Components**: Vue components in `components/` directory (e.g., `Counter.vue`)
- **Deployment**: Configured for both Netlify (`netlify.toml`) and Vercel (`vercel.json`)
- **Demo Application**: Go web server (`sample/server.go`) with HTML/JS frontend for testing demonstrations
- **Test Suite**: Go Playwright tests (`sample/demo_test.go`) showcasing browser automation patterns

## Development Commands

### Installation and Setup
```bash
pnpm install
```

### Development
```bash
pnpm dev
```
This starts the development server and opens the presentation at http://localhost:3030

### Build
```bash
pnpm build
```
Builds the presentation for production deployment

### Export
```bash
pnpm export
```
Exports the presentation to static files or PDF

### Demo Application (Go)
```bash
# Navigate to sample directory
cd sample

# Install Go dependencies
go mod tidy

# Install Playwright browser drivers (required before running tests)
go install github.com/playwright-community/playwright-go/cmd/playwright@v0.5200.0
playwright install --with-deps

# Run the demo web server
go run server.go
```
Server runs on http://localhost:8080 and serves a simple HTML page with interactive button for testing.

### Run Playwright Tests
```bash
# From sample directory, with server running
go test -v

# Run specific test
go test -v -run TestButtonClick
```
Tests require the demo server to be running on localhost:8080.

## Key Files

### Presentation Files
- `slides.md`: Main presentation content with slide configuration and content
- `components/Counter.vue`: Vue component demonstrating counter functionality
- `package.json`: Dependencies include Slidev CLI and themes
- `pages/imported-slides.md`: Additional slide content
- `snippets/external.ts`: TypeScript code snippets for the presentation

### Demo Application Files
- `sample/server.go`: Simple Go web server serving HTML page with interactive elements
- `sample/demo_test.go`: Comprehensive Playwright test suite with 5 test cases covering button interactions, element visibility, page title, and response time
- `sample/go.mod`: Go module dependencies including playwright-go and testify

## Presentation Structure

The slides cover:
1. Introduction to browser automation testing
2. Comparison between Selenium and Playwright
3. Playwright setup and basic usage in Go
4. Practical test examples
5. Key features and benefits
6. Live demo walkthrough

## Theme and Styling

- Uses Slidev's default theme with custom gradient background
- Dark color scheme optimized for presentations
- Responsive design with grid layouts for content organization

## Test Architecture

The demo application follows a practical testing pattern:

1. **Server Setup**: Simple Go HTTP server serves a single-page application with interactive elements
2. **Test Structure**: Each test function handles full Playwright lifecycle (start → launch browser → create page → test → cleanup)
3. **Test Coverage**: 
   - Button click interactions and text changes
   - Element visibility and existence checks
   - Page title and heading verification
   - Response time performance testing
   - Message cycling through predefined array

Tests use `testify/require` for critical assertions and `testify/assert` for validation checks.