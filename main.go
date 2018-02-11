package main

import (
        "log"
        "net/http"
        _ "net/http/pprof"
        "runtime"
        "time"
        "sync"
)

func bigBytes() *[]byte {
        s := make([]byte, 1000000)
        return &s
}

func main() {
        var wg sync.WaitGroup

        go func() {
                log.Println(http.ListenAndServe("localhost:6060", nil))
        }()

        var mem runtime.MemStats
        runtime.ReadMemStats(&mem)
        log.Println(mem.Alloc)
        log.Println(mem.TotalAlloc)
        log.Println(mem.HeapAlloc)
        log.Println(mem.HeapSys)

        for {
                s := bigBytes()
                if s == nil {
                        log.Println("oh noes")
                }
                time.Sleep(1)
        }

        runtime.ReadMemStats(&mem)
        log.Println(mem.Alloc)
        log.Println(mem.TotalAlloc)
        log.Println(mem.HeapAlloc)
        log.Println(mem.HeapSys)

        wg.Add(1)
        wg.Wait()

}
