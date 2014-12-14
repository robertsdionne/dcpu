#include <iostream>
#include <string>

#include "lexer.hpp"

namespace dcpu {

  const auto Lexer::kAdvancedOpcodeRegex = std::regex(
      "(jsr|int|iag|ias|rfi|iaq|hwn|hwq|hwi)", std::regex_constants::icase);
  const auto Lexer::kBasicOpcodeRegex = std::regex(
      "(set|add|sub|mul|mli|div|dvi|mod|mdi|and|bor|xor|shr|asr"
          "|shl|ifb|ifc|ife|ifn|ifg|ifa|ifl|ifu|adx|sbx|sti|std)", std::regex_constants::icase);
  const auto Lexer::kBinaryRegex = std::regex("0b([01]{1,16})");
  const auto Lexer::kColonRegex = std::regex(":");
  const auto Lexer::kCommaRegex = std::regex(",");
  const auto Lexer::kCommentRegex = std::regex("#\\s*(.*)");
  const auto Lexer::kDataRegex = std::regex("\\.data", std::regex_constants::icase);
  const auto Lexer::kDecimalRegex = std::regex("([0-9]+)");
  const auto Lexer::kHexadecimalRegex = std::regex("0x([0-9a-fA-F]{1,4})");
  const auto Lexer::kIdentifierRegex = std::regex("(\\w+)");
  const auto Lexer::kLeftBracketRegex = std::regex("\\[");
  const auto Lexer::kMinusRegex = std::regex("-");
  const auto Lexer::kNullptrRegex = std::regex("nullptr", std::regex_constants::icase);
  const auto Lexer::kPeekRegex = std::regex("peek", std::regex_constants::icase);
  const auto Lexer::kPickRegex = std::regex("pick", std::regex_constants::icase);
  const auto Lexer::kPlusRegex = std::regex("\\+");
  const auto Lexer::kPopRegex = std::regex("pop", std::regex_constants::icase);
  const auto Lexer::kPushRegex = std::regex("push", std::regex_constants::icase);
  const auto Lexer::kRightBracketRegex = std::regex("\\]");
  const auto Lexer::kSemicolonRegex = std::regex(";");
  const auto Lexer::kStringRegex = std::regex("\"(.*)\"");
  const auto Lexer::kWhitespaceRegex = std::regex("(\\s+)");

  std::ostream &operator <<(std::ostream &out, Token::Type type) {
    switch (type) {
      case Token::Type::kAdvancedOpcode: return out << "kAdvancedOpcode";
      case Token::Type::kBasicOpcode: return out << "kBasicOpcode";
      case Token::Type::kBinary: return out << "kBinary";
      case Token::Type::kColon: return out << "kColon";
      case Token::Type::kComma: return out << "kComma";
      case Token::Type::kComment: return out << "kComment";
      case Token::Type::kData: return out << "kData";
      case Token::Type::kDecimal: return out << "kDecimal";
      case Token::Type::kHexadecimal: return out << "kHexadecimal";
      case Token::Type::kIdentifier: return out << "kIdentifier";
      case Token::Type::kInvalid: return out << "kInvalid";
      case Token::Type::kLeftBracket: return out << "kLeftBracket";
      case Token::Type::kMinus: return out << "kMinus";
      case Token::Type::kNullptr: return out << "kNullptr";
      case Token::Type::kPeek: return out << "kPeek";
      case Token::Type::kPick: return out << "kPick";
      case Token::Type::kPlus: return out << "kPlus";
      case Token::Type::kPop: return out << "kPop";
      case Token::Type::kPush: return out << "kPush";
      case Token::Type::kRightBracket: return out << "kRightBracket";
      case Token::Type::kSemicolon: return out << "kSemicolon";
      case Token::Type::kString: return out << "kString";
      case Token::Type::kWhitespace: return out << "kWhitespace";
      default: return out << "?";
    }
  }

  std::ostream &operator <<(std::ostream &out, Token token) {
    return out << "Token {" << std::endl
        << "  type: " << token.type << std::endl
        << "  value: \"" << token.value << "\"" << std::endl
        << "}";
  }

  Lexer::Lexer(const std::string &input) : input(input), position(input.begin()) {}

  Token Lexer::EatToken() {
    auto new_position = std::string::const_iterator();
    auto result = SeeToken(&new_position);
    position = new_position;
    std::cout << "Eating " << result << std::endl;
    return result;
  }

  Token Lexer::SeeToken(std::string::const_iterator *new_position) const {
    using Type = Token::Type;
    if (position == input.cend()) {
      return Token{Type::kExhausted};
    }
    auto match = std::smatch();
    if (Matches(kAdvancedOpcodeRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kAdvancedOpcode, match[1].str()};
    } else if (Matches(kBasicOpcodeRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kBasicOpcode, match[1].str()};
    } else if (Matches(kBinaryRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kBinary, match[1].str()};
    } else if (Matches(kColonRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kColon, match.str()};
    } else if (Matches(kCommaRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kComma, match.str()};
    } else if (Matches(kCommentRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kComment, match[1].str()};
    } else if (Matches(kDataRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kData, match.str()};
    } else if (Matches(kHexadecimalRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kHexadecimal, match[1].str()};
    } else if (Matches(kDecimalRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kDecimal, match[1].str()};
    } else if (Matches(kLeftBracketRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kLeftBracket, match.str()};
    } else if (Matches(kMinusRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kMinus, match.str()};
    } else if (Matches(kNullptrRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kNullptr, match.str()};
    } else if (Matches(kPeekRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kPeek, match.str()};
    } else if (Matches(kPickRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kPick, match.str()};
    } else if (Matches(kPlusRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kPlus, match.str()};
    } else if (Matches(kPopRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kPop, match.str()};
    } else if (Matches(kPushRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kPush, match.str()};
    } else if (Matches(kRightBracketRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kRightBracket, match.str()};
    } else if (Matches(kSemicolonRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kSemicolon, match.str()};
    } else if (Matches(kStringRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kString, match[1].str()};
    } else if (Matches(kIdentifierRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kIdentifier, match[1].str()};
    } else if (Matches(kWhitespaceRegex, &match)) {
      MaybeAdvance(match, new_position);
      return Token{Type::kWhitespace, match[1].str()};
    } else {
      return Token{Type::kInvalid};
    }
  }

}  // namespace dcpu
