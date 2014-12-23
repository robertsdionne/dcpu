#ifndef RSD_PROGRAM_HPP_
#define RSD_PROGRAM_HPP_

#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>
#include <glm/glm.hpp>
#include <unordered_map>
#include <vector>

namespace rsd {

  class Shader;

  class Program {
  public:
    Program();

    virtual ~Program();

    GLuint get_handle() const;

    void CompileAndLink();

    void Create(const std::vector<Shader *> &&shaders);

    GLint GetAttributeLocation(const std::string &name);

    GLint GetUniformLocation(const std::string &name);
    
    void Uniformsi(const std::unordered_map<std::string, int> &&uniforms);

    void Uniformsf(const std::unordered_map<std::string, float> &&uniforms);
    
    void Uniforms2f(const std::unordered_map<std::string, glm::vec2> &&uniforms);
    
    void Uniforms3f(const std::unordered_map<std::string, glm::vec3> &&uniforms);

    void Uniforms4f(const std::unordered_map<std::string, glm::vec4> &&uniforms);

    void UniformsMatrix4f(const std::unordered_map<std::string, const glm::mat4 *> &&uniforms);

    void Use();

  private:
    void MaybeOutputLinkerError();

  private:
    std::vector<Shader *> shaders;
    GLuint handle;
  };

}  // namespace rsd

#endif  // RSD_PROGRAM_HPP_
