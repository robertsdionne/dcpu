#ifndef DCPU_HARDWARE_HPP_
#define DCPU_HARDWARE_HPP_

namespace dcpu {

  class Dcpu;

  class Hardware {
  public:
    Hardware() = default;

    virtual ~Hardware() = default;

    virtual void Connect(Dcpu *dcpu) = 0;

    virtual void Execute() = 0;

    virtual void HandleHardwareInterrupt() = 0;
  };

}  // namespace dcpu

#endif  // DCPU_HARDWARE_HPP_
