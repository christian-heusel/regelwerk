//go:build linux && !rw_midna
// +build linux,!rw_midna

package main

var loops = []controlLoop{
	&krkStudioBoxen{},
}
