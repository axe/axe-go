## Statements
- All renderables have current and previous state and their respective game times
- When you type the currently focused object gets key events
- When you scroll the hovered object gets scroll events
- When you have pointer events the
- Input is only handled when a stage is settled and updating

## Questions
- Queue input events and process at a frequency?
- Ignore input events in queue that are irrelevant? (2 mouse moves can be 1)

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