#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>
#include <checks.hpp>
#include <glm/gtc/matrix_transform.hpp>

#include "dcpurenderer.hpp"
#include "dsl.hpp"

namespace dcpu {

  using namespace dsl;

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

    model_view = glm::mat4(
      1.0f, 0.0f, 0.0f, 0.0f,
      0.0f, 1.0f, 0.0f, 0.0f,
      0.0f, 0.0f, 1.0f, 0.0f,
      0.0f, -0.8f, -1.5f, 1.0f
    );

    dcpu.Connect(&display);

    Dsl d;
    d.label("main")
        .set(x, 0x8000)
        .set(y, 84)
        .set(a, 1)
        .hwi(0)

      .label("update")
        .add(d["x"], d["vx"])
        .add(d["y"], d["vy"])
        .add(d["z"], d["vz"])

        .set(a, "x")
        .set(b, "vx")
        .jsr("clip_coordinate_above")
        .jsr("clip_coordinate_below")

        .set(a, "y")
        .set(b, "vy")
        .jsr("clip_coordinate_above")
        .jsr("clip_coordinate_below")

        .set(a, "z")
        .set(b, "vz")
        .jsr("clip_coordinate_above")
        .jsr("clip_coordinate_below")

        .set(i, 0x8000)
        .set(j, "boxes")
        .jsr("display_boxes")
        .jsr("display_boundary")

        .set(pc, "update")

      .label("clip_coordinate_above")
        .ife(d[a], 223)
          .set(pc, pop)
        .ifu(d[a], 223)
          .set(pc, pop)
        .set(d[a], 223)
        .mli(d[b], -1)
        .set(pc, pop)

      .label("clip_coordinate_below")
        .ife(d[a], 0)
          .set(pc, pop)
        .ifa(d[a], 0)
          .set(pc, pop)
        .set(d[a], 0)
        .mli(d[b], -1)
        .set(pc, pop)

      .label("display_boxes")
        .ife(j, "boxes_end")
          .set(pc, pop)
        .set(a, d["y"])
        .shl(a, 8)
        .set(b, d["x"])
        .and_(b, 0x00FF)
        .bor(a, b)
        .add(a, d[j])
        .sti(d[i], a)
        .set(a, d["z"])
        .and_(a, 0x00FF)
        .add(a, d[j])
        .sti(d[i], a)
        .set(pc, "display_boxes")

      .label("display_boundary")
        .ife(j, "boundary_end")
          .set(pc, pop)
        .sti(d[i], d[j])
        .set(pc, "display_boundary")

      .label("x").data(0)
      .label("y").data(0)
      .label("z").data(0)
      .label("vx").data(1)
      .label("vy").data(2)
      .label("vz").data(3)

      .label("boxes")
        .data(
            0x0000, 0x0700,
            0x2000, 0x0700,
            0x2020, 0x0700,
            0x0020, 0x0700,
            0x0000, 0x0700,

            0x0000, 0x0720,
            0x0000, 0x0700,
            0x2000, 0x0700,
            0x2000, 0x0720,
            0x2020, 0x0720,
            0x2020, 0x0700,
            0x0020, 0x0700,
            0x0020, 0x0720,

            0x0000, 0x0720,
            0x2000, 0x0720,
            0x2020, 0x0720,
            0x0020, 0x0720,
            0x0000, 0x0720,
            0x0000, 0x0020,

            0x3000, 0x0000,
            0x3000, 0x0700,
            0x5000, 0x0700,
            0x5020, 0x0700,
            0x3020, 0x0700,
            0x3000, 0x0700,

            0x3000, 0x0720,
            0x3000, 0x0700,
            0x5000, 0x0700,
            0x5000, 0x0720,
            0x5020, 0x0720,
            0x5020, 0x0700,
            0x3020, 0x0700,
            0x3020, 0x0720,

            0x3000, 0x0720,
            0x5000, 0x0720,
            0x5020, 0x0720,
            0x3020, 0x0720,
            0x3000, 0x0720,
            0x3000, 0x0020)
      .label("boxes_end")

      .label("boundary")
        .data(
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
            0x0000, 0x07FF)
      .label("boundary_end")

      .Assemble(dcpu.memory_begin());
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
    dcpu.ExecuteInstructions(1666);
    display.Execute();
  }
}
