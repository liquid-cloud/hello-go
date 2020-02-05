# Tested using Istio `1.4.3`.

### Istio using HTTPS
```
# siege -c128 -d1 -t 10S https://https-hello-go.<redacted>
** SIEGE 4.0.4
** Preparing 128 concurrent users for battle.
The server is now under siege...
Lifting the server siege...
Transactions:                    702 hits
Availability:                 100.00 %
Elapsed time:                   9.05 secs
Data transferred:               0.01 MB
Response time:                  0.12 secs
Transaction rate:              77.57 trans/sec
Throughput:                     0.00 MB/sec
Concurrency:                    9.27
Successful transactions:         785
Failed transactions:               0
Longest transaction:            0.38
Shortest transaction:           0.07
```

### Istio using HTTP
```
# siege -c128 -d1 -t 10S http://http-hello-go.<redacted>
** SIEGE 4.0.4
** Preparing 128 concurrent users for battle.
The server is now under siege...
Lifting the server siege...
Transactions:                   2310 hits
Availability:                 100.00 %
Elapsed time:                   9.30 secs
Data transferred:               0.02 MB
Response time:                  0.03 secs
Transaction rate:             248.39 trans/sec
Throughput:                     0.00 MB/sec
Concurrency:                    7.26
Successful transactions:        2310
Failed transactions:               0
Longest transaction:            0.16
Shortest transaction:           0.02
```

### Nginx using HTTPS
```
# siege -c128 -d1 -t 10S https://hello-go.<redacted>
** SIEGE 4.0.4
** Preparing 128 concurrent users for battle.
The server is now under siege...
Lifting the server siege...
Transactions:                   2246 hits
Availability:                 100.00 %
Elapsed time:                   9.28 secs
Data transferred:               0.02 MB
Response time:                  0.04 secs
Transaction rate:             242.03 trans/sec
Throughput:                     0.00 MB/sec
Concurrency:                   10.08
Successful transactions:        2246
Failed transactions:               0
Longest transaction:            0.31
Shortest transaction:           0.02
```

### Nginx using HTTP
```
# siege -c128 -d1 -t 10S http://hello-go.<redacted>
** SIEGE 4.0.4
** Preparing 128 concurrent users for battle.
The server is now under siege...
Lifting the server siege...
Transactions:                   2411 hits
Availability:                 100.00 %
Elapsed time:                   9.43 secs
Data transferred:               0.02 MB
Response time:                  0.02 secs
Transaction rate:             255.67 trans/sec
Throughput:                     0.00 MB/sec
Concurrency:                    4.80
Successful transactions:        2411
Failed transactions:               0
Longest transaction:            0.18
Shortest transaction:           0.00
```
