package cli

import (
    "fmt"

    "github.com/AlecAivazis/survey/v2"
    "github.com/ayushsharma74/boiler/internal/config"
)

// ParseFlagsAndPrompt parses the CLI flags and interactively prompts the user for missing input.
func ParseFlagsAndPrompt(repos map[string]config.RepositoryInfo) (string, string, string, error) {
    var projectName string
    prompt := &survey.Input{
        Message: "Enter the project name",
    }
	err := survey.AskOne(prompt, &projectName)
	if err != nil {
		return "", "", "", fmt.Errorf("error during project name input: %w",err)
	}
    repoChoice, err := promptForRepo("Choose a boilerplate:", repos, "basic-express")
    if err != nil {
        return "", "", "", fmt.Errorf("error during repo selection: %w",err)
    }

    database, err := promptForChoice("Choose a database:", []string{"mongodb", "postgresql", "none"}, "none")
    if err != nil {
        return "", "", "", fmt.Errorf("error during database selection: %w", err)
    }

    return projectName, repoChoice, database, nil
}

func promptForRepo(prompt string, repos map[string]config.RepositoryInfo, defaultChoice string) (string, error){
    var repoKeys []string
    for key := range repos {
        repoKeys = append(repoKeys, key)
    }

    var selection string
    promptStruct := &survey.Select{
        Message: prompt,
        Options: repoKeys,
        Default: defaultChoice,
    }
	err := survey.AskOne(promptStruct, &selection)
	if err != nil {
		return "", fmt.Errorf("error during prompting: %w", err)
	}
    return selection, nil
}

func promptForChoice(prompt string, choices []string, defaultChoice string) (string, error) {
    var selection string
    promptStruct := &survey.Select{
        Message: prompt,
        Options: choices,
        Default: defaultChoice,
    }
	err := survey.AskOne(promptStruct, &selection)
	if err != nil {
		return "", fmt.Errorf("error during prompting: %w", err)
	}
    return selection, nil
}