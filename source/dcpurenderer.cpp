#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>
#include <checks.hpp>

#include "dcpurenderer.hpp"

namespace dcpu {

  void DcpuRenderer::Change(int width, int height) {
    glViewport(0, 0, width, height);
  }

  void DcpuRenderer::Create() {
    glClearColor(0.0f, 0.0f, 0.0f, 0.0f);
    vertex_shader.CreateFromFile(GL_VERTEX_SHADER, "source/vertex.glsl");
    fragment_shader.CreateFromFile(GL_FRAGMENT_SHADER, "source/fragment.glsl");
    program.Create({&vertex_shader, &fragment_shader});
    program.CompileAndLink();
    vertex_format.Create({
      {"vertex_position", GL_FLOAT, 4},
      {"vertex_color", GL_FLOAT, 4}
    });

    beam_buffer.Create(GL_ARRAY_BUFFER);
    beam_vertex_array.Create();
    vertex_format.Apply(beam_vertex_array, program);

    line_buffer.Create(GL_ARRAY_BUFFER);
    line_vertex_array.Create();
    vertex_format.Apply(line_vertex_array, program);

    CHECK_STATE(!glGetError());

    dcpu.Connect(&display);
  }

  void DcpuRenderer::Render() {
    dcpu.ExecuteInstructions(1666);
    display.Execute();

    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);

    program.Use();

    auto beam_drawable = display.BeamView();
    beam_buffer.Data(beam_drawable.data_size(), beam_drawable.data.data(), GL_STREAM_DRAW);
    beam_vertex_array.Bind();
    glDrawArrays(beam_drawable.element_type, 0, beam_drawable.element_count);

    auto line_drawable = display.LineView();
    line_buffer.Data(line_drawable.data_size(), line_drawable.data.data(), GL_STREAM_DRAW);
    line_vertex_array.Bind();
    glDrawArrays(line_drawable.element_type, 0, line_drawable.element_count);

    CHECK_STATE(!glGetError());
  }
}
