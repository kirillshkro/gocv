#include "stitching.h"

PStitcher Stitcher_Create(StitcherMode mode)
{
    auto _mode = static_cast<cv::Stitcher::Mode>(mode);
    auto stitcher = new cv::Ptr<cv::Stitcher>(cv::Stitcher::create(_mode));
    return stitcher;
}

StitcherStatus Stitcher_Stitch(PStitcher stitcher, Mats images, Mat pano)
{
    auto _stitcher = *stitcher;
    std::vector<cv::Mat> _images;
    for (int i = 0; i < images.length; i++) {
        _images.push_back(*images.mats[i]);
    }
    cv::Mat _pano;
    StitcherStatus status = static_cast<StitcherStatus>(_stitcher->stitch(_images, _pano));
    *pano = _pano;
    return status;
}

StitcherStatus Stitcher_StitchWithMasks(PStitcher stitcher, Mats images, Mats masks, Mat pano)
{
    auto _stitcher = *stitcher;
    std::vector<cv::Mat> _images;
    cv::Mat _pano;
    for (int i = 0; i < images.length; i++) {
        _images.push_back(*images.mats[i]);
    }
    std::vector<cv::Mat> _masks;
    for (int i = 0; i < masks.length; i++) {
        _masks.push_back(*masks.mats[i]);
    }
    StitcherStatus status =
        static_cast<StitcherStatus>(_stitcher->stitch(_images, _masks, _pano));
    *pano = _pano;
    return status;
}
