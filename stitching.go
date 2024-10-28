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

type Stitcher_Status struct {
	p C.StitcherStatus
}

const (
	Stitcher_Status_OK = iota
	Stitcher_Status_ERR_NEED_MORE_IMGS
	Stitcher_Status_ERR_HOMOGRAPHY_EST_FAIL
	Stitcher_Status_ERR_CAMERA_PARAMS_ADJUST_FAIL
)

func Stitcher_Create(mode Stitcher_Mode) Stitcher {
	return Stitcher{p: unsafe.Pointer(C.Stitcher_Create(mode.p))}
}