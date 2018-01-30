package main

import (
	"context"
	"fmt"
	_ "net/http/pprof" // include pprop
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	sparta "github.com/mweagle/Sparta"
	spartaCF "github.com/mweagle/Sparta/aws/cloudformation"
	"github.com/sirupsen/logrus"
)

/*
• func ()
• func () error
• func (TIn), error
• func () (TOut, error)
• func (context.Context) error
• func (context.Context, TIn) error
• func (context.Context) (TOut, error)
• func (context.Context, TIn) (TOut, error)
*/

// Standard AWS λ function
func helloWorld(ctx context.Context) (string, error) {
	logger, loggerOk := ctx.Value(sparta.ContextKeyLogger).(*logrus.Logger)
	if loggerOk {
		logger.Info("Access structured logger")
	}
	contextLogger, contextLoggerOk := ctx.Value(sparta.ContextKeyRequestLogger).(*logrus.Entry)
	if contextLoggerOk {
		contextLogger.Info("Request scoped log")
	} else if loggerOk {
		logger.Warn("Failed to access scoped logger")
	} else {
		fmt.Printf("Failed to access any logger")
	}
	return "Hello World 🌏", nil
}

////////////////////////////////////////////////////////////////////////////////
// Main
func main() {
	lambdaFn := sparta.HandleAWSLambda("Hello World",
		helloWorld,
		sparta.IAMRoleDefinition{})

	sess := session.Must(session.NewSession())
	awsName, awsNameErr := spartaCF.UserAccountScopedStackName("MyHelloWorldStack",
		sess)
	if awsNameErr != nil {
		fmt.Print("Failed to create stack name\n")
		os.Exit(1)
	}
	// Sanitize the name so that it doesn't have any spaces
	var lambdaFunctions []*sparta.LambdaAWSInfo
	lambdaFunctions = append(lambdaFunctions, lambdaFn)

	err := sparta.Main(awsName,
		"Simple Sparta application that demonstrates core functionality",
		lambdaFunctions,
		nil,
		nil)
	if err != nil {
		os.Exit(1)
	}
}
