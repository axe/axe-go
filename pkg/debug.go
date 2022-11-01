package axe

import (
	"fmt"
	"time"

	"github.com/axe/axe-go/pkg/ui"
)

type DebugSystem struct {
	Logs      []DebugLog
	Snapshots []DebugSnapshot
	Graphs    []DebugGraph
	Events    []DebugEvent
	Enabled   bool
}

var _ GameSystem = &DebugSystem{}

func (ds *DebugSystem) Trigger(ev DebugEvent) {}
func (ds *DebugSystem) Init(game *Game) error { return nil }
func (ds *DebugSystem) Update(game *Game)     {}
func (ds *DebugSystem) Destroy()              {}

func (ds *DebugSystem) LogError(err error) {
	// TODO log in DebugLog
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

type DebugLog struct {
	Severity int
	Message  string
	Data     any
}
type DebugSnapshot struct {
	At      time.Time
	Elapsed time.Duration
}
type DebugGraph struct {
	Placement ui.Placement
	Database  *StatDatabase
	Set       *StatSet
}
type DebugProfiler struct{}

func (prof *DebugProfiler) Begin(ev string) {}
func (prof *DebugProfiler) End()            {}

type StatDatabase struct {
	Name    string
	Updated time.Time
	Visible bool
	Event   *DebugEvent
	Enabled bool
	Sets    []StatSet
}
type StatPoint struct {
	Total int
	Sum   float64
	Min   float64
	Max   float64
}
type StatSet struct {
	Index        int
	Description  string
	Interval     time.Duration
	PointerTime  time.Time
	PointerIndex int
	Points       []StatPoint
}
type DebugEvent struct {
	Id       int
	Name     string
	Parent   *DebugEvent
	Depth    int
	Children []*DebugEvent
	Sibling  *DebugEvent
}
