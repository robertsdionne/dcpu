#ifndef DCPU_DCPURENDERER_HPP_
#define DCPU_DCPURENDERER_HPP_

#include <renderer.hpp>

namespace dcpu {

  class DcpuRenderer : public rsd::Renderer {
  public:
    DcpuRenderer() = default;

    virtual ~DcpuRenderer() = default;

    void Change(int width, int height) override;

    void Create() override;

    void Render() override;
  };

}  // namespace dcpu

#endif  // DCPU_DCPURENDERER_HPP_
