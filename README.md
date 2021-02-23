[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE)
[![Build status](https://github.com/hoto/template-go-web-server/workflows/Build%20and%20test/badge.svg?branch=master)](https://github.com/hoto/template-go-web-server/actions)
[![Release](https://img.shields.io/github/release/hoto/template-go-web-server.svg?style=flat-square)](https://github.com/hoto/template-go-web-server/releases/latest)
[![Powered By: goreleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser/goreleaser)

# Golang REST web server template repository

Uses [echo](https://github.com/labstack/echo) as the web framework.

---

Clone and replace `template-go-web-server` with the name of your project.

Github action will use `goreleaser` to automatically releases formula to homebrew repo at [hoto/homebrew-repo](https://github.com/hoto/homebrew-repo) on every git tag.

### Installation
    
Mac:

    brew install hoto/repo/template-go-web-server

Mac or Linux:

    sudo curl -L \
      "https://github.com/hoto/template-go-web-server/releases/download/1.0.0/template-go-web-server_1.0.0_$(uname -s)_$(uname -m)" \
       -o /usr/local/bin/template-go-web-server

    sudo chmod +x /usr/local/bin/template-go-web-server
    
Or manually download binary from [releases](https://github.com/hoto/template-go-web-server/releases).
    
### Run

    template-go-web-server --help
    template-go-web-server --version
    template-go-web-server --port 5000

    curl localhost:5000

### Development

Build and test:

    go get github.com/hoto/template-go-web-server
    
    make dependencies build test
    
Build binary:

     make build
    ./bin/template-go-web-server

Run with arguments:

    make args="myargs" run

Install to global golang bin directory:

    make install
    template-go-web-server
    
    