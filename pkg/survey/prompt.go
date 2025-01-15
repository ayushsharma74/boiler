package survey
import (
    "fmt"
    "github.com/AlecAivazis/survey/v2"
)

func AskOne(p survey.Prompt, response interface{}, opts ...survey.AskOpt) error {
    err := survey.AskOne(p, response, opts...)
    if err != nil {
        return fmt.Errorf("error during prompt: %w", err)
    }
    return nil
}