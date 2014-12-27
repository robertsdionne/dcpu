#include <cmath>
#include <drawable.hpp>
#include <glm/glm.hpp>

#include "dcpu.hpp"
#include "mackaparsuspendedparticleexciterdisplay.hpp"

namespace dcpu {

  const std::unordered_map<Word, MackaparSuspendedParticleExciterDisplay::Vertex>
      MackaparSuspendedParticleExciterDisplay::kOrigins = {
    {4, {0.5f * std::cos(kTheta0), 0.0f, 0.5f * std::sin(kTheta0), 0.0f, 0.0f, 0.0f, 1.0f}},
    {2, {0.5f * std::cos(kTheta1), 0.0f, 0.5f * std::sin(kTheta1), 0.0f, 0.0f, 0.0f, 1.0f}},
    {1, {0.5f * std::cos(kTheta2), 0.0f, 0.5f * std::sin(kTheta2), 0.0f, 0.0f, 0.0f, 1.0f}}
  };

  rsd::Drawable MackaparSuspendedParticleExciterDisplay::BeamView() {
    auto data = std::vector<float>{};
    auto count = 0;
    auto alpha = 1.0f / std::sqrt(number_of_vertices) / 2.0f;
    for (auto i = 0; i < number_of_vertices; ++i) {
      Word *memory_map = dcpu->address(memory_map_offset);
      auto word1 = memory_map[2 * i + 1];
      auto word3 = memory_map[2 * i + 3];
      auto color0 = (word1 >> kColorShift) & kColorMask;
      auto color1 = (word3 >> kColorShift) & kColorMask;
      auto v0 = BuildVertex(memory_map[2 * i + 0], memory_map[2 * i + 1], alpha);
      auto v1 = BuildVertex(memory_map[2 * i + 2], memory_map[2 * i + 3], alpha);
      auto v0_black = Vertex{v0.x, v0.y, v0.z, 0.0f, 0.0f, 0.0f, v0.a};
      auto v1_black = Vertex{v1.x, v1.y, v1.z, 0.0f, 0.0f, 0.0f, v1.a};
      for (auto c = 1; c <= 4; c <<= 1) {
        if ((color0 & c)  > 0 || (color1 & c) > 0) {
          count += 1;
          auto r = static_cast<float>(4 == c);
          auto g = static_cast<float>(2 == c);
          auto b = static_cast<float>(1 == c);
          auto v0_prime = Vertex{v0.x, v0.y, v0.z, r, g, b, v0.a};
          auto v1_prime = Vertex{v1.x, v1.y, v1.z, r, g, b, v1.a};
          auto v0_final = (color0 & c) ? v0_prime : v0_black;
          auto v1_final = (color1 & c) ? v1_prime : v1_black;
          auto origin = kOrigins.at(c);
          origin.a = alpha;
          data.insert(data.cend(), {
            v0_final.x, v0_final.y, v0_final.z, v0_final.r, v0_final.g, v0_final.b, v0_final.a,
            v1_final.x, v1_final.y, v1_final.z, v1_final.r, v1_final.g, v1_final.b, v1_final.a,
            origin.x, origin.y, origin.z, origin.r, origin.g, origin.b, origin.a
          });
        }
      }
    }
    return {data, GL_TRIANGLES, 3 * count};
  }

  MackaparSuspendedParticleExciterDisplay::Vertex
      MackaparSuspendedParticleExciterDisplay::BuildVertex(Word word0, Word word1, float alpha) {
    auto x = (word0 & kCoordinateMask) / kByteMaximum - 0.5f;
    auto z = (word0 >> kYShift) / kByteMaximum - 0.5f;
    auto y = (word1 & kCoordinateMask) / kByteMaximum + 0.5f;
    auto color = (word1 >> kColorShift) & kColorMask;
    auto r = (0x4 & color) ? 1.0f : 0.0f;
    auto g = (0x2 & color) ? 1.0f : 0.0f;
    auto b = (0x1 & color) ? 1.0f : 0.0f;
    auto theta = static_cast<float>(M_PI / 180.0f * angle);
    auto cos0 = std::cos(theta);
    auto sin0 = std::sin(theta);
    return {x * cos0 + z * sin0, y, x * sin0 - z * cos0, r, g, b, alpha};
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
    auto data = std::vector<float>{};
    for (auto i = 0; i < number_of_vertices; ++i) {
      Word *memory_map = dcpu->address(memory_map_offset);
      auto vertex = BuildVertex(memory_map[2 * i], memory_map[2 * i + 1]);
      data.insert(data.cend(), {
        vertex.x, vertex.y, vertex.z, vertex.r, vertex.g, vertex.b, vertex.a
      });
    }
    return {data, GL_LINE_STRIP, number_of_vertices};
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
