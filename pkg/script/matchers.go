package script

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type MatchChars struct {
	id    int
	Chars Set[rune]
	Not   bool
}

func (m MatchChars) ID() int { return m.id }

func (m MatchChars) String() string {
	chars := strings.Split(string(m.Chars.Values()), "")
	sort.Strings(chars)

	if m.Not {
		return "[^" + strings.Join(chars, "") + "]"
	} else {
		return "[" + strings.Join(chars, "") + "]"
	}
}

func (m MatchChars) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	p := input.Read()
	if p == nil {
		return false, nil
	}
	var matches bool
	if out.Rule.ignoreCase {
		upper := unicode.ToUpper(p.Value)
		lower := unicode.ToLower(p.Value)
		matches = (m.Chars.Has(upper) || m.Chars.Has(lower)) != m.Not
	} else {
		matches = m.Chars.Has(p.Value) != m.Not
	}
	// if matches {
	// 	out.setMatcher(p.Index, m, 0)
	// }
	return matches, nil
}

func isRange(a, b rune) bool {
	return ((isBetween(a, '0', '8') && isBetween(b, a+1, '9')) ||
		(isBetween(a, 'a', 'y') && isBetween(b, a+1, 'z')) ||
		(isBetween(a, 'A', 'Y') && isBetween(b, a+1, 'Z')))
}

func isBetween(v, min, max rune) bool {
	return v >= min && v <= max
}

type MatchAny struct {
	id int
}

func (m MatchAny) ID() int { return m.id }

func (m MatchAny) String() string {
	return "."
}

func (m MatchAny) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	r := input.Read()
	if r == nil {
		return false, nil
	}
	// out.setMatcher(r.Index, m, 0)
	return true, nil
}

type MatchRepeatMaybe struct {
	id          int
	RepeatMaybe Matcher
}

func (m MatchRepeatMaybe) ID() int { return m.id }

func (m MatchRepeatMaybe) String() string {
	return m.RepeatMaybe.String() + "*"
}

func (m MatchRepeatMaybe) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	return matchRepeat(rs, input, out, m.RepeatMaybe, 0, math.MaxInt, m)
}

type MatchRepeat struct {
	id     int
	Repeat Matcher
}

func (m MatchRepeat) ID() int { return m.id }

func (m MatchRepeat) String() string {
	return m.Repeat.String() + "+"
}

func (m MatchRepeat) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	return matchRepeat(rs, input, out, m.Repeat, 1, math.MaxInt, m)
}

type MatchRepeatRange struct {
	id          int
	RepeatRange Matcher
	Min         int
	Max         int
}

func (m MatchRepeatRange) ID() int { return m.id }

func (m MatchRepeatRange) String() string {
	min := ""
	max := ""
	if m.Min != 0 {
		min = strconv.FormatInt(int64(m.Min), 10)
	}
	if m.Max != math.MaxInt {
		max = strconv.FormatInt(int64(m.Max), 10)
	}
	if min == max {
		return m.RepeatRange.String() + "{" + min + "}"
	}
	return m.RepeatRange.String() + "{" + min + "," + max + "}"
}

func (m MatchRepeatRange) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	return matchRepeat(rs, input, out, m.RepeatRange, m.Min, m.Max, m)
}

type MatchMaybe struct {
	id    int
	Maybe Matcher
}

func (m MatchMaybe) ID() int { return m.id }

func (m MatchMaybe) String() string { return m.Maybe.String() + "?" }

func (m MatchMaybe) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	return matchRepeat(rs, input, out, m.Maybe, 0, 1, m)
}

type MatchLiteral struct {
	id      int
	Literal rune
}

func (m MatchLiteral) ID() int { return m.id }

func (m MatchLiteral) String() string {
	if RuneToRuleTokenKind(m.Literal) != RuleTokenLiteral {
		return "\\" + string(m.Literal)
	} else {
		return string(m.Literal)
	}
}

func (m MatchLiteral) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	p := input.Read()
	var matched bool
	if out.Rule.ignoreCase {
		upper := unicode.ToUpper(m.Literal)
		lower := unicode.ToLower(m.Literal)
		matched = p != nil && (p.Value == upper || p.Value == lower)
	} else {
		matched = p != nil && p.Value == m.Literal
	}
	// if matched {
	// 	out.setMatcher(p.Index, m, 0)
	// }
	return matched, nil
}

type MatchOneOf struct {
	id      int
	OneOf   []Buffer[Matcher]
	Capture int
	Reset   bool
	Not     bool
}

func (m MatchOneOf) ID() int { return m.id }

func (m MatchOneOf) String() string {
	oneOfs := make([]string, len(m.OneOf))
	for i, oneOf := range m.OneOf {
		oneOfs[i] = matchersString(oneOf.data)
	}
	prefix := ""
	switch true {
	case m.Not:
		prefix = "?!"
	case m.Capture == -1:
		prefix = "?:"
	case m.Reset:
		prefix = "?="
	}
	return "(" + prefix + strings.Join(oneOfs, "|") + ")"
}

func (m MatchOneOf) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	start := input.Pos()

	for i := range m.OneOf {
		oneOf := m.OneOf[i]
		input.Reset(start)

		match, err := matchSlice(rs, input, out, oneOf.data)
		if err != nil {
			return false, err
		}
		if match {
			end := input.Pos()
			if m.Reset {
				input.Reset(start)
			}
			if m.Capture != -1 {
				out.Captures[m.Capture] = append(out.Captures[m.Capture], newRange(start, end))
			}
			matched := !m.Not
			// if matched {
			// 	out.setMatcher(start.Index, m, 0)
			// }
			return matched, nil
		} else if !out.Rule.greedy {
			out.clearGreedyAfter(start)
		}
	}
	input.Reset(start)

	return m.Not, nil
}

type MatchRule struct {
	id    int
	Rule  string
	Alias string
}

func (m MatchRule) ID() int { return m.id }

func (m MatchRule) String() string {
	if m.Rule == m.Alias {
		return "<" + m.Rule + ">"
	} else {
		return "<" + m.Alias + ":" + m.Rule + ">"
	}
}

func (m MatchRule) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	rule := rs.Rule(m.Rule)
	if rule == nil {
		return false, fmt.Errorf("Rule %s not defined in set", m.Rule)
	}
	matched, err := rule.Match(rs, input, out.Input, out)
	if matched != nil && !rule.ignore {
		matched.Name = m.Alias
		out.addChild(matched)
	}
	return matched != nil, err
}

type MatchReference struct {
	id       int
	Captured int
}

func (m MatchReference) ID() int { return m.id }

func (m MatchReference) String() string {
	return "\\" + strconv.FormatInt(int64(m.Captured), 10)
}

func (m MatchReference) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	// start := input.Pos()
	var captureValue []rune
	if capture := out.lastCapture(m.Captured); capture != nil {
		captureValue = capture.GetRunes(*input)
	}
	if len(captureValue) > 0 {
		for _, expected := range captureValue {
			next := input.Read()
			if next == nil || next.Value != expected {
				return false, nil
			}
		}
	}
	// out.setMatcher(start.Index, m, 0)
	return true, nil
}

type MatchLineStart struct {
	id int
}

func (m MatchLineStart) ID() int { return m.id }

func (m MatchLineStart) String() string {
	return "^"
}

func (m MatchLineStart) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	pos := input.Pos()
	matched := pos.Column == 0
	// if matched {
	// out.setMatcher(pos.Index, m, 0)
	// }
	return matched, nil
}

type MatchLineEnd struct {
	id int
}

func (m MatchLineEnd) ID() int { return m.id }

func (m MatchLineEnd) String() string {
	return "$"
}

func (m MatchLineEnd) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	p := input.Peek()
	matched := p == nil || p.Value == '\n'
	// if matched {
	// 	out.setMatcher(p.Index, m, 0)
	// }
	return matched, nil
}

type MatchShortcut struct {
	id       int
	key      rune
	Shortcut Matcher
}

func (m MatchShortcut) ID() int { return m.id }

func (m MatchShortcut) String() string {
	return `\` + string(m.key)
}

func (m MatchShortcut) Match(rs RuleSet, input *RuneBuffer, out *RuleMatch) (bool, error) {
	return m.Shortcut.Match(rs, input, out)
}

func matchRepeat(rs RuleSet, input *RuneBuffer, out *RuleMatch, matcher Matcher, min int, max int, forMatcher Matcher) (bool, error) {
	repeatMax := max
	start := input.Pos()
	if !out.Rule.greedy {
		greedyMatches := out.getGreedyMatches(start.Index, forMatcher)
		if greedyMatches >= 0 {
			if rs.Logger.OnGetGreedyMatches != nil {
				rs.Logger.OnGetGreedyMatches(input, out, forMatcher, matcher, min, max, greedyMatches, start)
			}

			repeatMax = min + greedyMatches
			if repeatMax < min {
				return false, nil
			}
		}
	}
	if repeatMax == 0 {
		return true, nil
	}

	logger := rs.loggerFor(input, out, matcher, start)

	matches := 0
	matchesMoved := 0
	last := start
	for {
		pos := input.Pos()
		match, err := logger(matcher.Match(rs, input, out))
		if err != nil {
			return false, err
		}
		next := input.Pos()
		if match {
			matches++
			if next != last {
				matchesMoved++
			}
			if matches == repeatMax || input.Ended() {
				break
			}
		} else {
			input.Reset(pos)
			break
		}
		if next == last {
			return logger(false, fmt.Errorf("%s in %s is repeating but not consuming input", matcher.String(), out.Rule.Name()))
		}
		last = next
	}

	matched := matches >= min && matches <= max
	if !out.Rule.greedy {
		matchedGreedy := matched && matchesMoved > 0
		if matchedGreedy {
			if rs.Logger.OnSetGreedyMatches != nil {
				rs.Logger.OnSetGreedyMatches(input, out, forMatcher, matcher, min, max, matches-min, start)
			}

			out.setGreedyMatches(start.Index, forMatcher, matches-min)
		}
	}

	return matched, nil
}

func matchSlice(rs RuleSet, input *RuneBuffer, out *RuleMatch, matchers []Matcher) (bool, error) {
	matcherPositions := make([]Position, len(matchers))
	matcherIndex := 0
	start := input.Pos()

	// TODO only add child to out

	for matcherIndex < len(matchers) {
		matcher := matchers[matcherIndex]
		matcherPositions[matcherIndex] = input.Pos()
		logger := rs.loggerFor(input, out, matcher, input.Pos())

		match, err := logger(matcher.Match(rs, input, out))
		if err != nil {
			return false, err
		}

		if !match {
			if out.Rule.greedy {
				input.Reset(start)

				return false, nil
			}

			resetTo, greediness := out.getGreedyMatchInRange(start, input.Pos())
			if resetTo != -1 {
				out.decreaseGreedy(resetTo)

				for matcherIndex > 0 && resetTo < matcherPositions[matcherIndex].Index {
					matcherIndex--
				}

				resetAt := matcherPositions[matcherIndex]

				input.Reset(resetAt)

				if rs.Logger.OnDecreaseGreedyMatches != nil {
					rs.Logger.OnDecreaseGreedyMatches(input, out, matchers[matcherIndex], resetTo, greediness)
				}

				continue
			} else {
				input.Reset(start)

				return false, nil
			}
		}
		matcherIndex++
	}

	return true, nil
}
