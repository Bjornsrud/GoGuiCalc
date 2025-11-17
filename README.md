m# GoGuiCalc

Just a simple GUI calculator written in Go with the Fyne toolkit.
Mainly a playground for me to learn basic Go and Fyne.

## Requirements

- Go 1.25 or newer
- Fyne GUI toolkit: fyne.io/fyne/v2 (v2.7.1 or newer, pulled via Go modules)
- On Windows: MSYS2 with the mingw-w64 64-bit toolchain (`gcc`) installed

Make sure `C:\msys64\mingw64\bin` is on your `PATH` so `gcc` is available.

## Running from source

From the project root:
go run ./cmd/calculator

## Building a Windows executable

go build -v -ldflags="-H=windowsgui" -o GoCalculate.exe ./cmd/calculator

This produces GoCalculate.exe in the project root.

## Adding an application icon

To embed an icon into the binary:
1. Place icon.png in the ./assets folder:
    assets/icon.png

2. Install the Fyne tools CLI (if you have not already):
    go install fyne.io/tools/cmd/fyne@latest

3. Generate the bundled resources file:
    fyne bundle assets/icon.png -o internal/gui/resources.go

4. Rebuild the executable.
