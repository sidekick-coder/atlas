package entry

import (
	"fmt"
	"strings"
	"unicode"
)

// tokenType represents the type of a lexical token.
type tokenType int

const (
	tokenIdent  tokenType = iota // identifier: key or value
	tokenOp                      // =, !=, >, <, >=, <=
	tokenAnd                     // &&
	tokenOr                      // ||
	tokenNot                     // !
	tokenLParen                  // (
	tokenRParen                  // )
	tokenEOF
)

type token struct {
	typ tokenType
	val string
}

func tokenize(input string) ([]token, error) {
	var tokens []token
	i := 0
	for i < len(input) {
		ch := rune(input[i])

		if unicode.IsSpace(ch) {
			i++
			continue
		}

		switch {
		case ch == '(':
			tokens = append(tokens, token{tokenLParen, "("})
			i++
		case ch == ')':
			tokens = append(tokens, token{tokenRParen, ")"})
			i++
		case ch == '&' && i+1 < len(input) && input[i+1] == '&':
			tokens = append(tokens, token{tokenAnd, "&&"})
			i += 2
		case ch == '|' && i+1 < len(input) && input[i+1] == '|':
			tokens = append(tokens, token{tokenOr, "||"})
			i += 2
		case ch == '!' && i+1 < len(input) && input[i+1] == '=':
			tokens = append(tokens, token{tokenOp, "!="})
			i += 2
		case ch == '>' && i+1 < len(input) && input[i+1] == '=':
			tokens = append(tokens, token{tokenOp, ">="})
			i += 2
		case ch == '<' && i+1 < len(input) && input[i+1] == '=':
			tokens = append(tokens, token{tokenOp, "<="})
			i += 2
		case ch == '=':
			tokens = append(tokens, token{tokenOp, "="})
			i++
		case ch == '>':
			tokens = append(tokens, token{tokenOp, ">"})
			i++
		case ch == '<':
			tokens = append(tokens, token{tokenOp, "<"})
			i++
		case ch == '!':
			tokens = append(tokens, token{tokenNot, "!"})
			i++
		case ch == '"' || ch == '\'':
			quote := ch
			i++ // skip opening quote
			j := i
			for j < len(input) && rune(input[j]) != quote {
				j++
			}
			if j >= len(input) {
				return nil, fmt.Errorf("unterminated string literal")
			}
			tokens = append(tokens, token{tokenIdent, input[i:j]})
			i = j + 1 // skip closing quote
		default:
			j := i
			for j < len(input) {
				c := rune(input[j])
				if unicode.IsSpace(c) || c == '(' || c == ')' ||
					c == '&' || c == '|' || c == '!' ||
					c == '=' || c == '>' || c == '<' {
					break
				}
				j++
			}
			if j == i {
				return nil, fmt.Errorf("unexpected character %q at position %d", ch, i)
			}
			tokens = append(tokens, token{tokenIdent, input[i:j]})
			i = j
		}
	}
	tokens = append(tokens, token{tokenEOF, ""})
	return tokens, nil
}

// QueryNode is an AST node produced by the query parser.
type QueryNode struct {
	kind string // "binary" | "condition" | "negate"

	// binary: AND / OR
	op    string
	left  *QueryNode
	right *QueryNode

	// condition: key op value  (op may be "EXISTS" when no operator is given)
	key      string
	operator string
	value    string
}

type queryParser struct {
	tokens []token
	pos    int
}

func (p *queryParser) peek() token {
	if p.pos >= len(p.tokens) {
		return token{tokenEOF, ""}
	}
	return p.tokens[p.pos]
}

func (p *queryParser) consume() token {
	t := p.peek()
	p.pos++
	return t
}

func (p *queryParser) parseExpr() (*QueryNode, error) {
	return p.parseOr()
}

func (p *queryParser) parseOr() (*QueryNode, error) {
	left, err := p.parseAnd()
	if err != nil {
		return nil, err
	}
	for p.peek().typ == tokenOr {
		p.consume()
		right, err := p.parseAnd()
		if err != nil {
			return nil, err
		}
		left = &QueryNode{kind: "binary", op: "OR", left: left, right: right}
	}
	return left, nil
}

func (p *queryParser) parseAnd() (*QueryNode, error) {
	left, err := p.parsePrimary()
	if err != nil {
		return nil, err
	}
	for p.peek().typ == tokenAnd {
		p.consume()
		right, err := p.parsePrimary()
		if err != nil {
			return nil, err
		}
		left = &QueryNode{kind: "binary", op: "AND", left: left, right: right}
	}
	return left, nil
}

func (p *queryParser) parsePrimary() (*QueryNode, error) {
	t := p.peek()

	if t.typ == tokenLParen {
		p.consume()
		node, err := p.parseExpr()
		if err != nil {
			return nil, err
		}
		if p.peek().typ != tokenRParen {
			return nil, fmt.Errorf("expected closing parenthesis")
		}
		p.consume()
		return node, nil
	}

	// !key  →  key does not exist
	if t.typ == tokenNot {
		p.consume()
		if p.peek().typ != tokenIdent {
			return nil, fmt.Errorf("expected identifier after '!'")
		}
		key := p.consume().val
		return &QueryNode{kind: "negate", key: key}, nil
	}

	if t.typ == tokenIdent {
		key := p.consume().val
		if p.peek().typ == tokenOp {
			op := p.consume().val
			if p.peek().typ != tokenIdent {
				return nil, fmt.Errorf("expected value after operator %q", op)
			}
			val := p.consume().val
			return &QueryNode{kind: "condition", key: key, operator: op, value: val}, nil
		}
		// bare key: check existence
		return &QueryNode{kind: "condition", key: key, operator: "EXISTS"}, nil
	}

	return nil, fmt.Errorf("unexpected token %q", t.val)
}

// ParseQuery parses a slice of filter strings into a single AST.
// Multiple strings are combined with implicit AND.
func ParseQuery(queries []string) (*QueryNode, error) {
	if len(queries) == 0 {
		return nil, nil
	}
	combined := strings.Join(queries, " && ")
	tokens, err := tokenize(combined)
	if err != nil {
		return nil, err
	}
	p := &queryParser{tokens: tokens}
	node, err := p.parseExpr()
	if err != nil {
		return nil, err
	}
	if p.peek().typ != tokenEOF {
		return nil, fmt.Errorf("unexpected token %q", p.peek().val)
	}
	return node, nil
}

// BuildSQL converts a QueryNode AST into a SQL condition string.
// Positional parameters are appended to params.
func BuildSQL(node *QueryNode, params *[]interface{}) (string, error) {
	if node == nil {
		return "1=1", nil
	}

	switch node.kind {
	case "binary":
		left, err := BuildSQL(node.left, params)
		if err != nil {
			return "", err
		}
		right, err := BuildSQL(node.right, params)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("(%s %s %s)", left, node.op, right), nil

	case "condition":
		if node.operator == "EXISTS" {
			*params = append(*params, node.key)
			return "EXISTS (SELECT 1 FROM entry_metas em WHERE em.entry_id = entries.id AND em.name = ?)", nil
		}
		isNumeric := node.operator == ">" || node.operator == "<" ||
			node.operator == ">=" || node.operator == "<="
		*params = append(*params, node.key, node.value)
		if isNumeric {
			return fmt.Sprintf(
				"EXISTS (SELECT 1 FROM entry_metas em WHERE em.entry_id = entries.id AND em.name = ? AND CAST(em.value AS NUMERIC) %s ?)",
				node.operator,
			), nil
		}
		return fmt.Sprintf(
			"EXISTS (SELECT 1 FROM entry_metas em WHERE em.entry_id = entries.id AND em.name = ? AND em.value %s ?)",
			node.operator,
		), nil

	case "negate":
		*params = append(*params, node.key)
		return "NOT EXISTS (SELECT 1 FROM entry_metas em WHERE em.entry_id = entries.id AND em.name = ?)", nil
	}

	return "", fmt.Errorf("unknown node kind %q", node.kind)
}
