package main

import (

	"logToFile/core"

	"github.com/pkg/errors"
)

var logger = core.NewFileLogger()
func main() {

	logger.Print("test")

	logger.Info().Str("Category", "search").Int("DuringTime", 80).Msg("Searching for things ...")

	err := errors.New("this is an error")

	logger.Error().Err(err).Msg("Err occurred")
	err = func3()
	if err != nil {
		logger.Error().Stack().Err(err).Msg("Error occurred (func3 error)")
	}
}

func func1() error {
	return errors.New("this is an error (func1)")
}

func func2() error {
	err := func1()
	if err != nil {
		return err
	}
	return nil
}

func func3() error {
	err := func2()
	if err != nil {
		return err
	}
	return nil
}
