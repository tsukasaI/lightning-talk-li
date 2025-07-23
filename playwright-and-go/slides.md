---
theme: default
background: linear-gradient(45deg, #1e3a8a, #3b82f6)
title: Playwright for E2E Testing with Go
info: |
  ## Playwright for E2E Testing with Go
  A lightning talk about using Playwright for browser automation testing with Golang
class: text-center text-white
highlighter: shiki
lineNumbers: false
drawings:
  enabled: false
transition: slide-left
favicon: https://go.dev/images/go-logo-blue.svg
colorSchema: dark
---

# Playwright for E2E Testing with Go

Modern Browser Automation Made Simple

<div class="pt-8">
  <span class="text-sm text-gray-300">Tsukasa INOUE</span>
</div>

<div class="pt-4">
  <span @click="$slidev.nav.next" class="px-2 py-1 rounded cursor-pointer" hover="bg-white bg-opacity-10">
    Let's dive in! <carbon:arrow-right class="inline"/>
  </span>
</div>

---

# Why Browser Automation Testing?

Modern web applications need comprehensive testing

- **Unit tests** - Test individual functions
- **Integration tests** - Test component interactions  
- **E2E tests** - Test complete user workflows

<div class="bg-blue-900 bg-opacity-50 p-4 rounded mt-4 text-white">
E2E testing catches issues that unit tests miss - real user interactions, browser quirks, timing issues
</div>

---

# Browser Automation Landscape

Popular tools for browser automation testing

<div class="grid grid-cols-2 gap-8 mt-8">

<div class="bg-gray-800 bg-opacity-50 p-4 rounded text-white">
<h3 class="text-lg font-bold mb-2">Selenium</h3>
<ul class="text-sm">
<li>✅ Industry standard since 2004</li>
<li>✅ Mature ecosystem</li>
<li>✅ Multi-language support</li>
<li>✅ Cross-browser support</li>
<li>❌ Complex setup</li>
<li>❌ Flaky tests (timing issues)</li>
</ul>
</div>

<div class="bg-blue-800 bg-opacity-50 p-4 rounded text-white">
<h3 class="text-lg font-bold mb-2">Playwright</h3>
<ul class="text-sm">
<li>✅ Modern & fast</li>
<li>✅ Auto-waiting built-in</li>
<li>✅ Simple cross-browser setup</li>
<li>✅ Network interception</li>
<li>❌ Newer (less mature)</li>
</ul>
</div>

</div>

<div class="grid grid-cols-2 gap-8 mt-4">

<div class="bg-green-800 bg-opacity-50 p-4 rounded text-white">
<h3 class="text-lg font-bold mb-2">Cypress</h3>
<ul class="text-sm">
<li>✅ Great developer experience</li>
<li>✅ Real-time debugging</li>
<li>❌ JavaScript/TypeScript only</li>
<li>❌ Limited cross-browser support</li>
</ul>
</div>

<div class="bg-purple-800 bg-opacity-50 p-4 rounded text-white">
<h3 class="text-lg font-bold mb-2">Puppeteer</h3>
<ul class="text-sm">
<li>✅ Chrome DevTools Protocol</li>
<li>✅ Fast execution</li>
<li>❌ Chrome/Chromium only</li>
<li>❌ JavaScript/Node.js only</li>
</ul>
</div>

</div>

---

# Browser Automation Options for Go

How Go developers can leverage browser automation

<div class="grid grid-cols-2 gap-8 mt-8">

<div class="bg-gray-800 bg-opacity-50 p-4 rounded text-white">
<h3 class="text-lg font-bold mb-2">Selenium with Go</h3>
<ul class="text-sm">
<li>✅ `tebeka/selenium` package</li>
<li>✅ Well-established in Go ecosystem</li>
<li>✅ Extensive documentation</li>
<li>❌ Complex WebDriver setup</li>
<li>❌ Manual wait handling required</li>
</ul>
</div>

<div class="bg-blue-800 bg-opacity-50 p-4 rounded text-white">
<h3 class="text-lg font-bold mb-2">Playwright with Go</h3>
<ul class="text-sm">
<li>✅ `playwright-community/playwright-go`</li>
<li>✅ Modern Go-friendly API</li>
<li>✅ Auto-waiting built-in</li>
<li>✅ Zero-config cross-browser testing</li>
</ul>
</div>

</div>

<div class="bg-blue-900 bg-opacity-50 p-4 rounded mt-6 text-white">
<strong>For Go developers:</strong> Playwright offers a more modern, Go-idiomatic approach with better reliability
</div>

---

# Selenium vs Playwright

Why Playwright is the better choice for Go

<div class="grid grid-cols-2 gap-8 mt-8">

<div>
<h3 class="text-lg font-bold mb-4 text-orange-600">Selenium Setup</h3>

```go
// Complex setup & browser management
caps := selenium.Capabilities{
    "browserName": "chrome",
}
driver, err := selenium.NewRemote(caps, "")

// Manual waits required
driver.FindElement(By.ID("submit")).Click()
time.Sleep(2 * time.Second)

// Manual element state checking
element := driver.FindElement(By.ID("button"))
if element.IsEnabled() {
    element.Click()
}
```

</div>

<div>
<h3 class="text-lg font-bold mb-4 text-blue-600">Playwright Setup</h3>

```go
// Simple setup
pw, err := playwright.Run()
browser, err := pw.Chromium.Launch()
page, err := browser.NewPage()

// Auto-waiting built-in
page.Click("#submit")

// Intelligent waiting - no manual checks
page.Click("#button") // Waits automatically
```

</div>

</div>

---

# Why Playwright for Go?

Perfect match for Go developers

- **Community Go bindings** - Well-maintained unofficial Go support
- **Auto-waiting** - No more flaky tests from timing issues
- **Unified cross-browser API** - Same code works across Chrome, Firefox, Safari, Edge
- **Fast execution** - Parallel test execution
- **Modern features** - Network interception, mobile emulation
- **Simple API** - Clean, intuitive Go interface

<div class="bg-yellow-900 bg-opacity-50 p-3 rounded mt-4 text-white text-sm">
<strong>Note:</strong> Go bindings are community-maintained, not officially supported by Microsoft like Node.js, Python, Java, and .NET
</div>

---

# Demo Application (server.go)

Simple Go web server for testing demonstrations

<div class="grid grid-cols-3 gap-6 mt-4">

<div class="col-span-2">

```go
func main() {
    http.HandleFunc("/", handleHome)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
    html := `<!DOCTYPE html>
<html><body>
    <h1>Playwright Demo App</h1>
    <div id="text-display">Hello, World!</div>
    <button id="change-button">Change Text</button>
    <script>
        // JavaScript cycles through messages on button click
        const messages = ["Hello, World!", "Welcome to Playwright Testing!", 
                         "Go + Playwright = Amazing!"];
        // Button click handler updates text content
    </script>
</body></html>`
    w.Write([]byte(html))
}
```

</div>

<div class="flex items-center">

<div class="bg-blue-600 bg-opacity-40 p-4 rounded-lg text-white border border-blue-400">
<strong>Key elements:</strong> Simple HTTP server serving HTML with interactive button and text display for testing. Go unit tests alone cannot test the text-display content updating since it requires browser JavaScript execution.
</div>

</div>

</div>

---

# Basic Test Structure

Foundation of a Playwright test (from demo_test.go)

```go
func TestButtonClick(t *testing.T) {
    // Start Playwright
    pw, err := playwright.Run()
    require.NoError(t, err, "Failed to start Playwright")
    defer pw.Stop()

    // Launch browser
    browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
        Headless: playwright.Bool(false),
    })
    require.NoError(t, err, "Failed to launch browser")
    defer browser.Close()

    // Create new page
    page, err := browser.NewPage()
    require.NoError(t, err, "Failed to create page")
    
    // Navigate to the web app
    _, err = page.Goto("http://localhost:8080")
    require.NoError(t, err, "Failed to navigate to localhost:8080")
}
```

---

# Practical E2E Test Example

Real-world button interaction testing (from demo_test.go)

```go
func TestButtonClick(t *testing.T) {
    // ... setup code ...
    
    // Verify initial text
    textDisplay := page.Locator("#text-display")
    initialText, err := textDisplay.TextContent()
    require.NoError(t, err, "Failed to get initial text")
    assert.Equal(t, "Hello, World!", initialText)

    // Click the button
    changeButton := page.Locator("#change-button")
    err = changeButton.Click()
    require.NoError(t, err, "Failed to click button")

    // Verify text has changed
    newText, err := textDisplay.TextContent()
    require.NoError(t, err, "Failed to get new text")
    assert.Equal(t, "Welcome to Playwright Testing!", newText)
}
```

---

# Key Playwright Features

What makes Playwright powerful

<div class="grid grid-cols-2 gap-8 mt-4">

<div>
<h3 class="text-lg font-bold mb-4">Auto-waiting</h3>
<ul class="text-sm space-y-2">
<li>✅ Waits for elements to be visible</li>
<li>✅ Waits for elements to be enabled</li>
<li>✅ Waits for network requests</li>
<li>✅ No more `time.Sleep()`</li>
</ul>
</div>

<div>
<h3 class="text-lg font-bold mb-4">Cross-browser</h3>
<ul class="text-sm space-y-2">
<li>🌐 Chrome/Chromium - Zero setup</li>
<li>🦊 Firefox - Automatically managed</li>
<li>🧭 Safari (WebKit) - Built-in driver</li>
<li>📱 Mobile browsers - Device emulation</li>
</ul>
</div>

</div>

<div class="mt-8">
<h3 class="text-lg font-bold mb-4">Advanced Features</h3>
<ul class="text-sm space-y-2">
<li>📸 Screenshots and videos</li>
<li>🌐 Network interception</li>
<li>📱 Mobile device emulation</li>
<li>⚡ Parallel execution</li>
</ul>
</div>

---

# Demo Test Suite

Our comprehensive test suite (demo_test.go)

<div class="grid grid-cols-2 gap-8 mt-8">

<div>
<h3 class="text-lg font-bold mb-4 text-blue-600">Test Coverage</h3>
<ul class="text-sm space-y-2">
<li>✅ **TestButtonClick** - Basic interaction</li>
<li>✅ **TestMultipleClicks** - Message cycling</li>
<li>✅ **TestButtonExists** - Element visibility</li>
<li>✅ **TestPageTitle** - Page metadata</li>
<li>✅ **TestResponseTime** - Performance</li>
</ul>
</div>

<div>
<h3 class="text-lg font-bold mb-4 text-green-600">What We Test</h3>
<ul class="text-sm space-y-2">
<li>🎯 Button functionality</li>
<li>📝 Text content changes</li>
<li>👁️ UI element visibility</li>
<li>⚡ Performance metrics</li>
<li>🔄 State transitions</li>
</ul>
</div>

</div>

<div class="mt-8 bg-blue-900 bg-opacity-50 p-4 rounded text-white">
<strong>Run tests:</strong> `go test -v` • <strong>Start server:</strong> `go run server.go`
</div>

---

# Live Demo Time!

Let's see our tests in action

<div class="grid grid-cols-2 gap-8 mt-8">

<div>
<h3 class="text-lg font-bold mb-4 text-blue-600">Demo Steps</h3>
<ol class="text-sm space-y-2">
<li>Start the Go server</li>
<li>Open the web app</li>
<li>Run Playwright tests</li>
<li>See automated testing magic!</li>
</ol>
</div>

<div>
<h3 class="text-lg font-bold mb-4 text-green-600">Commands</h3>
<div class="text-sm space-y-2">
<div><code>go run server.go</code></div>
<div><code>go test -v</code></div>
<div><code>go test -v TestButtonClick</code></div>
</div>
</div>

</div>

---

# Key Takeaways

What to remember about Playwright + Go

- **Easy setup** - Simple installation and configuration
- **Auto-waiting** - No more flaky tests from timing issues  
- **Effortless cross-browser** - Test across all major browsers without driver management
- **Go-friendly** - Clean API that feels natural in Go
- **Fast execution** - Parallel tests, quick feedback

<div class="bg-blue-900 bg-opacity-50 p-4 rounded mt-8 text-white">
<strong>Start small:</strong> Pick one critical user flow and automate it with Playwright
</div>

---
layout: center
class: text-center text-white
---

# Thank you for listening!

<div class="mt-8">
<h3 class="text-lg mb-4">Resources</h3>
<ul class="text-sm space-y-2">
<li>🔗 <a href="https://playwright.dev/go" target="_blank">playwright.dev/go</a></li>
<li>📚 <a href="https://github.com/playwright-community/playwright-go" target="_blank">GitHub: playwright-community/playwright-go</a></li>
<li>📖 <a href="https://pkg.go.dev/github.com/playwright-community/playwright-go" target="_blank">Go Package Documentation</a></li>
</ul>
</div>