#ifndef RSD_DRAWABLE_HPP_
#define RSD_DRAWABLE_HPP_

#include <GLFW/glfw3.h>
#include <vector>

namespace rsd {

  struct Drawable {
    std::vector<float> data;
    GLenum element_type;
    GLsizei element_count;

    size_t data_size() const;
  };

}  // namespace rsd

#endif  // RSD_DRAWABLE_HPP_
