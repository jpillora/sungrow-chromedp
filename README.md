# sungrow-chromedp

`sungrow-chromedp` is a small CLI tool to drive your local chrome installation to login to Sungrow website and extract your current solar output in kwh

[![GoDev](https://img.shields.io/static/v1?label=godoc&message=reference&color=00add8)](https://pkg.go.dev/github.com/jpillora/sungrow-chromedp)
[![CI](https://github.com/jpillora/sungrow-chromedp/workflows/CI/badge.svg)](https://github.com/jpillora/sungrow-chromedp/actions?workflow=CI)

### Install

**Binaries**

<!-- NOTE: these badges only work on public repos -->

[![Releases](https://img.shields.io/github/release/jpillora/sungrow-chromedp.svg)](https://github.com/jpillora/sungrow-chromedp/releases)
[![Releases](https://img.shields.io/github/downloads/jpillora/sungrow-chromedp/total.svg)](https://github.com/jpillora/sungrow-chromedp/releases)

Download [the latest pre-compiled binaries here](https://github.com/jpillora/sungrow-chromedp/releases/latest) or install it now with `curl https://i.jpillora.com/jpillora/sungrow-chromedp! | bash`

**Source**

```sh
$ go install github.com/jpillora/sungrow-chromedp@latest
```

### Usage

```
sungrow-chromedp --help
```

```
  Usage: sungrow-chromedp [options]

  Options:
  --portal, -p        sungrow portal (default https://www.isolarcloud.com, env PORTAL)
  --email, -e         sungrow username (env EMAIL)
  --pass              sungrow email (env PASSWORD)
  --timeout, -t       overall timeout (default 30s)
  --no-headless, -n   disable headless mode
  --show-network, -s  show network requests
  --debug, -d         show chromedp actions
  --version, -v       display version
  --help, -h          display help

  Version:
    0.0.0-src
```
### TODO

* Support multiple plants
* Page eval to make arbitrary API calls

### LICENSE

See [LICENSE](./LICENSE)