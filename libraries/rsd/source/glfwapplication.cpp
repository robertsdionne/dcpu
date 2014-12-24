#include <GLXW/glxw.h>
#include <GLFW/glfw3.h>
#include <glm/glm.hpp>
#include <iostream>

#include "checks.hpp"
#include "glfwapplication.hpp"
#include "mouse.hpp"
#include "renderer.hpp"

namespace rsd {

  GlfwApplication *GlfwApplication::instance = nullptr;

  GlfwApplication::GlfwApplication(
      int argument_count, char *arguments[], int width, int height, int samples,
      const std::string &title, Renderer &renderer, Mouse &mouse,
      int context_version_major, int context_version_minor)
  : window(nullptr), argument_count(argument_count), arguments(arguments), width(width),
  height(height), samples(samples), title(title), renderer(renderer), mouse(mouse),
  context_version_major(context_version_major), context_version_minor(context_version_minor) {
    instance = this;
  }

  GlfwApplication::~GlfwApplication() {
    instance = nullptr;
  }

  void GlfwApplication::HandleKeyboard(GLFWwindow *window, int key,
                                       int scancode, int action, int mods) {
    if (instance) {
      switch (action) {
        case GLFW_PRESS:
        case GLFW_REPEAT: {
          if (GLFW_KEY_ESCAPE == key) {
            glfwSetWindowShouldClose(window, GL_TRUE);
          }
          // instance->keyboard.OnKeyDown(key);
          break;
        }
        case GLFW_RELEASE: {
          // instance->keyboard.OnKeyUp(key);
          break;
        }
      }
    }
  }

  void GlfwApplication::HandleMouseButton(GLFWwindow *window, int button, int action, int mods) {
    if (instance) {
      switch (action) {
        case GLFW_PRESS: {
          instance->mouse.OnButtonDown(button);
          break;
        }
        case GLFW_RELEASE: {
          instance->mouse.OnButtonUp(button);
          break;
        }
      }
    }
  }

  void GlfwApplication::HandleReshape(GLFWwindow *window, int width, int height) {
    if (instance) {
      instance->renderer.Change(width, height);
    }
  }

  int GlfwApplication::Run() {
    CHECK_STATE(glfwInit() != -1);
    glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, context_version_major);
    glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, context_version_minor);
    glfwWindowHint(GLFW_OPENGL_FORWARD_COMPAT, GL_TRUE);
    glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE);
    if (samples) {
      glfwWindowHint(GLFW_SAMPLES, samples);
    }
    window = glfwCreateWindow(width, height, title.c_str(), nullptr, nullptr);
    CHECK_STATE(window != nullptr);
    glfwSetKeyCallback(window, HandleKeyboard);
    glfwSetMouseButtonCallback(window, HandleMouseButton);
    glfwSetFramebufferSizeCallback(window, HandleReshape);
    glfwMakeContextCurrent(window);
    CHECK_STATE(!glxwInit());
    glfwSwapInterval(1);
    renderer.Create();
    int framebuffer_width, framebuffer_height;
    glfwGetFramebufferSize(window, &framebuffer_width, &framebuffer_height);
    HandleReshape(window, framebuffer_width, framebuffer_height);
    while (!glfwWindowShouldClose(window)) {
      renderer.Render();
      // keyboard.Update();
      mouse.Update();
      double x, y;
      glfwGetCursorPos(window, &x, &y);
      mouse.OnCursorMove(glm::vec2(x, y));
      glfwSwapBuffers(window);
      glfwPollEvents();
    }
    glfwTerminate();
    return 0;
  }

}  // namespace rsd
