package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/pre-issue-access-token", handlePreIssueAccessToken)

	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// handlePreIssueAccessToken handles the POST request for pre-issue access token action
func handlePreIssueAccessToken(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var requestBody RequestBody
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "invalid_request", "Failed to parse request body")
		return
	}

	// Log the incoming request for debugging
	log.Printf("Received request - FlowID: %s, RequestID: %s, ActionType: %s",
		requestBody.FlowID, requestBody.RequestID, requestBody.ActionType)
	if len(requestBody.Event.Request.AdditionalParams) > 0 {
		log.Printf("Addtional params: %s", requestBody.Event.Request.AdditionalParams[0].Value)
	}

	// Validate action type
	if requestBody.ActionType != "PRE_ISSUE_ACCESS_TOKEN" {
		sendErrorResponse(w, http.StatusBadRequest, "invalid_action_type",
			"Expected actionType to be PRE_ISSUE_ACCESS_TOKEN")
		return
	}

	// Process the request and generate a response
	response := processTokenRequest(requestBody)

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// processTokenRequest processes the token request and returns operations to perform
func processTokenRequest(req RequestBody) SuccessResponse {
	operations := []Operation{}

	// Modify aud claim in the access token according to the first additional param
	operations = append(operations, Operation{
		Op:    "replace",
		Path:  "/accessToken/claims/aud/-",
		Value: req.Event.Request.AdditionalParams[0].Value[0], 
	})

	return SuccessResponse{
		ActionStatus: "SUCCESS",
		Operations:   operations,
	}
}

// sendErrorResponse sends an error response
func sendErrorResponse(w http.ResponseWriter, statusCode int, errorMsg, errorDesc string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		ActionStatus:     "ERROR",
		ErrorMessage:     errorMsg,
		ErrorDescription: errorDesc,
	})
}
