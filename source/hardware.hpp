#ifndef DCPU_HARDWARE_HPP_
#define DCPU_HARDWARE_HPP_

namespace dcpu {

  using Word = unsigned short;

  class Dcpu;

  class Hardware {
  public:
    Hardware() = default;

    virtual ~Hardware() = default;

    virtual void Connect(Dcpu *dcpu) = 0;

    virtual void Execute() = 0;

    virtual unsigned int GetId() const = 0;

    virtual unsigned int GetManufacturerId() const = 0;

    virtual Word GetVersion() const = 0;

    virtual void HandleHardwareInterrupt() = 0;
  };

}  // namespace dcpu

#endif  // DCPU_HARDWARE_HPP_
