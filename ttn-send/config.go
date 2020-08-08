package main

import (
    "log"
    "strconv"
    "github.com/pelletier/go-toml"
)

type Config struct {
    appID        string
    appAccessKey string
    deviceName   string
    devicePort   uint
}

func getConfig(c string) Config {

    cfg := Config{
        appID: "my-cool-app",
        appAccessKey: "ttn-account-v2.some-more-or-less-random-string",
        deviceName: "my-cool-device",
        devicePort: uint(3),
    }
    t, e := toml.LoadFile(c)
    if e != nil {
        log.Printf("No configuration: %v\n", e)
        return cfg
    }

    if t.Has("appID") {
        cfg.appID = t.Get("appID").(string)
    }

    if t.Has("appAccessKey") {
        cfg.appAccessKey = t.Get("appAccessKey").(string)
    }

    if t.Has("deviceName") {
        cfg.deviceName = t.Get("deviceName").(string)
    }

    if t.Has("devicePort") {
        port := t.Get("devicePort").(string)
        devicePort, err := strconv.ParseInt(port, 0, 16)
        if err != nil {
            log.Printf("Error decoding devicePort: %v, using default %v\n", err, cfg.devicePort)
        } else {
            cfg.devicePort = uint(devicePort)
        }
    }

    return cfg
}

