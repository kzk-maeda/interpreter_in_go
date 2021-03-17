package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

// New 関数
// Lexer構造体を返却する
// input: 解析するコードの実態が格納
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken 関数
// readPositionにある文字に対応するtokenを返却する
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	
	l.skipWhitespace()
	
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal) // 予約語かどうかを判定
			return tok
			} else if isDigit(l.ch) {
				tok.Type = token.INT
				tok.Literal = l.readNumber()
				return tok
				} else {
					tok = newToken(token.ILLEGAL, l.ch)
				}
				
			}
			l.readChar()
			return tok
}

// ********************************
// ******** Private Method ********
// ********************************

// newToken 関数
// readPositionにあるLiteralに対応するtokenを返却する
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readChar 関数
// Lexer構造体のinput（コードの実態）を1byteずつ読んでいく
// Lexer.position: 現在読んでいる文字
// Lexer.readPosition: 次に読む文字
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// readIdentifier 関数
// positionにある文字がLetterだった場合、Letterが区切られるまで続けて読み取る
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber 関数
// positionにある文字がNumberだった場合、区切られるまで続けて読み取る
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace 関数
// 空白文字系（スペース、タブ、改行）を読み飛ばす
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// isLetter 関数
// 大文字・小文字の英字と_を受け取った場合trueを返す
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit 関数
// 数字を受け取った場合trueを返す
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}