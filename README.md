# Bench

Dipping into benchmarking HTTP servers. No, I do not pretend to know what I'm
doing, but I'm just interested in experimenting with tools and programming
languages.

## TL;DR

| Language | Requests/sec | Transfer/sec |
| -------- | ------------ | ------------ |
| Ruby     | 3,025.77     | 0.97MB       |
| Go       | 84,765.94    | 14.51MB      |
| Java     | 106,226.41   | 12.81MB      |

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
    Latency    45.76ms   37.02ms 978.51ms   98.04%
    Req/Sec   255.92    110.30   630.00     66.84%
  91087 requests in 30.10s, 29.26MB read
  Socket errors: connect 157, read 20069, write 0, timeout 0
  Non-2xx or 3xx responses: 98
Requests/sec:   3025.77
Transfer/sec:      0.97MB
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
    Latency     2.81ms  368.92us   8.62ms   80.64%
    Req/Sec     7.11k     3.04k   14.42k    55.78%
  2551744 requests in 30.10s, 436.84MB read
  Socket errors: connect 157, read 97, write 0, timeout 0
Requests/sec:  84765.94
Transfer/sec:     14.51MB
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

## Note from the Author

If you look at each sample code, the idiomatic Ruby is both clear and extremely
concise. It's definitely a programmer's best friend.
