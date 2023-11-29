// ux is a package that supplies basic components built on the ui package.
//
// The components in the ux package are structured in the following way:
// - XSettings: common options for the X component that can be specified globally
// - X: the X component options that can be built into a component
// - XBase: the built X component
//
// The options and built component can be passed as children to other components,
// anything that implements HasComponent can be. When passing options (which would
// be the most common use case) it is built by the parent while the parent is being
// built.
//
// Each Kind can have a Template defined, and each component has a kind or can be
// given a custom kind that dictates which template it uses. For example ux defines
// one kind of Button but you can define `const KindFAB ux.Kind = 45` and then pass
// `Kind: KindFAB` to `Button` and also define the FAB template on the theme.
package ux
