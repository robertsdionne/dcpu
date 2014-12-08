cmake_minimum_required(VERSION 2.8.12)
project(dcpu CXX)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++11 -stdlib=libc++")

add_subdirectory(libraries/gtest-1.7.0)

include_directories(
  SYSTEM
  libraries/gtest-1.7.0/include
  libraries/protobuf/src
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
