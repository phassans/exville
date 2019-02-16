package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/phassans/exville/engines"

	"github.com/phassans/banana/helper"
	"github.com/phassans/banana/shared"
)

type (
	hresp struct {
		Message string    `json:"message,omitempty"`
		Error   *APIError `json:"error,omitempty"`
	}
)

func (rtr *router) newImageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := shared.GetLogger()

		err := r.ParseMultipartForm(10000000)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(hresp{Error: NewAPIError(err)})
			return
		}

		m := r.MultipartForm
		images := m.File["images"]
		userID := r.FormValue("userId")

		logger = logger.With().Str("endpoint", "/uploadimage").Logger()
		logger.Info().Msgf("upload image request")

		if err := Validate(images); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(hresp{Error: NewAPIError(err)})
			return
		}

		for i, _ := range images {
			//for each fileheader, get a handle to the actual file
			file, err := images[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//create destination file making sure the path is writeable.
			dst, err := os.Create("upload_images/" + images[i].Filename)
			defer dst.Close()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				err = json.NewEncoder(w).Encode(hresp{Error: NewAPIError(err)})
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				err = json.NewEncoder(w).Encode(hresp{Error: NewAPIError(err)})
				return
			}
		}

		var uid int64
		if uid, err = strconv.ParseInt(userID, 10, 64); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(hresp{Error: NewAPIError(err)})
			return
		}

		err = rtr.engines.UpdateUserWithImage(engines.UserID(uid), engines.ImageName("upload_images/"+images[0].Filename))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode(hresp{Error: NewAPIError(err)})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(hresp{Message: fmt.Sprintf("success uploaded image!")})
		return
	}
}

func Validate(images []*multipart.FileHeader) error {
	if len(images) == 0 {
		return helper.ValidationError{Message: fmt.Sprint("submit happy hour failed, missing images!")}
	}
	return nil
}
