#include <chrono>

#include "clock.hpp"
#include "dcpu.hpp"

namespace dcpu {

  void Clock::Execute() {
    if (interval) {
      auto now = std::chrono::high_resolution_clock::now();
      if (last_tick + std::chrono::duration<double>(interval / kClockFrequency) <= now) {
        last_tick = now;
        ticks += 1;
        if (message) {
          dcpu->Interrupt(message);
        }
      }
    }
  }

  void Clock::Connect(Dcpu *dcpu) {
    this->dcpu = dcpu;
  }

  void Clock::HandleHardwareInterrupt() {
    if (dcpu) {
      switch (dcpu->register_a) {
        case 0: {
          SetInterval(dcpu->register_b);
          break;
        }
        case 1: {
          dcpu->register_c = GetElapsedTicks();
          break;
        }
        case 2: {
          SetSoftwareInterrupt(dcpu->register_b);
          break;
        }
        default: break;
      }
    }
  }

  void Clock::SetSoftwareInterrupt(Word message) {
    this->message = message;
  }

  void Clock::SetInterval(Word interval) {
    this->interval = interval;
    ticks = 0;
    last_tick = std::chrono::high_resolution_clock::now();
  }

  Word Clock::GetElapsedTicks() {
    return ticks;
  }

}  // namespace dcpu
