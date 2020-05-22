package reflections

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				Ciry string
			}{"Codd", "London"},
			[]string{"Codd", "London"},
		},
		{
			"Struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Codd", 28},
			[]string{"Codd"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{28, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{28, "London"},
			},
			[]string{"Chris", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Seoul"},
			},
			[]string{"London", "Seoul"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			fmt.Println(reflect.ValueOf(test))
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
	/*
		expected := "Chris"
		var got []string

		x := struct {
			Name string
		}{expected}

		walk(x, func(input string) {
			got = append(got, input)
		})

		fmt.Println(x)
		fmt.Println(got)
		if len(got) != 1 {
			t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
		}
		if got[0] != expected {
			t.Errorf("got %q want %q", got[0], expected)
		}
	*/
}
