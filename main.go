package main

import (
    "fmt"

    "github.com/k0tletka/spigot-coreprotect-prometheus-exporter/config"
)

func main() {

    _, err := config.GetConfiguration()

    if err != nil {
        panic(err)
    }

    fmt.Println("Sample code")
}
