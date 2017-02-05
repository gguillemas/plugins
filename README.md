# Plugins

## Description

Small library for building plugin-based applications in Go.

This repository also includes a tool for assisting in building plugins in build systems.

## Requirements

This package requires Go 1.8 or newer.

A release candidate version of Go 1.8 can be downloaded by running:

```
$ go get golang.org/x/build/version/go1.8rc3
```

## Example

A simple example can be found in the "_examples" directory.

The example plugins can be compiled with the provided tool:

```
$ plugins-build -g go1.8rc3 -o /tmp/hello/plugins/ _examples/hello/plugins/
```

And then loaded by the example application:

```
$ go1.8rc3 run _examples/hello/hello.go /tmp/hello/plugins/
Available plugins:
        NAME    VERSION DESCRIPTION
        english 0.0.1   This plugin says "hello" in English.
        spanish 0.0.1   This plugin says "hello" in Spanish.
```

```
$ go1.8rc3 run _examples/hello/hello.go /tmp/hello/plugins/ spanish
Hola!
```

```
$ go1.8rc3 run _examples/hello/hello.go /tmp/hello/plugins/ english
Hello!
```
