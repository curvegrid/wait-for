package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var timeoutFlag int
var intervalFlag int
var addressFlag string

func init() {
	flag.IntVar(&timeoutFlag, "timeout", 0, "timeout (in second) before the program exits with an error code")
	flag.IntVar(&intervalFlag, "interval", 1, "minimum interval (in second) between connection attempts")

	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	addressFlag = flag.Arg(0)
}

func main() {
	d := &net.Dialer{}
	t := time.NewTimer(0)
	i := time.Duration(intervalFlag) * time.Second

	ctx := context.Background()
	if timeoutFlag > 0 {
		ctx, _ = context.WithTimeout(ctx, time.Duration(timeoutFlag)*time.Second)
	}

dialer:
	for {
		select {
		case <-ctx.Done():
			break dialer
		default:
		}

		select {
		case <-ctx.Done():
			break dialer
		case <-t.C:
			t.Reset(i)
			if _, err := d.DialContext(ctx, "tcp", addressFlag); err == nil {
				printStderr("successfully connected\n")
				os.Exit(0)
			}
			printStderr("connection failed\n")
		}
	}

	printStderr("timed out\n")
	os.Exit(1)
}

func printStderr(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}

func usage() {
	printStderr("Usage: %s [OPTIONS] host:port\n", os.Args[0])
	flag.PrintDefaults()
}
