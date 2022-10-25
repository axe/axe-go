package react

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	changes := ""

	v := Val(0)
	v.Set(2)

	Watch(func() int {
		return v.Get() * 2
	}, func(value int) {
		changes += fmt.Sprintf("%v ", value)
	})

	if changes != "4 " {
		t.Errorf("Changes is not '4 ' as expected.")
	}

	v.Set(45)

	if changes != "4 90 " {
		t.Errorf("Changes is not '4 90 ' as expected.")
	}

	triple := Computed(func() int {
		return v.Get() * 3
	})

	if triple.Get() != 135 {
		t.Errorf("triple is not 135 as expected.")
	}

	tripleDouble := Computed(func() int {
		return triple.Get() * 2
	})

	if tripleDouble.Get() != 270 {
		t.Errorf("tripleDouble is not 270 as expected.")
	}

	v.Set(3)

	if changes != "4 90 6 " {
		t.Errorf("Changes is not '4 90 3 ' as expected.")
	}
	if triple.Get() != 9 {
		t.Errorf("triple is not 6 as expected.")
	}
	if tripleDouble.Get() != 18 {
		t.Errorf("tripleDouble is not 18 as expected.")
	}

	v.Detach()
	v.Set(23)

	if changes != "4 90 6 " {
		t.Errorf("Changes is not '4 90 3 ' as expected.")
	}
	if triple.Get() != 9 {
		t.Errorf("triple is not 6 as expected.")
	}
	if tripleDouble.Get() != 18 {
		t.Errorf("tripleDouble is not 18 as expected.")
	}
}
