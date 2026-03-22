package main

type AIProvider interface {
	GenerateDesign(prompt string) (string, error)

	GetProviderName() string
}
