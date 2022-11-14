# axe-go
Axe Game Engine for Go

### Use
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
