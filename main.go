package main

import "github.com/hasura/ndc-sdk-go/connector"

func main() {
	if err := connector.Start[Configuration, State](
		&Connector{},
		connector.WithMetricsPrefix("ndc-es-go"),
		connector.WithDefaultServiceName("ndc-es-go"),
	); err != nil {
		panic(err)
	}
}
