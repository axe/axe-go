package ease

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/axe/axe-go/pkg/id"
)

// Formats:
// easing
// easing-modifier
// easing*scale
// easing-modifier*scale
// bezier(mx1,my1,mx2,my2)
// bezier(mx1,my1,mx2,my2)*scale
var Regex = regexp.MustCompile(`^(?i)(bezier\(\s*(-?\d*\.?\d*)\s*,\s*(-?\d*\.?\d*)\s*,\s*(-?\d*\.?\d*)\s*,\s*(-?\d*\.?\d*)\s*\)|([^-]*))(|\s*-\s*([^*]+))(|\s*\*\s*(-?\d*\.?\d*))`)

func MustParse(s string) Easing {
	easing, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return easing
}

func Parse(s string) (Easing, error) {
	matches := Regex.FindStringSubmatch(s)
	if matches == nil {
		return nil, fmt.Errorf("invalid easing %s", s)
	}

	var easing Easing
	if matches[1] != "" {
		mx1, _ := strconv.ParseFloat(matches[2], 32)
		my1, _ := strconv.ParseFloat(matches[3], 32)
		mx2, _ := strconv.ParseFloat(matches[4], 32)
		my2, _ := strconv.ParseFloat(matches[5], 32)

		easing = Bezier{
			MX1: float32(mx1),
			MY1: float32(my1),
			MX2: float32(mx2),
			MY2: float32(my2),
		}
	} else {
		easing = Easings.Get(id.Maybe(strings.ToLower(matches[6])))
		if easing == nil {
			return nil, fmt.Errorf("easing not found %s", matches[6])
		}
	}

	if matches[7] != "" {
		modifier := Modifiers.Get(id.Maybe(strings.ToLower(matches[7])))
		if modifier.Fn == nil {
			return nil, fmt.Errorf("easing modifier not found %s", matches[7])
		}

		easing = modifier.Modify(easing)
	}

	if matches[9] != "" {
		scale, _ := strconv.ParseFloat(matches[9], 32)

		easing = Scaled{
			Scale:  float32(scale),
			Easing: easing,
		}
	}

	return easing, nil
}
