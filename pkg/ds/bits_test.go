package ds

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		bits     Bits
		expected uint32
	}{
		{
			name:     "0",
			bits:     NewBits(0),
			expected: 63,
		},
		{
			name:     "1",
			bits:     NewBits(1),
			expected: 63,
		},
		{
			name:     "63",
			bits:     NewBits(63),
			expected: 63,
		},
		{
			name:     "64",
			bits:     NewBits(64),
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
		bits     Bits
		ensure   uint32
		expected uint32
	}{
		{
			name:     "64",
			bits:     NewBits(0),
			ensure:   64,
			expected: 63 + 64,
		},
		{
			name:     "63",
			bits:     NewBits(0),
			ensure:   63,
			expected: 63,
		},
		{
			name:     "127",
			bits:     NewBits(63),
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
	b := NewBits(127)

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
