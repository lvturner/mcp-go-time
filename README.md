# mcp-go-time
Simple MCP time functions

Provides the following tools
* get_local_time - to get the local time
* get_local_date - to get the local date

# Usage
Compile it then add it to your MCP host of choice!
```sh
go build
```

Example MCP JSON
```json
{
  "name": "Time",
  "key": "TimeServer",
  "description": "Provides basic functions for dealing with time",
  "command": "/path/to/mcp-go-time/mcp-go-time"
}
```
