package clientdependency

import (
    "fmt"
    "github.com/jfrog-solutiontest/baddependency"
)

func clientdependency () {
    fmt.Println ("Hello World, main here")
    baddependency.ModuleTestPackageName()
}
