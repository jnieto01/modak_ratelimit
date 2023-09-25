package i18n

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"path/filepath"
	"os"


	"modak_ratelimit/internal/app/entity"
	"modak_ratelimit/internal/app/utils/logger"
	
	
)



var (	

	supportedLang =map[string]language.Tag {
		"en": language.English,
		"es": language.Spanish,
	}  

	MessageMissing = entity.DataError{ ID: -1, Message: "MessageMissing key is not in the file of language" }
	Minutes = " (Minutes) "
    Seconds = " (Seconds)"

	// general error (0 - 9)
	NotError = entity.DataError{ ID: 0, Message: "" }
	InternalServerError = entity.DataError{ ID: 1, Message: "Internal server error"}  

	// communicacitions error (10-19)
	ErrorMiddlewareQueryParams =  entity.DataError{ ID: 10, Message: "Error in the request parameters" }

	// RateLimited (20-29)
	SuspendedService =  entity.DataError{ ID: 20, Message: "Temporarily suspended service" }
	UnsupportedFlow = entity.DataError{ ID: 21, Message: "Unsupported flow" }
)


func SetLanguage(lang string) error{

	var err error

	// Use default values
	if (lang==""){	
		logger.Info("Use lang=en by defult because it is Empty")
		return nil
	}

	filePath:= getPathofFile() + lang +".json"

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

	Minutes, err  = localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "Minutes",
			Other: "Minutes",
		},
	})
	if err != nil {
		return err
	}

	Seconds, err  = localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "Seconds",
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

	InternalServerError.Message, err  = localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "InternalServerError",
			Other: "Internal server error",
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

	SuspendedService.Message, err  = localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "SuspendedService",
			Other: "Temporarily suspended service",
		},
	})
	if err != nil {
		return err
	}

	UnsupportedFlow.Message, err  = localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "UnsupportedFlow",
			Other: "Temporarily suspended service",
		},
	})
	if err != nil {
		return err
	}




	return err
}



func getPathofFile()(string){

	dir, err := os.Getwd()
	if err != nil {
		logger.Error("Error directory i18n:", err)
		return ""
	}

	fileName := "en.json"
	path := ""

	// Llamada a la función Walk para buscar el archivo
	err = filepath.Walk(dir, func(ruta string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
			
		// Verificar si el nombre del archivo coincide
		if info.Name() == fileName {
			path = ruta
			path = path[ : len(path) - len(fileName)]
		}
		
		
		return nil
	})

	if err != nil {
		logger.Error("Error i18n file Path:", err)
	}

	return path
}