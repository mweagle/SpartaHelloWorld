package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	sparta "github.com/mweagle/Sparta"
	spartaCF "github.com/mweagle/Sparta/aws/cloudformation"
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

	fmt.Fprint(w, "Hello World 🌍")
}

////////////////////////////////////////////////////////////////////////////////
// Main
func main() {

	lambdaFn := sparta.NewLambda(sparta.IAMRoleDefinition{},
		helloWorld,
		nil)

	// Sanitize the name so that it doesn't have any spaces
	stackName := spartaCF.UserScopedStackName("SpartaHelloWorld")
	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFunctions = append(lambdaFunctions, lambdaFn)
	err := sparta.Main(stackName,
		stackName,
		lambdaFunctions,
		nil,
		nil)
	if err != nil {
		os.Exit(1)
	}
}
