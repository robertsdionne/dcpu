#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>
#include <iostream>
#include <fstream>
#include <string>
#include <streambuf>

#include "checks.hpp"
#include "shader.hpp"

namespace rsd {

  Shader::Shader() : type(), sources(), handle() {}

  Shader::~Shader() {
    if (handle) {
      glDeleteShader(handle);
      handle = 0;
    }
  }

  GLuint Shader::get_handle() const {
    return handle;
  }

  void Shader::Compile() {
#ifdef WIN32
    GLchar **source_code = new GLchar*[sources.size()];
    GLint *lengths = new GLint[sources.size()];
#else
    GLchar *source_code[sources.size()];
    GLint lengths[sources.size()];
#endif

    {
      int index = 0;
      for (auto &source : sources) {
        source_code[index] = new char[source.size() + 1];
        std::copy(source.begin(), source.end(), source_code[index]);
        source_code[index][source.size()] = '\0';
        lengths[index] = static_cast<GLint>(source.size());
        ++index;
      }
    }

    glShaderSource(handle, static_cast<GLsizei>(sources.size()), source_code, lengths);

    for (int i = 0; i < sources.size(); ++i) {
      delete[] source_code[i];
      source_code[i] = nullptr;
    }
#ifdef WIN32
    delete[] source_code;
    delete[] lengths;
#endif

    glCompileShader(handle);
    MaybeOutputCompilerError();
  }

  void Shader::CreateFromFile(GLenum type, const std::string &filename) {
    Create(type, {ReadFile(filename)});
  }

  void Shader::Create(GLenum type, const std::vector<std::string> &&sources) {
    if (handle) {
      glDeleteShader(handle);
      handle = 0;
    }
    this->type = type;
    this->sources = sources;
    handle = glCreateShader(type);
  }

  void Shader::MaybeOutputCompilerError() {
    GLint success;
    glGetShaderiv(handle, GL_COMPILE_STATUS, &success);
    if (!success) {
      GLchar info_log[kMaxInfoLogLength];
      GLsizei length;
      glGetShaderInfoLog(handle, kMaxInfoLogLength, &length, info_log);
      if (length) {
        FAIL(info_log);
      } else {
        FAIL("Failed to compile shader.");
      }
    }
  }

  std::string Shader::ReadFile(const std::string &filename) {
    std::ifstream file(filename);
    CHECK_STATE(file.good());
    std::string content;
    file.seekg(0, std::ios::end);
    content.reserve(file.tellg());
    file.seekg(0, std::ios::beg);
    content.assign(std::istreambuf_iterator<char>(file), std::istreambuf_iterator<char>());
    return content;
  }

}  // namespace rsd
