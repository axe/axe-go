package axe

import (
	"context"
	"fmt"
	"time"
)

type SQL string

type Order struct {
	Field string
	Desc  bool
}

type Results[V any] interface {
	Offset() uint
	Limit() uint
	Total() uint
	Next() (V, error)
	Close()
}

type QueryLogic interface {
	BuildQuery(q SQL) (SQL, error)
}

type Query struct {
	Query  QueryLogic
	Offset *uint
	Limit  *uint
	Order  []Order
}

type Storage[E any, PK comparable] interface {
	Get(ctx context.Context, pk PK) (E, error)
	Create(ctx context.Context, e E) (PK, error)
	Update(ctx context.Context, e E, pk PK) error
	Delete(ctx context.Context, pk PK) error

	Search(ctx context.Context, q Query) (Results[E], error)
	First(ctx context.Context, q Query) (*E, error)
	All(ctx context.Context, q Query) ([]E, error)
	Exists(ctx context.Context, q Query) (bool, error)
}

type Table struct{}

type ORMView struct {
	Fields []ORMViewField
}
type ORMViewField struct {
	Field   string
	SubView string
}

// User{ID: 23}.Load(ctx, UserViewEmployee)
type User struct {
	Table    `t:"Users" pk:"UserID"`
	ID       string             `c:"UserID"`
	Name     string             `c:"EmpName" v:"*"` // on all views
	Email    string             `c:"Email" v:"withEmail,withEmployee"`
	Employee OneToOne[Employee] `c:"EmpID" v:"withEmployee(basic)"` // just on this view
}
type Employee struct {
	Table    `t:"Employees" pk:"EmpID"`
	ID       string             `c:"EmpID"`
	Name     string             `c:"EmpName" v:"*"`
	DOB      time.Time          `c:"DOB" v:"basic"`
	User     OneToOne[User]     `fc:"UserID"`
	Contacts OneToMany[Contact] `fc:"EmpID"`
}
type Contact struct {
	Table    `t:"Contacts" pk:"ContactID"`
	ID       string              `c:"ContactID"`
	Name     string              `c:"ContactName"`
	Employee ManyToOne[Employee] `c:"EmpID"`
}

type ManyToOne[V any] struct {
	instance *V
	value    any
	column   string
}

func (fk ManyToOne[V]) Get() V {
	return *fk.instance
}

type OneToOne[V any] struct {
	instance      *V
	value         any
	column        string
	foreignColumn string
}

func (fk OneToOne[V]) Get() V {
	return *fk.instance
}

type OneToMany[V any] struct {
	values        []V
	foreignColumn string
}

func (m OneToMany[V]) Get() []V {
	return m.values
}

type Store struct{}

func Load[M any](s *Store, model M, view string) M { return model }

func ORM() {
	s := &Store{}
	u := Load(s, User{ID: "73738"}, "withEmployee")

	fmt.Println(u.Employee.Get().Name)
}
