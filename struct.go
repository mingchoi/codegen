package main

import "strings"

type Config struct {
	Tasks []Task
}

type Task struct {
	Name   string
	Fields map[string]Field
	Files  []Template
}

type Template struct {
	TemplatePath string
	ExportPath   string
}

type Field struct {
	Type    string
	Message string
	Value   string
	Map     map[string]string
	Case    string
}

type RenderContext struct {
	Template Template
	Fields   map[string]Field
}

func (ctx *RenderContext) Print(name string) string {
	field, ok := ctx.Fields[name]
	if !ok {
		return "Error variable not found"
	}
	return field.Value
}

func (ctx *RenderContext) Equal(fieldName string, value string) bool {
	field, ok := ctx.Fields[fieldName]
	if !ok {
		return false
	}
	return field.Value == value
}

func (ctx *RenderContext) Constant(name string) string {
	field, ok := ctx.Fields[name]
	if !ok {
		return "ERROR_VARIABLE_NOT_FOUND"
	}
	result := strings.ToUpper(field.Value)
	result = strings.Replace(result, " ", "_", -1)
	return result
}

func (ctx *RenderContext) Title(name string) string {
	field, ok := ctx.Fields[name]
	if !ok {
		return "ErrorVariableNotFound"
	}
	str := strings.Title(field.Value)
	strs := strings.Split(str, " ")
	result := strings.Join(strs, "")
	return result
}

func (ctx *RenderContext) Camel(name string) string {
	field, ok := ctx.Fields[name]
	if !ok {
		return "errorVariableNotFound"
	}
	str := strings.Title(field.Value)
	strs := strings.Split(str, " ")
	result := strings.Join(strs, "")
	return strings.ToLower(result[0:1]) + result[1:]
}

func (ctx *RenderContext) Dash(name string) string {
	field, ok := ctx.Fields[name]
	if !ok {
		return "error-variable-not-found"
	}
	str := strings.ToLower(field.Value)
	strs := strings.Split(str, " ")
	result := strings.Join(strs, "-")
	return result
}
