#ifndef RSD_RENDERER_HPP_
#define RSD_RENDERER_HPP_

#include "interface.hpp"

namespace rsd {

  class Renderer {
    DECLARE_INTERFACE(Renderer);

  public:
    virtual void Change(int width, int height) = 0;

    virtual void Create() = 0;

    virtual void Render() = 0;
  };

}  // namespace rsd

#endif  // RSD_RENDERER_HPP_
