#include <drawable.hpp>
#include <glm/glm.hpp>

#include "dcpu.hpp"
#include "mackaparsuspendedparticleexciterdisplay.hpp"

namespace dcpu {

  rsd::Drawable MackaparSuspendedParticleExciterDisplay::BeamView() {
    return {};
  }

  void MackaparSuspendedParticleExciterDisplay::Execute() {
    current_state = number_of_vertices ? State::kRunning : State::kNoData;
    auto new_angle = glm::mix(angle, target_angle, kLerpAlpha);
    current_state = new_angle - angle > kTurnEpsilon ? State::kTurning : current_state;
    angle = new_angle;
  }

  void MackaparSuspendedParticleExciterDisplay::HandleHardwareInterrupt() {
    if (dcpu) {
      switch (dcpu->register_a) {
        case 0: {
          Poll(&dcpu->register_b, &dcpu->register_c);
          break;
        }
        case 1: {
          Map(dcpu->register_x, dcpu->register_y);
          break;
        }
        case 2: {
          RotateTo(dcpu->register_x);
          break;
        }
        default: break;
      }
    }
  }

  void MackaparSuspendedParticleExciterDisplay::Poll(Word *register_b, Word *register_c) {
    *register_b = static_cast<Word>(current_state);
    *register_c = static_cast<Word>(last_error);
    last_error = Error::kNone;
  }

  rsd::Drawable MackaparSuspendedParticleExciterDisplay::LineView() {
    return {};
  }

  void MackaparSuspendedParticleExciterDisplay::Map(Word register_x, Word register_y) {
    memory_map_offset = register_x;
    number_of_vertices = register_y;
    last_error = number_of_vertices > 128 ? Error::kBroken : Error::kNone;
  }

  void MackaparSuspendedParticleExciterDisplay::RotateTo(Word register_x) {
    target_angle = register_x;
  }

}  // namespace dcpu
