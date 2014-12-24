#ifndef DCPU_CLOCK_HPP_
#define DCPU_CLOCK_HPP_

#include <chrono>

#include "dcpu.hpp"
#include "hardware.hpp"

namespace dcpu {

  class Clock : public Hardware {
  public:
    static constexpr float kClockFrequency = 60.0;
    static constexpr unsigned int kId = 0x12d0b402;
    static constexpr unsigned int kManufacturerId = 0x00000000;
    static constexpr Word kVersion = 0x1;

    Clock() = default;

    virtual ~Clock() = default;

    void Connect(Dcpu *dcpu) override;

    void Execute() override;

    inline unsigned int GetId() const override {
      return kId;
    }

    inline unsigned int GetManufacturerId() const override {
      return kManufacturerId;
    }

    inline Word GetVersion() const override {
      return kVersion;
    }

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
