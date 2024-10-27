#include "stitching.h"

PStitcher Stitcher_Create(StitcherMode mode)
{
    auto stitcher = new cv::Ptr<cv::Stitcher>(cv::Stitcher::create(mode));
    return stitcher;
}
