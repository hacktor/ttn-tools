# The things Network tools

In this repository a few command line tools for sending/receiving data to and from IoT devices on The Things Network.

# ttn-send

Command line tool for sending bytes to a port on a TTN LoraWan device.

```bash
[laptop] ~/go/src/github.com/hacktor/ttn-tools/ttn-send$ ./ttn-send -msg 1337
 DEBUG ttn-sdk: Connecting to discovery...      Address=discovery.thethings.network:1900
 DEBUG ttn-sdk: Connected to discovery          Address=discovery.thethings.network:1900
 DEBUG ttn-sdk: Finding handler...             
 DEBUG rpc-client: call done                    auth-type=key duration=74.108293ms method=/discovery.Discovery/GetByAppID service-name=hacktor-ttn-app service-version=2.0.5
 DEBUG ttn-sdk: Connecting to MQTT...           Address=tcp://eu.thethings.network:1883
 DEBUG ttn-sdk: Connected to MQTT               Address=tcp://eu.thethings.network:1883
  INFO mqtt: connected                         
  INFO Published downlink message 1337         
 DEBUG ttn-sdk: Disconnecting from MQTT...     
 DEBUG mqtt: disconnecting                     

```
