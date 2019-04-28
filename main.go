package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/manifoldco/promptui"

	"github.com/BurntSushi/toml"
)

func main() {
	// Load Config
	conf := loadConfig("config.toml")

	// Prompt for selecting task
	selectedTaskName := promptSelectTask(conf)

	// Find task entity by task name
	task, err := findTaskByName(conf, selectedTaskName)
	if err != nil {
		panic(err)
	}

	// Prompt for task info
	filledTask := promptTaskFields(task, os.Stdin)

	// ExecuteTask
	executeTask(filledTask)
}

func loadConfig(path string) Config {
	var conf Config
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		panic(err)
	}
	return conf
}

func promptSelectTask(conf Config) string {
	taskList := make([]string, 0)
	for _, v := range conf.Tasks {
		taskList = append(taskList, v.Name)
	}
	prompt := promptui.Select{
		Label: "Select task",
		Items: taskList,
	}
	_, result, err := prompt.Run()
	if err != nil {
		panic(err)
	}
	return result
}

func findTaskByName(conf Config, name string) (Task, error) {
	for _, v := range conf.Tasks {
		if v.Name == name {
			return v, nil
		}
	}
	return Task{}, errors.New("Task not found")
}

func promptTaskFields(task Task, input *os.File) Task {
	reader := bufio.NewReader(input)
	// Collect key and sort
	keys := make([]string, 0)
	for k := range task.Fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Prompt for fields
	for _, k := range keys {
		v := task.Fields[k]
		if v.Type == "StringInput" {
			// Print Message
			fmt.Print(v.Message + ": ")

			// Read input
			inputText, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			inputText = strings.TrimSpace(inputText)

			// Convert text
			if mappedText, ok := task.Fields[k].Map[inputText]; ok {
				inputText = mappedText
			}

			// Update model
			task.Fields[k] = Field{
				Type:    v.Type,
				Message: v.Message,
				Value:   inputText,
				Map:     v.Map,
			}
		}
	}
	return task
}

func executeTask(task Task) {
	for _, file := range task.Files {
		// Load and parse template file
		tmplByte, err := ioutil.ReadFile(file.TemplatePath)
		if err != nil {
			panic(err)
		}
		tmpl, err := template.New(file.TemplatePath).Parse(string(tmplByte))
		if err != nil {
			panic(err)
		}

		// Parse export path
		ePath, err := template.New(file.TemplatePath + "_exportpath").Parse(file.ExportPath)
		if err != nil {
			panic(err)
		}

		// Ready Context
		ctx := RenderContext{
			Template: file,
			Fields:   task.Fields,
		}

		// Render export path
		var b bytes.Buffer
		wtr := bufio.NewWriter(&b)
		err = ePath.Execute(wtr, &ctx)
		if err != nil {
			panic(err)
		}
		wtr.Flush()
		exportPath := b.String()

		// Make Directory
		dir, _ := filepath.Split(file.ExportPath)
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}

		// Create File
		wf, err := os.Create(exportPath)
		if err != nil {
			panic(exportPath)
		}

		// Render file
		err = tmpl.Execute(wf, &ctx)
		if err != nil {
			panic(err)
		}
	}
}
