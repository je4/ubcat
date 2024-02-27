package main

import (
	"context"
	"crypto/tls"
	"flag"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	configdata "github.com/je4/ubcat/v2/config"
	"github.com/je4/ubcat/v2/pkg/config"
	"github.com/je4/ubcat/v2/pkg/index"
	"github.com/je4/utils/v2/pkg/zLogger"
	"github.com/rs/zerolog"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"time"
)

var configfile = flag.String("config", "", "location of toml configuration file")

type LoggingHttpElasticClient struct {
	c      http.Client
	logger zLogger.ZLogger
}

func (l LoggingHttpElasticClient) RoundTrip(request *http.Request) (*http.Response, error) {
	// Log the http request dump
	requestDump, err := httputil.DumpRequest(request, true)
	if err != nil {
		l.logger.Debug().Err(err).Msg("cannot dump request")
	}
	l.logger.Debug().Msgf("reqDump: %s", string(requestDump))
	return l.c.Do(request)
}

func main() {
	flag.Parse()

	var cfgFS fs.FS
	var cfgFile string
	if *configfile != "" {
		cfgFS = os.DirFS(filepath.Dir(*configfile))
		cfgFile = filepath.Base(*configfile)
	} else {
		cfgFS = configdata.ConfigFS
		cfgFile = "test.toml"
	}

	conf := &TestConfig{
		LogLevel: "DEBUG",
		ElasticSearch: config.ElasticSearchConfig{
			Debug: false,
		},
	}

	if err := LoadTestConfig(cfgFS, cfgFile, conf); err != nil {
		log.Fatalf("cannot load toml from [%v] %s: %v", cfgFS, cfgFile, err)
	}

	// create logger instance
	var out io.Writer = os.Stdout
	if conf.LogFile != "" {
		fp, err := os.OpenFile(conf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Panicf("cannot open logfile %s: %v", conf.LogFile, err)
		}
		defer fp.Close()
		out = fp
	}

	output := zerolog.ConsoleWriter{Out: out, TimeFormat: time.RFC3339}
	_logger := zerolog.New(output).With().Timestamp().Logger()
	_logger.Level(zLogger.LogLevel(conf.LogLevel))
	var logger zLogger.ZLogger = &_logger

	elasticConfig := elasticsearch.Config{
		APIKey:    string(conf.ElasticSearch.ApiKey),
		Addresses: conf.ElasticSearch.Endpoint,

		// Retry on 429 TooManyRequests statuses
		//
		RetryOnStatus: []int{502, 503, 504, 429},

		// Retry up to 5 attempts
		//
		MaxRetries: 5,

		Logger: &elastictransport.ColorLogger{Output: os.Stdout},
		//		Transport: doer,
	}
	if conf.ElasticSearch.Debug {
		doer := LoggingHttpElasticClient{
			c: http.Client{
				// Load a trusted CA here, if running in production
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
			},
			logger: logger,
		}
		elasticConfig.Transport = doer
	}
	elastic, err := elasticsearch.NewTypedClient(elasticConfig)
	if err != nil {
		logger.Fatal().Err(err)
	}

	indexClient := index.NewClient("alma-full", elastic)
	var identifiers = []string{"oai:alma.41SLSP_UBS:9954051500105504"}
	docs, err := indexClient.GetDocuments(context.Background(), identifiers...)
	if err != nil {
		logger.Fatal().Err(err).Msgf("cannot get documents %v", identifiers)
	}
	logger.Info().Msgf("docs: %v", docs)
}
