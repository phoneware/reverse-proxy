# reverse-proxy

### Prerequisites

* golang >= 1.4.2

## Installation

1. Clone `git clone git@github.com:phoneware/reverse-proxy.git && cd reverse-proxy`
1. Compile `go install`

## Usage

Reverse Proxy Requires two Environment Variables: `PORT` and `DEST_HOST`.

### Run with

`PORT=5000 DEST_HOST=https://azphoneware.com reverse-proxy`
