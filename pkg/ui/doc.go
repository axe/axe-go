package ui

// In this package there contains elements to render a UI
//
// Concepts
// - Component		an interactable element like a button, textbox, checkbox, etc
// - Layer			a component can have any number of layers that make up the visual like shadows, borders, images, text, icons, etc
//   - Border		a line drawn around an shaped area
//   - Shadow		an inner or outer border with a gradient
//   - Solid		a solid background (color, texture, gradient) with
// - Background		a layer can be a color, textured, a gradient, etc
// - Placement      describes how a layer or component is placed in its parent component and reacts to resizing
// -
//

// TODO
// - Dirty and only place & rerender when dirty
//   - Dirty can happen on resize or when a component is marked dirty (manually or any variable changes)
// - Add max size to component
// - Add min size to component
// - Add margin & padding
// - Add cursors
// - Add popovers
// - Component templates
// - Styles that trickle down during render (pass Theme, Disabled, etc)
