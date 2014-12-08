set(
  DCPU_SOURCES
  assembler.cpp
  clock.cpp
  dcpu.cpp
  disassembler.cpp
  dsl.cpp
  disassembler.cpp
  dsl.cpp
  lexer.cpp
  parser.cpp
  generated/program.pb.cc
)

add_custom_command(
  OUTPUT generated/ generated/program.pb.cc generated/program.pb.h
  COMMAND mkdir -p generated
  COMMAND ../libraries/protobuf/src/protoc --proto_path . --cpp_out generated/ program.proto
  DEPENDS protobuf_target
)

add_library(dcpu_library ${DCPU_SOURCES})

add_executable(dcpu main.cpp ${DCPU_SOURCES})

target_link_libraries(dcpu libprotobuf.a ncurses)
