package main

import(
    "fmt"
    "time"
    "runtime"
)

func main(){
    for i:=0; i<10; i++ {
        go func(){
            fmt.Println("in the go routine")
        }()
        fmt.Println("in the main routine")
        time.Sleep(100 * time.Millisecond)
    }

    fmt.Printf("NumCPU %d\n", runtime.NumCPU())
    fmt.Printf("NumGoroutine %d\n", runtime.NumGoroutine())
}

