### Race Detector

* Go Provides race detector tool for finding race conditions in Go code

* Binary needs to be race enabled
* When racy behaviour is detected a warning is printed
* Race enabled binary will 10 times slower and consume 10 times more memory
* Integration tests and load tests are good candidates to test with binary with race enabled

    ```shellscript
    > go test -race mypkg // test the package
    > go run -race mysrc.go // compile and run the program
    > go build -race mycmd // build the command
    > go install -race mypkg // install the package
    ```