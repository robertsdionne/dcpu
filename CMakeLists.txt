cmake_minimum_required (VERSION 2.8.12)
project(rsd)

add_definitions(-DGL_GLEXT_PROTOTYPES)

add_subdirectory(libraries/glfw EXCLUDE_FROM_ALL)
add_subdirectory(libraries/glxw EXCLUDE_FROM_ALL)

if (NOT WIN32)
  add_definitions(-std=c++11)
endif()

include_directories(
  libraries/glfw/include
  libraries/glm
  libraries/glxw/include
  include
)

add_subdirectory(source)
