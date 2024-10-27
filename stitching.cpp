#include "stitching.h"

PStitcher Stitcher_Create(StitcherMode mode)
{
    auto _mode = static_cast<cv::Stitcher::Mode>(mode);
    auto stitcher = new cv::Ptr<cv::Stitcher>(cv::Stitcher::create(_mode));
    return stitcher;
}
