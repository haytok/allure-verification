package main

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestSum(t *testing.T) {

	runner.Run(t, "TestSum 0", func(pt provider.T) {
		actual := sum(3, 5)
		expected := 8

		pt.Assert().Equal(expected, actual)
	})

	runner.Run(t, "TestSum 1", func(pt provider.T) {
		actual := sum(3, 5)
		expected := 9

		pt.Assert().Equal(expected, actual)
	})
}
