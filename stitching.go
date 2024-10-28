package gocv

//#include "stitching.h"
import "C"
import "unsafe"

type Stitcher struct {
	p C.PStitcher
}

type Stitcher_Mode struct {
	p C.StitcherMode
}

type Stitcher_Status struct {
	p C.StitcherStatus
}

type OMat struct {
	p C.Mat
}

const (
	Stitcher_Status_OK = iota
	Stitcher_Status_ERR_NEED_MORE_IMGS
	Stitcher_Status_ERR_HOMOGRAPHY_EST_FAIL
	Stitcher_Status_ERR_CAMERA_PARAMS_ADJUST_FAIL
)

func Stitcher_Create(mode Stitcher_Mode) Stitcher {
	return Stitcher{p: C.Stitcher_Create(mode.p)}
}

func (s Stitcher) Stitch(images []Mat) (Mat, Stitcher_Status) {
	var cimg OMat
	var cstatus C.StitcherStatus
	var images_c C.Mats = C.Mats{
		length: C.int(len(images)),
		mats: (*C.Mat)(unsafe.Pointer(&images[0])),
	}
	cstatus = C.Stitcher_Stitch(s.p, images_c, cimg.p)
	return Mat{p: cimg.p }, Stitcher_Status{p: cstatus}
}