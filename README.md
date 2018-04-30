go-template-asset-pipeline
=======================
An example project using Go standard library templates with a modern Webpack asset pipeline.

Table of Contents
-----------------

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)

Features
--------

- **Webpack Asset Pipeline**
  - Bundled assets
  - Hashed filenames for easy caching.
- **Go standard library templates**
  - Simple, familiar template language
  - Template function to retrieve hashed filenames from the webpack manifest
- **Chi router**
  - Idiomatic router (can be easily removed if you prefer)
- **Asset caching middleware**
- **Builds into single binary with Packr (assets included)**
- **Frontend**
  - Bootstrap 4
  - JQuery

Prerequisites
--------
- **Go 1.8+**
  - github.com/go-chi/chi
  - github.com/gobuffalo/packr
- **Node 8.11.1+**
- **Yarn** (or npm if you prefer)

Getting Started
--------
All yarn commands can be replaced with npm if you prefer.

**Install dependencies**
```bash
go get -d ./...
yarn install --dev
```
**Build**
```bash
yarn build (or yarn dev)
packr build
```
**Run**
```bash
./go-template-asset-pipeline.exe
```
Running on localhost:3000
