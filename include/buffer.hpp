#ifndef RSD_BUFFER_HPP_
#define RSD_BUFFER_HPP_

#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>

namespace rsd {

  class Buffer {
  public:
    Buffer();

    virtual ~Buffer();

    GLuint get_handle() const;

    void Bind();

    void Create(GLenum target);

    void Data(GLsizeiptr size, const GLvoid *data, GLenum usage);

    void SubData(GLintptr offset, GLsizeiptr size, const GLvoid * data);

  private:
    GLenum target;
    GLuint handle;
  };

}  // namespace rsd

#endif  // RSD_BUFFER_HPP_
