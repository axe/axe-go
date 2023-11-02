package ui

type Flags uint64

func (f Flags) Any() bool                    { return f != 0 }
func (f Flags) None() bool                   { return f == 0 }
func (f Flags) Exactly(other Flags) bool     { return f == other }
func (f Flags) Is(other Flags) bool          { return (f & other) != 0 }
func (f Flags) Not(other Flags) bool         { return (f & other) == 0 }
func (f Flags) All(other Flags) bool         { return (f & other) == other }
func (f Flags) NotAll(other Flags) bool      { return (f & other) != other }
func (f Flags) WithRemove(other Flags) Flags { return f & ^other }
func (f Flags) WithAdd(other Flags) Flags    { return f | other }
func (f *Flags) Remove(other Flags)          { *f = *f & ^other }
func (f *Flags) Add(other Flags)             { *f = *f | other }
