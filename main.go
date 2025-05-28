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

	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func dateToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(fmt.Sprintf("The current local date is %s", time.Now().Format("2006-01-02"))), nil
}

func timeToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(fmt.Sprintf("The current local time is %s", time.Now().Format("15:04:05"))), nil
}