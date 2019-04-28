package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	expected := Config{
		Tasks: []Task{
			Task{
				Name: "First Template",
				Fields: map[string]Field{
					"ClassName": Field{
						Type:    "StringInput",
						Message: "Class Name",
					},
					"MyVariable": Field{
						Type:  "String",
						Value: "MY_VARIABLE",
					},
				},
				Files: []Template{
					Template{
						TemplatePath: "test/template/template.java",
						ExportPath:   "test/output/out.java",
					},
				},
			},
		},
	}

	// Start Test
	conf := loadConfig("test/config_testcase.toml")

	if !reflect.DeepEqual(conf, expected) {
		t.Error("Loaded Config not equal to expectation.")
		t.Logf("Expected: %+v", expected)
		t.Logf("Actual: %+v", conf)
	}
}

func TestFindTaskByName(t *testing.T) {
	conf := Config{
		Tasks: []Task{
			Task{Name: "First Task"},
			Task{Name: "Second Task"},
			Task{Name: "Third Task"},
			Task{Name: "Fourth Task"},
		},
	}
	expected := Task{Name: "Third Task"}

	// Start Test
	task, err := findTaskByName(conf, "Third Task")
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(task, expected) {
		t.Error("Finding wrong task")
		t.Logf("Expected: %+v", expected)
		t.Logf("Actual: %+v", task)
	}
}

func TestPromptTaskFields(t *testing.T) {
	task := Task{
		Name: "First Template",
		Fields: map[string]Field{
			"ClassName": Field{
				Type:    "StringInput",
				Message: "Class Name",
			},
			"MyVariable": Field{
				Type:  "String",
				Value: "MY_VARIABLE",
			},
			"MyMap": Field{
				Type:    "StringInput",
				Message: "Class Type (M=Model, V=View, C=Controller)",
				Map: map[string]string{
					"M": "Model",
					"V": "View",
					"C": "Controller",
				},
			},
		},
		Files: []Template{
			Template{
				TemplatePath: "template/template.java",
				ExportPath:   "src/out.java",
			},
		},
	}

	expected := Task{
		Name: "First Template",
		Fields: map[string]Field{
			"ClassName": Field{
				Type:    "StringInput",
				Message: "Class Name",
				Value:   "qwerasdf",
			},
			"MyVariable": Field{
				Type:  "String",
				Value: "MY_VARIABLE",
			},
			"MyMap": Field{
				Type:    "StringInput",
				Message: "Class Type (M=Model, V=View, C=Controller)",
				Value:   "Model",
				Map: map[string]string{
					"M": "Model",
					"V": "View",
					"C": "Controller",
				},
			},
		},
		Files: []Template{
			Template{
				TemplatePath: "template/template.java",
				ExportPath:   "src/out.java",
			},
		},
	}

	// Mock os.Stdin by io.utilTempFile
	tmpFile, _ := ioutil.TempFile("", "example")
	defer os.Remove(tmpFile.Name())
	tmpFile.Write([]byte("qwerasdf\nM\n\n"))
	tmpFile.Seek(0, 0)

	// Start Test
	result := promptTaskFields(task, tmpFile)

	if !reflect.DeepEqual(result, expected) {
		t.Error("Loaded Config not equal to expectation.")
		t.Logf("Expected: %+v", expected)
		t.Logf("Actual: %+v", result)
	}
}

func TestExecuteTask(t *testing.T) {
	task := Task{
		Name: "First Template",
		Fields: map[string]Field{
			"ClassName": Field{
				Type:    "StringInput",
				Message: "Class Name",
				Value:   "MyClass",
			},
			"MyVariable": Field{
				Type:  "String",
				Value: "MY_VARIABLE",
			},
		},
		Files: []Template{
			Template{
				TemplatePath: "test/template/template.java",
				ExportPath:   "test/output/{{.Title \"ClassName\"}}.java",
			},
		},
	}

	// Clear environment
	err := os.RemoveAll("test/output/MyClass.java")
	if err != nil {
		t.Error(err)
	}

	// Start Test
	executeTask(task)

	// Read File
	fileByte, err := ioutil.ReadFile("test/output/MyClass.java")
	if err != nil {
		t.Error(err)
	}

	fileContent := string(fileByte)

	lines := []string{
		"// MyClass",
		"public class MyClass {",
		"private static final String MY_VARIABLE = \"HelloWorld\";",
		"System.out.println(MY_VARIABLE);",
	}

	for _, v := range lines {
		if !strings.Contains(fileContent, v) {
			t.Errorf("Expected \"%s\" in output file", v)
		}
	}
}
