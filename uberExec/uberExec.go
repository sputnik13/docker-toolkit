package main

import (
	"os"
	"os/exec"
	"strings"
	"syscall"
	"text/template"
)

func generateConfig(configTemplatePath string, configFilePath string) {
	envMap := make(map[string]string)
	for _, envVar := range os.Environ() {
		split := strings.SplitN(envVar, "=", 2)
		envMap[split[0]] = split[1]
	}

	environment := struct {
		Environment map[string]string
	}{
		envMap,
	}

	t, err := template.ParseFiles(configTemplatePath)
	if err != nil {
		panic(err)
	}

	if configFilePath == "" {
		configFilePath = "config.out"
	}
	f, err := os.Create(configFilePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = t.Execute(f, environment)
	if err != nil {
		panic(err)
	}
}

func main() {
	configTemplatePath := os.Getenv("UBEREXEC_CONFIGTEMPLATE")
	configFilePath := os.Getenv("UBEREXEC_CONFIGPATH")

	if configTemplatePath != "" {
		generateConfig(configTemplatePath, configFilePath)
	}

	if len(os.Args) > 1 {
		binary, lookErr := exec.LookPath(os.Args[1])
		if lookErr != nil {
			panic(lookErr)
		}
		syscall.Exec(binary, os.Args[1:], os.Environ())
	}
}
