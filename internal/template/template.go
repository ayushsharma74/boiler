package template

import (
    "fmt"
    "os"
    "path/filepath"
    "text/template"
)
// updatePackageJson updates the package json
func UpdatePackageJson(projectPath string, database string) error{
    packageJsonPath := filepath.Join(projectPath, "package.json")
    packageJsonFile, err := os.OpenFile(packageJsonPath, os.O_RDWR, 0644)
    if err != nil {
        return fmt.Errorf("error opening package.json: %w",err)
    }
    defer packageJsonFile.Close()

    tmpl, err := template.ParseFiles(packageJsonPath)
    if err != nil {
        return fmt.Errorf("error parsing package.json: %w",err)
    }

    var dependencies map[string]interface{}
    switch database {
        case "mongodb":
            dependencies = map[string]interface{} {
                "mongoose": "^8.0.3",
            }
        case "postgresql":
            dependencies = map[string]interface{}{
                "pg": "^8.11.3",
            }
        default:
            dependencies = nil
    }

    if dependencies != nil {
        err = tmpl.Execute(packageJsonFile, map[string]interface{}{"dependencies":dependencies})
        if err != nil {
            return fmt.Errorf("error updating package.json %w",err)
        }
    }
    return nil

}