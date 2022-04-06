# goop2
This project is meant to implement an optimization toolbox with nice default features like MATLAB's YALMIP. It is a spiritual successor to the goop library from MIT's Distributed Robotics Lab. This version is meant to be written in pure Go and makes use of my own low-level library for interfacing with Gurobi.

## Installation

Warning: The setup script is designed to only work on Mac OS X. If you are interested in using this on a Windows machine, then there are no guarantees that it will work.

### Installation in Module

1. Use a "-d" `go get -d github.com/kwesiRutledge/gurobi.go/gurobi`. Pay attention to which version appears in your terminal output.
2. Enter Go's internal installation of gurobi.go. For example, run `cd ~/go/pkg/mod/github.com/kwesi\!rutledge/gurobi.go@v0.0.0-20220103225839-e6367b1d0f27` where the suffix is the version number from the previous output.
3. Run go generate with sudo privileges from this installation. `sudo go generate`.

### Development Installation

1. Clone the library using `git clone github.com/kwesiRutledge/gurobi.go `
2. Run the setup script from inside the cloned repository: `go generate`.

## Examples

There are several examples in the test files. Check out `lp_solve_test.go` for an example.