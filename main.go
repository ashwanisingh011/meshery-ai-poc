package main

import (
	"fmt"
	"strings"
)

func processUserRequest(ai AIProvider, prompt string) {
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("🚀 Routing request to: %s\n", ai.GetProviderName())

	// 1. FIRST we declare and generate the design
	design, err := ai.GenerateDesign(prompt)
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	// 2. THEN we print it (after it has been generated!)
	fmt.Printf("✅ Success! Generated Design Payload:\n%s\n", design)
	fmt.Println(strings.Repeat("-", 60))
}

func main() {
	fmt.Println("🤖 Starting Meshery AI Adapter 'Bring Your Own Model' Prototype...")
	fmt.Println()

	// Scenario 1: User A wants to use Cloud AI and provides an API key
	cloudAI := &OpenAIProvider{
		APIKey: "sk-fake-key-12345",
	}

	// Scenario 2: User B is in a secure bank and uses a local model
	localAI := &OllamaProvider{
		EndPointURL: "http://localhost:11434", // Fixed the capital P here!
		ModelName:   "llama3",
	}

	testPrompt := "Create a secure Istio service mesh configuration."

	// Watch how seamlessly the system swaps between the two models
	processUserRequest(cloudAI, testPrompt)
	processUserRequest(localAI, testPrompt)
}
