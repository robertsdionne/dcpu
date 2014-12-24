#ifndef DCPU_DCPURENDERER_HPP_
#define DCPU_DCPURENDERER_HPP_

#include <buffer.hpp>
#include <drawable.hpp>
#include <glm/glm.hpp>
#include <program.hpp>
#include <renderer.hpp>
#include <shader.hpp>
#include <vertexarray.hpp>
#include <vertexformat.hpp>

#include "dcpu.hpp"
#include "mackaparsuspendedparticleexciterdisplay.hpp"

namespace dcpu {

  class DcpuRenderer : public rsd::Renderer {
  public:
    DcpuRenderer() = default;

    virtual ~DcpuRenderer() = default;

    void Change(int width, int height) override;

    void Create() override;

    void Render() override;

  private:
    Dcpu dcpu;
    MackaparSuspendedParticleExciterDisplay display;
    rsd::Shader vertex_shader, fragment_shader;
    rsd::Program program;
    rsd::VertexFormat vertex_format;
    rsd::VertexArray beam_vertex_array, line_vertex_array;
    rsd::Buffer beam_buffer, line_buffer;
    glm::mat4 model_view, projection;
  };

}  // namespace dcpu

#endif  // DCPU_DCPURENDERER_HPP_
