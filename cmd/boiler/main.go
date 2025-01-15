package main

import (
	"fmt"
	"os"

    "github.com/ayushsharma74/boiler/internal/cli"
    "github.com/ayushsharma74/boiler/internal/config"
    "github.com/ayushsharma74/boiler/internal/git"
    "github.com/ayushsharma74/boiler/internal/template"
)


func main() {
	projectName, repoChoice, database, err := cli.ParseFlagsAndPrompt(config.Repositories)
    if err != nil {
        fmt.Println("Error during CLI interaction:", err)
        os.Exit(1)
    }
    fmt.Println("\nSummary:")
    fmt.Println("Project Name:", projectName)
    fmt.Println("Boilerplate:",repoChoice)
	fmt.Println("Database:", database)
	
    projectPath := "."+string(os.PathSeparator)+projectName
    err = git.CloneRepository(config.Repositories[repoChoice].URL,projectPath)
	if err != nil {
        fmt.Println("Error cloning repository:", err)
        os.Exit(1)
    }

    if database != "none" {
        err = template.UpdatePackageJson(projectPath, database)
        if err != nil {
            fmt.Println("Error updating package.json:", err)
            os.Exit(1)
        }
    }
	fmt.Println("Project created successfully at:", projectPath)
}