### ECS Design

Features:
- Control over entity/component order (for hierarchy). Ex: If you update the transform of a parent, the child entities can update their transform after.
- A hierarchy can be established so if you remove a parent entity all dependent entities are removed as well.
- Common components are ones where most entities have it so it should be stored efficiently.
- Sparse components are for components less frequently used where the entity needs to store where in the component data it's value lies.
- Shared components are components where there are a few different instances shared between many entities. Ex: A collider component could share volume and behaviors for similar entities like a bullet.
- Fixed entities have a known set of components at creation and won't have any components added or removed. Allows for performance improvements.
- Component system can iterate through component instances on update.
- Entity system can iterate through entity instances on update.
- Systems define which components they care about, so when it comes to processing all systems we can decipher which ones we need to run first and which we can run in parallel. Component systems that are marked isolated can be processed in parallel first. After that non-intersecting groups are formed and executed in parallel until all systems are handled. Or the world can be marked synchronous meaning no parallel logic is ran. Systems can also be given a priority which is respected above the parallel logic.





World
  Component[C]
  Type[T]
    TypeComponent[T,C]
      get(t) c
  DataStorage[D]

  Data[D]



