package main

import (
	"context"
	"testing"

	sparta "github.com/mweagle/Sparta"
	"github.com/sirupsen/logrus"
)

func TestHelloWorld(t *testing.T) {
	logger, loggerErr := sparta.NewLogger("info")
	if loggerErr != nil {
		t.Fatal("Failed to initialize logger: ", loggerErr)
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx,
		sparta.ContextKeyLogger,
		logger,
	)
	loggerEntry := logger.WithFields(logrus.Fields{
		"Test": "Field",
	})
	ctx = context.WithValue(ctx,
		sparta.ContextKeyRequestLogger,
		loggerEntry)

	hello, helloErr := helloWorld(ctx)
	if helloErr != nil {
		t.Fatalf("Failed to call helloWorld()")
	}
	if len(hello) <= 0 {
		t.Fatalf("helloWorld() returned an empty value")
	}
}
