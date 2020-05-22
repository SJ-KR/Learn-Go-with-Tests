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
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Boz", "Bar"},
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

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains := func(t *testing.T, got []string, want string) {
			t.Helper()
			for _, s := range got {
				if s == want {
					return
				}
			}
			t.Errorf("expected %+v to contain %q but it didn't", got, want)
		}

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)
		go func() {
			aChannel <- Profile{33, "Ber"}
			aChannel <- Profile{34, "Kato"}
			close(aChannel)
		}()
		var got []string
		want := []string{"Ber", "Kato"}
		walk(aChannel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	})
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
