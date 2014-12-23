#ifndef RSD_MOUSE_HPP_
#define RSD_MOUSE_HPP_

#include <chrono>
#include <glm/glm.hpp>
#include <unordered_map>

namespace rsd {

  class Mouse {
  public:
    Mouse() = default;

    virtual ~Mouse() = default;

    glm::vec2 get_cursor_position() const;

    float GetButtonVelocity(int button);

    glm::vec2 GetCursorVelocity();

    bool HasCursorMoved() const;

    bool IsButtonDown(int button);

    void OnButtonDown(int button);

    void OnButtonUp(int button);

    void OnCursorMove(glm::vec2 position);

    void Update();

  private:
    glm::vec2 cursor_position, previous_cursor_position;
    std::unordered_map<int, bool> buttons, previous_buttons;
    std::chrono::high_resolution_clock::time_point last_update_time;
    float dt;
  };

}  // namespace rsd

#endif  // RSD_MOUSE_HPP_
