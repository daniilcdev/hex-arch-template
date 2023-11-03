package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	adapter := NewAdapter()
	answer, err := adapter.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(2))
}

func TestSubtraction(t *testing.T) {
	adapter := NewAdapter()
	answer, err := adapter.Subtraction(5, 3)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(2))
}

func TestMultiplication(t *testing.T) {
	adapter := NewAdapter()
	answer, err := adapter.Multiplication(2, 2)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(4))
}

func TestDivision(t *testing.T) {
	adapter := NewAdapter()
	answer, err := adapter.Division(4, 2)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(2))
}
