Simple demo of creating a unrar and zip command (separate binaries).  
It uses the archiver library: https://github.com/mholt/archiver

To ensure a correct build after cloning this repository make sure you have `dep` installed.  
Run `dep ensure` within the root of this folder to vendor the required libraries.

Then run `go build .` inside each of the `cmd` folders.