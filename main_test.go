package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReProcess(t *testing.T) {
	result := reProcess(`<span id="c-oneonones" level="1" class=" c-oneonones competency"><a title="go to competency github page" class="github-link" target="_blank" href="https://github.com/SearchSpring/competencies/blob/master/competencies/one-on-ones.md"><i class="fab fa-github"></i></a> one on ones <a title="add this competency to the google sheet for tracking" style="display:none" class="drive-link" href="javascript:;"><i class="fab fa-google-drive"></i></a></span><span id="c-hiring" level="1" class=" c-hiring competency"><a title="go to competency github page" class="github-link" target="_blank" href="https://github.com/SearchSpring/competencies/blob/master/competencies/hiring.md"><i class="fab fa-github"></i></a> hiring <a title="add this competency to the google sheet for tracking" style="display:none" class="drive-link" href="javascript:;"><i class="fab fa-google-drive"></i></a></span><span id="c-interviewing" level="1" class=" c-interviewing competency"><a title="go to competency github page" class="github-link" target="_blank" href="https://github.com/SearchSpring/competencies/blob/master/competencies/interviewing.md"><i class="fab fa-github"></i></a> interviewing <a title="add this competency to the google sheet for tracking" style="display:none" class="drive-link" href="javascript:;"><i class="fab fa-google-drive"></i></a></span><span id="c-leading" level="1" class=" c-leading competency"><a title="go to competency github page" class="github-link" target="_blank" href="https://github.com/SearchSpring/competencies/blob/master/competencies/leading.md"><i class="fab fa-github"></i></a> leading <a title="add this competency to the google sheet for tracking" style="display:none" class="drive-link" href="javascript:;"><i class="fab fa-google-drive"></i></a></span><span id="c-careerdevelopment" level="1" class=" c-careerdevelopment competency"><a title="go to competency github page" class="github-link" target="_blank" href="https://github.com/SearchSpring/competencies/blob/master/competencies/career-development.md"><i class="fab fa-github"></i></a> career development <a title="add this competency to the google sheet for tracking" style="display:none" class="drive-link" href="javascript:;"><i class="fab fa-google-drive"></i></a></span><span id="c-visionandstrategy" level="1" class="missing c-visionandstrategy competency"><a title="go to competency github page" class="github-link" target="_blank" href="https://github.com/SearchSpring/competencies/new/master/competencies"><i class="fab fa-github"></i></a> vision and strategy <a title="add this competency to the google sheet for tracking" style="display:none" class="drive-link" href="javascript:;"><i class="fab fa-google-drive"></i></a></span><span id="c-deployments" level="2" class=" c-deployments competency"><a title="go to competency`)
	require.Equal(t, "", result)
}

func TestProcessInherited(t *testing.T) {
	result, err := processInherits("#original \n##Skills\n<inherit doc=\"../test1.md\"/>\n   a  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(result)
	require.Equal(t, "#original \n##Skills\n#### <a href=\"../test1.html\">test 1</a>\n<skills>\nbreakdancing\n</skills>\n\n#### <a href=\"../test2.html\">test 2</a>\n<skills>\nfigure skating\n</skills>\n\n#### <a href=\"../test3.html\">test 3</a>\n<skills>\nkung fu\n</skills>\n   a", result)
}

func TestProcessInheritedAndMarkdownification(t *testing.T) {
	html, title, err := processHTML("test1.md")
	if err != nil {
		panic(err)
	}
	require.Equal(t, "test 1", title)
	require.Equal(t, "<h1>test 1</h1>\n\n<h2>Test document with sub skills</h2>\n\n<div class=\"skill-group\"><a class=\"bad\" href=\"https://github.com/searchspring/competencies/blob/master/competencies/breakdancing.md\">breakdancing</a></div>\n\n<h2>comment that shouldn&rsquo;t appear</h2>\n\n<h4><a href=\"../test2.html\">test 2</a></h4>\n\n<div class=\"skill-group\"><a class=\"bad\" href=\"https://github.com/searchspring/competencies/blob/master/competencies/figure-skating.md\">figure skating</a></div>\n\n<h4><a href=\"../test3.html\">test 3</a></h4>\n\n<div class=\"skill-group\"><a class=\"bad\" href=\"https://github.com/searchspring/competencies/blob/master/competencies/kung-fu.md\">kung fu</a></div>\n\n<p>a</p>\n", string(html))
}
