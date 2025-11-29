# GoGuiCalc

GoGuiCalc is a simple graphical calculator written in Go using the Fyne GUI toolkit.
The project was created as a personal learning exercise to explore Go, GUI programming, and application packaging.

## Features

- Fully functional basic calculator
- Decimal support
- Clear (AC) button
- Color-themed buttons
- Large, readable display
- UI implemented entirely in Fyne (v2.7.1)

## Requirements

- Go 1.20 or newer
- Windows: C compiler required. Fyne uses OpenGL bindings through CGO, so Windows builds require a working GCC toolchain. Recommended (and tested):
TDM-GCC (64-bit)
https://jmeubank.github.io/tdm-gcc/
- CGO must be enabled

## Important

go-gl compatibility note

The latest version of github.com/go-gl/gl is not compatible with current Fyne releases.
This project includes a replace directive in go.mod to pin a working version:

replace github.com/go-gl/gl => github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6

When cloning the repository on a fresh machine, run:

go clean -modcache
go mod tidy

## Running the Application

From the project root:
go run ./cmd/calculator

## Building a Windows executable

go build -v -ldflags="-H=windowsgui" -o GoCalculate.exe ./cmd/calculator

This produces GoCalculate.exe in the project root.

## License

MIT License (see LICENSE file).

## Author

Christian Bjørnsrud
GitHub: https://github.com/Bjornsrud