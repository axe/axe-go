package ds

import (
	"testing"
)

func initBits(on []uint32) Bits {
	b := NewBits(on[len(on)-1] + 1)
	for _, onIndex := range on {
		b.Set(onIndex, true)
	}
	return b
}

func TestSparseList(t *testing.T) {

	type ListBuilder func(list *SparseList[string])

	type Expected struct {
		items []string
		free  Bits
	}

	tests := []struct {
		name     string
		build    ListBuilder
		expected Expected
	}{
		{
			"+1 +3 +5",
			func(l *SparseList[string]) {
				l.Add("1")
				l.Add("3")
				l.Add("5")
			},
			Expected{
				[]string{"1", "3", "5"},
				NewBits(0),
			},
		},
		{
			"+1 +3 +5 -3 +7",
			func(l *SparseList[string]) {
				l.Add("1")
				x := l.Add("3")
				l.Add("5")
				l.Free(x)
				l.Add("7")
			},
			Expected{
				[]string{"1", "7", "5"},
				NewBits(0),
			},
		},
		{
			"+1 +3 +5 -1 -3 -5",
			func(l *SparseList[string]) {
				a := l.Add("1")
				b := l.Add("3")
				c := l.Add("5")
				l.Free(a)
				l.Free(b)
				l.Free(c)
			},
			Expected{
				[]string{},
				initBits([]uint32{0, 1, 2}),
			},
		},
		{
			"+1 +3 +5 -5 +7",
			func(l *SparseList[string]) {
				l.Add("1")
				l.Add("3")
				c := l.Add("5")
				l.Free(c)
				l.Add("7")
			},
			Expected{
				[]string{"1", "3", "7"},
				NewBits(0),
			},
		},
		{
			"+1 +3 +5 -5 +7 -3",
			func(l *SparseList[string]) {
				l.Add("1")
				b := l.Add("3")
				c := l.Add("5")
				l.Free(c)
				l.Add("7")
				l.Free(b)
			},
			Expected{
				[]string{"1", "7"},
				initBits([]uint32{1}),
			},
		},
		{
			"+1 +3 +5 -5 compress",
			func(l *SparseList[string]) {
				l.Add("1")
				b := l.Add("3")
				l.Add("5")
				l.Free(b)
				l.Compress(true, nil)
			},
			Expected{
				[]string{"1", "5"},
				NewBits(0),
			},
		},
		{
			"take",
			func(l *SparseList[string]) {
				a, _ := l.Take()
				*a = "1"

				b, bi := l.Take()
				*b = "3"

				c, _ := l.Take()
				*c = "5"

				l.Free(bi)
			},
			Expected{
				[]string{"1", "5"},
				initBits([]uint32{1}),
			},
		},
	}

	for _, test := range tests {
		list := NewSparseList[string](10)

		test.build(&list)

		actual := list.Values()

		text := ""
		for _, actualValue := range actual {
			text += "," + actualValue
		}
		t.Logf("Test '%s': %s", test.name, text)

		if len(actual) != len(test.expected.items) {
			t.Errorf("Test '%s' actual(%d) and expected(%d) are different lengths", test.name, len(actual), len(test.expected.items))
		} else {
			for i, actualValue := range actual {
				if actualValue != test.expected.items[i] {
					t.Errorf("Test '%s' at index %d failed, actual %s <> expected %s", test.name, i, actualValue, test.expected.items[i])
				}
			}
		}

		if !list.free.SameOnes(test.expected.free) {
			t.Errorf("Test '%s' actual free(%+v) and expected free(%+v) are different", test.name, list.free.values, test.expected.free.values)
		}
	}
}
