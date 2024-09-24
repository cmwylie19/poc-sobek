package main

import (
	"fmt"
	"log"
	"os"

	"github.com/grafana/sobek"
)

// https://github.com/grafana/sobek/?tab=readme-ov-file
func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: %s <controller-path> <module-path> <hash>", os.Args[0])
	}

	controllerPath := os.Args[1]
	modulePath := os.Args[2]
	hash := os.Args[3]

	vm := sobek.New()

	_, err := vm.RunString(fmt.Sprintf(`const jsProgramPath = "%s";`, modulePath))
	if err != nil {
		log.Fatalf("Error setting jsProgramPath: %v", err)
	}

	_, err = vm.RunString(fmt.Sprintf(`const hash = "%s";`, hash))
	if err != nil {
		log.Fatalf("Error setting hash: %v", err)
	}
	controllerScript, err := os.ReadFile(controllerPath)
	if err != nil {
		log.Fatalf("Failed to read controller.js file %s: %v", controllerPath, err)
	}

	_, err = vm.RunString(string(controllerScript))
	if err != nil {
		log.Fatalf("Failed to execute controller.js: %v", err)
	}

	fmt.Println("Controller script executed successfully")
}
