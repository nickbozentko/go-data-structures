package stack_test

import (
	"testing"

	"github.com/nickbozentko/go-data-structures/stack"
)

func TestStack(t *testing.T) {
	myStack := stack.NewStack[int]()

	if myStack.Length() != 0 {
		t.Error("stack length is not 0 after creation")
	} else {
		t.Log("stack length is 0 after creation")
	}

	myStack.Push(1)
	myStack.Push(2)
	myStack.Push(3)

	if myStack.Length() != 3 {
		t.Errorf("stack length is not 3 after pushing 3 elements")
	} else {
		t.Log("stack length is 3 after pushing 3 elements")
	}

	three, ok := myStack.Peek()
	if !ok {
		t.Error("peek unsuccessful")
	} else {
		t.Log("peek ok")
	}
	if three != 3 {
		t.Error("should have peeked 3")
	} else {
		t.Log("peeked 3")
	}

	threeAlso, ok := myStack.Pop()
	if !ok {
		t.Error("pop unsuccessful")
	} else {
		t.Log("pop ok")
	}
	if threeAlso != 3 {
		t.Error("should have popped 3")
	} else {
		t.Log("popped 3")
	}
	if myStack.Length() != 2 {
		t.Errorf("length is wrong after pop")
	}

	two, ok := myStack.Pop()
	if !ok {
		t.Error("pop unsuccessful")
	} else {
		t.Log("pop ok")
	}
	if two != 2 {
		t.Error("should have popped 2")
	} else {
		t.Log("popped 2")
	}

	myStack.Pop()

	_, ok = myStack.Peek()
	if ok {
		t.Errorf("peek empty stack should not be ok")
	} else {
		t.Log("peek empty stack not ok, as expected")
	}

	_, ok = myStack.Pop()
	if ok {
		t.Errorf("pop empty stack should not be ok")
	} else {
		t.Log("pop empty stack not ok, as expected")
	}
}
