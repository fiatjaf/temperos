package main

import (
	"crypto/tls"
	"github.com/go-martini/martini"
	"github.com/hoisie/mustache"
	"github.com/martini-contrib/cors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	m := martini.Classic()

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Origin", "Accept", "X-Requested-With"},
	}))

	// custom http insecure client
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	insecureClient := &http.Client{Transport: tr}

	m.Get("/", func(res http.ResponseWriter, req *http.Request) {
		http.Redirect(res, req, "https://github.com/fiatjaf/temperos/blob/master/README.md", 302)
	})

	m.Get("/**", func(params martini.Params, res http.ResponseWriter, req *http.Request) {
		qs := req.URL.Query()
		log.Print(params["_1"])
		log.Print(qs)

		// parse url
		target := strings.ToLower(params["_1"])
		if strings.HasPrefix(target, "http://") == false &&
			strings.HasPrefix(target, "https://") == false {
			target = "http://" + target
		}
		t, err := url.Parse(target)
		if err != nil {
			log.Print(err)
			http.Error(res, "Oops!", http.StatusNotFound)
			return
		}

		// important variables
		target = t.String()
		domain := t.Host + t.Path
		// remove ending slash from domain
		if strings.HasSuffix(domain, "/") {
			domain = domain[:len(domain)-1]
		}

		// get response from target url
		response, err := insecureClient.Get(target)
		if err != nil {
			log.Print(err)
			http.Error(res, "Oops!", http.StatusNotFound)
			return
		}

		// the content-type of the response is very important
		res.Header().Set("Content-Type", response.Header.Get("Content-Type"))

		// the params must be strings, not lists (as they are right now)
		context := make(map[string]string)
		for k, v := range qs {
			context[k] = v[0]
		}

		// render the template with this context
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		template := string(contents)
		rendered := mustache.Render(template, context)
		res.Write([]byte(rendered))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	m.RunOnAddr("0.0.0.0:" + port)
}
