## Statements
- All renderables have current and previous state and their respective game times
- When you type the currently focused object gets key events
- Input is only handled when a stage is settled and updating

## Questions
- Queue input events and process at a frequency?
- Ignore input events in queue that are irrelevant? (2 mouse moves can be 1)

## Language rules
- word:         `[a-z][_a-z0-9]*`
- string:       `"[^"]*"`
- whole:        `[-+]?[0-9,]+`
- decimal:      `[-+]?[0-9,]*\.`
- bool:         `(true|false)`
- args:         `{value}(  ,  {value})*`
- argsNamed:    `{name:word}  =  {a:value}(  ,  {name:word}  =  {a:value})*`
- param:        `{name:word} {type:word}`
- params:       `{param}(  ,  {param})*`
- func:         `\(  {params}  \)->\{  {body}  \}`
- expr:         `{token:word}(\.  {next:expr}  |\(  \)|\(  {args}  \)|\(  {argsNamed}  \))?`
- value:        `({string}|{whole}|{decimal}|{expr}|{bool}|{func})`
- define:       `var {name:word} {type:word}? (= {value})?`
- if:           `if {condition:value} then {body} end`
- while:        `while {condition:value} then {body} end`
- each:         `each {item:word} in {collection:expr} then {body} end`
- set:          `{settable:expr}  =  {value}`
- return:       `return(  {value})?`
- body:         `(({define}|{if}|{expr}|{while}|{each}|{set}|{return})  )*`

## Language rules rules
- `[]` = characters within are valid for matching, if a hyphen is between two characters in a sequence its a range
- `.` = any character
- `*` = 0 or more for previous token
- `+` = 1 or more for previous token
- `?` = 0 or 1 for previous token
- `{min,max}` = min to max for previous token
- `(a|b|c)` = one of the options
- `{refName:ruleName}` = a rule has a subset rule that's named. if the name is reused in a rule it is a list of values
- `{ruleName}` = a rule has a subset rule that's named the ruleName
- `\` = escapes [](){}*+?.
- ` ` = 1 or more whitespace characters
- `  ` = 0 or more whitespace characters
- any other characters are considered a literal

## Debugging

- Memory/pools should all be registered in a global list (particles, UI, etc) to be watched. 
  - In debug mode memory/pools could support dynamical reallocation
  - In debug mode memory/pool information is viewable at runtime and can print out at the end (avg, min, max, std dev)
- Any path that points to an object or property can be used to watch a value.
  - This watcher can be started and stopped at any time.
  - This watcher if watching a numeric value can use a fixed size round-robin store that describes the stats. This can be visualized as a chart.
  - The watcher can be set to print out the current value each frame or at a frequency to somewhere (screen, log, etc)
  - The value being watched may not exist (could be game/play/bullets/0 which may only exist when at least one bullet exists)
- Debug events can be scattered throughout the code surrounding anything where elapsed time should be tracked.
- Logs have levels, and categories.
- Debug entities can have draw logic and types. This can be text above unit heads, bounding boxes, etc. This is a component that can be attached to an entity.
- Maybe each component has debug related data for inspection & rendering? Maybe the type system also has a debug render logic. So the bounds might have a debug renderer. Or a character can have debug render attached to it. Maybe you can see all things that has debug rendering and you can turn on/off them all or specific ones.
- Debug rendering:
  - Depth testing is disabled
  - Line, sphere, AABB, OBB, cross, circle, axis, text, triangle
  - Should only render X number of close/appliable renderables to avoid too many. 
- In game console
  - run commands (custom and engine provided)
    - inspect data values
    - set values
    - hide/show graphs
    - what data values should be watched
  
## Memory

- We need to track memory usage and use it to drive how often things are unloaded for new requests
- This can be divided up by type or intermixed. For example, if 1gb is our memory limit we could assign 100mb for audio.
- A world can be broken up into chunks where chunks contain entities that are loaded when the chunk is requested. Automatic chunk loading and unloading can be provided given the known views into the world and bounds that describe the chunks. 
- Maybe there's a chunk component on an entity which stores a chunk ID - so an entity can move between chunks and when it's time to unload a chunk we get all entities with that ID and remove them and save their current state. We also would have hooks for entities when they are loaded, reloaded, or saved in the chunk loading processes.
- Entities references assets that may or may not be loaded so loading and unloading chunks may affect which assets are loaded. 

## Clocks

- Defined timers/timelines so updatable things can be easily controlled. 
- There is a real-time clock that is in line with the computer's clock. 
  - The UI is typically tied to this clock. 
- There is a game-time default clock that can be used to pause things in a game. 
- Clocks can also have separate dates. 
- A clock can have a delta time which is always in real-world seconds but is scalable. 
- There is an elapsed time that matches the speed of the clock, where something like an hour can elapse in 3 real-world minutes. 
- Each updatable object can have which clock they use.
- Each updatable object type can have a default clock they use.
- Time based conditions work on the clock date & time and not delta-time (s). These conditions can say "at 9am trigger" and the 9am is in clock time.
- Clocks can have their date & time set or frozen. But the clock can still "run" (update the delta-time).
- Clocks could have a max elapsed time - so during slow frames they don't "jump" forward in time too much. (governing)
- Clocks for elapsed time could use a running average that over time yields correct elapsed times. This solves the issue where we update the current frame based on the duration of the last time. Doing that can resolve the issues that arrive when we double update something which causes us to continually do that the following frames thinking we need to.

## Level-of-detail

Can be applied to anything, not just meshes & textures.
Given a view (camera + scene) and an object with LOD states return which LOD state should be used.
A LOD state can be an asset or generated. Precreated or on demand.
A LOD state could be:
- A mesh to render with
- A texture to render with
- An AI implementation to use
- Audio to play
- Networking strategy to use
  - ex: only sync object position and not animation state
  - ex: use 2 bytes instead of 4

## Assets

An asset can be loaded on certain events:
- Start of game
- Before scene starts
- After scene starts
- When it needs to be used (may cause "wait" for rendering/use)
- When it could be used soon (it belongs to an entity that is probably going to be within view)

An asset can be referenced by multiple things going through their life cycles. When an asset is loaded during runtime each entity using the asset will have a reference counted value so the asset knows when it's being used.

An asset can be unloaded on certain events:
- End of game
- After scene ends
- Before scene ends
- When it is no longer referenced
- When it is no longer referenced and X has elapsed
- When it is no longer referenced and it doesn't appear it will be in the near future

An asset load group can exist in space and when it collides with the asset volume oriented in each view's camera and cause assets to be loaded if they are not already.

Asset manager keeps track of memory usage and can have a max set and when it needs to load an asset it can unload lower priority assets. Priority is lowest if the asset is not required right now. Or the objects using it are all far away.

Some resources (videos and long playing audio) are streamed into memory.

Assets can be marked for certain platforms. Ex: you use a simpler model for the phone version of the game vs console vs PC.

Maybe a given number of assets are related to an entity - and certain assets may keep the entity held back from being marked as live. Maybe each component needs IsReady logic?

## Jobs

- At the start of each frame each game system is asked to queue up jobs.
- A job...
  - Has a type & ID
  - Has a min & max wait
  - Can be profiled to calculate real costs
  - Can have a min & max frequency
  - Can be notified on certain events (run, cancel, expired, miss)
  - Can have a next job that is automatically ran or queued after a job instance runs.
  - Has a before and after group
  - Has a before and after type
  - Has a before and after ID
- A job can have dependencies (to certain job types or job instances) that must be finished before it completes
- A job can have dependents - so it has to run before other types/instances.
- A job can have a lifetime
  - It is removed on completion
  - It is removed at end of frame
  - It is not removed
  - It is removed after X seconds of not being picked up
  - It is cancelled/replaced
- A job can have an estimated cost
  - If the job system has a targeted cost per frame this can be used to determine what it can consider
- A job can have a priority
  - This is the first thing used to prioritize it in it's group
- A job system can use dependencies to dynamically calculate groups (topological sort) or can use defined group orders.
- A job system can be sequential or concurrent.
- The job system will...
  1. Collect new jobs
  2. For applicable jobs add to active list
  3. Add all required jobs to frame jobs
  4. Sort applicable jobs and grab highest priority to add to frame jobs until budget is met. Same priority jobs can be compared based on which one is most over-due to run.
  5. Assign groups if auto-group is enabled
  6. Executing jobs one group at a time. If concurrent, the jobs in a group can be executed at the same time.
  7. A system (like physics) can define multiple groups to control ordering and each component instance can keep a group ID or the data for each group can be stored separately.

## Steering
Concepts
- Filter given a subject and potential object returns true or false
- Force: given a subject and dt - returns a velocity and weight
- Behavior: applies a force to a subject
- Accumulator: a force which is built from multiple sources
- Spatial: given a query, filter, and options - iterate over objects and 

Forces: `(steer.Subject, float32) steer.Force`
- Path (towards future point along path)
- Alignment (average direction - maybe weighted by something? (ex:distance))
- Arrive (end up at target)
- Avoid Obstacles (lookahead and if any intercept, move tangentially)
- Away (opposite direction of target)
- Average/Cohesion (towards average spot - maybe weighted?)
- Constant (given force)
- Containment (stay inside shape)
- Direction (towards direction)
- Dodge (avoid entities subject may intercept)
- Drive (thrusting, breaking, turning)
- Face (subject faces target)
- Flow (field of forces - closest one is used)
- Follower (towards a an area relative to another entity)
- Match (average velocity of nearby - maybe weighted by something?)
- Separation (away from entities around it)
- To (towards target)
- Wander (towards point on imaginary circle away from subject where the angle randomly changes each update)
- Modifier (another force if enabled, scaled, and can update at a frequency, can have max & min magnitude)

Accumulators:
- Average (of all forces)
- Context (X number of slots around subject where each might hold a force that's good(+) or bad(-) - prioritize the bad things?)
- First (first force)
- Max (accumulate forces up to a certain magnitude, ignore the others)

Constraints:
- Turning (max rotation per second)
- Zero Velocity Threshold (if velocity is this near to zero, consider it zero)
- Dual (two constraints)

Filter: `(steer.Subject, steer.Entity) bool`
- And
- Or
- InFront (if entity is in front/back of subject)
- Proximity (if entity is with-in/out of a certain ranges)
- View (FOV)

Targets: `(steer.Subject) steer.Entity`
- Average (average position of entities in space that pass a test)
- Chain (looks at one target, if none return looks at another)
- Closest (closest valid entity in space)
- Facing (another target if subject is in front or back of it)
- Filtered (only if another target matches a condition)
- Future (towards the interception point of the subject and another target)
- InLine (anywhere between two targets)
- Interpose (X% between two targets)
- Local (only if another target is within X)
- Offset (relative to another target - add possibly rotated vector)
- Relative (process relative to another subject, not the main one)
- Slowest (query space - one with smallest velocity)
- Weakest (query space - one it could intercept the soonest)

user defines subject state type
user defines complete behavior template that uses a specific state tpe. when a subject's behavior is set the template generates an instance. it may be shared if it's entirely stateless, or it may be stateful.
subject state is a generic float array used to store data. scripting can set values by defined names & types.

- force = direction & magnitude, direction established by two points. magnitude can be distance, scaled, negated, etc
- output of system is a force, something translates that into an animation or something


## Saving
Saving can happen on request or actively during game play.
It can be tied to just the game install, the profile locally, the profile remotely, or a specific world/level.
Some savable values can natively support dirty checking so saving is only done when a value changes OR changes can be saved on a frequency. The developer should be able to specify what objects, properties, etc of a level are savable.

## Serialization
As much game configuration, data, and logic could be encoded in files.

```xml
<game name="Test GLFW" enable-debug="true" fixed-draw-frequency="166.6ms" first-stage="cube">
    <world-settings 
        entity-capacity="2048" 
        entity-stage-capacity="128" 
        average-component-per-entity="4" 
        delete-on-destroy="true"
    />
    <windows>
        <window title="Test GLFW Main Window" placement="centered(720,480)">
            <placement type="centered" width="720" height="480" />
        </window>
    </windows>
    <stages>
        <stage name="cube">
            <assets>
                <asset name="cube model" uri="cube.obj" />
            </assets>
            <actions>
                <main>
                    <close type="key" key="escape" />
                    <down type="key" key="z" />
                    <undo type="key" key="z" cmdctrl="true" />
                    <pasteUp type="key" key="v" ctrl="true" up-only="true" />
                    <pressA type="key" key="a" press-interval="250ms" first-press-delay="1s" />
                    <logInput type="key" key="c" />
                    <delete type="key" key="backspace" />
                </main>
            </actions>
            <views3></views3>
            <scenes3>
                <enable require="game,scene,ecs">
                    scene.world.enable(
                        ecs.dataSettings(capacity=2048, stageCapacity=128),
                        // components
                        TAG, MESH, TRANSFORM3, AUDIO, ACTION, LIGHT, LOGIC, INPUT
                    )
                </enable>
                <load require="game,scene,ecs,debug">
                    var e = ecs.new()
                    var logInput = false

                    TAG.set(e, "cube")
                    MESH.set(e, mesh(ref="cube model"))
                    TRANSFORM3.set(e, newTransform4(
                        position = vec4f(x=0, y=0, z=-3, w=0),
                        scale = vec4f(x=1, y=1, z=1, w=0)
                    ))
                    LOGIC.set(e, (e, ctx)->{
                        var dt = game.state.updateTimer.elapsed.seconds
                        var transform = TRANSFORM3.get(e)
                        var rot = transform.rotation

                        rot.x = rot.x + dt*6
                        rot.y = rot.x + dt*4
                        transform.rotation = rot
                    })
                    LIGHT.set(e, light(
                        diffuse = colorf(r=1, g=1, b=1, a=1),
                        ambient = colorf(r=0.5, g=0.5, b=0.5, a=1),
                        position = vec4f(x=-5, y=5, z=10)
                    ))
                    ACTION.set(e, (action)->{
                        if action.name=("close") then
                            game.running = false
                        else if action.name=("logInput") then
                            logInput = !logInput
                        else if action.name=("delete") then
                            e.delete()
                        else
                            var inputNames = list()
                            if action.data.inputs then
                                loop action.data.inputs(in)
                                    inputNames.add(in.name)
                                end
                            end
                            debug.print("%s %0.1f (priority=%d, inputs=%s)", action.name, action.data.value, action.priority, inputNames.join(","))
                        end
                        return true
                    })
                    INPUT.set(e, input.systemEvents(
                        inputChange = (input)->{
                            debug.print("%s changed to %v\n", input.name, input.value)
                        }
                    ))
                </enable>
            </scenes3>
        </stage>
    </stages>
</game>
```


## Loop
- Get input
- Update time system
- Update stage system (intro, outro, update)
- Update current settled stage
    - Handle input
    - Pre-update cameras
    - Update scenes
        - Do work based on schedule, priority, availability, etc
            - Animations
            - Assets
            - Logic
            - Systems
            - Space (moved spatial entities are reindexed)
            - Networking (read & write packets, queue data)
            - Particles
            - Audio
            - Physics
            - Steering Behaviors
            - Path Finding
            - Navigation
            - State Machines
            - Scripts
            - Custom
    - Post-update cameras
- Render?
    - Update cameras
    - For each window, for each view
        - Render current state or interpolate

