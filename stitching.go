package gocv

//#include "stitching.h"
import "C"
import "unsafe"

type Stitcher struct {
	p unsafe.Pointer
}

type Stitcher_Mode struct {
	p C.StitcherMode
}

func Stitcher_Create(mode Stitcher_Mode) Stitcher {
	return Stitcher{p: unsafe.Pointer(C.Stitcher_Create(mode.p))}
}