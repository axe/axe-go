package script

import (
	"errors"
	"fmt"
	"strings"
)

var ErrEndOfInput = errors.New("end of input")
var ErrReadToEnd = errors.New("read to fail")
var ErrCharsNoEnd = errors.New("missing ] at end of chars")
var ErrRangeNoDelimiter = errors.New("missing , in range")
var ErrRangeNoEnd = errors.New("missing } at end of range")
var ErrRangeInvalid = errors.New("range has invalid numbers, min is optional but must be 0 or more and max is optional and must be more than or equal to min")
var ErrUnexpectedToken = errors.New("unexpected token")
var ErrRuleNoEnd = errors.New("missing > at end of rule")
var ErrOneOfNoEnd = errors.New("missing ) at end of one of")
var ErrOneOfEnd = errors.New("one of end")
var ErrOneOfDelimiter = errors.New("one of delimiter")

// ## Language rules rules
// - `[]` = characters within are valid for matching, if a hyphen is between two characters in a sequence its a range
// - `.` = any character
// - `\` = escapes [](){}*+?.^$!
// - `\shortcut` = is a matcher in a shortcut map in the rule set
// - `^` = start of line
// - `$` = end of line
// - `*` = 0 or more of previous matcher
// - `+` = 1 or more of previous matcher
// - `?` = 0 or 1 of previous matcher
// - `{n}` = n of previous matcher
// - `{min,max}` = min to max of previous matcher. min defaults to 0, max defaults to inf.
// - `(a|b|c)` = one of the options
// - `(?=a)` = matches if a matches and doesn't read it
// - `(?!a)` = matches if a does not match and doesn't read it
// - `(?:a)` = matches if a matches but doesn't capture it
// - `\n` = matches group "n" captured in rule
// - `<refName:ruleName>` = a rule has a subset rule that's named. if the name is reused in a rule it is a list of values
// - `<ruleName>` = a rule has a subset rule that's named the ruleName
// - any other characters are considered a literal
type RuleTokenKind rune

var (
	RuleTokenEOF                  RuleTokenKind = 0
	RuleTokenEscape               RuleTokenKind = '\\'
	RuleTokenAny                  RuleTokenKind = '.'
	RuleTokenRepeatMaybe          RuleTokenKind = '*'
	RuleTokenRepeat               RuleTokenKind = '+'
	RuleTokenRepeatRangeStart     RuleTokenKind = '{'
	RuleTokenRepeatRangeEnd       RuleTokenKind = '}'
	RuleTokenRepeatRangeDelimiter RuleTokenKind = ','
	RuleTokenMaybe                RuleTokenKind = '?'
	RuleTokenLiteral              RuleTokenKind = 'l'
	RuleTokenShortcut             RuleTokenKind = 's'
	RuleTokenOneOfStart           RuleTokenKind = '('
	RuleTokenOneOfEnd             RuleTokenKind = ')'
	RuleTokenOneOfDelimiter       RuleTokenKind = '|'
	RuleTokenCharsStart           RuleTokenKind = '['
	RuleTokenCharsEnd             RuleTokenKind = ']'
	RuleTokenCharsRange           RuleTokenKind = '-'
	RuleTokenCharsNot             RuleTokenKind = '^'
	RuleTokenRuleStart            RuleTokenKind = '<'
	RuleTokenRuleEnd              RuleTokenKind = '>'
	RuleTokenRuleDelimiter        RuleTokenKind = ':'
	RuleTokenLineStart            RuleTokenKind = '^'
	RuleTokenLineEnd              RuleTokenKind = '$'
)

func RuneToRuleTokenKind(k rune) RuleTokenKind {
	v := RuleTokenKind(k)
	switch v {
	case RuleTokenEscape,
		RuleTokenAny,
		RuleTokenRepeatMaybe,
		RuleTokenRepeat,
		RuleTokenRepeatRangeStart,
		RuleTokenRepeatRangeEnd,
		RuleTokenRepeatRangeDelimiter,
		RuleTokenMaybe,
		RuleTokenOneOfStart,
		RuleTokenOneOfEnd,
		RuleTokenOneOfDelimiter,
		RuleTokenCharsStart,
		RuleTokenCharsEnd,
		RuleTokenCharsRange,
		RuleTokenCharsNot,
		RuleTokenRuleStart,
		RuleTokenRuleEnd,
		RuleTokenRuleDelimiter,
		RuleTokenLineStart,
		RuleTokenLineEnd:
		return v
	default:
		return RuleTokenLiteral
	}
}

func StdShortcuts() map[rune]Matcher {
	return map[rune]Matcher{
		'f': matchCharsFromString("\f", -1),
		'n': matchCharsFromString("\n", -1),
		'r': matchCharsFromString("\r", -1),
		't': matchCharsFromString("\t", -1),
		'v': matchCharsFromString("\v", -1),
		'0': matchCharsFromString("\x00", -1),
		's': matchCharsFromString(" \f\n\r\t\v\u00A0\u2028\u2029", -1),
		'S': matchCharsFromString("^ \f\n\r\t\v\u00A0\u2028\u2029", -1),
		'w': matchCharsFromString("a-zA-Z0-9_", -1),
		'W': matchCharsFromString("^a-zA-Z0-9_", -1),
		'd': matchCharsFromString("0-9", -1),
		'D': matchCharsFromString("^0-9", -1),
	}
}

type RuleToken struct {
	Kind RuleTokenKind
	Rune
}

type LogFn func(string, ...any)

type Logger struct {
	OnGetGreedyMatches      func(input *RuneBuffer, out *RuleMatch, repeater, repeating Matcher, repeatMin, repeatMax, greedyMatches int, start Position)
	OnSetGreedyMatches      func(input *RuneBuffer, out *RuleMatch, repeater, repeating Matcher, repeatMin, repeatMax, greedyMatches int, start Position)
	OnDecreaseGreedyMatches func(input *RuneBuffer, out *RuleMatch, resetToMatcher Matcher, resetTo int, greediness int)
	OnMatchError            func(input *RuneBuffer, out *RuleMatch, m Matcher, err error)
	OnMatch                 func(input *RuneBuffer, out *RuleMatch, m Matcher, start Position)
	OnNotMatch              func(input *RuneBuffer, out *RuleMatch, m Matcher, start Position)
}

func (l *Logger) Quiet(greedy, matchError, notMatch, match bool) {
	if greedy {
		l.OnGetGreedyMatches = nil
		l.OnSetGreedyMatches = nil
		l.OnDecreaseGreedyMatches = nil
	}
	if matchError {
		l.OnMatchError = nil
	}
	if notMatch {
		l.OnNotMatch = nil
	}
	if match {
		l.OnMatch = nil
	}
}

type RuleSet struct {
	rules     map[string]Rule
	matcherID int

	Shortcuts map[rune]Matcher
	Logger    Logger
}

func NewStdRuleSet(rules map[string]RulePattern) (RuleSet, error) {
	return NewRuleSet(rules, StdShortcuts())
}

func NewRuleSet(rules map[string]RulePattern, shortcuts map[rune]Matcher) (RuleSet, error) {
	rs := RuleSet{
		rules:     make(map[string]Rule),
		matcherID: 1,
		Shortcuts: shortcuts,
	}
	for ruleName, ruleInput := range rules {
		parsed, err := ruleInput.Parse()
		if err != nil {
			return rs, err
		}
		rule := Rule{
			name:       ruleName,
			shortcuts:  shortcuts,
			matcherID:  &rs.matcherID,
			ignoreCase: parsed.IgnoreCase,
			greedy:     parsed.Greedy,
			ignore:     parsed.Ignore,
		}
		err = rule.Parse(parsed.Pattern)
		if err != nil {
			return rs, err
		}
		rs.rules[strings.ToLower(ruleName)] = rule
	}
	return rs, nil
}

func NewFuncLogger(log LogFn) Logger {
	return Logger{
		OnSetGreedyMatches: func(input *RuneBuffer, out *RuleMatch, repeater, repeating Matcher, repeatMin, repeatMax, greedyMatches int, start Position) {
			log("%s: %d greedy matches for %s found at %v\n", out.Path(), greedyMatches, repeating.String(), start)
		},
		OnGetGreedyMatches: func(input *RuneBuffer, out *RuleMatch, repeater, repeating Matcher, repeatMin, repeatMax, greedyMatches int, start Position) {
			log("%s: greedy match %s %d found at %v\n", out.Path(), repeating.String(), greedyMatches, start)
		},
		OnDecreaseGreedyMatches: func(input *RuneBuffer, out *RuleMatch, resetToMatcher Matcher, resetTo, greediness int) {
			log("%s: no match, greedy match of %d found at %d - resetting to matcher %v at %v\n", out.Path(), greediness, resetTo, resetToMatcher, input.Pos())
		},
		OnMatchError: func(input *RuneBuffer, out *RuleMatch, m Matcher, err error) {
			log("%s: %s errored: %v\n", out.Path(), m, err)
		},
		OnMatch: func(input *RuneBuffer, out *RuleMatch, m Matcher, start Position) {
			matched := input.Get(start.Index, input.Pos().Index)
			log("%s: %s matched at %v '%s'\n", out.Path(), m, input.Pos(), matched)
		},
		OnNotMatch: func(input *RuneBuffer, out *RuleMatch, m Matcher, start Position) {
			log("%s: %s not matched at %v\n", out.Path(), m, input.Pos())
		},
	}
}

func NewStdLogger() Logger {
	return NewFuncLogger(func(s string, a ...any) {
		fmt.Printf(s, a...)
	})
}

func (rs RuleSet) Rule(name string) *Rule {
	rule, exists := rs.rules[strings.ToLower(name)]
	if !exists {
		return nil
	}
	return &rule
}

func (rs RuleSet) Parse(input string) (*RuleMatch, error) {
	for ruleName := range rs.rules {
		match, _ := rs.ParseRule(ruleName, input)
		if match != nil {
			return match, nil
		}
	}
	return nil, fmt.Errorf("no matched rule found")
}

func (rs RuleSet) ParseRule(ruleName string, inputText string) (*RuleMatch, error) {
	rule := rs.Rule(ruleName)
	if rule == nil {
		return nil, fmt.Errorf("rule %s does not exist", ruleName)
	}

	runes := NewRuneBuffer(inputText)
	input := &Input{
		Matchers: make([]MatcherMatch, runes.Len()+1),
		Runes:    runes,
		cycles:   make(Set[cycleState]),
	}

	match, err := rule.Match(rs, runes, input, nil)
	if err != nil {
		return nil, err
	}
	if match == nil {
		last := runes.Max().Index - 1
		return nil, fmt.Errorf("unexpected '%s' at %v", string(inputText[last]), runes.Max())
	}

	return match, nil
}

func (rs RuleSet) loggerFor(input *RuneBuffer, out *RuleMatch, m Matcher, start Position) func(bool, error) (bool, error) {
	return func(matched bool, err error) (bool, error) {
		if err != nil {
			if rs.Logger.OnMatchError != nil {
				rs.Logger.OnMatchError(input, out, m, err)
			}
		} else if matched {
			if rs.Logger.OnMatch != nil {
				rs.Logger.OnMatch(input, out, m, start)
			}
		} else {
			if rs.Logger.OnNotMatch != nil {
				rs.Logger.OnNotMatch(input, out, m, start)
			}
		}
		return matched, err
	}
}

type Position struct {
	Index  int
	Column int
	Line   int
}

func (p Position) String() string {
	if p.Line == 0 {
		return fmt.Sprintf("(i=%d)", p.Index)
	}
	return fmt.Sprintf("(i=%d, line=%d, col=%d)", p.Index, p.Line, p.Column)
}

type Matcher interface {
	String() string
	ID() int
	Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error)
}

type MatchResult struct {
	Matched bool
	// A matcher that matched can have this number of alternative states.
	// For a repeating matcher this represents how many greedy matches over the
	// minimum it had. For a matcher with multiple options this represents the
	// number of options it never tried. This can be passed down to the Match
	// method with a decreased amount to force the matcher to try the alternatives
	// because the original match was not
	Alternatives *int
	// The children of this match result.
	Children []MatchResult
}

func (mr *MatchResult) ChildAt(i int) *MatchResult {
	if mr == nil || i < 0 || i >= len(mr.Children) {
		return nil
	}
	return &mr.Children[i]
}

func (mr *MatchResult) SetChild(i int, child MatchResult) {
	oldSize := len(mr.Children)
	newSize := i + 1
	if newSize > oldSize {
		mr.Children = append(mr.Children, make([]MatchResult, newSize-oldSize)...)
	} else if newSize < oldSize {
		mr.Children = mr.Children[:newSize]
	}
	mr.Children[i] = child
}

func (mr MatchResult) WithAlternatives(alternatives int) MatchResult {
	copy := mr
	copy.Alternatives = &alternatives
	return copy
}

func (mr *MatchResult) TakeAlternative() bool {
	if len(mr.Children) > 0 {
		for i := range mr.Children {
			if mr.Children[i].TakeAlternative() {
				return true
			}
		}
	}
	if mr.Alternatives != nil && *mr.Alternatives > 0 {
		*mr.Alternatives--
		return true
	}
	return false
}

var (
	MatchYes = MatchResult{Matched: true}
	MatchNo  = MatchResult{Matched: false}
)

type Range struct {
	Start Position
	End   Position
}

func newRange(start Position, end Position) Range {
	return Range{
		Start: start,
		End:   end,
	}
}

func (c Range) Get(in RuneBuffer) string {
	return in.Get(c.Start.Index, c.End.Index)
}

func (c Range) GetRunes(in RuneBuffer) []rune {
	return in.GetRunes(c.Start.Index, c.End.Index)
}

func (c Range) String(in string) string {
	return in[c.Start.Index:c.End.Index]
}

// A rule matched over some range of input
type RuleMatch struct {
	// The range over what this matched.
	Range
	// The alias/rule that created this node.
	Name string
	// The rule that was matched
	Rule *Rule
	// The matchers that matched in this rule
	Matchers []MatcherMatch
	// The captured groups in this rule
	Captures map[int][]Range
	// The parent of this node.
	Parent *RuleMatch
	// The children of this node as they occurred in the input.
	Children []*RuleMatch
	// The name-to-output nodes
	ChildIndex map[string][]int
	// Global input dat
	Input *Input
}

func newRuleMatch(rule *Rule, input *Input, parent *RuleMatch) *RuleMatch {
	return &RuleMatch{
		Name:       rule.name,
		Rule:       rule,
		Captures:   make(map[int][]Range),
		Children:   make([]*RuleMatch, 0),
		ChildIndex: make(map[string][]int),
		Input:      input,
		Parent:     parent,
	}
}

func (n *RuleMatch) Path() string {
	prefix := ""
	if n.Parent != nil && n.Parent.Parent != nil {
		prefix = n.Parent.Path() + "."
	}
	return prefix + n.Name
}

func (n *RuleMatch) addChild(child *RuleMatch) {
	n.ChildIndex[child.Name] = append(n.ChildIndex[child.Name], len(n.Children))
	n.Children = append(n.Children, child)
}

func (n RuleMatch) lastCapture(capture int) *Range {
	if captures, exists := n.Captures[capture]; exists {
		return &captures[len(captures)-1]
	}
	return nil
}

func (n *RuleMatch) setGreedyMatches(index int, m Matcher, greedyMatches int) {
	matchers := n.Input.Matchers
	if index < len(matchers) {
		matchers[index] = MatcherMatch{
			Matcher:       m,
			GreedyMatches: greedyMatches,
			RuleMatch:     n,
		}
	}
}

func (n *RuleMatch) getGreedyMatches(index int, m Matcher) int {
	matchers := n.Input.Matchers
	if index >= len(matchers) {
		return -1
	}
	if matchers[index].Matcher == nil {
		return -1
	}
	if matchers[index].Matcher.ID() == m.ID() {
		return matchers[index].GreedyMatches
	}
	return -1
}

func (n *RuleMatch) decreaseGreedy(index int) {
	matchers := n.Input.Matchers
	matchers[index].GreedyMatches--
}

func (n *RuleMatch) getGreedyMatchInRange(start, end Position) (reset int, greedyMatches int) {
	matchers := n.Input.Matchers
	i := end.Index - 1
	reset = -1
	for i >= start.Index {
		if matchers[i].GreedyMatches > 0 {
			reset = i - 1
			greedyMatches = matchers[i].GreedyMatches
			matcher := matchers[i].Matcher.ID()

			state := cycleState{
				matcher:    matcher,
				reset:      reset,
				greediness: greedyMatches,
			}
			if n.Input.cycles.Has(state) {
				reset = -1
			} else {
				n.Input.cycles.Add(state)
			}

			return
		}
		i--
	}
	return
}

func (n *RuleMatch) clearGreedyAfter(pos Position) {
	stop := n.Input.Runes.Max()
	if stop.Index > pos.Index {
		matchers := n.Input.Matchers

		for i := pos.Index; i < stop.Index; i++ {
			matchers[i].GreedyMatches = -1
		}
	}
}

// A matcher matched over
type MatcherMatch struct {
	// The matcher that matched
	Matcher Matcher
	// The index of this node in its matchers repitition logic. This is stored here
	// for when the matching logic needs to backtrack on its greediness to find a match.
	GreedyMatches int
	// The RuleMatch this matcher belongs to
	RuleMatch *RuleMatch
}

func tokensToRunes(tokens []RuleToken) []rune {
	if len(tokens) == 0 {
		return []rune{}
	}
	runes := make([]rune, len(tokens))
	for i, t := range tokens {
		runes[i] = t.Value
	}
	return runes
}

func tokensToString(tokens []RuleToken) string {
	return string(tokensToRunes(tokens))
}

func matchersString(matchers []Matcher) string {
	s := ""
	for _, m := range matchers {
		s += m.String()
	}
	return s
}
