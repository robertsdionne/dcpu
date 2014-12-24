#ifndef RSD_GLFWAPPLICATION_HPP_
#define RSD_GLFWAPPLICATION_HPP_

#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>
#include <string>

#include "application.hpp"

namespace rsd {

  class Mouse;
  class Renderer;

  class GlfwApplication : public Application {
  public:
    GlfwApplication(int argument_count, char *arguments[], int width, int height, int samples,
                    const std::string &title, Renderer &renderer, Mouse &mouse);

    virtual ~GlfwApplication();

    virtual int Run() override;

  protected:
    static void HandleKeyboard(GLFWwindow *window, int key, int scancode, int action, int mods);

    static void HandleMouseButton(GLFWwindow *window, int button, int action, int mods);

    static void HandleReshape(GLFWwindow *window, int width, int height);

    static GlfwApplication *instance;

  private:
    GLFWwindow *window;
    int argument_count;
    char **arguments;
    int width, height, samples;
    const std::string title;
    Renderer &renderer;
    Mouse &mouse;
  };

}  // namespace rsd

#endif  // RSD_GLFWAPPLICATION_HPP_
