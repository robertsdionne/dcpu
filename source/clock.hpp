#ifndef DCPU_CLOCK_HPP_
#define DCPU_CLOCK_HPP_

#include <chrono>

#include "dcpu.hpp"
#include "hardware.hpp"

namespace dcpu {
  
  class Clock : public Hardware {
  public:
    static constexpr float kClockFrequency = 60.0;

    Clock() = default;

    virtual ~Clock() = default;

    void Connect(Dcpu *dcpu) override;

    void Execute() override;

    void HandleHardwareInterrupt() override;

    void SetSoftwareInterrupt(Word message);

    void SetInterval(Word interval);

    Word GetElapsedTicks();

  private:
    std::chrono::high_resolution_clock::time_point last_tick{};
    Dcpu *dcpu = nullptr;
    Word interval = 0, message = 0, ticks = 0;
  };

}  // namespace dcpu

#endif  // DCPU_CLOCK_HPP_
