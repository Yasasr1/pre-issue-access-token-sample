# Pre-Issue Access Token Action Server

A simple Go server that implements the WSO2 Identity Server pre-issue access token action API contract.

## Overview

This server exposes a single POST endpoint (`/pre-issue-access-token`) that handles pre-issue access token events from WSO2 Identity Server. It extracts values from request parameters and uses them to modify the access token's audience claim before the token is issued.

## Features

- Handles pre-issue access token events
- Extracts values from request's `additionalParams`
- Modifies the audience (`aud`) claim of the access token dynamically
- Returns appropriate success, failed, or error responses
- Includes comprehensive logging for debugging

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

1. Clone or navigate to the repository:
```bash
cd /Users/yasasramanayake/go-server
```

2. Run the server:
```bash
go run .
```

The server will start on port 8080.

## API Endpoint

### POST /pre-issue-access-token

Handles pre-issue access token events and modifies the audience claim based on the first additional parameter value.

**Request Body Example:**
```json
{
  "flowId": "Ec1wMjmiG8",
  "requestId": "6783a180f817075313ea3fb01d79c221",
  "actionType": "PRE_ISSUE_ACCESS_TOKEN",
  "event": {
    "request": {
      "grantType": "authorization_code",
      "clientId": "1u31N7of6gCNR9FqkG1neSlsF_Qa",
      "scopes": ["read", "write"]
    },
    "tenant": {
      "id": "2",
      "name": "bar.com"
    },
    "user": {
      "id": "e204849c-4ec2-41f1-8ff7-ec1ebff02821"
    },
    "accessToken": {
      "tokenType": "JWT",
      "claims": [
        {"name": "sub", "value": "e204849c-4ec2-41f1-8ff7-ec1ebff02821"},
        {"name": "iss", "value": "https://localhost:9443/t/foo.com/oauth2/"},
        {"name": "aud", "value": ["original-audience"]},
        {"name": "expires_in", "value": 3600}
      ],
      "scopes": ["read"]
    }
  },
  "allowedOperations": [
    {
      "op": "add",
      "paths": ["/accessToken/claims", "/accessToken/scopes"]
    },
    {
      "op": "replace",
      "paths": ["/accessToken/claims", "/refreshToken/claims/expires_in"]
    }
  ]
}
```

**Success Response Example:**
```json
{
  "actionStatus": "SUCCESS",
  "operations": [
    {
      "op": "replace",
      "path": "/accessToken/claims/aud/-",
      "value": "custom-audience-value"
    }
  ]
}
```

The server extracts the first value from `request.additionalParams[0].value[0]` and uses it to replace the audience claim in the access token.

## Current Implementation

The `processTokenRequest` function in `main.go` currently:

1. **Extracts the first additional parameter value** from the request
2. **Replaces the audience claim** with the extracted value
3. **Logs** the additional parameters for debugging

### Customization Examples

You can extend this function to add more complex logic:

```go
// Add custom claim based on grant type
if req.Event.Request.GrantType == "client_credentials" {
    operations = append(operations, Operation{
        Op:   "add",
        Path: "/accessToken/claims/-",
        Value: CustomClaim{
            Name:  "app_type",
            Value: "service",
        },
    })
}

// Add multiple audience values from different params
if len(req.Event.Request.AdditionalParams) > 1 {
    for i, param := range req.Event.Request.AdditionalParams {
        if len(param.Value) > 0 {
            operations = append(operations, Operation{
                Op:    "add",
                Path:  "/accessToken/claims/aud/-",
                Value: param.Value[0],
            })
        }
    }
}

// Modify expiry based on client ID
if req.Event.Request.ClientID == "trusted-client-id" {
    operations = append(operations, Operation{
        Op:    "replace",
        Path:  "/accessToken/claims/expires_in",
        Value: 7200, // 2 hours for trusted clients
    })
}
```

## Testing

You can test the endpoint using curl:

```bash
curl -X POST http://localhost:8080/pre-issue-access-token \
  -H "Content-Type: application/json" \
  -d '{
    "actionType": "PRE_ISSUE_ACCESS_TOKEN",
    "event": {
      "request": {
        "grantType": "authorization_code",
        "clientId": "test-client",
        "additionalParams": [
          {
            "name": "custom_audience",
            "value": ["https://api.example.com"]
          }
        ]
      },
      "tenant": {
        "id": "1",
        "name": "example.com"
      },
      "accessToken": {
        "tokenType": "JWT",
        "claims": [
          {"name": "sub", "value": "user123"},
          {"name": "aud", "value": ["default-audience"]},
          {"name": "expires_in", "value": 3600}
        ]
      }
    },
    "allowedOperations": [
      {"op": "replace", "paths": ["/accessToken/claims"]}
    ]
  }'
```

## Response Types

The server can return three types of responses:

1. **SUCCESS**: Operations were processed successfully
2. **FAILED**: The token issuance should be denied (e.g., invalid scope)
3. **ERROR**: An error occurred processing the request (HTTP 400, 401, or 500)

## Project Structure

```
.
├── main.go       # Server implementation and request handler
├── models.go     # Data structures matching the OpenAPI spec
├── go.mod        # Go module file
└── README.md     # This file
```

## License

This is a sample implementation for demonstration purposes.
# pre-issue-access-token-sample
