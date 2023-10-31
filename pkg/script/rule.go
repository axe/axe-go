package script

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

type RulePattern string

func (r RulePattern) Parse() (parsed RulePatternParsed, err error) {
	parsed.Pattern = string(r)
	if r[0] == '/' {
		i := len(r) - 1
		for i >= 1 {
			switch r[i] {
			case 'i':
				parsed.IgnoreCase = true
			case 'g':
				parsed.Greedy = true
			case 'n':
				parsed.Ignore = true
			case '/':
				parsed.Pattern = parsed.Pattern[1:i]
				return
			}
			i--
		}
		err = fmt.Errorf("invalid rule %s, if it starts with a / it must end with / possible followed by one of the options: i, g, n", parsed.Pattern)
	}
	return
}

type RulePatternParsed struct {
	Pattern    string
	IgnoreCase bool
	Greedy     bool
	Ignore     bool
}

type Rule struct {
	name        string
	ignoreCase  bool
	ignore      bool
	greedy      bool
	shortcuts   map[rune]Matcher
	input       RuneBuffer
	tokens      Buffer[RuleToken]
	matchers    Buffer[Matcher]
	bufferStack Stack[*Buffer[Matcher]]
	captures    int
	matcherID   *int
}

func (r Rule) String() string {
	return string(r.input.buffer.data)
}

func (r *Rule) Match(rs RuleSet, runes *RuneBuffer, input *Input, parent *RuleMatch) (*RuleMatch, error) {
	ruleMatch := newRuleMatch(r, input, parent)
	ruleMatch.Start = runes.Pos()

	match, err := matchSlice(rs, runes, ruleMatch, r.matchers.data)
	if err != nil || !match {
		return nil, err
	}
	ruleMatch.End = runes.Pos()

	return ruleMatch, nil
}

func (r Rule) Name() string {
	return r.name
}

func (r Rule) Ignore() bool {
	return r.ignore
}

func (r *Rule) Parse(input string) error {
	r.input.Set(input)

	r.buildTokens()

	err := r.buildMatches()

	if err == nil && !r.tokens.Ended() {
		err = fmt.Errorf("not all tokens consumed in %s, read up to %v", r.Name(), r.tokens.Pos())
	}

	return err
}

func (r *Rule) buildTokens() {
	r.tokens.Clear(r.input.Len())

	for tk := r.readToken(); tk.Kind != RuleTokenEOF; tk = r.readToken() {
		r.tokens.Add(tk)
	}
}

func (r *Rule) readToken() RuleToken {
	c, k := r.read()
	return RuleToken{Rune: c, Kind: k}
}

func (r *Rule) read() (c Rune, k RuleTokenKind) {
	c, k = r.next()
	if k == RuleTokenEscape {
		c, _ = r.next()
		if shortcut, ok := r.shortcuts[c.Value]; ok && shortcut != nil {
			k = RuleTokenShortcut
		} else if isBetween(c.Value, '0', '9') {
			k = RuleTokenEscape
		} else {
			k = RuleTokenLiteral
		}
	}
	return
}

func (r *Rule) next() (c Rune, k RuleTokenKind) {
	cp := r.input.Read()
	if cp != nil {
		c = *cp
		k = RuneToRuleTokenKind(c.Value)
	}
	return
}

func (r *Rule) nextMatcherID() int {
	id := *r.matcherID
	*r.matcherID++
	return id
}

func (r *Rule) buildMatches() error {
	r.matchers.Clear(24)
	r.bufferStack.Push(&r.matchers)

	for {
		m, err := r.readMatch(false)
		if err != nil {
			if err == ErrEndOfInput {
				break
			}
			return err
		}
		if m == nil {
			return fmt.Errorf("no matcher could be determined at %v", r.tokens.Pos())
		}
		r.matchers.Add(m)
	}
	return nil
}

func (r *Rule) readMatch(oneOf bool) (Matcher, error) {
	p := r.tokens.Read()
	if p == nil {
		return nil, ErrEndOfInput
	}
	switch p.Kind {
	case RuleTokenAny:
		return r.readAny()
	case RuleTokenLineStart:
		return r.readLineStart()
	case RuleTokenLineEnd:
		return r.readLineEnd()
	case RuleTokenShortcut:
		return r.readShortcut(p.Value)
	case RuleTokenEscape:
		return r.readReference(p.Value)
	case RuleTokenLiteral,
		RuleTokenRepeatRangeDelimiter,
		RuleTokenRuleDelimiter,
		RuleTokenCharsRange,
		RuleTokenCharsEnd,
		RuleTokenRepeatRangeEnd,
		RuleTokenRuleEnd:
		return r.readLiteral(p.Value)
	case RuleTokenCharsStart:
		return r.readChars()
	case RuleTokenMaybe:
		return r.readMaybe()
	case RuleTokenRepeat:
		return r.readRepeat()
	case RuleTokenRepeatMaybe:
		return r.readRepeatMaybe()
	case RuleTokenRepeatRangeStart:
		return r.readRepeatRange()
	case RuleTokenOneOfStart:
		return r.readOneOf()
	case RuleTokenRuleStart:
		return r.readRule()
	case RuleTokenOneOfEnd:
		if oneOf {
			return nil, ErrOneOfEnd
		} else {
			return r.readLiteral(p.Value)
		}
	case RuleTokenOneOfDelimiter:
		if oneOf {
			return nil, ErrOneOfDelimiter
		} else {
			return r.readLiteral(p.Value)
		}
	default:
		return nil, errors.Join(ErrUnexpectedToken, fmt.Errorf("%s", string(p.Value)))
	}
}

func (r *Rule) readTo(kind RuleTokenKind, match bool) (tokens []RuleToken, end bool) {
	tokens = make([]RuleToken, 0)
	for t := r.tokens.Read(); t != nil; t = r.tokens.Read() {
		if (t.Kind == kind) == match {
			return
		}
		tokens = append(tokens, *t)
	}
	end = true
	return
}

func (r *Rule) readValidStop(valid, stop Set[rune]) (tokens []RuleToken, stopped *RuleToken) {
	tokens = make([]RuleToken, 0)
	for t := r.tokens.Read(); t != nil; t = r.tokens.Read() {
		if stop.Has(t.Value) {
			stopped = t
			return
		}
		if !valid.Has(t.Value) {
			break
		}
		tokens = append(tokens, *t)
	}
	return
}

func (r *Rule) readAny() (MatchAny, error) {
	return MatchAny{id: r.nextMatcherID()}, nil
}

func (r *Rule) readShortcut(key rune) (Matcher, error) {
	return MatchShortcut{id: r.nextMatcherID(), key: key, Shortcut: r.shortcuts[key]}, nil
}

func (r *Rule) readLineStart() (MatchLineStart, error) {
	return MatchLineStart{id: r.nextMatcherID()}, nil
}

func (r *Rule) readLineEnd() (MatchLineEnd, error) {
	return MatchLineEnd{id: r.nextMatcherID()}, nil
}

func (r *Rule) readLiteral(literal rune) (MatchLiteral, error) {
	return MatchLiteral{id: r.nextMatcherID(), Literal: literal}, nil
}

func (r *Rule) readReference(first rune) (MatchReference, error) {
	capture := []rune{first}
	for more := r.tokens.Peek(); more != nil && isBetween(more.Rune.Value, '0', '9'); more = r.tokens.Peek() {
		capture = append(capture, more.Value)
		r.tokens.Read()
	}
	captureIndex, _ := strconv.ParseInt(string(capture), 10, 64)

	return MatchReference{id: r.nextMatcherID(), Captured: int(captureIndex)}, nil
}

func (r *Rule) readChars() (Matcher, error) {
	charsWithRanges, end := r.readTo(RuleTokenCharsEnd, true)
	if end {
		return nil, ErrCharsNoEnd
	}
	return matchCharsFromTokens(charsWithRanges, r.nextMatcherID()), nil
}

func matchCharsFromString(input string, id int) MatchChars {
	r := Rule{}
	r.input.Set(input)
	r.buildTokens()
	return matchCharsFromTokens(r.tokens.data, id)
}

func matchCharsFromTokens(charsWithRanges []RuleToken, id int) MatchChars {
	chars := NewSet[rune]()
	last := len(charsWithRanges) - 1
	not := false
	for i, c := range charsWithRanges {
		if i == 0 && c.Kind == RuleTokenCharsNot {
			not = true
			continue
		}
		if c.Kind == RuleTokenCharsRange && i > 0 && i < last && isRange(charsWithRanges[i-1].Value, charsWithRanges[i+1].Value) {
			for j := charsWithRanges[i-1].Value + 1; j < charsWithRanges[i+1].Value; j++ {
				chars.Add(j)
			}
		} else {
			chars.Add(c.Value)
		}
	}
	return MatchChars{id: id, Chars: chars, Not: not}
}

func (r *Rule) removeLast() *Matcher {
	last := r.bufferStack.Peek()
	if last == nil {
		return nil
	}
	return (*last).RemoveLast()
}

func (r *Rule) readMaybe() (Matcher, error) {
	inner := r.removeLast()
	if inner == nil {
		return nil, errors.New("Unexpected ?")
	}
	return MatchMaybe{id: r.nextMatcherID(), Maybe: *inner}, nil
}

func (r *Rule) readRepeat() (Matcher, error) {
	inner := r.removeLast()
	if inner == nil {
		return nil, errors.New("Unexpected +")
	}
	return MatchRepeat{id: r.nextMatcherID(), Repeat: *inner}, nil
}

func (r *Rule) readRepeatMaybe() (Matcher, error) {
	inner := r.removeLast()
	if inner == nil {
		return nil, errors.New("Unexpected *")
	}
	return MatchRepeatMaybe{id: r.nextMatcherID(), RepeatMaybe: *inner}, nil
}

var (
	repeatRangeMinStop = NewSet([]rune(",}")...)
	repeatRangeMaxStop = NewSet([]rune("}")...)
	repeatRangeValid   = NewSet([]rune("0123456789")...)
)

func (r *Rule) readRepeatRange() (Matcher, error) {
	pos := r.tokens.Pos()

	var minTokens, maxTokens []RuleToken
	var stop *RuleToken

	minTokens, stop = r.readValidStop(repeatRangeValid, repeatRangeMinStop)
	if stop == nil {
		return nil, fmt.Errorf("invalid {} starting at %v", pos)
	}
	if stop.Value == '}' {
		if len(minTokens) == 0 {
			return nil, fmt.Errorf("empty {} starting at %v", pos)
		}
		maxTokens = minTokens
	} else {
		maxTokens, stop = r.readValidStop(repeatRangeValid, repeatRangeMaxStop)
		if stop == nil {
			return nil, fmt.Errorf("invalid {} starting at %v", pos)
		}
	}

	minString := tokensToString(minTokens)
	maxString := tokensToString(maxTokens)
	min := 0
	max := math.MaxInt

	if len(minString) > 0 {
		minParsed, err := strconv.ParseInt(minString, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid min %s in range at %v", minString, pos)
		}
		min = int(minParsed)
	}
	if len(maxString) > 0 {
		maxParsed, err := strconv.ParseInt(maxString, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid max %s in range at %v", maxString, pos)
		}
		max = int(maxParsed)
	}

	if max < min {
		return nil, ErrRangeInvalid
	}
	inner := r.removeLast()
	if inner == nil {
		return nil, ErrRangeInvalid
	}
	return MatchRepeatRange{id: r.nextMatcherID(), RepeatRange: *inner, Min: min, Max: max}, nil
}

var (
	ruleAliasStop = NewSet([]rune(":>")...)
	ruleNameStop  = NewSet('>')
	ruleValid     = NewSet([]rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_.-0123456789")...)
)

func (r *Rule) readRule() (Matcher, error) {
	pos := r.tokens.Pos()

	var nameTokens, aliasTokens []RuleToken
	var stop *RuleToken

	nameTokens, stop = r.readValidStop(ruleValid, ruleAliasStop)
	if stop == nil {
		return nil, fmt.Errorf("invalid <> rule at %v", pos)
	}
	if stop.Value == '>' {
		if len(nameTokens) == 0 {
			return nil, fmt.Errorf("invalid <> rule at %v", pos)
		}
		aliasTokens = nameTokens
	} else {
		aliasTokens = nameTokens
		nameTokens, stop = r.readValidStop(ruleValid, ruleNameStop)
		if stop == nil {
			return nil, fmt.Errorf("invalid <%s:> rule at %v", tokensToString(aliasTokens), pos)
		}
		if len(nameTokens) == 0 {
			return nil, fmt.Errorf("invalid <> rule at %v", pos)
		}
	}

	rule := tokensToString(nameTokens)
	alias := tokensToString(aliasTokens)

	return MatchRule{id: r.nextMatcherID(), Rule: rule, Alias: alias}, nil
}

func (r *Rule) readOneOf() (Matcher, error) {
	pos := r.tokens.Pos()

	oneOf := MatchOneOf{
		OneOf:   make([]Buffer[Matcher], 0),
		Capture: r.captures,
	}
	r.captures++

	special := r.tokens.Peek()
	if special != nil && special.Value == '?' {
		r.tokens.Read()
		kind := r.tokens.Read()
		if kind == nil {
			return nil, fmt.Errorf("unexpected end of input at %v", pos)
		}
		switch kind.Value {
		case ':':
			oneOf.Capture = -1
			r.captures--
		case '!':
			oneOf.Reset = true
			oneOf.Not = true
		case '=':
			oneOf.Reset = true
		default:
			return nil, fmt.Errorf("expected :!= after ? at %v", pos)
		}
	}

	current := NewBuffer[Matcher](0)
	r.bufferStack.Push(&current)

	for {
		matcher, err := r.readMatch(true)
		switch err {
		case nil:
			current.Add(matcher)
		case ErrOneOfDelimiter:
			oneOf.OneOf = append(oneOf.OneOf, current)
			r.bufferStack.Pop()
			current = NewBuffer[Matcher](0)
			r.bufferStack.Push(&current)
		case ErrOneOfEnd:
			oneOf.OneOf = append(oneOf.OneOf, current)
			oneOf.id = r.nextMatcherID()
			r.bufferStack.Pop()
			return oneOf, nil
		case ErrEndOfInput:
			return nil, ErrOneOfNoEnd
		default:
			return nil, err
		}
	}
}
