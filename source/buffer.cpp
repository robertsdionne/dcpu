#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>

#include "buffer.hpp"

namespace rsd {

  Buffer::Buffer() : target(), handle() {}

  Buffer::~Buffer() {
    if (handle) {
      glDeleteBuffers(1, &handle);
      handle = 0;
    }
  }

  GLuint Buffer::get_handle() const {
    return handle;
  }

  void Buffer::Bind() {
    glBindBuffer(target, handle);
  }

  void Buffer::Create(GLenum target) {
    this->target = target;
    glGenBuffers(1, &handle);
    Bind();
  }

  void Buffer::Data(GLsizeiptr size, const GLvoid *data, GLenum usage) {
    Bind();
    glBufferData(target, size, data, usage);
  }

  void Buffer::SubData(GLintptr offset, GLsizeiptr size, const GLvoid *data) {
    Bind();
    glBufferSubData(target, offset, size, data);
  }

}  // namespace rsd
