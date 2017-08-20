package osdetector

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"io/ioutil"
	"launchpad.net/golang-user-agent"
	"log"
	"net/http"
	"regexp"
)

type OSDetectorConfig struct {
	ServeAddress      string
	CassandraHosts    []string
	CassandraKeyspace string
	TemplateFile      string
}

type OSCount struct {
	OS    string
	Count int
}

type OSDetector struct {
	config   OSDetectorConfig
	session  *gocql.Session
	template string
}

func NewOSDetector(config OSDetectorConfig) *OSDetector {
	handler := new(OSDetector)
	handler.config = config

	cluster := gocql.NewCluster(config.CassandraHosts...)
	cluster.Keyspace = config.CassandraKeyspace
	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	handler.session = session

	data, err := ioutil.ReadFile(config.TemplateFile)

	if err != nil {
		log.Fatal(err)
	}

	handler.template = string(data)

	return handler
}

func (od *OSDetector) Handler(w http.ResponseWriter, r *http.Request) {
	uaString := r.Header.Get("User-Agent")
	ua := user_agent.New(uaString)
	osName := ua.OSInfo().Name

	if err := od.session.Query(`UPDATE browsers.browser_counts SET counter = counter + 1 WHERE os = ?`, osName).Exec(); err != nil {
		log.Panic(err)
	}

	var os string
	var counter int
	counts := make([]OSCount, 0)

	iter := od.session.Query(`SELECT os, counter FROM browsers.browser_counts`).Iter()

	for iter.Scan(&os, &counter) {
		counts = append(counts, OSCount{OS: os, Count: counter})
	}
	if err := iter.Close(); err != nil {
		log.Panic(err)
	}

	log.Printf(`user-agent: "%s", os: "%s"`, uaString, osName)

	re := regexp.MustCompile("REPLACE_DATA")
	jsonData, _ := json.Marshal(counts)
	result := re.ReplaceAllString(od.template, string(jsonData))

	w.Write([]byte(result))
}

func (od *OSDetector) Serve() {
	http.HandleFunc("/", od.Handler)

	log.Printf("Starting server at %s", od.config.ServeAddress)
	log.Fatal(http.ListenAndServe(od.config.ServeAddress, nil))
}
