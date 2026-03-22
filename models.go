package main

import "fmt"

// Cloud AI implementation (Open AI)

type OpenAIProvider struct {
	APIKey string
}

func (o *OpenAIProvider) GenerateDesign(prompt string) (string, error) {
	fmt.Printf("[OpenAI Cloud] Processing prompt: '%s'...\n", prompt)
	mockDesign := `{"name": "cloud-service-mesh", "components": [{"type": "Istio", "version": "1.20"}]}`
	return mockDesign, nil
}

func (o *OpenAIProvider) GetProviderName() string {
	return "OpenAI (GPT-4)"
}

// Local AI Implementation (Ollama)

type OllamaProvider struct {
	EndPointURL string
	ModelName   string
}

func (l *OllamaProvider) GenerateDesign(prompt string) (string, error) {
	fmt.Printf("💻 [Ollama Local] Spinning up local GPU for model '%s'...\n", l.ModelName)
	fmt.Printf("💻 [Ollama Local] Processing prompt: '%s'...\n", prompt)

	mockDesign := `{"name": "local-envoy-filter", "components": [{"type": "Envoy", "version": "1.28"}]}`
	return mockDesign, nil
}

func (l *OllamaProvider) GetProviderName() string {
	return fmt.Sprintf("Local Ollama (%s)", l.ModelName)
}
