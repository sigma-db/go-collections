package collections

import (
	"testing"
	"time"

	"golang.org/x/exp/slices"
)

func TestNestedLoopJoin(t *testing.T) {
	const secondsPerDay = 24 * 60 * 60 // seconds in a day
	date := func(year int, month time.Month, day int) uint32 {
		duration := time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Unix()
		return uint32(duration) / secondsPerDay
	}

	type employee struct {
		id         uint32
		name       string
		birthday   uint32 // days since epoch (1970-01-01)
		salary     float64
		department uint32 // department.id
	}

	type department struct {
		id   uint32
		name string
		head uint32 // employee.id
	}

	employees := []employee{
		{1, "John", date(1990, 2, 1), 1000, 1},
		{2, "Mary", date(1991, 1, 27), 2000, 1},
		{3, "Bob", date(1992, 3, 10), 3000, 2},
		{4, "Anne", date(1993, 9, 19), 4000, 2},
		{5, "Tom", date(1994, 11, 2), 5000, 3},
		{6, "Jerry", date(1995, 1, 8), 6000, 3},
	}

	departments := []department{
		{1, "Sales", 1},
		{2, "Marketing", 3},
		{3, "IT", 5},
	}

	employee_department := func(e *employee) *uint32 { return &e.department }
	department_id := func(d *department) *uint32 { return &d.id }

	got := NestedLoopJoin(employees, departments, employee_department, department_id)
	want := []Pair[employee, department]{
		{employees[0], departments[0]},
		{employees[1], departments[0]},
		{employees[2], departments[1]},
		{employees[3], departments[1]},
		{employees[4], departments[2]},
		{employees[5], departments[2]},
	}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
