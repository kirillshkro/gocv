#ifndef _OPENCVV3_STITCHING_H_
#define _OPENCVV3_STITCHING_H_
#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/stitching.hpp>
extern "C" {
#endif

#include "core.h"

typedef enum {
    PANORAMA = 0,
    SCANS = 1
} StitcherMode;

typedef enum {
    OK,
    ERROR_NEED_MORE_IMGS,
    ERROR_HOMOGRAPHY_EST_FAIL,
    ERROR_CAMERA_PARAMS_ADJUST_FAIL
} StitcherStatus;
#ifdef __cplusplus
typedef cv::Ptr<cv::Stitcher>* PStitcher;
#else
typedef void* PStitcher;
#endif

PStitcher Stitcher_Create(StitcherMode mode);

#ifdef __cplusplus
}
#endif
#endif