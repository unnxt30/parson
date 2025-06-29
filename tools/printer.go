package tools

import (
	"fmt"
	"strings"

	"github.com/unnxt30/parson/pkg/analysis"
)

// PrettyPrinter formats and displays tokens in a readable way
type PrettyPrinter struct {
	tokens []analysis.Token
}

// NewPrettyPrinter creates a new pretty printer instance
func NewPrettyPrinter(tokens []analysis.Token) *PrettyPrinter {
	return &PrettyPrinter{tokens: tokens}
}

// Print displays all tokens in a formatted table
func (pp *PrettyPrinter) Print() {
	if len(pp.tokens) == 0 {
		fmt.Println("No tokens to display")
		return
	}

	fmt.Println("TOKEN ANALYSIS")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%-4s %-15s %-20s %-10s %-10s %-6s\n",
		"#", "TYPE", "VALUE", "START", "END", "LINE")
	fmt.Println(strings.Repeat("-", 80))

	for i, token := range pp.tokens {
		value := pp.formatValue(token.Value)
		fmt.Printf("%-4d %-15s %-20s %-10d %-10d %-6d\n",
			i+1, token.Type, value, token.Start, token.End, token.Line)
	}

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("Total tokens: %d\n", len(pp.tokens))
}

// PrintCompact displays tokens in a more compact format
func (pp *PrettyPrinter) PrintCompact() {
	if len(pp.tokens) == 0 {
		fmt.Println("No tokens to display")
		return
	}

	fmt.Print("Tokens: ")
	for i, token := range pp.tokens {
		if i > 0 {
			fmt.Print(" → ")
		}
		fmt.Printf("%s", token.Type)
		if token.Value != "" && token.Type != analysis.EOF {
			fmt.Printf("('%s')", pp.formatValue(token.Value))
		}
	}
	fmt.Println()
}

// PrintByLine groups tokens by line number
func (pp *PrettyPrinter) PrintByLine() {
	if len(pp.tokens) == 0 {
		fmt.Println("No tokens to display")
		return
	}

	lineGroups := make(map[int][]analysis.Token)
	maxLine := 0

	// Group tokens by line
	for _, token := range pp.tokens {
		lineGroups[token.Line] = append(lineGroups[token.Line], token)
		if token.Line > maxLine {
			maxLine = token.Line
		}
	}

	fmt.Println("TOKENS BY LINE")
	fmt.Println(strings.Repeat("=", 60))

	for line := 1; line <= maxLine; line++ {
		if tokens, exists := lineGroups[line]; exists {
			fmt.Printf("Line %d: ", line)
			for i, token := range tokens {
				if i > 0 {
					fmt.Print(" ")
				}
				fmt.Printf("%s", token.Type)
				if token.Value != "" && token.Type != analysis.EOF {
					fmt.Printf("('%s')", pp.formatValue(token.Value))
				}
			}
			fmt.Println()
		}
	}
	fmt.Println(strings.Repeat("=", 60))
}

// formatValue formats token values for display, handling special characters
func (pp *PrettyPrinter) formatValue(value string) string {
	if value == "" {
		return "<empty>"
	}

	// Replace special characters with readable representations
	formatted := strings.ReplaceAll(value, "\n", "\\n")
	formatted = strings.ReplaceAll(formatted, "\t", "\\t")
	formatted = strings.ReplaceAll(formatted, "\r", "\\r")

	// Truncate long values
	if len(formatted) > 18 {
		formatted = formatted[:15] + "..."
	}

	return formatted
}

// PrintSummary displays a summary of token types and their counts
func (pp *PrettyPrinter) PrintSummary() {
	if len(pp.tokens) == 0 {
		fmt.Println("No tokens to display")
		return
	}

	typeCounts := make(map[analysis.TokenType]int)
	for _, token := range pp.tokens {
		typeCounts[token.Type]++
	}

	fmt.Println("TOKEN SUMMARY")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("%-20s %s\n", "TOKEN TYPE", "COUNT")
	fmt.Println(strings.Repeat("-", 40))

	for tokenType, count := range typeCounts {
		fmt.Printf("%-20s %d\n", tokenType, count)
	}

	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("Total: %d tokens\n", len(pp.tokens))
}
