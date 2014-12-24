#version 410 core

uniform mat4 model_view;
uniform mat4 projection;

in vec4 vertex_position;
in vec4 vertex_color;

out vec4 color;

void main() {
  gl_Position = projection * model_view * vertex_position;
  color = vertex_color;
}
