package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 1. Define what we expect the user to send us
type AIRequest struct {
	Model  string `json:"model"` // "cloud" or "local"
	Prompt string `json:"prompt"`
}

// 2. Define what we will send back to the user
type AIResponse struct {
	Provider string `json:"provider"`
	Design   string `json:"design"`
	Error    string `json:"error,omitempty"`
}

// 3. This is our API Handler (Similar to how Meshery works)
func generateHandler(w http.ResponseWriter, r *http.Request) {
	// We only want to accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read the JSON body sent by the user
	var req AIRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Initialize our two mock providers
	cloudAI := &OpenAIProvider{APIKey: "sk-fake-key-12345"}
	localAI := &OllamaProvider{EndPointURL: "http://localhost:11434", ModelName: "llama3"}

	// Route to the correct model based on the JSON request!
	var selectedAI AIProvider
	if req.Model == "cloud" {
		selectedAI = cloudAI
	} else if req.Model == "local" {
		selectedAI = localAI
	} else {
		http.Error(w, "Unknown model. Please send 'cloud' or 'local'", http.StatusBadRequest)
		return
	}

	// Generate the design
	design, err := selectedAI.GenerateDesign(req.Prompt)

	// Prepare the response payload
	resp := AIResponse{
		Provider: selectedAI.GetProviderName(),
	}

	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Design = design
	}

	// Send the JSON back to the user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	fmt.Println("🤖 Starting Meshery AI BYOM API on http://localhost:8080...")

	// Create an API endpoint called /api/design
	http.HandleFunc("/api/design", generateHandler)

	// Start the server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
