#ifndef DCPU_MACKAPARSUSPENDEDPARTICLEEXCITERDISPLAY_HPP_
#define DCPU_MACKAPARSUSPENDEDPARTICLEEXCITERDISPLAY_HPP_

#include <drawable.hpp>
#include <unordered_map>

#include "hardware.hpp"

namespace dcpu {

  using Word = unsigned short;

  class Dcpu;

  class MackaparSuspendedParticleExciterDisplay : public Hardware {
  public:
    enum class Error : Word {
      kNone = 0x0000,
      kBroken = 0xFFFF
    };

    enum class State : Word {
      kNoData = 0x0000,
      kRunning = 0x0001,
      kTurning = 0x0002
    };

    static constexpr unsigned int kId = 0x42babf3c;
    static constexpr unsigned int kManufacturerId = 0x1eb37e91;
    static constexpr Word kVersion = 0x0003;

    MackaparSuspendedParticleExciterDisplay() = default;

    virtual ~MackaparSuspendedParticleExciterDisplay() = default;

    rsd::Drawable BeamView();

    inline void Connect(Dcpu *dcpu) override {
      this->dcpu = dcpu;
    }

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

    void Poll(Word *register_b, Word *register_c);

    rsd::Drawable LineView();

    void Map(Word register_x, Word register_y);

    void RotateTo(Word register_x);

    void Override();

  private:
    struct Vertex {
      float x, y, z;
      float r, g, b, a;
    };

    static constexpr float kAngularVelocity = 50.0f;
    static constexpr float kByteMaximum = 255.0f;
    static constexpr Word kCoordinateMask = 0x00FF;
    static constexpr Word kColorMask = 0x0007;
    static constexpr Word kColorShift = 8;
    static constexpr Word kIntensityMask = 0x0001;
    static constexpr Word kIntensityShift = 9;
    static constexpr float kLerpAlpha = 0.01f;
    static constexpr float kTurnEpsilon = 1e-2f;
    static constexpr Word kYShift = 8;
    static constexpr float kOffset = -4.0f * M_PI / 3.0f + M_PI / 2.0f;
    static constexpr float kTheta0 = 0.0f + kOffset;
    static constexpr float kTheta1 = 2.0f * M_PI / 3.0f + kOffset;
    static constexpr float kTheta2 = 4.0f * M_PI / 3.0f + kOffset;
    static const std::unordered_map<Word, Vertex> kOrigins;

    Vertex BuildVertex(Word word0, Word word1, float alpha = 0.8f);

    Dcpu *dcpu = nullptr;
    State current_state = State::kNoData;
    Error last_error = Error::kNone;
    Word memory_map_offset = 0x0000, number_of_vertices = 0;
    float angle = 0.0f, target_angle = 90.0f;
  };

}  // namespace dcpu

#endif  // DCPU_MACKAPARSUSPENDEDPARTICLEEXCITERDISPLAY_HPP_
