# Dragon

<!-- MarkdownTOC autolink=true autoanchor=true-->

- [start dragon](#start-dragon)
- [ab performance](#ab-performance)

<!-- /MarkdownTOC -->



![CI Status](https://travis-ci.org/azerothyang/dragon.svg?branch=master)

 Dragon 🐲 is a lightweight high performance web framework with [Go](https://golang.org/) for the feature and comfortable develop.
 
# components 
1. [httprouter](https://github.com/julienschmidt/httprouter). HttpRouter is a lightweight high performance HTTP request router.
2. [GORM](https://github.com/jinzhu/gorm). The fantastic ORM library for Golang, aims to be developer friendly.

# start dragon
 dragon is mvc go framework. So you can hack with controller and model. 
 __config__:
 > dir release/conf,   .env ***debug/prod***
 > or you can set env DRAGON  in docker
 
 __build__: 
 >Just compile your src file and move bin file to directory dragon/release/
 
 >you might develop with [fswatch](https://github.com/codeskyblue/fswatch) for hot rebuilding.
 just create .fsw.json config:
 ```
{
  "desc": "Auto generated by fswatch [dragon]",
  "triggers": [
    {
      "name": "",
      "pattens": [
        "**/*.go"
      ],
      "env": {
        "DEBUG": "1"
      },
      "cmd": "go fmt dragon... && go build -o ./release/dragon && ./release/dragon",
      "shell": true,
      "delay": "100ms",
      "stop_timeout": "500ms",
      "signal": "KILL",
      "kill_signal": ""
    }
  ],
  "watch_paths": [
    "."
  ],
  "watch_depth": 10
}
 
 ```
 > just run cmd
 ```
 $ fswatch
 ```
 
 __config__:
> dragon/release/conf/  

# set ide env
if you use goland ide, can set
```
Output directory: <project_dir>/release
Working directory: <project_dir>/release
```
>Because default dragon exec file is output in release dir, so the dirs are resolved according to release dir.

>After that, feel free to run your app and start hacking :)

# ab performance
 ``` 
 cpu: 1 core, ram: 1 G
 Server Software:        nginx/1.12.2
 Server Hostname:        test.com
 Server Port:            80
 
 Document Path:          /
 Document Length:        13 bytes
 
 Concurrency Level:      100
 Time taken for tests:   9.341 seconds
 Complete requests:      100000
 Failed requests:        0
 Write errors:           0
 Total transferred:      17700000 bytes
 HTML transferred:       1300000 bytes
 Requests per second:    10705.75 [#/sec] (mean)
 Time per request:       9.341 [ms] (mean)
 Time per request:       0.093 [ms] (mean, across all concurrent requests)
 Transfer rate:          1850.51 [Kbytes/sec] received
 
 Connection Times (ms)
               min  mean[+/-sd] median   max
 Connect:        0    1   1.2      1       9
 Processing:     0    8   2.7      8      30
 Waiting:        0    7   2.5      6      29
 Total:          0    9   2.9      9      31
 
 Percentage of the requests served within a certain time (ms)
   50%      9
   66%     10
   75%     11
   80%     11
   90%     13
   95%     14
   98%     16
   99%     18
  100%     31 (longest request)

 ```

app will read DRAGON env first to check prod,dev or test mode

