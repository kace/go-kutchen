package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"time"
)
//https://stackoverflow.com/questions/25819736/why-do-browsers-inefficiently-make-2-requests-here

var addr = flag.String("addr", ":80", "http service address")
//var request string = "https://api.giphy.com/v1/gifs/random?api_key=<API_KEY>&tag=\"chocolate cake\"&rating=R"

//https://fonts.google.com/specimen/Nova+Cut?selection.family=Nova+Cut
var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<link href="https://fonts.googleapis.com/css?family=Nova+Cut" rel="stylesheet">
<body>
<h1 style="text-align: center; font-family: 'Nova Cut'; cursive;">In search of CAKE?</h1>
<iframe src= {{ . }} width="100%" height="500" frameBorder="0">
  <p>Your browser does not support iframes.</p>
</iframe>
</body>
</html>
`))

type Gif struct {
    // TODO GIPHY API random endpoint not returning embed_url
    URL string `json:"embed_url"`
    ID string `json:"id"`
}

type singleResult struct {
    Data *Gif `json:"data"`
}

func giphy_req(r *http.Request)(string) {

    resp, err := http.Get("https://api.giphy.com/v1/gifs/random?api_key=<API_KEY>&tag=\"chocolate+cake\"&rating=G")

    if err != nil {
        // error handler
        //return err
    }

    if resp.StatusCode != 200{
        // return response error
        //return resp.Status
    }

    data, _ := ioutil.ReadAll(resp.Body)
    //fmt.Println(string(data))

    result := singleResult{}
    err = json.Unmarshal(data, &result)
    if err != nil {

    }

    // TODO GIPHY API random endpoint not returning embed_url
    //return result.Data.URL

    // random = ID -> getByID = embed_url
    req := fmt.Sprintf("https://api.giphy.com/v1/gifs/%s?api_key=<API_KEY>", result.Data.ID)
    fmt.Println(time.Now(), "\t", r.RemoteAddr, "\t", req)
    resp, err = http.Get(req)
    data, _ = ioutil.ReadAll(resp.Body)
    result = singleResult{}
    err = json.Unmarshal(data, &result)
    return result.Data.URL
}

func home(w http.ResponseWriter, r *http.Request) {
	homeTemplate.Execute(w, giphy_req(r))
}


func main() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
