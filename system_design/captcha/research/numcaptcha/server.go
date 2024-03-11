package numcaptcha

import (
	"bytes"
	"net/http"
	"path"
)

type captchaHandler struct {
	imgWidth  int
	imgHeight int
}

// Server returns a handler that serves HTTP requests with image representations of captchas.
// file name part must contain a captcha id, file extension - its format PNG or WAV
//
// For example, for file name "LBm5vMjHDtdUfaWYXiQX.png" it serves an image captcha
// with id "LBm5vMjHDtdUfaWYXiQX"
//
// To serve a captcha as a download file, the URL must be constructed in such a way as
// if the file to serve is in the "download" subdirectory:
// "/download/LBm5vMjHDtdUfaWYXiQX.wav"
//
// To reload captcha (get a different solution for the same captcha id), append
// "?reload=x" to URL, where x may be anything (for example, current type or a random
// number to make browsers refetch an image instead of loading it from cache).
func Server(imgWidth, imgHeight int) http.Handler {
	return &captchaHandler{
		imgWidth:  imgWidth,
		imgHeight: imgHeight,
	}
}

func (h *captchaHandler) serve(w http.ResponseWriter, r *http.Request, id, ext string, download bool) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = content
		/// WriteImage()
	}

	return nil
}

func (h *captchaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	if r.FormValue("reload") != "" {
		// Reload(id)
	}
	download := path.Base(dir) == "download"
	if h.serve(w, r, id, ext, download) == ErrNotFound {
		http.NotFound(w, r)
	}
}
