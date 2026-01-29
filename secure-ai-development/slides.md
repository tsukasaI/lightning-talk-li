---
theme: seriph
title: Secure Development with AI Coding Assistants
info: |
  ## Secure Development with AI Coding Assistants
  Understanding security risks when coding with AI
class: text-center
transition: slide-left
mdc: true
---

# Secure Development with AI Coding Assistants

Understanding security risks when coding with AI

---
layout: two-cols-header
---

# Why I Talk About This

I use Claude Code **a lot**

::left::

<img src="/asset/claude-billing.png" class="h-50" />

Billing

::right::

<img src="/asset/daily-usage.png" class="h-50" />

Cost calculated by ccusage

---

# Table of Contents

- Claude Code Quick Introduction
- Code Vulnerabilities
- Execution & Third-party Risks
- Takeaways

---
layout: center
---

# Claude Code Quick Introduction

---

# What is Claude Code?

**Agentic coding tool** that runs in your terminal (CLI)

- Understands your codebase context
- Edits files, runs commands, searches code
- Requires your approval for actions

https://code.claude.com/docs/en/overview

---

# Claude Code Configuration

```
~/.claude/settings.json          # Global settings
.claude/settings.json            # Project-level settings (overrides global)
```

```json
{
  "permissions": {
    "allow": ["Bash(npm *)"],
    "deny": ["Read(.env*)"]
  },
  "hooks": { ... }
}
```

---
layout: center
---

# Code Vulnerabilities

---

# What Can Go Wrong

AI can write vulnerable code - just like humans

```go
// Bad: SQL Injection
query := fmt.Sprintf("SELECT * FROM users WHERE id = %s", userID)
// Good: db.Query("SELECT * FROM users WHERE id = ?", userID)

// Bad: XSS
fmt.Fprintf(w, "<div>Hello, %s</div>", username)
// Good: tmpl.Execute(w, username)  // html/template auto-escapes

// Bad: Path Traversal
path := filepath.Join("/data", userInput)  // userInput: "../etc/passwd"
// Good: filepath.Clean() + validate within base directory
```

---

# How to Review

### Use Custom Command

```
/security-review
```

Add to `.claude/commands/security-review.md`:

```markdown
Review the code changes for security vulnerabilities:
- SQL injection, XSS, CSRF, path traversal
- Hardcoded secrets
- Insecure dependencies
```

---
layout: center
---

# Execution & Third-party Risks

---

# What Can Go Wrong

### Command Injection

- AI suggests shell commands that may be dangerous
- Piped commands, `rm -rf`, `curl | sh`, etc.
- Prompt injection can trick AI into running malicious commands

### Untrusted Third-party Tools (MCP)

- MCP skills can inject hidden instructions
- Skills may exfiltrate code/secrets to external servers
- Users grant broad permissions without review
- "Allow all" becomes the path of least resistance

---

# How to Reduce Risk

### Block Dangerous Commands

```json
{ "permissions": { "deny": ["Bash(rm -rf*)"] } }
```

When a denied command is attempted:

<img src="/asset/denied-command.png" class="w-120" />

---
layout: two-cols-header
---

# Reduce Risk for Network

::left::

### Built-in Protections

- `curl`, `wget` blocked by default
- Write access restricted to working directory
- Use `/sandbox` for isolation

### Without deny rule for curl

<img src="/asset/curl-command.png" class="w-full" />

::right::

### Whitelist Network Access

```json
{
  "permissions": {
    "allow": ["Bash(curl http://localhost:*)"],
    "deny": ["Bash(curl *)"]
  }
}
```

---
layout: center
---

# Takeaways

---

# Key Takeaways

**This is not a new concept** - similar to classical development risks

### Always review what AI wants to do

Don't blindly trust generated code or suggested commands

<br>

### Don't trust third-party skills

Even if they look useful - audit before installation

<br>

### Learn up-to-date best practices

AI tools evolve fast - so do the risks

---

# References

- **Claude Code Documentation** - https://code.claude.com/docs
- **Claude Code Security** - https://code.claude.com/docs/en/security
- **Security Reviews** - https://support.claude.com/en/articles/11932705-automated-security-reviews-in-claude-code
- **MCP (Model Context Protocol)** - https://modelcontextprotocol.io
- **ccusage** - https://github.com/ryoppippi/ccusage
- **OWASP Top 10** - https://owasp.org/Top10

---
layout: center
---

# Thank You for Listening

If you'd like me to present more about Claude Code, feel free to let me know
