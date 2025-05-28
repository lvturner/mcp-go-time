package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"Time Server",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	dateTool := mcp.NewTool("get_local_date",
		mcp.WithDescription("Get the current local date"),
	)
	s.AddTool(dateTool, dateToolHandler)

	timeTool := mcp.NewTool("get_local_time",
		mcp.WithDescription(("Get the current local time")),
	)
	s.AddTool(timeTool, timeToolHandler)
}

func dateToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(fmt.Sprintf("The current local date is %s", time.DateOnly)), nil
}

func timeToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(fmt.Sprintf("The current local date is %s", time.TimeOnly)), nil
}