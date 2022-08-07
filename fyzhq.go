package fyzhq

import (
	"github.com/mozillazg/go-pinyin"
	"github.com/tidwall/gjson"
)

func Transform(text string, lan string) (result string) {
	p := pinyin.Pinyin(text, pinyin.NewArgs())
	if len(p) == 0 {
		return text
	}
	switch lan {
	case "jp":
		for _, v := range p {
			result += gjson.Get(worddictStr, "Japanese."+v[0]).String()
		}
	case "en":
		for _, v := range p {
			result += gjson.Get(worddictStr, "English."+v[0]).String()
		}
	case "fr":
		for _, v := range p {
			result += gjson.Get(worddictStr, "French."+v[0]).String()
		}
	case "gm":
		for _, v := range p {
			result += gjson.Get(worddictStr, "German."+v[0]).String()
		}
	case "ru":
		for _, v := range p {
			result += gjson.Get(worddictStr, "Russian."+v[0]).String()
		}
	case "kr":
		for _, v := range p {
			result += gjson.Get(worddictStr, "Korean."+v[0]).String()
		}
	case "th":
		for _, v := range p {
			result += gjson.Get(worddictStr, "Thai."+v[0]).String()
		}
	default:
		result = text
	}
	return
}
