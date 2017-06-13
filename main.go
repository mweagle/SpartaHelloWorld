package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"strings"

	"github.com/Sirupsen/logrus"
	sparta "github.com/mweagle/Sparta"
)

// Standard AWS λ function
func helloWorld(event *json.RawMessage,
	context *sparta.LambdaContext,
	w http.ResponseWriter,
	logger *logrus.Logger) {

	configuration, _ := sparta.Discover()

	logger.WithFields(logrus.Fields{
		"Discovery": configuration,
	}).Info("Custom resource request")

	fmt.Fprint(w, "Hello World")
}

////////////////////////////////////////////////////////////////////////////////
// Main
func main() {

	lambdaFn := sparta.NewLambda(sparta.IAMRoleDefinition{},
		helloWorld,
		nil)

	userName := os.Getenv("USER")
	if "" == userName {
		userName = os.Getenv("USERNAME")
	}
	// Sanitize the name so that it doesn't have any spaces
	canonicalName := strings.Replace(userName, " ", "", -1)
	stackName := fmt.Sprintf("SpartaHelloWorld-%s", canonicalName)
	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFunctions = append(lambdaFunctions, lambdaFn)
	err := sparta.Main(stackName,
		fmt.Sprintf("Sparta %s for %s", stackName, userName),
		lambdaFunctions,
		nil,
		nil)
	if err != nil {
		os.Exit(1)
	}
}
