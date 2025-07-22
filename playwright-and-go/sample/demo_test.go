package main

import (
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestButtonClick(t *testing.T) {
	// Start Playwright
	pw, err := playwright.Run()
	require.NoError(t, err, "Failed to start Playwright")
	defer pw.Stop()

	// Launch browser
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false), // Set to true for headless mode
	})
	require.NoError(t, err, "Failed to launch browser")
	defer browser.Close()

	// Create new page
	page, err := browser.NewPage()
	require.NoError(t, err, "Failed to create page")

	// Navigate to the web app
	_, err = page.Goto("http://localhost:8080")
	require.NoError(t, err, "Failed to navigate to localhost:8080")

	// Verify initial text
	textDisplay := page.Locator("#text-display")
	initialText, err := textDisplay.TextContent()
	require.NoError(t, err, "Failed to get initial text")
	assert.Equal(t, "Hello, World!", initialText, "Initial text should be 'Hello, World!'")

	// Click the button
	changeButton := page.Locator("#change-button")
	err = changeButton.Click()
	require.NoError(t, err, "Failed to click button")

	// Verify text has changed
	newText, err := textDisplay.TextContent()
	require.NoError(t, err, "Failed to get new text")
	assert.Equal(t, "Welcome to Playwright Testing!", newText, "Text should change after button click")

	// Click button again to test cycling
	err = changeButton.Click()
	require.NoError(t, err, "Failed to click button second time")

	thirdText, err := textDisplay.TextContent()
	require.NoError(t, err, "Failed to get third text")
	assert.Equal(t, "Go + Playwright = Amazing!", thirdText, "Text should cycle through messages")
}

func TestMultipleClicks(t *testing.T) {
	// Start Playwright
	pw, err := playwright.Run()
	require.NoError(t, err, "Failed to start Playwright")
	defer pw.Stop()

	// Launch browser in headless mode for faster execution
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	require.NoError(t, err, "Failed to launch browser")
	defer browser.Close()

	// Create new page
	page, err := browser.NewPage()
	require.NoError(t, err, "Failed to create page")

	// Navigate to the web app
	_, err = page.Goto("http://localhost:8080")
	require.NoError(t, err, "Failed to navigate to localhost:8080")

	// Expected messages in order
	expectedMessages := []string{
		"Hello, World!",
		"Welcome to Playwright Testing!",
		"Go + Playwright = Amazing!",
		"Browser automation is fun!",
		"Testing made easy!",
		"Click me again!",
		"Hello, World!", // Should cycle back to first message
	}

	textDisplay := page.Locator("#text-display")
	changeButton := page.Locator("#change-button")

	// Test cycling through all messages
	for i, expectedMsg := range expectedMessages {
		currentText, err := textDisplay.TextContent()
		require.NoError(t, err, "Failed to get text content at step %d", i)
		assert.Equal(t, expectedMsg, currentText, "Message at step %d should be '%s'", i, expectedMsg)

		// Click button to go to next message (except for the last iteration)
		if i < len(expectedMessages)-1 {
			err = changeButton.Click()
			require.NoError(t, err, "Failed to click button at step %d", i)
		}
	}
}

func TestButtonExists(t *testing.T) {
	// Start Playwright
	pw, err := playwright.Run()
	require.NoError(t, err, "Failed to start Playwright")
	defer pw.Stop()

	// Launch browser
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	require.NoError(t, err, "Failed to launch browser")
	defer browser.Close()

	// Create new page
	page, err := browser.NewPage()
	require.NoError(t, err, "Failed to create page")

	// Navigate to the web app
	_, err = page.Goto("http://localhost:8080")
	require.NoError(t, err, "Failed to navigate to localhost:8080")

	// Test that elements exist
	textDisplay := page.Locator("#text-display")
	changeButton := page.Locator("#change-button")

	// Check if elements are visible
	isTextVisible, err := textDisplay.IsVisible()
	require.NoError(t, err, "Failed to check text display visibility")
	assert.True(t, isTextVisible, "Text display should be visible")

	isButtonVisible, err := changeButton.IsVisible()
	require.NoError(t, err, "Failed to check button visibility")
	assert.True(t, isButtonVisible, "Button should be visible")

	// Check button text
	buttonText, err := changeButton.TextContent()
	require.NoError(t, err, "Failed to get button text")
	assert.Equal(t, "Change Text", buttonText, "Button should have correct text")
}

func TestPageTitle(t *testing.T) {
	// Start Playwright
	pw, err := playwright.Run()
	require.NoError(t, err, "Failed to start Playwright")
	defer pw.Stop()

	// Launch browser
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	require.NoError(t, err, "Failed to launch browser")
	defer browser.Close()

	// Create new page
	page, err := browser.NewPage()
	require.NoError(t, err, "Failed to create page")

	// Navigate to the web app
	_, err = page.Goto("http://localhost:8080")
	require.NoError(t, err, "Failed to navigate to localhost:8080")

	// Check page title
	title, err := page.Title()
	require.NoError(t, err, "Failed to get page title")
	assert.Equal(t, "Playwright Demo App", title, "Page title should be correct")

	// Check h1 heading
	heading := page.Locator("h1")
	headingText, err := heading.TextContent()
	require.NoError(t, err, "Failed to get heading text")
	assert.Equal(t, "Playwright Demo App", headingText, "H1 heading should be correct")
}

func TestResponseTime(t *testing.T) {
	// Start Playwright
	pw, err := playwright.Run()
	require.NoError(t, err, "Failed to start Playwright")
	defer pw.Stop()

	// Launch browser
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	require.NoError(t, err, "Failed to launch browser")
	defer browser.Close()

	// Create new page
	page, err := browser.NewPage()
	require.NoError(t, err, "Failed to create page")

	// Measure response time
	start := time.Now()
	_, err = page.Goto("http://localhost:8080")
	require.NoError(t, err, "Failed to navigate to localhost:8080")
	responseTime := time.Since(start)

	// Assert response time is reasonable (less than 5 seconds)
	assert.Less(t, responseTime, 5*time.Second, "Page should load within 5 seconds")

	// Test button click response time
	changeButton := page.Locator("#change-button")
	textDisplay := page.Locator("#text-display")

	start = time.Now()
	err = changeButton.Click()
	require.NoError(t, err, "Failed to click button")

	// Wait for text to change
	_, err = textDisplay.TextContent()
	require.NoError(t, err, "Failed to get text after click")
	clickResponseTime := time.Since(start)

	assert.Less(t, clickResponseTime, 1*time.Second, "Button click should respond within 1 second")
}
