package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFunctions = append(lambdaFunctions, lambdaFn)
	sparta.Main("SpartaHelloWorld",
		fmt.Sprintf("Test HelloWorld resource command"),
		lambdaFunctions,
		nil,
		nil)
}
