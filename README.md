[//]: # (center header with logo)
<h1 align="center">
  Priestess Infrastructure
</h1>

<div align="center">

A collection of tools and scripts to manage the infrastructure of the Priestess project.

[![Build Status](https://app.travis-ci.com/priestess-dev/infra.svg)](https://travis-ci.org/priestess-dev/infra)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)


</div>

## ✨ Features

- `cache`: package for managing cache servers
  - [x] `redis`
  - [ ] `memcached`
- `http`: package for making http request, client, server
- `oauth`: oauth flow implementation
  - [x] `github`
  - [ ] `google`
  - [ ] `...`
- `thridparty`: package for third-party services sdk/clients
  - [x] `github`
  - [ ] `...`
- `utils`: package for common utilities
  - [x] `random`: random string, number, etc
  - [ ] `...`

## 📦 Install

```bash
go get github.com/priestess-dev/infra@v1.0.0
```