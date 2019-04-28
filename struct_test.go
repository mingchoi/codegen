package main

import "testing"

type RenderContextCaseTestCase struct {
	Title    string
	Expected string
	Actual   string
}

func TestRenderContextCase(t *testing.T) {
	ctx := RenderContext{
		Fields: map[string]Field{
			"Test": Field{
				Value: "My example Word",
			},
		},
	}

	cases := []RenderContextCaseTestCase{
		RenderContextCaseTestCase{
			Title:    "Camel",
			Expected: "myExampleWord",
			Actual:   ctx.Camel("Test"),
		},
		RenderContextCaseTestCase{
			Title:    "Title",
			Expected: "MyExampleWord",
			Actual:   ctx.Title("Test"),
		},
		RenderContextCaseTestCase{
			Title:    "Dash",
			Expected: "my-example-word",
			Actual:   ctx.Dash("Test"),
		},
		RenderContextCaseTestCase{
			Title:    "Constant",
			Expected: "MY_EXAMPLE_WORD",
			Actual:   ctx.Constant("Test"),
		},
	}

	for _, v := range cases {
		if v.Actual != v.Expected {
			t.Errorf("Result of RenderContext.%s() not equal to expectation.", v.Title)
			t.Logf("Expected: %+v", v.Expected)
			t.Logf("Actual: %+v", v.Actual)
		}
	}
}

type RenderContextEqualTestCase struct {
	Title    string
	Expected bool
	Actual   bool
}

func TestRenderContextEqual(t *testing.T) {
	ctx := RenderContext{
		Fields: map[string]Field{
			"Type": Field{
				Value: "Model",
			},
		},
	}

	cases := []RenderContextEqualTestCase{
		RenderContextEqualTestCase{
			Title:    "Equal Case",
			Expected: true,
			Actual:   ctx.Equal("Type", "Model"),
		},
		RenderContextEqualTestCase{
			Title:    "Not Equal Case",
			Expected: false,
			Actual:   ctx.Equal("Type", "View"),
		},
	}

	for _, v := range cases {
		if v.Actual != v.Expected {
			t.Errorf("Result of \"%s\" on RenderContext.Equal() not equal to expectation.", v.Title)
			t.Logf("Expected: %+v", v.Expected)
			t.Logf("Actual: %+v", v.Actual)
		}
	}
}
