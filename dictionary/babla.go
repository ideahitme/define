package dictionary

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

var (
	babla           = "http://en.bab.la/dictionary/german-english/%s"
	speechPartRegex = regexp.MustCompile(`\{(vb|m|f|n|adj\.|adv\.)\}`)
	speechPartMap   = map[string]string{
		"vb":   "verb",
		"m":    "noun masculine",
		"f":    "noun female",
		"n":    "noun neutral",
		"adj.": "adjective",
		"adv.": "adverb",
	}
	xpathResultList                = "div.quick-results div.quick-result-entry"
	xpathResultListItemWord        = ".quick-result-option a"
	xpathResultListItemSpeechPart  = ".quick-result-option span.suffix"
	xpathResultListItemDefinitions = ".quick-result-overview .sense-group-results li a"

	//colors for pretty print
	cTitle      = color.New(color.FgRed).Add(color.Underline)
	cGeneral    = color.New(color.FgBlue)
	cSpeechpart = color.New(color.FgYellow)
	cDefinition = color.New(color.FgGreen)
)

// Babla implements Dictionary interface
// extracts data from bab.la website
type Babla struct {
}

// Print queries babla endpoint and prints word definitions
func (b *Babla) Print(toDefine string) error {
	url := fmt.Sprintf(babla, toDefine)
	//extract html doc
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}

	i := 0
	doc.Find(xpathResultList).Each(func(_ int, sel *goquery.Selection) {
		wordSel := sel.Find(xpathResultListItemWord).First() //the same as word to be defined
		speechPartSel := sel.Find(xpathResultListItemSpeechPart).First()
		definitionSel := sel.Find(xpathResultListItemDefinitions)
		if isSelectionEmpty(wordSel) || isSelectionEmpty(speechPartSel) || isSelectionEmpty(definitionSel) {
			return
		}
		speechPartMatches := speechPartRegex.FindStringSubmatch(speechPartSel.Text())
		if len(speechPartMatches) < 2 {
			return
		}
		speechPart := speechPartMap[speechPartMatches[1]]
		i++
		cTitle.Printf("Definition: %d\n", i)
		cGeneral.Printf("%s: ", wordSel.Text())
		cSpeechpart.Println(speechPart)

		//extract definitions
		definitionsSlice := []string{}
		sel.Find(xpathResultListItemDefinitions).Each(func(_ int, defSel *goquery.Selection) {
			definitionsSlice = append(definitionsSlice, defSel.Text())
		})
		definitions := strings.Join(definitionsSlice, ",")
		cDefinition.Println(definitions)
		cGeneral.Println("===================")
	})

	return nil
}

func isSelectionEmpty(sel *goquery.Selection) bool {
	return sel.Nodes == nil
}
