package main

import (
        "log"
        "net/http"
        _ "net/http/pprof"
        "runtime"
        "time"
        "sync"
)

type getBytes struct {
        data []byte
}

func newBytes() getBytes {
        size := 1024 * 1024
        gb := getBytes{
                data: make([]byte, size),
        }
        return gb

}


func main() {
        var wg sync.WaitGroup

        go func() {
                log.Println(http.ListenAndServe("localhost:6060", nil))
        }()

        var mem runtime.MemStats
        i := 0
        gbs := make([]getBytes, 0)
        for {
                i++
                gbs = append(gbs, newBytes())
                runtime.ReadMemStats(&mem)
                log.Printf("Alloc:%d | TotalAlloc:%d | HeapAlloc:%d | HeapSys:%d", mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
                time.Sleep(time.Duration(500)*time.Millisecond)
        }
        wg.Add(1)
        wg.Wait()

}
