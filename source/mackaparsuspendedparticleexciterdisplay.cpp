#include "dcpu.hpp"
#include "mackaparsuspendedparticleexciterdisplay.hpp"

namespace dcpu {

  void MackaparSuspendedParticleExciterDisplay::Connect(Dcpu *dcpu) {}

  void MackaparSuspendedParticleExciterDisplay::Execute() {}

  unsigned int MackaparSuspendedParticleExciterDisplay::GetId() const {
    return 0;
  }

  unsigned int MackaparSuspendedParticleExciterDisplay::GetManufacturerId() const {
    return 0;
  }

  Word MackaparSuspendedParticleExciterDisplay::GetVersion() const {
    return 0;
  }

  void MackaparSuspendedParticleExciterDisplay::HandleHardwareInterrupt() {}

}  // namespace dcpu
