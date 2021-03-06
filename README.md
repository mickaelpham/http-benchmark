# Bench

Dipping into benchmarking HTTP servers. No, I do not pretend to know what I'm
doing, but I'm just interested in experimenting with tools and programming
languages.

## TL;DR

| Language | Framework                                        | Requests/sec | Transfer/sec |
| -------- | ------------------------------------------------ | ------------ | ------------ |
| Java     | [Vert.x](https://vertx.io/)                      | 106,226.41   | 12.81MB      |
| Go       | [net/http](https://golang.org/pkg/net/http/)     | 85,673.56    | 13.44MB      |
| Crystal  | [router.cr](https://github.com/tbrand/router.cr) | 58,458.47    | 8.45MB       |
| Ruby     | [Sinatra](http://sinatrarb.com/)                 | 3,857.18     | 747.73KB     |

## The Setup

Each directory contains idiomatic code (well, at least to my level of
understanding) to create a HTTP server, read the `quotes.txt` file (once), and
respond to `GET /franklin-says` with a random quote from Benjamin Franklin.

## Measurements

The test leverages [wg/wrk](https://github.com/wg/wrk) to benchmark each server
independently.

```
brew install wrk
```

It was executed on a MacBook Pro with the following specs:

```
Model Name:                 MacBook Pro
Model Identifier:           MacBookPro15,2
Processor Name:             Quad-Core Intel Core i7
Processor Speed:            2.7 GHz
Number of Processors:       1
Total Number of Cores:      4
L2 Cache (per Core):        256 KB
L3 Cache:                   8 MB
Hyper-Threading Technology: Enabled
Memory:                     16 GB
```

The command executed for each test was:

```
wrk -t12 -c400 -d30s http://localhost:<APPLICATION_PORT>/franklin-says
```

| Application | Port |
| ----------- | ---- |
| Ruby        | 4567 |
| Go          | 8001 |
| Java        | 8080 |

### Ruby

Start the server:

```
cd ruby-sample
bundle exec ruby app.rb
```

Collect the benchmark results:

```
❯ wrk -t12 -c400 -d30s http://localhost:4567/franklin-says
Running 30s test @ http://localhost:4567/franklin-says
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    33.82ms    7.56ms  67.99ms   61.80%
    Req/Sec   323.14     93.25   600.00     68.58%
  116080 requests in 30.09s, 21.98MB read
  Socket errors: connect 157, read 22441, write 0, timeout 0
Requests/sec:   3857.18
Transfer/sec:    747.73KB
```

### Go

Compile and start the server:

```
cd go-sample
go build app.go
./app
```

Collect the benchmark results:

```
❯ wrk -t12 -c400 -d30s http://localhost:8001/franklin-says
Running 30s test @ http://localhost:8001/franklin-says
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.77ms  443.03us  10.75ms   81.90%
    Req/Sec     7.20k     3.47k   80.14k    62.29%
  2578764 requests in 30.10s, 404.58MB read
  Socket errors: connect 157, read 100, write 0, timeout 0
Requests/sec:  85673.56
Transfer/sec:     13.44MB
```

### Java

Using Java 11:

```
❯ java -version
openjdk version "11.0.8" 2020-07-14
OpenJDK Runtime Environment AdoptOpenJDK (build 11.0.8+10)
OpenJDK 64-Bit Server VM AdoptOpenJDK (build 11.0.8+10, mixed mode)
```

Compile and start the server:

```
cd vertx-sample
./gradlew shadowJar
java -jar build/libs/vertx-sample-all.jar
```

Collect the benchmark results:

```
❯ wrk -t12 -c400 -d30s http://localhost:8080/franklin-says
Running 30s test @ http://localhost:8080/franklin-says
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.31ms    2.58ms 122.99ms   98.98%
    Req/Sec     8.90k     3.93k   16.72k    62.38%
  3197896 requests in 30.10s, 385.76MB read
  Socket errors: connect 157, read 200, write 0, timeout 0
Requests/sec: 106226.41
Transfer/sec:     12.81MB
```

### Crystal

Install dependencies, compile and start the server:

```
cd crystal-sample
shards install
crystal app.cr
```

Collect the benchmark results:

```
❯ wrk -t12 -c400 -d30s http://localhost:8002/franklin-says
Running 30s test @ http://localhost:8002/franklin-says
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.08ms  324.72us  13.65ms   89.59%
    Req/Sec     5.35k     2.86k    9.35k    69.73%
  1759658 requests in 30.10s, 254.24MB read
  Socket errors: connect 157, read 98, write 0, timeout 0
Requests/sec:  58458.47
Transfer/sec:      8.45MB
```

## Note from the Author

If you look at each sample code, the idiomatic Ruby is both clear and extremely
concise. It's definitely a programmer's best friend.
