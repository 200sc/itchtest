package iserve

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	itchio "github.com/itchio/go-itchio"
)

type Endpoint string

const (
	Me                       Endpoint = "HandleMe"
	MyGames                  Endpoint = "HandleMyGames"
	CredentialsSubkey        Endpoint = "HandleCredentialsSubkey"
	GameUploads              Endpoint = "HandleGameUploads"
	DownloadSession          Endpoint = "HandleDownloadSession"
	BuildFiles               Endpoint = "HandleBuildFiles"
	FinalizeBuildFile        Endpoint = "HandleFinalizeBuildFile"
	BuildFileDownload        Endpoint = "HandleBuildFileDownload"
	CreateBuild              Endpoint = "HandleCreateBuild"
	CreateRediffBuildFailure Endpoint = "HandleCreateRediffBuildFailure"
	WharfStatus              Endpoint = "HandleWharfStatus"
	WharfBuilds              Endpoint = "HandleWharfBuilds"
	WharfChannels            Endpoint = "HandleWharfChannels"
	GetChannel               Endpoint = "HandleGetChannel"
	ListGameUploads          Endpoint = "HandleListGameUploads"
	DownloadUploadBuild      Endpoint = "HandleDownloadUploadBuild"
	FindUpgrade              Endpoint = "HandleFindUpgrade"
	UploadDownload           Endpoint = "HandleUploadDownload"
)

func New() *Server {

	s := &Server{}
	mux := mux.NewRouter()

	routes := map[string]http.Handler{
		"{key}/me":                                         http.HandlerFunc(s.HandleMe),
		"{key}/my-games":                                   http.HandlerFunc(s.HandleMyGames),
		"{key}/credentials/subkey":                         http.HandlerFunc(s.HandleCredentialsSubkey),
		"{key}/game/{game}/uploads":                        http.HandlerFunc(s.HandleGameUploads),
		"{key}/game/{game}/download":                       http.HandlerFunc(s.HandleDownloadSession),
		"{key}/wharf/builds/{build}/files":                 http.HandlerFunc(s.HandleBuildFiles),
		"{key}/wharf/builds/{build}/files/{file}":          http.HandlerFunc(s.HandleFinalizeBuildFile),
		"{key}/wharf/builds/{build}/files/{file}/download": http.HandlerFunc(s.HandleBuildFileDownload),
		"{key}/wharf/builds/{build}/events":                http.HandlerFunc(s.HandleCreateBuild),
		"{key}/wharf/builds/{build}/failures/rediff":       http.HandlerFunc(s.HandleCreateRediffBuildFailure),
		"{key}/wharf/status":                               http.HandlerFunc(s.HandleWharfStatus),
		"{key}/wharf/builds":                               http.HandlerFunc(s.HandleWharfBuilds),
		"{key}/wharf/channels":                             http.HandlerFunc(s.HandleWharfChannels),
		"{key}/wharf/channels/{channel}":                   http.HandlerFunc(s.HandleGetChannel),
		"{key}/download-key/{downloadKey}/uploads":         http.HandlerFunc(s.HandleListGameUploads),
		"{key}/upload/{upload}/download/builds/{build}":    http.HandlerFunc(s.HandleDownloadUploadBuild),
		"{key}/upload/{upload}/upgrade/{build}":            http.HandlerFunc(s.HandleFindUpgrade),
		"{key}/upload/{upload}/download":                   http.HandlerFunc(s.HandleUploadDownload),
	}

	for route, hand := range routes {
		route = "/api/1/" + route
		mux.Handle(route, hand)
	}

	s.Router = mux
	s.testResponses = make(map[Endpoint]interface{})

	return s
}

type Server struct {
	testResponses map[Endpoint]interface{}
	*mux.Router
}

func (s *Server) StoreTestResponse(key Endpoint, val interface{}) {
	s.testResponses[key] = val
}

func (s *Server) WriteResponse(w http.ResponseWriter, resp interface{}) {
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(503)
	}
	w.Write(data)
}

func (s *Server) HandleWharfStatus(w http.ResponseWriter, req *http.Request) {
	resp := itchio.WharfStatusResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleWharfBuilds(w http.ResponseWriter, req *http.Request) {
	resp := itchio.NewBuildResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleWharfChannels(w http.ResponseWriter, req *http.Request) {
	resp := itchio.ListChannelsResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleMe(w http.ResponseWriter, req *http.Request) {
	resp := itchio.GetMeResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleMyGames(w http.ResponseWriter, req *http.Request) {
	resp := itchio.ListMyGamesResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleCredentialsSubkey(w http.ResponseWriter, req *http.Request) {
	resp := itchio.SubkeyResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleGameUploads(w http.ResponseWriter, req *http.Request) {
	resp := itchio.GameUploadsResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleUploadDownload(w http.ResponseWriter, req *http.Request) {
	resp := itchio.UploadDownloadResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleGetChannel(w http.ResponseWriter, req *http.Request) {
	resp := itchio.GetChannelResponse{}
	vars := mux.Vars(req)
	key := vars["key"]
	if key == "testkey" {
		fmt.Println("Got test key")
		if v, ok := s.testResponses[GetChannel]; ok {
			s.WriteResponse(w, v)
			return
			// Can't check this because butler uses a vendored dependency here
			// so itchio.GetChannelResponse != itchio.GetChannelResponse
			// //resp, ok = v.(itchio.GetChannelResponse)
			// if !ok {
			// 	fmt.Println("Stored wrong type for GetChannel")
			// 	fmt.Println(v, reflect.TypeOf(v))
			// } else {
			// 	fmt.Println("Retrieved test value for GetChannel")
			// }
		}
	}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleBuildFiles(w http.ResponseWriter, req *http.Request) {
	resp := itchio.ListBuildFilesResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleFinalizeBuildFile(w http.ResponseWriter, req *http.Request) {
	resp := itchio.FinalizeBuildFileResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleBuildFileDownload(w http.ResponseWriter, req *http.Request) {
	resp := itchio.DownloadBuildFileResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleDownloadUploadBuild(w http.ResponseWriter, req *http.Request) {
	resp := itchio.DownloadUploadBuildResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleCreateBuild(w http.ResponseWriter, req *http.Request) {
	resp := itchio.CreateBuildEventResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleCreateRediffBuildFailure(w http.ResponseWriter, req *http.Request) {
	resp := itchio.CreateBuildFailureResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleListGameUploads(w http.ResponseWriter, req *http.Request) {
	resp := itchio.ListGameUploadsResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleFindUpgrade(w http.ResponseWriter, req *http.Request) {
	resp := itchio.FindUpgradeResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
func (s *Server) HandleDownloadSession(w http.ResponseWriter, req *http.Request) {
	resp := itchio.NewDownloadSessionResponse{}
	// Todo: populate response
	s.WriteResponse(w, resp)
}
