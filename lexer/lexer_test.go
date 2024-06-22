package lexer

import (
	tok "github.com/lbwise/interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);
!-/*5;
5 < 10 > 5;
if (5 < 10) {
return true;
} else {
return false;
}
10 == 10;
10 != 9;
`

	tests := []struct {
		expectedType    tok.TokenType
		expectedLiteral string
	}{
		{tok.LET, "let"},
		{tok.IDENT, "five"},
		{tok.ASSIGN, "="},
		{tok.INT, "5"},
		{tok.SEMICOLON, ";"},
		{tok.LET, "let"},
		{tok.IDENT, "ten"},
		{tok.ASSIGN, "="},
		{tok.INT, "10"},
		{tok.SEMICOLON, ";"},
		{tok.LET, "let"},
		{tok.IDENT, "add"},
		{tok.ASSIGN, "="},
		{tok.FUNCTION, "fn"},
		{tok.LPAREN, "("},
		{tok.IDENT, "x"},
		{tok.COMMA, ","},
		{tok.IDENT, "y"},
		{tok.RPAREN, ")"},
		{tok.LBRACE, "{"},
		{tok.IDENT, "x"},
		{tok.PLUS, "+"},
		{tok.IDENT, "y"},
		{tok.SEMICOLON, ";"},
		{tok.RBRACE, "}"},
		{tok.SEMICOLON, ";"},
		{tok.LET, "let"},
		{tok.IDENT, "result"},
		{tok.ASSIGN, "="},
		{tok.IDENT, "add"},
		{tok.LPAREN, "("},
		{tok.IDENT, "five"},
		{tok.COMMA, ","},
		{tok.IDENT, "ten"},
		{tok.RPAREN, ")"},
		{tok.SEMICOLON, ";"},
		{tok.BANG, "!"},
		{tok.MINUS, "-"},
		{tok.SLASH, "/"},
		{tok.ASTERISK, "*"},
		{tok.INT, "5"},
		{tok.SEMICOLON, ";"},
		{tok.INT, "5"},
		{tok.LT, "<"},
		{tok.INT, "10"},
		{tok.GT, ">"},
		{tok.INT, "5"},
		{tok.SEMICOLON, ";"},
		{tok.IF, "if"},
		{tok.LPAREN, "("},
		{tok.INT, "5"},
		{tok.LT, "<"},
		{tok.INT, "10"},
		{tok.RPAREN, ")"},
		{tok.LBRACE, "{"},
		{tok.RETURN, "return"},
		{tok.TRUE, "true"},
		{tok.SEMICOLON, ";"},
		{tok.RBRACE, "}"},
		{tok.ELSE, "else"},
		{tok.LBRACE, "{"},
		{tok.RETURN, "return"},
		{tok.FALSE, "false"},
		{tok.SEMICOLON, ";"},
		{tok.RBRACE, "}"},
		{tok.INT, "10"},
		{tok.EQ, "=="},
		{tok.INT, "10"},
		{tok.SEMICOLON, ";"},
		{tok.INT, "10"},
		{tok.NOT_EQ, "!="},
		{tok.INT, "9"},
		{tok.SEMICOLON, ";"},
		{tok.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
