m# GoGuiCalc

Just a simple GUI calculator written in Go with the Fyne toolkit.
Mainly a playground for me to learn basic Go and Fyne.

## Requirements

- Go 1.25 or newer
- Fyne v1.7.0 or newer
- On Windows: MSYS2 with the mingw-w64 64-bit toolchain (`gcc`) installed

Make sure `C:\msys64\mingw64\bin` is on your `PATH` so `gcc` is available.

## Running from source

From the project root:

go run ./cmd/calculator

## Building a Windows executable

go build -v -ldflags="-H=windowsgui" -o GoCalculate.exe ./cmd/calculator

## Add icon in ./assets

fyne bundle assets/icon.png -o internal/gui/resources.go
