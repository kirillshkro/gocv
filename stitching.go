package gocv

//#include "stitching.h"
import "C"
import "unsafe"

type Stitcher struct {
	p C.PStitcher
}

type StitcherMode struct {
	p C.StitcherMode
}

type StitcherStatus struct {
	error
	p C.StitcherStatus
}

func (s StitcherStatus) Error() string {
	switch s.p {
	case Stitcher_Status_ERR_NEED_MORE_IMGS:
		return "ERR_NEED_MORE_IMGS"
	case Stitcher_Status_ERR_HOMOGRAPHY_EST_FAIL:
		return "ERR_HOMOGRAPHY_EST_FAIL"
	case Stitcher_Status_ERR_CAMERA_PARAMS_ADJUST_FAIL:
		return "ERR_CAMERA_PARAMS_ADJUST_FAIL"
	}
	return "OK"
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

func StitcherCreate(mode StitcherMode) Stitcher {
	return Stitcher{p: C.Stitcher_Create(mode.p)}
}

func (s Stitcher) Stitch(images []Mat) (Mat, StitcherStatus) {
	var cimg OMat
	var cstatus C.StitcherStatus
	var images_c C.Mats = C.Mats{
		length: C.int(len(images)),
		mats: (*C.Mat)(unsafe.Pointer(&images[0])),
	}
	cstatus = C.Stitcher_Stitch(s.p, images_c, cimg.p)
	return Mat{p: cimg.p }, StitcherStatus{p: cstatus}
}

func (s Stitcher) StitchWithMasks(images []Mat, masks []Mat) (Mat, StitcherStatus) {
	var cimg OMat
	var cstatus C.StitcherStatus
	var images_c C.Mats = C.Mats{
		length: C.int(len(images)),
		mats: (*C.Mat)(unsafe.Pointer(&images[0])),
	}
	var masks_c C.Mats = C.Mats{
		length: C.int(len(masks)),
		mats: (*C.Mat)(unsafe.Pointer(&masks[0])),
	}
	cstatus = C.Stitcher_StitchWithMasks(s.p, images_c, masks_c, cimg.p)
	return Mat{p: cimg.p }, StitcherStatus{p: cstatus}
}