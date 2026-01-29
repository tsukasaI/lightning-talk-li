---
theme: seriph
title: Garbage Collector in Go
info: |
  ## Slidev Starter Template
  Presentation slides for developers.

  Learn more at [Sli.dev](https://sli.dev)
# apply unocss classes to the current slide
class: text-center

# slide transition: https://sli.dev/guide/animations.html#slide-transitions
transition: slide-left
# enable MDC Syntax: https://sli.dev/features/mdc
mdc: true
---

table of content

---

memory management in a program

C allocate and free memory

forgot free memory

some engineers free same memory twice and crush the program

---

is that engineer's work?

We should focus on business logic, design and other things

---

Garbage Collector

dynamically free memory

---

is that all happy?

-> No

During GC operation, the application stops

Stop the world (STW)

---

go gc basic

concurrent mark sweep

---

tri color

sweep

STW

---

# Go GC Workflow

```mermaid
flowchart TD
    Start([GC Trigger]) --> Init[Initialize GC]
    Init --> MarkStart[Mark Phase Start]

    MarkStart --> AllWhite[All objects marked WHITE]
    AllWhite --> RootGray[Mark root objects as GRAY]

    RootGray --> GrayExists{Any GRAY objects exist?}

    GrayExists -->|Yes| SelectGray[Select a GRAY object]
    SelectGray --> MarkBlack[Mark selected object BLACK]
    MarkBlack --> ScanRefs[Scan object's references]
    ScanRefs --> MarkRefsGray[Mark referenced WHITE objects as GRAY]
    MarkRefsGray --> GrayExists

    GrayExists -->|No| MarkComplete[Mark phase complete]
    MarkComplete --> SweepStart[Sweep Phase Start]

    SweepStart --> ScanHeap[Scan heap memory]
    ScanHeap --> CheckColor{Object color?}

    CheckColor -->|WHITE| FreeMemory[Free object memory]
    CheckColor -->|BLACK| KeepObject[Keep object, reset to WHITE]

    FreeMemory --> NextObject{More objects?}
    KeepObject --> NextObject

    NextObject -->|Yes| ScanHeap
    NextObject -->|No| GCComplete[GC Complete]

    GCComplete --> End([Program continues])

    %% Styling
    classDef whiteObj fill:#ffffff,stroke:#000000,stroke-width:2px
    classDef grayObj fill:#808080,stroke:#000000,stroke-width:2px,color:#ffffff
    classDef blackObj fill:#000000,stroke:#ffffff,stroke-width:2px,color:#ffffff
    classDef gcPhase fill:#e1f5fe,stroke:#0277bd,stroke-width:2px
    classDef decision fill:#fff3e0,stroke:#f57c00,stroke-width:2px

    class AllWhite,FreeMemory whiteObj
    class RootGray,MarkRefsGray grayObj
    class MarkBlack,KeepObject blackObj
    class MarkStart,SweepStart gcPhase
    class GrayExists,CheckColor,NextObject decision
```

---

# Object States in Go GC

- **WHITE**: Potentially unreachable objects (candidates for collection)
- **GRAY**: Objects discovered but not yet scanned for references
- **BLACK**: Objects confirmed as reachable and fully scanned

---

problem

many overhead


---

```go

type sample struct{}

func NewValue() sample {
	s := sample{}
	return s // <- stack
}

func NewPointer() *sample {
	s := sample{}
	return &s // <- heap
}


```
