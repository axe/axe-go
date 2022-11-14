package ecs

import "github.com/axe/axe-go/pkg/util"

type Tree struct {
	entity   *Entity
	parent   *Entity
	children []*Entity
	getTree  func(e *Entity) *Tree
}

func NewTree(e *Entity, getTree func(e *Entity) *Tree) Tree {
	return Tree{
		entity:  e,
		getTree: getTree,
	}
}

func (t *Tree) Entity() *Entity {
	return t.entity
}

func (t *Tree) Parent() *Entity {
	return t.parent
}

func (t *Tree) Children() []*Entity {
	return t.children
}

func (t *Tree) SetParent(parent *Entity) {
	parentTree := t.getTree(t.parent)
	setParent(t, parentTree)
}

func (t *Tree) AddChild(child *Entity) {
	childTree := t.getTree(child)
	setParent(childTree, t)
}

func (t *Tree) RemoveChild(child *Entity) {
	childTree := t.getTree(child)
	setParent(childTree, nil)
}

func (t *Tree) Delete() {
	t.entity.Delete()
	t.DeleteChildren()
}

func (t *Tree) DeleteChildren() {
	if t.children != nil {
		for _, child := range t.children {
			childTree := t.getTree(child)
			childTree.Delete()
		}
	}
}

func setParent(tree *Tree, parent *Tree) {
	if tree.parent != nil {
		parent.children = util.SliceRemove(parent.children, tree.entity)
	}
	tree.parent = parent.entity
	if tree.parent != nil {
		parent.children = append(parent.children, tree.entity)
	}
}
