package main

type Config struct {
	Tasks  []Task
	Fields map[string]Field
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
	return ToConstantCase(field.Value)
}

func (ctx *RenderContext) Underline(name string) string {
	field, ok := ctx.Fields[name]
	if !ok {
		return "error_variable_not_found"
	}
	return ToUnderlineCase(field.Value)
}

func (ctx *RenderContext) Title(name string) string {
	field, ok := ctx.Fields[name]
	if !ok {
		return "ErrorVariableNotFound"
	}
	return ToTitleCase(field.Value)
}

func (ctx *RenderContext) Camel(name string) string {
	field, ok := ctx.Fields[name]
	if !ok {
		return "errorVariableNotFound"
	}
	return ToCamelCase(field.Value)
}

func (ctx *RenderContext) Dash(name string) string {
	field, ok := ctx.Fields[name]
	if !ok {
		return "error-variable-not-found"
	}
	return ToDashCase(field.Value)
}

func (ctx *RenderContext) Path(name string) string {
	field, ok := ctx.Fields[name]
	if !ok {
		return "error/variable/not/found"
	}
	return ToPathCase(field.Value)
}

func (ctx *RenderContext) Package(name string) string {
	field, ok := ctx.Fields[name]
	if !ok {
		return "error.variable.not.found"
	}
	return ToPackageCase(field.Value)
}
