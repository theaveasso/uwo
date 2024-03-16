package main

import "log"

func main() {
    app := newApp()

    if err := app.start(); err != nil {
        log.Println(err)
    }
}
