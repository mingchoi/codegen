package main

import (
	"reflect"
	"testing"
)

func TestSplitWord(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Split", args{"Hello Happy World"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{"Hello_Happy_World"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{"HELLO_HAPPY_WORLD"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{"Hello-Happy-World"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{"Hello.Happy.World"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{"Hello/Happy/World"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{"Hello\\Happy\\World"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{"helloHappyWorld"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{"HelloHappyWorld"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{".HelloHappy World"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{"HelloHappy World_"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{"-Hello._/:Happy World_"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{".-Hello._/:Happy World_"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{".  Hello._/:Happy World_"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{".  Hello._/:Happy World_* -"}, []string{"Hello", "Happy", "World"}},
		{"Split", args{"1Hello._/:HApPY2 World3"}, []string{"1", "Hello", "H", "Ap", "Py2", "World3"}},
		{"Split", args{" 9h37Y2_"}, []string{"9", "H37", "Y2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitWord(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToConstantCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ToConstantCase", args{"HappyHelloWorld"}, "HAPPY_HELLO_WORLD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToConstantCase(tt.args.str); got != tt.want {
				t.Errorf("ToConstantCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUnderlineCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ToUnderlineCase", args{"HappyHelloWorld"}, "happy_hello_world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUnderlineCase(tt.args.str); got != tt.want {
				t.Errorf("ToUnderlineCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToTitleCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ToTitleCase", args{"Happy Hello World"}, "HappyHelloWorld"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToTitleCase(tt.args.str); got != tt.want {
				t.Errorf("ToTitleCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToCamelCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ToCamelCase", args{"Happy Hello World"}, "happyHelloWorld"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelCase(tt.args.str); got != tt.want {
				t.Errorf("ToCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDashCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ToCamelCase", args{"Happy Hello World"}, "happy-hello-world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDashCase(tt.args.str); got != tt.want {
				t.Errorf("ToDashCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPathCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ToPathCase", args{"Happy Hello World"}, "happy/hello/world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPathCase(tt.args.str); got != tt.want {
				t.Errorf("ToPathCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToPackageCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"ToPackageCase", args{"Happy Hello World"}, "happy.hello.world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToPackageCase(tt.args.str); got != tt.want {
				t.Errorf("ToPackageCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
