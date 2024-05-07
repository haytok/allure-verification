package main

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestSum(t *testing.T) {

	runner.Run(t, "TestSum", func(pt provider.T) {
		actual := sum(3, 5)
		expected := 8

		pt.Assert().Equal(expected, actual)
	})
}
