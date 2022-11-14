package ds

import "testing"

func initBits(on []uint32) Bits[uint32] {
	b := NewBits[uint32](on[len(on)-1] + 1)
	for _, onIndex := range on {
		b.Set(onIndex, true)
	}
	return b
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		bits     Bits[uint32]
		expected uint32
	}{
		{
			name:     "0",
			bits:     NewBits[uint32](0),
			expected: 63,
		},
		{
			name:     "1",
			bits:     NewBits[uint32](1),
			expected: 63,
		},
		{
			name:     "63",
			bits:     NewBits[uint32](63),
			expected: 63,
		},
		{
			name:     "64",
			bits:     NewBits[uint32](64),
			expected: 63 + 64,
		},
	}

	for _, test := range tests {
		actual := test.bits.Max()
		if actual != test.expected {
			t.Errorf("%s: Actual capacity %d not as expected: %d", test.name, actual, test.expected)
		}
	}
}

func TestEnsureMax(t *testing.T) {
	tests := []struct {
		name     string
		bits     Bits[uint32]
		ensure   uint32
		expected uint32
	}{
		{
			name:     "64",
			bits:     NewBits[uint32](0),
			ensure:   64,
			expected: 63 + 64,
		},
		{
			name:     "63",
			bits:     NewBits[uint32](0),
			ensure:   63,
			expected: 63,
		},
		{
			name:     "127",
			bits:     NewBits[uint32](63),
			ensure:   127,
			expected: 127,
		},
	}

	for _, test := range tests {
		test.bits.EnsureMax(test.ensure)
		actual := test.bits.Max()
		if actual != test.expected {
			t.Errorf("%s: Ensured capacity %d not as expected: %d", test.name, actual, test.expected)
		}
	}
}
func TestOperations(t *testing.T) {
	b := NewBits[uint32](127)

	if b.FirstOn() != -1 {
		t.Errorf("Unexpected on in bits")
	}

	b.Set(1, false)

	if b.FirstOn() != -1 {
		t.Errorf("Unexpected on in bits after fake set")
	}

	b.Set(1, true)

	if b.FirstOn() != 1 {
		t.Errorf("First on expected to be 1")
	}

	b.Set(1, false)

	if b.FirstOn() != -1 {
		t.Errorf("First on expected to be back to -1")
	}

	if b.Ons() != 0 {
		t.Errorf("Bits are on but shouldn't be.")
	}

	for i := uint32(0); i < 64; i++ {
		if b.Get(i) == true {
			t.Errorf("Bit %d should be off", i)
		}
	}

	b.Set(63, true)

	if b.Ons() != 1 {
		t.Errorf("Bits are not 1")
	}

	if b.FirstOn() != 63 {
		t.Errorf("First on expected to be back to 63")
	}

	if b.Get(63) != true {
		t.Errorf("Bit 63 expected to be on")
	}

	if b.Get(64) != false {
		t.Errorf("Bit 64 expected to be off")
	}

	b.Set(63, false)
	b.Set(64, true)

	if b.Ons() != 1 {
		t.Errorf("Bits are not 1")
	}

	if b.FirstOn() != 64 {
		t.Errorf("First on expected to be back to 64")
	}

	if b.Get(64) != true {
		t.Errorf("Bit 64 expected to be on")
	}
}

func TestOnAfter(t *testing.T) {
	tests := []struct {
		name     string
		bits     Bits[uint32]
		after    int32
		expected int32
	}{
		{
			name:     "empty",
			bits:     NewBits[uint32](0), //initBits([]uint32{}),
			after:    0,
			expected: -1,
		},
		{
			name:     "filled after -1",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    -1,
			expected: 0,
		},
		{
			name:     "filled after 0",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    0,
			expected: 1,
		},
		{
			name:     "filled after 1",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    1,
			expected: 2,
		},
		{
			name:     "filled after 2",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    2,
			expected: 3,
		},
		{
			name:     "filled after 3",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    3,
			expected: 4,
		},
		{
			name:     "filled after 4",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    4,
			expected: -1,
		},
		{
			name:     "sparse after 2",
			bits:     initBits([]uint32{2, 10, 54}),
			after:    2,
			expected: 10,
		},
		{
			name:     "sparse after 10",
			bits:     initBits([]uint32{2, 10, 54}),
			after:    10,
			expected: 54,
		},
		{
			name:     "sparse after 54",
			bits:     initBits([]uint32{2, 10, 54}),
			after:    54,
			expected: -1,
		},
		{
			name:     "really sparse after 0 (124 on)",
			bits:     initBits([]uint32{124}),
			after:    0,
			expected: 124,
		},
		{
			name:     "really sparse after 123 (124 on)",
			bits:     initBits([]uint32{124}),
			after:    123,
			expected: 124,
		},
		{
			name:     "really sparse after 124 (124 on)",
			bits:     initBits([]uint32{124}),
			after:    124,
			expected: -1,
		},
		{
			name:     "border test 1",
			bits:     initBits([]uint32{64}),
			after:    63,
			expected: 64,
		},
		{
			name:     "border test 2",
			bits:     initBits([]uint32{64}),
			after:    64,
			expected: -1,
		},
		{
			name:     "border test 3",
			bits:     initBits([]uint32{65}),
			after:    64,
			expected: 65,
		},
		{
			name:     "border test 4",
			bits:     initBits([]uint32{65, 128}),
			after:    65,
			expected: 128,
		},
	}

	for _, test := range tests {
		actual := test.bits.OnAfter(test.after)
		if actual != test.expected {
			t.Errorf("%s: OnAfter %d not as expected: %d", test.name, actual, test.expected)
		}
	}
}

func TestOffAfter(t *testing.T) {
	tests := []struct {
		name     string
		bits     Bits[uint32]
		after    int32
		expected int32
	}{
		{
			name:     "empty",
			bits:     NewBits[uint32](0), //initBits([]uint32{}),
			after:    0,
			expected: 1,
		},
		{
			name:     "filled after 0",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    0,
			expected: 5,
		},
		{
			name:     "filled after 1",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    1,
			expected: 5,
		},
		{
			name:     "filled after 2",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    2,
			expected: 5,
		},
		{
			name:     "filled after 3",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    3,
			expected: 5,
		},
		{
			name:     "filled after 4",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    4,
			expected: 5,
		},
		{
			name:     "filled after 5",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    5,
			expected: 6,
		},
		{
			name:     "filled after 5",
			bits:     initBits([]uint32{0, 1, 2, 3, 4}),
			after:    64,
			expected: -1,
		},
		{
			name:     "sparse after 2",
			bits:     initBits([]uint32{2, 10, 54}),
			after:    2,
			expected: 3,
		},
		{
			name:     "sparse after 2b",
			bits:     initBits([]uint32{2, 10, 54}),
			after:    1,
			expected: 3,
		},
		{
			name:     "sparse after 10",
			bits:     initBits([]uint32{2, 10, 54}),
			after:    10,
			expected: 11,
		},
		{
			name:     "sparse after 10b",
			bits:     initBits([]uint32{2, 10, 54}),
			after:    9,
			expected: 11,
		},
		{
			name:     "sparse after 54",
			bits:     initBits([]uint32{2, 10, 54}),
			after:    54,
			expected: 55,
		},
		{
			name:     "sparse after 55",
			bits:     initBits([]uint32{2, 10, 54}),
			after:    55,
			expected: 56,
		},
		{
			name:     "really sparse after 0 (124 on)",
			bits:     initBits([]uint32{124}),
			after:    0,
			expected: 1,
		},
		{
			name:     "really sparse after 123 (124 on)",
			bits:     initBits([]uint32{124}),
			after:    123,
			expected: 125,
		},
		{
			name:     "really sparse after 124 (124 on)",
			bits:     initBits([]uint32{124}),
			after:    124,
			expected: 125,
		},
		{
			name:     "border test 1",
			bits:     initBits([]uint32{64}),
			after:    63,
			expected: 65,
		},
		{
			name:     "border test 2",
			bits:     initBits([]uint32{64}),
			after:    64,
			expected: 65,
		},
		{
			name:     "border test 3",
			bits:     initBits([]uint32{65}),
			after:    64,
			expected: 66,
		},
		{
			name:     "border test 4",
			bits:     initBits([]uint32{65, 128}),
			after:    65,
			expected: 66,
		},
	}

	for _, test := range tests {
		actual := test.bits.OffAfter(test.after)
		if actual != test.expected {
			t.Errorf("%s: OnAfter %d not as expected: %d", test.name, actual, test.expected)
		}
	}
}
