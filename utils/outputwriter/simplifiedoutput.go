package outputwriter

import (
	"fmt"
	"strings"

	"github.com/jfrog/jfrog-cli-security/utils/severityutils"
)

const (
	simpleSeparator = ", "
)

type SimplifiedOutput struct {
	MarkdownOutput
}

func (smo *SimplifiedOutput) Separator() string {
	return simpleSeparator
}

func (smo *SimplifiedOutput) SeverityIcon(severity severityutils.Severity) string {
	return severityutils.GetSeverityIcon(severity)
}

func (smo *SimplifiedOutput) FormattedSeverity(severity, _ string) string {
	return severity
}

func (smo *SimplifiedOutput) Image(source ImageSource) string {
	return MarkAsBold(GetSimplifiedTitle(source))
}

func (smo *SimplifiedOutput) MarkInCenter(content string) string {
	return content
}

func (smo *SimplifiedOutput) MarkAsDetails(summary string, subTitleDepth int, content string) string {
	delimiter := "\n"
	if subTitleDepth == 0 {
		delimiter = ": "
	}
	return fmt.Sprintf("%s%s%s", smo.MarkAsTitle(summary, subTitleDepth), delimiter, content)
}

func (smo *SimplifiedOutput) MarkAsTitle(title string, subTitleDepth int) string {
	if subTitleDepth == 0 {
		return title
	}
	return fmt.Sprintf("%s\n%s %s\n%s", SectionDivider(), strings.Repeat("#", subTitleDepth), title, SectionDivider())
}
