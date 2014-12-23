#ifndef RSD_VERTEXARRAY_HPP_
#define RSD_VERTEXARRAY_HPP_

#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>

namespace rsd {

  class VertexArray {
  public:
    VertexArray();
    
    virtual ~VertexArray();

    void Bind();

    void Create();

    void EnableVertexAttribArray(GLuint index);

    void VertexAttribPointer(GLuint index, GLint size, GLenum type,
                             GLboolean normalized, GLsizei stride, const GLvoid *pointer);

  private:
    GLuint handle;
  };

}  // namespace rsd

#endif  // RSD_VERTEXARRAY_HPP_
