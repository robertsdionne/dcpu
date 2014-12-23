#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>
#include <vector>

#include "program.hpp"
#include "vertexarray.hpp"
#include "vertexformat.hpp"

namespace rsd {

  const std::unordered_map<GLenum, GLint> VertexFormat::kTypeSizes = {
    {GL_BYTE, sizeof(char)},
    {GL_DOUBLE, sizeof(double)},
    {GL_FLOAT, sizeof(float)},
    {GL_INT, sizeof(int)},
    {GL_SHORT, sizeof(short)}
  };

  void VertexFormat::Apply(VertexArray &array, Program &program) const {
    GLsizei stride = 0;
    for (auto &attribute : attributes) {
      stride += kTypeSizes.at(attribute.type) * attribute.size;
    }
    int offset = 0;
    for (auto &attribute : attributes) {
      auto attribute_location = program.GetAttributeLocation(attribute.name);
      array.EnableVertexAttribArray(attribute_location);
      array.VertexAttribPointer(attribute_location,
                                attribute.size, attribute.type, false, stride,
                                reinterpret_cast<GLvoid *>(offset));
      offset += kTypeSizes.at(attribute.type) * attribute.size;
    }
  }

  void VertexFormat::Create(std::vector<Attribute> &&attributes) {
    this->attributes = attributes;
  }

}  // namespace rsd
