# Plugins

## Description

Small library for building plugin-based applications in Go.

This repository also includes a tool for assisting in building plugins in build systems.

**NOTE: This package requires Go version 1.8 or newer.**

# Example

A usage example can be found in the "_examples" directory.

A release candidate version of Go 1.8 can be downloaded by running:

```
$ go get golang.org/x/build/version/go1.8rc3
```

Then, the example plugins can be compiled with:

```
$ plugins-build -g go1.8rc3 -o /tmp/hello/plugins _examples/hello/plugins/
```

And, finally, loaded by the example application:

```
$ go1.8rc3 run _examples/hello/hello.go /tmp/hello/plugins
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
