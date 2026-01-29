# Go Garbage Collection Workflow

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

## Object States in Go GC

- **WHITE**: Potentially unreachable objects (candidates for collection)
- **GRAY**: Objects discovered but not yet scanned for references  
- **BLACK**: Objects confirmed as reachable and fully scanned

## Phases

1. **Mark Phase**: Traverse object graph, marking reachable objects
2. **Sweep Phase**: Free memory of unmarked (WHITE) objects