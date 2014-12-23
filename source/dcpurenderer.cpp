#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>

#include "dcpurenderer.hpp"

namespace dcpu {

  void DcpuRenderer::Change(int width, int height) {
    glViewport(0, 0, width, height);
  }

  void DcpuRenderer::Create() {}

  void DcpuRenderer::Render() {}
}
