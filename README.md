# axe-go
Axe Game Engine for Go

Features
- [ ] Localization
- [ ] A game engine that can 

### Networking
- [ ] Take inputs for user and update characters & camera
  - There are inputs. This input is transmitted to server and also applied locally.
  - The input is applied to the character state & camera, other entities can be created (ex: bullets)
  - The character state drives the animation
  - 


### Core
- [ ] Predicate
- [ ] Consumer
- [ ] Compare
- [ ] Factory `[T any] Create()`
- [ ] MemoryFactory
- [ ] CanExpire `IsExpired(), Expire()`
- [ ] CanRemove `OnRemove(), Remove()`
- [ ] ExpireCollection `RemoveExpired(), `
- [ ] HasData `Get/SetData(any)`
- [ ] Match
- [ ] Flags
- [ ] FlagsMatch
- [ ] HasFlags `Flags()`
- [ ] Anchor
- [ ] Placement
- [ ] Collection
- [ ] ListArray
- [ ] ListSorted
- [ ] ListLinked
- [ ] ListCircular
- [ ] Memory
- [ ] MemoryPool
- [ ] MemoryHeap
- [ ] Registry

### Geometry
- [ ] Vec
- [ ] Matrix
- [ ] Shape
- [ ] Bounds
- [ ] Range
- [ ] Line
- [ ] Segment
- [ ] Sphere
- [ ] Plane
- [ ] HasShape
- [ ] ShapeCalculator

### Math
- [ ] Plot
- [ ] Randoms
- [ ] Probability
- [ ] Noise
- [ ] NoiseSimplex
- [ ] NoisePerlin
- [ ] Calculator

### Paths
- [ ] X




### Using
- Load preferences
- Initialize systems (audio, graphics, input, window)
- Load/define all stages
  - Stage has
    - assets needed to start
    - window(s) size & placement
    - all worlds/scenes
    - views
  - A stage has to be fully initialized before it can be switched to. The current stage has access to the progress of loading the next scene.
- Set/load starting stage
- Reset graphics & window systems based on current stage's requirements
- Loop
  - Update input system
  - If next stage queued, update status
  - Update current stage
    - Update worlds/scenes
      - Worlds have their own job system and budgets
  - Process graphics
    - Are we ready to draw?
    - Yes, get render state of current stage
      - Based on current update state, or interpolated
    - Render to targets (textures, windows) based on stage views
