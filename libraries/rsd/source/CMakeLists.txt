add_library(rsd buffer.cpp drawable.cpp glfwapplication.cpp mouse.cpp program.cpp shader.cpp
    vertexarray.cpp vertexformat.cpp)

if (NOT WIN32)
  target_link_libraries(rsd dl glfw glxw ${GLFW_LIBRARIES})
else()
  target_link_libraries(rsd glfw glxw ${GLFW_LIBRARIES})
endif()
