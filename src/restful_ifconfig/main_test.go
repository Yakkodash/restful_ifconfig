package main

import (
  "testing"
  "net/http"
  "io/ioutil"
  "regexp"
  "encoding/json"
)

func isJSON( s string ) bool {
  var js json.RawMessage
  return json.Unmarshal([]byte(s), &js) == nil
}

func TestHelpHandler( t *testing.T ) {

  go StartServer( )

  resp, err := http.Get( "http://localhost" + PORT + "/help" )
  defer resp.Body.Close( )

  if err != nil {
    t.Error( "Failed GET /help: ", err )
  }

  if resp.StatusCode != http.StatusOK {
    t.Error( "Failed GET /help: response status not OK" )
  }
}

func TestListHandler( t *testing.T ) {

  resp, err := http.Get( "http://localhost" + PORT + "/list" )
  defer resp.Body.Close( )

  body, err := ioutil.ReadAll( resp.Body )

  match, _ := regexp.MatchString( "^lo*", string( body ) ) // Because there is most likely a loopback

  if err != nil {
    t.Error( "Failed GET /list: ", err )
  }

  if resp.StatusCode != http.StatusOK {
    t.Error( "Failed GET /list: response status not OK" )
  }

  if match != true {
    t.Error( "Failed GET /list: no loopback interface in response" )
  }
}

func TestListHandlerVerbose( t *testing.T ) {

  resp, err := http.Get( "http://localhost" + PORT + "/list?verbose=true" )
  defer resp.Body.Close( )

  body, err := ioutil.ReadAll( resp.Body )

  matchlo, _ := regexp.MatchString( ".*lo*", string( body ) ) // Check that there is a set field
  matchmtu, _ := regexp.MatchString( ".*MTU*", string( body ) )

  if err != nil {
    t.Error( "Failed GET /list: ", err )
  }

  if resp.StatusCode != http.StatusOK {
    t.Error( "Failed GET /list: response status not OK" )
  }

  if matchlo != true {
    t.Error( "Failed GET /list: no loopback interface in response" )
  }

  if matchmtu != true {
    t.Error( "Failed GET /list: no MTU field in response" )
  }
}

func TestListHandlerVerboseJson( t *testing.T ) {
  resp, err := http.Get("http://localhost" + PORT + "/list?verbose=true&json=true")
  defer resp.Body.Close( )

  body, err := ioutil.ReadAll( resp.Body )

  matchlo, _ := regexp.MatchString( ".*lo*", string( body ) ) // Check that there is a set field
  matchmtu, _ := regexp.MatchString( ".*MTU*", string( body ) )

  if err != nil {
    t.Error( "Failed GET /list: ", err )
  }

  if resp.StatusCode != http.StatusOK {
    t.Error( "Failed GET /list: response status not OK" )
  }

  if matchlo != true {
    t.Error( "Failed GET /list: no loopback interface in response" )
  }

  if matchmtu != true {
    t.Error( "Failed GET /list: no MTU field in response" )
  }

}

