#ifndef _OPENCVV3_STITCHING_H_
#define _OPENCVV3_STITCHING_H_
#ifdef __cplusplus
#include <opencv2/opencv.hpp>
#include <opencv2/stitching.hpp>
extern "C" {
#endif

#include "core.h"
#ifdef __cplusplus
typedef cv::Stitcher::Mode StitcherMode;
typedef cv::Ptr<cv::Stitcher>* PStitcher;
#else
typedef int StitcherMode;
typedef void* Stitcher;
#endif

PStitcher Stitcher_Create(StitcherMode mode = cv::Stitcher::PANORAMA);

#ifdef __cplusplus
}
#endif
#endif