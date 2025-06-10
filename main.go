package main

import (
	"context"
	"fmt"
	"os"
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

	dateAndTimeTool := mcp.NewTool("get_local_datetime",
		mcp.WithDescription("Get the current local date and time in ISO 8601 format"),
	)
	s.AddTool(dateAndTimeTool, dateAndTimeToolHandler)

	httpServer := server.NewStreamableHTTPServer(s)
	if err := httpServer.Start(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))); err != nil {
		fmt.Println("Failed to start HTTP server:", err)
	}
	fmt.Println("Started server")
}

func dateToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(fmt.Sprintf("The current local date is %s", time.Now().Format("2006-01-02"))), nil
}

func timeToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(fmt.Sprintf("The current local time is %s", time.Now().Format("15:04:05"))), nil
}

func dateAndTimeToolHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(fmt.Sprintf("The current local date and time is %s", time.Now())), nil
}