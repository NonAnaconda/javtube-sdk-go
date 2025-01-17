package translate

import (
	"time"

	translater "github.com/zijiren233/google-translater"
)

func GoogleFreeTranslate(q, source, target string) (result string, err error) {
	var data translater.Translated
	if data, err = translater.Translate(q, translater.TranslationParams{
		From:       source,
		To:         target,
		Retry:      2,
		RetryDelay: 3 * time.Second,
	}); err != nil {
		result = q
	} else {
		result = data.Text
	}
	return
}
