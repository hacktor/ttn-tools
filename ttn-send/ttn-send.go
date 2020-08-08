package main

import (
    "flag"
    "encoding/hex"
    ttnsdk "github.com/TheThingsNetwork/go-app-sdk"
    ttnlog "github.com/TheThingsNetwork/go-utils/log"
    "github.com/TheThingsNetwork/go-utils/log/apex"
    "github.com/TheThingsNetwork/ttn/core/types"
)

const (
    clientName = "hacktor-ttn-app"
)

func main() {

    var (
        conf string
        appID string
        appAccessKey string
        deviceName string
        devicePort uint
        msg string
    )

    cfg := getConfig("ttn.toml")

    flag.StringVar(&conf, "conf", "ttn.toml", "path to TOML configuration file")
    flag.StringVar(&appID, "appID", cfg.appID, "TTN Application ID")
    flag.StringVar(&appAccessKey, "appAccessKey", cfg.appAccessKey, "Application Access Key")
    flag.StringVar(&deviceName, "deviceName", cfg.deviceName, "Device name")
    flag.UintVar(&devicePort, "devicePort", cfg.devicePort, "Output port on device")
    flag.StringVar(&msg, "msg", "deadbeef", "Hexadecimal message")
    flag.Parse()

    log := apex.Stdout() // We use a cli logger at Stdout
    log.MustParseLevel("debug")
    ttnlog.Set(log) // Set the logger as default for TTN

    payload, err := hex.DecodeString(msg)
    if err != nil {
        log.WithError(err).Fatalf("%v: msg is not a hexadecimal string", clientName)
    }

    // Create a new SDK configuration for the public community network
    config := ttnsdk.NewCommunityConfig(clientName)
    config.ClientVersion = "2.0.5" // The version of the application

    // Create a new SDK client for the application
    client := config.NewClient(appID, appAccessKey)

    // Make sure the client is closed before the function returns
    defer client.Close()

    // Start Publish/Subscribe client (MQTT)
    pubsub, err := client.PubSub()
    if err != nil {
        log.WithError(err).Fatalf("%v: could not get application pub/sub", clientName)
    }

    // Make sure the pubsub client is closed before the function returns
    defer pubsub.Close()

    // Get a publish/subscribe client scoped to deviceName
    myDevicePubSub := pubsub.Device(deviceName)

    // Make sure the pubsub client for this device is closed before the function returns
    defer myDevicePubSub.Close()

    // Publish downlink message
    err = myDevicePubSub.Publish(&types.DownlinkMessage{
        PayloadRaw: payload,
        FPort:      uint8(devicePort),
        Schedule:   types.ScheduleLast, // allowed values: "replace" (default), "first", "last"
        Confirmed:  true,               // can be left out, default is false
    })
    if err != nil {
        log.WithError(err).Fatalf("%v: could not schedule downlink message", clientName)
    }
    log.Info("Published downlink message " + msg)
}
