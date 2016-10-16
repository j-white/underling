Underling [![CircleCI](https://circleci.com/gh/j-white/underling.svg?style=svg)](https://circleci.com/gh/j-white/underling) [![Go Report Card](https://goreportcard.com/badge/j-white/underling)](https://goreportcard.com/report/j-white/underling)
================

Underling is an alternative implementation of the OpenNMS Minion written in Go.

It was built with the aim of:

1. Helping validate the communication channels between the OpenNMS server and the Minion processes.
2. Providing a runtime with smaller disk and memory utilization.
3. Showing how alternate implementations of the Minion could be built without using Java.

## What's currently supported?

* Making SNMP v1/v2c/v3 requests (no traps)

## What's planned?

* Running these detectors:
    * ICMP
    * SNMP
* Running these monitors:
    * ICMP
    * SNMP

## Running from *master*

Here's a guide for hacking and building the binaries from master.

### Dependencies

- Go 1.6

### Get Code

```bash
go get github.com/j-white/underling
```

### Building

```
cd $GOPATH/src/github.com/j-white/underling
go build .
```

### Running

```
cp underling.yaml.sample underling.yaml
./underling
```

## Setting up OpenNMS

Underling uses the STOMP protocol for communicating with OpenNMS's ActiveMQ service.

To enable STOMP, add the following to `etc/opennms-activemq.xml` in the `<transportConnectors>` section:

```xml
<transportConnector name="stomp" uri="stomp+nio://0.0.0.0:61613"/>
```

## Mobile Bindings

```
gomobile bind -target android -o underling.aar -javapkg go.underling ./bindings
```


