package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itchio/go-itchio"
)

var (
	port = flag.String("port", ":80", "")
)

// Todo: differ responses based on POST or GET
// Todo: handle arguments following question mark

var (
	routes = map[string]http.Handler{
		"{key}/me":                                         http.HandlerFunc(HandleMe),
		"{key}/my-games":                                   http.HandlerFunc(HandleMyGames),
		"{key}/credentials/subkey":                         http.HandlerFunc(HandleCredentialsSubkey),
		"{key}/game/{game}/uploads":                        http.HandlerFunc(HandleGameUploads),
		"{key}/game/{game}/download":                       http.HandlerFunc(HandleDownloadSession),
		"{key}/wharf/builds/{build}/files":                 http.HandlerFunc(HandleBuildFiles),
		"{key}/wharf/builds/{build}/files/{file}":          http.HandlerFunc(HandleFinalizeBuildFile),
		"{key}/wharf/builds/{build}/files/{file}/download": http.HandlerFunc(HandleBuildFileDownload),
		"{key}/wharf/builds/{build}/events":                http.HandlerFunc(HandleCreateBuild),
		"{key}/wharf/builds/{build}/failures/rediff":       http.HandlerFunc(HandleCreateRediffBuildFailure),
		"{key}/wharf/status":                               http.HandlerFunc(HandleWharfStatus),
		"{key}/wharf/builds":                               http.HandlerFunc(HandleWharfBuilds),
		"{key}/wharf/channels":                             http.HandlerFunc(HandleWharfChannels),
		"{key}/wharf/channels/{channel}":                   http.HandlerFunc(HandleGetChannel),
		"{key}/download-key/{downloadKey}/uploads":         http.HandlerFunc(HandleListGameUploads),
		"{key}/upload/{upload}/download/builds/{build}":    http.HandlerFunc(HandleDownloadUploadBuild),
		"{key}/upload/{upload}/upgrade/{build}":            http.HandlerFunc(HandleFindUpgrade),
		"{key}/upload/{upload}/download":                   http.HandlerFunc(HandleUploadDownload),
	}
)

func WriteResponse(w http.ResponseWriter, resp interface{}) {
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(503)
	}
	w.Write(data)
}

func HandleWharfStatus(w http.ResponseWriter, req *http.Request) {
	resp := itchio.WharfStatusResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleWharfBuilds(w http.ResponseWriter, req *http.Request) {
	resp := itchio.NewBuildResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleWharfChannels(w http.ResponseWriter, req *http.Request) {
	resp := itchio.ListChannelsResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleMe(w http.ResponseWriter, req *http.Request) {
	resp := itchio.GetMeResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleMyGames(w http.ResponseWriter, req *http.Request) {
	resp := itchio.ListMyGamesResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleCredentialsSubkey(w http.ResponseWriter, req *http.Request) {
	resp := itchio.SubkeyResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleGameUploads(w http.ResponseWriter, req *http.Request) {
	resp := itchio.GameUploadsResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleUploadDownload(w http.ResponseWriter, req *http.Request) {
	resp := itchio.UploadDownloadResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleGetChannel(w http.ResponseWriter, req *http.Request) {
	resp := itchio.GetChannelResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleBuildFiles(w http.ResponseWriter, req *http.Request) {
	resp := itchio.ListBuildFilesResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleFinalizeBuildFile(w http.ResponseWriter, req *http.Request) {
	resp := itchio.FinalizeBuildFileResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleBuildFileDownload(w http.ResponseWriter, req *http.Request) {
	resp := itchio.DownloadBuildFileResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleDownloadUploadBuild(w http.ResponseWriter, req *http.Request) {
	resp := itchio.DownloadUploadBuildResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleCreateBuild(w http.ResponseWriter, req *http.Request) {
	resp := itchio.CreateBuildEventResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleCreateRediffBuildFailure(w http.ResponseWriter, req *http.Request) {
	resp := itchio.CreateBuildFailureResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleListGameUploads(w http.ResponseWriter, req *http.Request) {
	resp := itchio.ListGameUploadsResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleFindUpgrade(w http.ResponseWriter, req *http.Request) {
	resp := itchio.FindUpgradeResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}
func HandleDownloadSession(w http.ResponseWriter, req *http.Request) {
	resp := itchio.NewDownloadSessionResponse{}
	// Todo: populate response
	WriteResponse(w, resp)
}

func main() {
	mux := mux.NewRouter()

	for route, hand := range routes {
		route = "/api/1/" + route
		mux.Handle(route, hand)
	}

	err := http.ListenAndServe(*port, mux)
	if err != nil {
		fmt.Println(err)
	}
}
