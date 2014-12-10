#ifndef DCPU_LEXER_HPP_
#define DCPU_LEXER_HPP_

#include <iostream>
#include <regex>
#include <string>

namespace dcpu {

  struct Token {
    enum class Type {
      kAdvancedOpcode,
      kBasicOpcode,
      kBinary,
      kColon,
      kComma,
      kComment,
      kData,
      kDecimal,
      kHexadecimal,
      kIdentifier,
      kInvalid,
      kLeftBracket,
      kMinus,
      kNullptr,
      kPeek,
      kPick,
      kPlus,
      kPop,
      kPush,
      kRightBracket,
      kSemicolon,
      kString,
      kWhitespace
    };
    Type type;
    std::string value;
  };

  std::ostream &operator <<(std::ostream &out, Token::Type type);
  std::ostream &operator <<(std::ostream &out, Token token);

  class Lexer {
  public:
    static const std::regex kAdvancedOpcodeRegex;
    static const std::regex kBasicOpcodeRegex;
    static const std::regex kBinaryRegex;
    static const std::regex kColonRegex;
    static const std::regex kCommaRegex;
    static const std::regex kCommentRegex;
    static const std::regex kDataRegex;
    static const std::regex kDecimalRegex;
    static const std::regex kHexadecimalRegex;
    static const std::regex kIdentifierRegex;
    static const std::regex kLeftBracketRegex;
    static const std::regex kMinusRegex;
    static const std::regex kNullptrRegex;
    static const std::regex kPeekRegex;
    static const std::regex kPickRegex;
    static const std::regex kPlusRegex;
    static const std::regex kPopRegex;
    static const std::regex kPushRegex;
    static const std::regex kRightBracketRegex;
    static const std::regex kSemicolonRegex;
    static const std::regex kStringRegex;
    static const std::regex kWhitespaceRegex;

  public:
    Lexer(const std::string &input);

    virtual ~Lexer() = default;

    Token EatToken();

    Token SeeToken(std::string::const_iterator *new_position = nullptr) const;

  private:
    inline void MaybeAdvance(
        const std::smatch &match, std::string::const_iterator *new_position) const {
      if (new_position) {
        *new_position = match.suffix().first;
      }
    }

    inline bool Matches(const std::regex &regex, std::smatch *match) const {
      return std::regex_search(position, input.end(), *match, regex) && !match->prefix().length();
    }

  private:
    const std::string &input;
    std::string::const_iterator position;
  };

}  // namespace dcpu



#endif  // DCPU_LEXER_HPP_
