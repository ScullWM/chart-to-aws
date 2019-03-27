package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type ScreenConfig struct {
	Httpserver struct {
		Port     string `yaml:"port"`
		Query    string `yaml:"query"`
		Selector string `yaml:"selector"`
		Output   string `yaml:"output"`
	}
	DomainScope string `yaml:"domain"`
	Aws         struct {
		ID     string `yaml:"id"`
		Secret string `yaml:"secret"`
		Token  string `yaml:"token"`
		Bucket string `yaml:"bucket"`
		Region string `yaml:"region"`
	}
}

type ScreenResponse struct {
	Bucket   string `json:"bucket"`
	Filepath string `json:"filepath"`
}

var screenConfig *ScreenConfig

func main() {
	config, err := readConfig()
	screenConfig = config

	if err != nil {
		log.Fatalln(err.Error())
	}

	http.HandleFunc("/", handleRequest)
	log.Printf("Start server: %s \n", ":"+screenConfig.Httpserver.Port)

	log.Printf("query term: %s \n", screenConfig.Httpserver.Query)
	log.Printf("selector term: %s \n", screenConfig.Httpserver.Selector)
	log.Printf("output term: %s \n", screenConfig.Httpserver.Output)

	http.ListenAndServe(screenConfig.Httpserver.Port, nil)
}

func readConfig() (*ScreenConfig, error) {
	config := &ScreenConfig{}

	if err := yamlUnmarshal("./config.yaml", config); err != nil {
		return nil, err
	}
	return config, nil
}

func yamlUnmarshal(path string, out interface{}) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(bytes, out); err != nil {
		return err
	}
	return nil
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get(screenConfig.Httpserver.Query)
	selector := r.URL.Query().Get(screenConfig.Httpserver.Selector)
	output := r.URL.Query().Get(screenConfig.Httpserver.Output)

	err := capture(r.Context(), query, selector, output)
	if err != nil {
		log.Fatal(err)
	}

	err = upload(r.Context(), output)
	if err != nil {
		log.Fatal(err)
	}

	// json resposne
	js, err := json.Marshal(ScreenResponse{Bucket: screenConfig.Aws.Bucket, Filepath: output})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send reponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(js)
}
