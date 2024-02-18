# axe-go
Axe Game Engine for Go

**Important aspects of Axe:**
- [Readable](#readable)
- [Editor](#editor)
- [Debugging](#debugging)
- [Modding](#modding)
- [Scripting](#scripting)
- [Level of detail](#level-of-detail)
- [Multi-dimensional](#multi-dimensional)
- [Real-time](#real-time)

### Readable
Data, configuration, & assets are stored in a human readable format for development. 
Storing everything in a human readable format accompanied with using a file versioning system (ie Git) allows for tracking changes and merging in a team environment. 
Tracking changes to specific builds is essential for tracking when bugs are introduced.
When an artifact is built data is compiled to a binary format.

### Editor
The editor is an integral part of the development experience. 
You define your screens, configuration, import assets, build your worlds and levels, and can extend the native types and add custom widgets to the editor.
When any external file changes the editor should be able to update when the developer wants it.

### Debugging
An assortment of debugging tools is essential for a real-time application. Data is constantly changing and you need tools to visualize, profile, record, and replay it for debugging and testing purposes. Being able to do so against any parts of the game data is essential to delivering an accurate and stable game!

### Modding
If a game developer wants their game to be moddable the engine should support it to whatever extent they want.

### Scripting
Scripting is a quick way to introduce custom logic that can be more easily changed during development and also supports moddability if desired.

### Level of detail
Using different assets or behaviors in different scenarios should be supported on everything. Traditional LOD for meshes and textures is supported as well as at the network level, AI, and audio to mention a few.

### Multi-dimensional
2d & 3d are equally prioritized.

### Real-time
Games are real-time applications, throughout play the requirements on the hardware can change drastically. A good game engine allows you to set priorities and requirements for all the logic that needs ran and visuals that need to be displayed.
