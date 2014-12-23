#ifndef RSD_SHADER_HPP_
#define RSD_SHADER_HPP_

#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>
#include <string>
#include <vector>

namespace rsd {

  class Shader {
  public:
    Shader();

    virtual ~Shader();

    GLuint get_handle() const;

    void CreateFromFile(GLenum type, const std::string &filename);

    void Create(GLenum type, const std::vector<std::string> &&sources);

    void Compile();

#ifdef WIN32
    static const GLsizei kMaxInfoLogLength = 4 * 1024;
#else
    static constexpr GLsizei kMaxInfoLogLength = 4 * 1024;
#endif

  private:
    void MaybeOutputCompilerError();

    std::string ReadFile(const std::string &filename);

  private:
    GLenum type;
    std::vector<std::string> sources;
    GLuint handle;
  };

}  // namespace rsd

#endif  // RSD_SHADER_HPP_
