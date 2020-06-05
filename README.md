# Airport

A simple program for redirecting stdout and stdin to network ports.

![Go version](https://img.shields.io/github/go-mod/go-version/willbarkoff/airport?logo=go&style=flat-square)
[![Latest release](https://img.shields.io/github/v/tag/willbarkoff/airport?label=latest%20release&sort=semver&style=flat-square)](https://github.com/willbarkoff/airport/releases)
[![License](https://img.shields.io/github/license/willbarkoff/airport?style=flat-square)](./LICENSE.md)

- [⬇️ **Download**](https://github.com/willbarkoff/airport/releases) 
- [🐛 **Report a bug**](https://github.com/willbarkoff/airport/issues/new)

---

## Examples

> ⚠️ `bash` is used in this example just as a simple program. **You should not use bash or another shell with airport, as it will expose your computer over the network**.

Using airport is super simple. It takes one argument, the command to run. For example:

```shell
$ airport /bin/bash
```

Then, in another terminal window:
```shell
$ nc localhost 8080
cat hello.txt
Hello, world
```

To specify a port, simply use the `address` flag.

```shell
$ airport --address :1234 /bin/bash
```

As of right now, `stdin` and `stdout` are redirected over the network, while `stderr` is output to the local computer's stderr. If you want you can change this with the `--redirectStderr` flag, for example:

```shell
$ airport --address :1234 --redirectStderr /bin/bash
```

## Usage
```shell
$ airport --help
Usage of airport:
  -address string
        the address to listen on (default ":8080")
  -redirectStderr
        redirects stderr over the network
```