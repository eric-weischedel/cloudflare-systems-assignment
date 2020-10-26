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

I ran this tool against several web APIs, running 1,000 requests for each one. All APIs tested return a small amount of JSON data.

| Endpoint                                                                               | Median Time | Slowest Time | Fastest Time |
| -------------------------------------------------------------------------------------- | ----------- | ------------ | ------------ |
| [Cloudflare](https://cloudflare-assignment.eric-weischedel.workers.dev/links)          | 13.4735ms   | 50.6964ms    | 11.7629ms    |
| [Heroku](https://hoodat-api.herokuapp.com/api)                                         | 18.5168ms   | 15.0490072s  | 16.9853ms    |
| [Google Firebase](https://fir-realtime-db-sample-6856c.firebaseio.com/categories.json) | 15.0199ms   | 15.044588s   | 13.0005ms    |

According to these results, Heroku and Google have a slowest time greater than 15 seconds! Perhaps this is due to a limiter. I decided to run the tool against a smaller sample size (30) to avoid the limiter.

| Endpoint                                                                               | Median Time | Slowest Time | Fastest Time |
| -------------------------------------------------------------------------------------- | ----------- | ------------ | ------------ |
| [Cloudflare](https://cloudflare-assignment.eric-weischedel.workers.dev/links)          | 13.9481ms   | 63.2818ms    | 12.869ms     |
| [Heroku](https://hoodat-api.herokuapp.com/api)                                         | 18.3276ms   | 55.2036ms    | 17.0059ms    |
| [Google Firebase](https://fir-realtime-db-sample-6856c.firebaseio.com/categories.json) | 14.9993ms   | 82.1894ms    | 13.882ms     |

I suspect these results are more accurate because the slowest times for Heroku and Google are more reasonable. According to the results, Cloudflare has the lowest median response time and the fastest response time. Bravo, Cloudflare.
