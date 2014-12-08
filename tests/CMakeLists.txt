include_directories(../source)

add_executable(
  alltests
  dcpu_test.cpp
  disassembler_test.cpp
)

target_link_libraries(alltests dcpu_library gtest gtest_main libprotobuf.a)
