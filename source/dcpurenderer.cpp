#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>
#include <checks.hpp>
#include <glm/gtc/matrix_transform.hpp>

#include "dcpurenderer.hpp"

namespace dcpu {

  void DcpuRenderer::Change(int width, int height) {
    glViewport(0, 0, width, height);
    auto aspect = static_cast<float>(width) / height;
    projection = glm::frustum(-0.1f * aspect, 0.1f * aspect, -0.1f, 0.1f, 0.1f, 1000.0f);
  }

  void DcpuRenderer::Create() {
    glClearColor(0.0f, 0.0f, 0.0f, 0.0f);
    glEnable(GL_BLEND);
    glDepthMask(GL_FALSE);
    glDisable(GL_CULL_FACE);
    glBlendFunc(GL_SRC_ALPHA, GL_ONE);
    // glLineWidth(2.0f);

    vertex_shader.CreateFromFile(GL_VERTEX_SHADER, "source/vertex.glsl");
    fragment_shader.CreateFromFile(GL_FRAGMENT_SHADER, "source/fragment.glsl");
    program.Create({&vertex_shader, &fragment_shader});
    program.CompileAndLink();
    vertex_format.Create({
      {"vertex_position", GL_FLOAT, 3},
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

    model_view = glm::mat4(
      1.0f, 0.0f, 0.0f, 0.0f,
      0.0f, 1.0f, 0.0f, 0.0f,
      0.0f, 0.0f, 1.0f, 0.0f,
      0.0f, -0.8f, -1.5f, 1.0f
    );
  }

  void DcpuRenderer::Render() {
    Update();

    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);

    program.Use();
    program.UniformsMatrix4f({
      {"model_view", &model_view},
      {"projection", &projection}
    });

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

  void DcpuRenderer::Update() {
    // dcpu.ExecuteInstructions(1666);
    display.Execute();

    x += vx;
    y += vy;
    z += vz;
    if (x < 0) {
      x = 0;
      vx *= -1;
    }
    if (x > 223) {
      x = 223;
      vx *= -1;
    }if (y < 0) {
      y = 0;
      vy *= -1;
    }
    if (y > 223) {
      y = 223;
      vy *= -1;
    }
    if (z < 0) {
      z = 0;
      vz *= -1;
    }
    if (z > 223) {
      z = 223;
      vz *= -1;
    }
    Word word0 = (y << 8) | x;
    Word word1 = z;
    auto data = {
      0x0000 + word0, 0x0700 + word1,
      0x2000 + word0, 0x0700 + word1,
      0x2020 + word0, 0x0700 + word1,
      0x0020 + word0, 0x0700 + word1,
      0x0000 + word0, 0x0700 + word1,

      0x0000 + word0, 0x0720 + word1,
      0x0000 + word0, 0x0700 + word1,
      0x2000 + word0, 0x0700 + word1,
      0x2000 + word0, 0x0720 + word1,
      0x2020 + word0, 0x0720 + word1,
      0x2020 + word0, 0x0700 + word1,
      0x0020 + word0, 0x0700 + word1,
      0x0020 + word0, 0x0720 + word1,

      0x0000 + word0, 0x0720 + word1,
      0x2000 + word0, 0x0720 + word1,
      0x2020 + word0, 0x0720 + word1,
      0x0020 + word0, 0x0720 + word1,
      0x0000 + word0, 0x0720 + word1,
      0x0000 + word0, 0x0020 + word1,

      0x3000 + word0, 0x0000 + word1,
      0x3000 + word0, 0x0700 + word1,
      0x5000 + word0, 0x0700 + word1,
      0x5020 + word0, 0x0700 + word1,
      0x3020 + word0, 0x0700 + word1,
      0x3000 + word0, 0x0700 + word1,

      0x3000 + word0, 0x0720 + word1,
      0x3000 + word0, 0x0700 + word1,
      0x5000 + word0, 0x0700 + word1,
      0x5000 + word0, 0x0720 + word1,
      0x5020 + word0, 0x0720 + word1,
      0x5020 + word0, 0x0700 + word1,
      0x3020 + word0, 0x0700 + word1,
      0x3020 + word0, 0x0720 + word1,

      0x3000 + word0, 0x0720 + word1,
      0x5000 + word0, 0x0720 + word1,
      0x5020 + word0, 0x0720 + word1,
      0x3020 + word0, 0x0720 + word1,
      0x3000 + word0, 0x0720 + word1,
      0x3000 + word0, 0x0020 + word1,

      0x0000, 0x0000,
      0x0000, 0x0700,
      0x2000, 0x0700,
      0x2000, 0x0000,
      0xDF00, 0x0000,
      0xDF00, 0x0700,
      0xFF00, 0x0700,
      0xFF20, 0x0700,
      0xFF20, 0x0000,
      0xFFDF, 0x0000,
      0xFFDF, 0x0700,
      0xFFFF, 0x0700,
      0xDFFF, 0x0700,
      0xDFFF, 0x0000,
      0x20FF, 0x0000,
      0x20FF, 0x0700,
      0x00FF, 0x0700,
      0x00DF, 0x0700,
      0x00DF, 0x0000,
      0x0020, 0x0000,
      0x0020, 0x0700,
      0x0000, 0x0700,
      0x0000, 0x0000,

      0x0000, 0x00FF,
      0x0000, 0x07FF,
      0x2000, 0x07FF,
      0x2000, 0x00FF,
      0xDF00, 0x00FF,
      0xDF00, 0x07FF,
      0xFF00, 0x07FF,
      0xFF20, 0x07FF,
      0xFF20, 0x00FF,
      0xFFDF, 0x00FF,
      0xFFDF, 0x07FF,
      0xFFFF, 0x07FF,
      0xDFFF, 0x07FF,
      0xDFFF, 0x00FF,
      0x20FF, 0x00FF,
      0x20FF, 0x07FF,
      0x00FF, 0x07FF,
      0x00DF, 0x07FF,
      0x00DF, 0x00FF,
      0x0020, 0x00FF,
      0x0020, 0x07FF,
      0x0000, 0x07FF
    };
    std::copy(data.begin(), data.end(), dcpu.memory_begin());
    display.Map(0, data.size() / 2);
  }
}
