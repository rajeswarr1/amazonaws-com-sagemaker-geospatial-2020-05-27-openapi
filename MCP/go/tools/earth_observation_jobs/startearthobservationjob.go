package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/config"
	"github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func StartearthobservationjobHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/earth-observation-jobs", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.StartEarthObservationJobOutput
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateStartearthobservationjobTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_earth-observation-jobs",
		mcp.WithDescription("Use this operation to create an Earth observation job."),
		mcp.WithString("ClientToken", mcp.Description("Input parameter: A unique token that guarantees that the call to this API is idempotent.")),
		mcp.WithString("ExecutionRoleArn", mcp.Required(), mcp.Description("Input parameter: The Amazon Resource Name (ARN) of the IAM role that you specified for the job.")),
		mcp.WithObject("InputConfig", mcp.Required(), mcp.Description("Input parameter: Input configuration information.")),
		mcp.WithObject("JobConfig", mcp.Required(), mcp.Description("Input parameter: The input structure for the JobConfig in an EarthObservationJob.")),
		mcp.WithString("KmsKeyId", mcp.Description("Input parameter: The Key Management Service key ID for server-side encryption.")),
		mcp.WithString("Name", mcp.Required(), mcp.Description("Input parameter: The name of the Earth Observation job.")),
		mcp.WithObject("Tags", mcp.Description("Input parameter: Each tag consists of a key and a value.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    StartearthobservationjobHandler(cfg),
	}
}
