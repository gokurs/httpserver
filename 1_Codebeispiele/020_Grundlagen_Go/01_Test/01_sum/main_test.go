package main

import "testing"

func TestSum1(t *testing.T) {
	got := sum(1, 2)
	if got != 3 {
		t.Errorf("sum(1,2)=%v, jedoch %v wurde erwartet", got, 3)
	}
	got = sum(10, 12)
	if got != 22 {
		t.Errorf("sum(10,12)=%v, jedoch %v wurde erwartet", got, 22)
	}
}

func TestSumN(t *testing.T) {
	got := sumN(1, 1, 1, 1)
	expect := 4
	if got != expect {
		t.Errorf("sumN(1,1,1,1)=%v, jedoch %v wurde erwartet", got, expect)
	}
	got = sumN(1, 1, 1, 1, 10)
	expect = 14
	if got != expect {
		t.Errorf("sumN(1,1,1,1)=%v, jedoch %v wurde erwartet", got, expect)
	}
}
