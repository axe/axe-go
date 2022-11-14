package axe

import "github.com/axe/axe-go/pkg/util"

type EntityTree struct {
	entity   *Entity
	parent   *Entity
	children []*Entity
	getTree  func(e *Entity) *EntityTree
}

func NewEntityTree(e *Entity, getTree func(e *Entity) *EntityTree) EntityTree {
	return EntityTree{
		entity:  e,
		getTree: getTree,
	}
}

func (t *EntityTree) Entity() *Entity {
	return t.entity
}

func (t *EntityTree) Parent() *Entity {
	return t.parent
}

func (t *EntityTree) Children() []*Entity {
	return t.children
}

func (t *EntityTree) SetParent(parent *Entity) {
	parentTree := t.getTree(t.parent)
	setParent(t, parentTree)
}

func (t *EntityTree) SetParentTree(parent *EntityTree) {
	setParent(t, parent)
}

func (t *EntityTree) AddChildTree(child *EntityTree) {
	setParent(child, t)
}

func (t *EntityTree) AddChild(child *Entity) {
	childTree := t.getTree(child)
	setParent(childTree, t)
}

func (t *EntityTree) RemoveChild(child *Entity) {
	childTree := t.getTree(child)
	setParent(childTree, nil)
}

func (t *EntityTree) Delete() {
	t.entity.Delete()
	t.DeleteChildren()
}

func (t *EntityTree) DeleteChildren() {
	if t.children != nil {
		for _, child := range t.children {
			childTree := t.getTree(child)
			childTree.Delete()
		}
	}
}

func setParent(tree *EntityTree, parent *EntityTree) {
	if tree.parent != nil {
		parent.children = util.SliceRemove(parent.children, tree.entity)
	}
	tree.parent = parent.entity
	if tree.parent != nil {
		parent.children = append(parent.children, tree.entity)
	}
}
