# Code Gen

This tool built for generating code template via CLI.

### Build & Run

go build && ./codegen

### Setup Config

Create a config.toml with following format and place next to ./codegen

```
[[tasks]]
name = "Create a MVC Class"

[tasks.fields]

    [tasks.fields.ClassName]
    type = "StringInput"
    message = "Class Name"

    [tasks.fields.MyVariable]
    type = "String"
    value = "MY_VARIABLE"

    [tasks.fields.ClassType]
    type = "StringInput"
    message = "Class Type (M=Model, V=View, C=Controller)"

    [tasks.fields.Map]
    M = "Model"
    V = "View"
    C = "Controller"

[[tasks.files]]
templatepath = "test/template/template.java"
exportpath = "test/output/{{.Title \"ClassName\"}}.java"
```

### Create Template

Template are renderred by [Go Template](https://golang.org/pkg/text/template/), like this:

```
{{if .Equal "ClassName" "MyClass"}}// MyClass{{end}}
public class {{.Print "ClassName"}} {
    private static final String {{.Constant "MyVariable"}} = "HelloWorld";
    public static void main(String args[]){
        System.out.println({{.Print "MyVariable"}});
    }
}
```

### Function availiable in template

| Name      | Description                          | Example                                   | Output           |
| --------- | ------------------------------------ | ----------------------------------------- | ---------------- |
| Print     | Print field value                    | {{.Print "MyField"}}                      | your exact input |
| Equal     | Check if field value equal to string | {{if .Equal "MyField" "foo"}} bar {{end}} | bar              |
| Title     | Turn value to title case             | {{.Title "MyField"}}                      | ConvertedInput   |
| Camel     | Turn value to camel case             | {{.Camel "MyField"}}                      | convertedInput   |
| Constant  | Turn value to constant case          | {{.Constant "MyField"}}                   | CONVERTED_INPUT  |
| Underline | Turn value to underline case         | {{.Underline "MyField"}}                  | converted_input  |
| Dash      | Turn value to dash case              | {{.Dash "MyField"}}                       | converted-input  |
| Path      | Turn value to path case              | {{.Path "MyField"}}                       | converted/input  |
| Package   | Turn value to package case           | {{.Package "MyField"}}                    | converted.input  |

# Roadmap

### 28/4/2019

- [x] Basic template generation
- [x] CLI menu & prompt
- [x] Template config format
- [x] Print variable to template
- [x] String conversion for template
      (Title, Constant, Camel, Dash)
- [x] If case for template (Value compare)

### 29/4/2019

- [x] Global variable

### 5/5/2019

- [x] More String conversion for template
      (package, file/path, under_line)

### TBC

- [ ] Windows CLI word missing bug
- [ ] Windows <? -> &lt;? bug
