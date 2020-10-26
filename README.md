# HTTP Client

A CLI to make HTTP requests to a given endpoint and profile the server's performance.

## Usage

1. Install [Go](https://golang.org/doc/install)

2. Run the program:

```
go run main.go <options>
```

## Options

| Flag        | Argument? | Description                              |
| ----------- | --------- | ---------------------------------------- |
| `--help`    |           | Shows help text                          |
| `--url`     | string    | Specifies the URL to send the request to |
| `--profile` | int       | Specifies the number of requests to make |

## Example

![Cloudflare screenshot](./screenshots/screenshot_cloudflare.PNG)

## Comparing Results Among Popular Sites
