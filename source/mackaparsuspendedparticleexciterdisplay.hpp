#ifndef DCPU_MACKAPARSUSPENDEDPARTICLEEXCITERDISPLAY_HPP_
#define DCPU_MACKAPARSUSPENDEDPARTICLEEXCITERDISPLAY_HPP_

#include "hardware.hpp"

namespace dcpu {

  using Word = unsigned short;

  class Dcpu;

  class MackaparSuspendedParticleExciterDisplay : public Hardware {
  public:
    MackaparSuspendedParticleExciterDisplay() = default;

    virtual ~MackaparSuspendedParticleExciterDisplay() = default;

    void Connect(Dcpu *dcpu) override;

    void Execute() override;

    unsigned int GetId() const override;

    unsigned int GetManufacturerId() const override;

    Word GetVersion() const override;

    void HandleHardwareInterrupt() override;
  };

}  // namespace dcpu

#endif  // DCPU_MACKAPARSUSPENDEDPARTICLEEXCITERDISPLAY_HPP_
