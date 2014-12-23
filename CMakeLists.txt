cmake_minimum_required(VERSION 2.8.12)
project(dcpu CXX)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -stdlib=libc++ -g")
add_definitions(-DGL_GLEXT_PROTOTYPES)

add_subdirectory(libraries/gtest-1.7.0)
add_subdirectory(libraries/rsd)

include_directories(
  SYSTEM
  libraries/gtest-1.7.0/include
  libraries/protobuf/src
  libraries/rsd/include
  libraries/rsd/libraries/glfw/include
  libraries/rsd/libraries/glm
  libraries/rsd/libraries/glxw
)

include(ExternalProject)
ExternalProject_Add(
  protobuf_target
  PREFIX libraries/protobuf
  SOURCE_DIR libraries/protobuf
  BINARY_DIR libraries/protobuf
  CONFIGURE_COMMAND ./autogen.sh; ./configure CC=${CMAKE_C_COMPILER} CXX=${CMAKE_CXX_COMPILER} CXXFLAGS=${CMAKE_CXX_FLAGS}
  BUILD_COMMAND $(MAKE) CC=${CMAKE_C_COMPILER} CXX=${CMAKE_CXX_COMPILER} CXXCPP="${CMAKE_CXX_COMPILER} -E" CXXFLAGS=${CMAKE_CXX_FLAGS}
  INSTALL_COMMAND true
)
link_directories(libraries/protobuf/src/.libs)

add_subdirectory(source)
add_subdirectory(tests)
