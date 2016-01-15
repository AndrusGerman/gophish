package main

/*
gophish - Open-Source Phishing Framework

The MIT License (MIT)

Copyright (c) 2013 Jordan Wright

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gophish/gophish/config"
	"github.com/gophish/gophish/controllers"
	"github.com/gophish/gophish/models"
	"github.com/gorilla/handlers"
)

var Logger = log.New(os.Stdout, " ", log.Ldate|log.Ltime|log.Lshortfile)

func main() {
	// Setup the global variables and settings
	err := models.Setup()
	if err != nil {
		fmt.Println(err)
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	// Start the web servers
	go func() {
		defer wg.Done()
		Logger.Printf("Starting admin server at http://%s\n", config.Conf.AdminURL)
		Logger.Fatal(http.ListenAndServe(config.Conf.AdminURL, handlers.CombinedLoggingHandler(os.Stdout, controllers.CreateAdminRouter())))
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		Logger.Printf("Starting phishing server at http://%s\n", config.Conf.PhishURL)
		Logger.Fatal(http.ListenAndServe(config.Conf.PhishURL, handlers.CombinedLoggingHandler(os.Stdout, controllers.CreatePhishingRouter())))
	}()
	wg.Wait()
}