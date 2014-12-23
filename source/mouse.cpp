#include <chrono>
#include <glm/glm.hpp>
#include <map>

#include "mouse.hpp"

namespace rsd {

  glm::vec2 Mouse::get_cursor_position() const {
    return cursor_position;
  }

  float Mouse::GetButtonVelocity(int button) {
    return (buttons[button] - previous_buttons[button]) * dt;
  }

  glm::vec2 Mouse::GetCursorVelocity() {
    return (cursor_position - previous_cursor_position) * dt;
  }

  bool Mouse::HasCursorMoved() const {
    return glm::vec2(0, 0) != cursor_position - previous_cursor_position;
  }

  bool Mouse::IsButtonDown(const int button) {
    return buttons[button];
  }

  void Mouse::OnButtonDown(const int button) {
    buttons[button] = true;
  }

  void Mouse::OnButtonUp(const int button) {
    buttons[button] = false;
  }

  void Mouse::OnCursorMove(const glm::vec2 position) {
    cursor_position = position;
  }

  void Mouse::Update() {
    previous_buttons = buttons;
    previous_cursor_position = cursor_position;
    auto now = std::chrono::high_resolution_clock::now();
    dt = std::chrono::duration_cast<std::chrono::duration<float>>(now - last_update_time).count();
    last_update_time = now;
  }

}  // namespace rsd
