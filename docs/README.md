# storage — documentation

  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />

## Overview

Package storage is togo's default filesystem storage provider. Blank-import
(or `togo install togo-framework/storage`) to register it with the kernel.

## Install

```bash
togo install togo-framework/storage
```

Set `STORAGE_DRIVER=<provider>` and install a driver (storage-s3, storage-r2, …).

## Configuration

Environment variables read by this plugin (extracted from the source):

_No environment variables read directly (uses the kernel/base config or the app DB)._

## Usage

```go
st := k.Storage
st.Put(ctx, "path/file.txt", data)
b, _ := st.Get(ctx, "path/file.txt")
url := st.Path("path/file.txt")
```

## Links

- Marketplace: https://to-go.dev/marketplace
- Source: https://github.com/togo-framework/storage
- README: ../README.md
