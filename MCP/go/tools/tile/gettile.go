package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/config"
	"github.com/amazon-sagemaker-geospatial-capabilities/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GettileHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		xVal, ok := args["x"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: x"), nil
		}
		x, ok := xVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: x"), nil
		}
		yVal, ok := args["y"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: y"), nil
		}
		y, ok := yVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: y"), nil
		}
		zVal, ok := args["z"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: z"), nil
		}
		z, ok := zVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: z"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["Arn"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("Arn=%v", val))
		}
		if val, ok := args["ExecutionRoleArn"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ExecutionRoleArn=%v", val))
		}
		if val, ok := args["ImageAssets"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ImageAssets=%v", val))
		}
		if val, ok := args["ImageMask"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("ImageMask=%v", val))
		}
		if val, ok := args["OutputDataType"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("OutputDataType=%v", val))
		}
		if val, ok := args["OutputFormat"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("OutputFormat=%v", val))
		}
		if val, ok := args["PropertyFilters"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("PropertyFilters=%v", val))
		}
		if val, ok := args["Target"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("Target=%v", val))
		}
		if val, ok := args["TimeRangeFilter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("TimeRangeFilter=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/tile/%s/%s/%s#Arn&ImageAssets&Target%s", cfg.BaseURL, x, y, z, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.GetTileOutput
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

func CreateGettileTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_tile_z_x_y#Arn&ImageAssets&Target",
		mcp.WithDescription("Gets a web mercator tile for the given Earth Observation job."),
		mcp.WithString("Arn", mcp.Required(), mcp.Description("The Amazon Resource Name (ARN) of the tile operation.")),
		mcp.WithString("ExecutionRoleArn", mcp.Description("The Amazon Resource Name (ARN) of the IAM role that you specify.")),
		mcp.WithArray("ImageAssets", mcp.Required(), mcp.Description("The particular assets or bands to tile.")),
		mcp.WithBoolean("ImageMask", mcp.Description("Determines whether or not to return a valid data mask.")),
		mcp.WithString("OutputDataType", mcp.Description("The output data type of the tile operation.")),
		mcp.WithString("OutputFormat", mcp.Description("The data format of the output tile. The formats include .npy, .png and .jpg.")),
		mcp.WithString("PropertyFilters", mcp.Description("Property filters for the imagery to tile.")),
		mcp.WithString("Target", mcp.Required(), mcp.Description("Determines what part of the Earth Observation job to tile. 'INPUT' or 'OUTPUT' are the valid options.")),
		mcp.WithString("TimeRangeFilter", mcp.Description("Time range filter applied to imagery to find the images to tile.")),
		mcp.WithNumber("x", mcp.Required(), mcp.Description("The x coordinate of the tile input.")),
		mcp.WithNumber("y", mcp.Required(), mcp.Description("The y coordinate of the tile input.")),
		mcp.WithNumber("z", mcp.Required(), mcp.Description("The z coordinate of the tile input.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GettileHandler(cfg),
	}
}
