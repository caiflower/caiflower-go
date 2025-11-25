package stack

import "testing"

func TestMyStack_Push(t *testing.T) {
	tests := []struct {
		name     string
		pushVals []int
		expected []int
	}{
		{
			name:     "push single element",
			pushVals: []int{1},
			expected: []int{1},
		},
		{
			name:     "push multiple elements",
			pushVals: []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "push empty",
			pushVals: []int{},
			expected: []int{},
		},
		{
			name:     "push negative numbers",
			pushVals: []int{-1, -2, -3},
			expected: []int{-1, -2, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MyStack{}
			for _, val := range tt.pushVals {
				s.Push(val)
			}

			if len(*s) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(*s))
			}

			for i, expected := range tt.expected {
				if (*s)[i] != expected {
					t.Errorf("Expected element at index %d to be %d, got %d", i, expected, (*s)[i])
				}
			}
		})
	}
}

func TestMyStack_Pop(t *testing.T) {
	tests := []struct {
		name        string
		pushVals    []int
		popCount    int
		expectedPop []int
		expectedLen int
	}{
		{
			name:        "pop single element",
			pushVals:    []int{1},
			popCount:    1,
			expectedPop: []int{1},
			expectedLen: 0,
		},
		{
			name:        "pop multiple elements",
			pushVals:    []int{1, 2, 3, 4, 5},
			popCount:    3,
			expectedPop: []int{5, 4, 3},
			expectedLen: 2,
		},
		{
			name:        "pop all elements",
			pushVals:    []int{10, 20, 30},
			popCount:    3,
			expectedPop: []int{30, 20, 10},
			expectedLen: 0,
		},
		{
			name:        "pop negative numbers",
			pushVals:    []int{-1, -2, -3},
			popCount:    2,
			expectedPop: []int{-3, -2},
			expectedLen: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MyStack{}
			for _, val := range tt.pushVals {
				s.Push(val)
			}

			poppedVals := make([]int, 0, tt.popCount)
			for i := 0; i < tt.popCount; i++ {
				poppedVals = append(poppedVals, s.Pop())
			}

			if len(poppedVals) != len(tt.expectedPop) {
				t.Errorf("Expected %d popped values, got %d", len(tt.expectedPop), len(poppedVals))
			}

			for i, expected := range tt.expectedPop {
				if poppedVals[i] != expected {
					t.Errorf("Expected popped value at index %d to be %d, got %d", i, expected, poppedVals[i])
				}
			}

			if s.Size() != tt.expectedLen {
				t.Errorf("Expected stack size %d after pops, got %d", tt.expectedLen, s.Size())
			}
		})
	}
}

func TestMyStack_Peek(t *testing.T) {
	tests := []struct {
		name     string
		pushVals []int
		expected int
	}{
		{
			name:     "peek single element",
			pushVals: []int{42},
			expected: 42,
		},
		{
			name:     "peek top of multiple elements",
			pushVals: []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "peek negative number",
			pushVals: []int{10, -20, 30},
			expected: 30,
		},
		{
			name:     "peek zero",
			pushVals: []int{1, 0},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MyStack{}
			for _, val := range tt.pushVals {
				s.Push(val)
			}

			result := s.Peek()
			if result != tt.expected {
				t.Errorf("Expected peek to return %d, got %d", tt.expected, result)
			}

			// Verify that peek doesn't modify the stack
			if s.Size() != len(tt.pushVals) {
				t.Errorf("Expected stack size to remain %d after peek, got %d", len(tt.pushVals), s.Size())
			}
		})
	}
}

func TestMyStack_Size(t *testing.T) {
	tests := []struct {
		name     string
		pushVals []int
		expected int
	}{
		{
			name:     "empty stack size",
			pushVals: []int{},
			expected: 0,
		},
		{
			name:     "single element size",
			pushVals: []int{1},
			expected: 1,
		},
		{
			name:     "multiple elements size",
			pushVals: []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "large stack size",
			pushVals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MyStack{}
			for _, val := range tt.pushVals {
				s.Push(val)
			}

			result := s.Size()
			if result != tt.expected {
				t.Errorf("Expected size %d, got %d", tt.expected, result)
			}
		})
	}
}

func TestMyStack_Integration(t *testing.T) {
	tests := []struct {
		name        string
		operations  []string
		values      []int
		expectedTop int
		expectedLen int
	}{
		{
			name:        "push and pop operations",
			operations:  []string{"push", "push", "pop", "push", "peek"},
			values:      []int{1, 2, 0, 3, 0}, // 0 for pop and peek operations
			expectedTop: 3,
			expectedLen: 2,
		},
		{
			name:        "complex operations",
			operations:  []string{"push", "push", "push", "pop", "pop", "push", "push"},
			values:      []int{10, 20, 30, 0, 0, 40, 50},
			expectedTop: 50,
			expectedLen: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MyStack{}
			valueIndex := 0

			for _, op := range tt.operations {
				switch op {
				case "push":
					s.Push(tt.values[valueIndex])
					valueIndex++
				case "pop":
					if s.Size() > 0 {
						s.Pop()
					}
					valueIndex++
				case "peek":
					if s.Size() > 0 {
						s.Peek()
					}
					valueIndex++
				}
			}

			if s.Size() != tt.expectedLen {
				t.Errorf("Expected final stack size %d, got %d", tt.expectedLen, s.Size())
			}

			if s.Size() > 0 {
				top := s.Peek()
				if top != tt.expectedTop {
					t.Errorf("Expected top element %d, got %d", tt.expectedTop, top)
				}
			}
		})
	}
}
