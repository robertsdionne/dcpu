#ifndef RSD_APPLICATION_HPP_
#define RSD_APPLICATION_HPP_

#include "interface.hpp"

namespace rsd {

  class Application {
    DECLARE_INTERFACE(Application);

  public:
    virtual int Run() = 0;
  };

}  // namespace rsd

#endif  // RSD_APPLICATION_HPP_
