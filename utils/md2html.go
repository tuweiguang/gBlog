package utils

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"regexp"
)

func MarkdownToHTML(md string) string {
	//myHTMLFlags := 0 |
	//	blackfriday.HTML_USE_XHTML |
	//	blackfriday.HTML_USE_SMARTYPANTS |
	//	blackfriday.HTML_SMARTYPANTS_FRACTIONS |
	//	blackfriday.HTML_SMARTYPANTS_DASHES |
	//	blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
	//
	//myExtensions := 0 |
	//	blackfriday.EXTENSION_NO_INTRA_EMPHASIS |
	//	blackfriday.EXTENSION_TABLES |
	//	blackfriday.EXTENSION_FENCED_CODE |
	//	blackfriday.EXTENSION_AUTOLINK |
	//	blackfriday.EXTENSION_STRIKETHROUGH |
	//	blackfriday.EXTENSION_SPACE_HEADERS |
	//	blackfriday.EXTENSION_HEADER_IDS |
	//	blackfriday.EXTENSION_BACKSLASH_LINE_BREAK |
	//	blackfriday.EXTENSION_DEFINITION_LISTS |
	//	blackfriday.EXTENSION_HARD_LINE_BREAK

	unsafe := blackfriday.Run([]byte(md))
	//html := bluemonday.UGCPolicy().AllowAttrs("value").OnElements("li").SanitizeBytes(unsafe)
	//html := bluemonday.StrictPolicy().SanitizeBytes(unsafe)

	html := bluemonday.UGCPolicy().AllowAttrs("title").Matching(regexp.MustCompile(`[\p{L}\p{N}\s\-_',:\[\]!\./\\\(\)&]*`)).Globally().SanitizeBytes(unsafe)
	return string(html)
	//renderer := blackfriday.HtmlRenderer(myHTMLFlags, "", "")
	//bytes := blackfriday.MarkdownOptions([]byte(md), renderer, blackfriday.Options{
	//	Extensions: myExtensions,
	//})
	//theHTML := string(bytes)
	//return bluemonday.UGCPolicy().Sanitize(theHTML)
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
