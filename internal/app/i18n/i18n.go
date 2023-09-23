package i18n

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"path/filepath"
	"strings"


	"modak_ratelimit/internal/app/entity"
	"modak_ratelimit/internal/app/utils/logger"
)



var (	

	basePath = "modak_ratelimit"

	supportedLang =map[string]language.Tag {
		"en": language.English,
		"es": language.Spanish,
	}  

	MessageMissing = entity.DataError{ ID: -1, Message: "MessageMissing key is not in the file of language" }

	// general error (0 - 10)
	NotError = entity.DataError{ ID: 0, Message: "" }

	// communicacitions error (10-20)
	ErrorMiddlewareQueryParams =  entity.DataError{ ID: 10, Message: "Error in the request parameters" }

)


func SetLanguage(lang string) error{

	var err error

	// Use default values
	if (lang==""){	
		logger.Info("Use lang=en by defult because it is Empty")
		return nil
	}


	absPath, err := filepath.Abs("./")
    if err != nil {
		// language not supported (lang="xx")
		// language file does not exist (delete)
		logger.Error("Error path", err)
        return err
    }

	logger.Info("directorio raiz de Actions: " + absPath)


	relPath := "/internal/app/i18n/locales/" + lang +".json"
	filePath := ""
	pos := strings.Index( absPath, basePath)
	if pos != -1 {
		afterBasePath := absPath[pos + len(basePath):]
		newPos := strings.Index( afterBasePath, "i18n")
		if newPos != -1 {
			relPath = "/" + afterBasePath[: newPos + len("i18n")] + "/locales/" + lang +".json"
		}
	  	absPath = absPath[:pos] + basePath
	  	filePath = absPath + relPath
   } else {
		logger.Error("Error absPath: " + absPath , err)
        return nil
   }


   logger.Info("archivo path : " + filePath)

	bundle := i18n.NewBundle(supportedLang[lang])	

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	if _, err = bundle.LoadMessageFile(filePath); err != nil {
		// Use default values
		logger.Error("Error filePath: " + filePath , err)
        return nil
    }

	localizer := i18n.NewLocalizer(bundle, lang)

	MessageMissing.Message, err  = localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "MessageMissing",
			Other: "MessageMissing key is not in the file of language",
		},
	})
	if err != nil {
		return err
	}


	NotError.Message, err  = localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "NotError",
			Other: "",
		},
	})
	if err != nil {
		return err
	}

	ErrorMiddlewareQueryParams.Message, err  = localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "ErrorMiddlewareQueryParams",
			Other: "Error in the request parameters",
		},
	})
	if err != nil {
		return err
	}



	return err
}