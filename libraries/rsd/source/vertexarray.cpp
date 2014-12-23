#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>

#include "vertexarray.hpp"

namespace rsd {

  VertexArray::VertexArray() : handle() {}

  VertexArray::~VertexArray() {
    if (handle) {
      glDeleteVertexArrays(1, &handle);
      handle = 0;
    }
  }

  void VertexArray::Bind() {
    glBindVertexArray(handle);
  }

  void VertexArray::Create() {
    glGenVertexArrays(1, &handle);
  }

  void VertexArray::EnableVertexAttribArray(GLuint index) {
    Bind();
    glEnableVertexAttribArray(index);
  }

  void VertexArray::VertexAttribPointer(GLuint index, GLint size, GLenum type, GLboolean normalized,
                                        GLsizei stride, const GLvoid *pointer) {
    Bind();
    glVertexAttribPointer(index, size, type, normalized, stride, pointer);
  }

}  // namespace rsd
