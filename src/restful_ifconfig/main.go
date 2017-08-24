package main

import (
  "fmt"
  "log"
  "net/http"
  "net"
  "encoding/json"
  "io/ioutil"
)

const DESC_PATH string = "./README.md"
const PORT string = ":32432"

type Device struct {
  NetDev          net.Interface
  HardwareAddr    net.HardwareAddr
  HardwareAddrStr string
  Flags           net.Flags
  FlagsStr        string
  Uniaddr         []net.Addr
  Muladdr         []net.Addr
}

func main( ) {
  log.Printf( "Server started" )
  StartServer( )
}

func StartServer( ) {
  http.HandleFunc( "/help", HelpHandler )
  http.HandleFunc( "/list", ListHandler )
  http.HandleFunc( "/", HelpHandler )

  http.ListenAndServe( PORT, nil )
}

func HelpHandler( w http.ResponseWriter, r *http.Request ) {
  log.Println( "Called help, method: " + r.Method )

  desc, err := ioutil.ReadFile( DESC_PATH )

  if err != nil {
    ErrorHandler( w, r, http.StatusInternalServerError, "Failed to read README file" )
    return
  } else {
    w.WriteHeader( http.StatusOK )
    fmt.Fprintf( w, string( desc ) )
  }
}

func ListHandler( w http.ResponseWriter, r *http.Request ) {
  r.ParseForm( )

  verbose_p := r.Form.Get( "verbose" )
  json_p := r.Form.Get( "json" )

  log.Printf( "Called list, method: %s, json: %s, verbose: %s\n", r.Method, json_p, verbose_p )

  l, err := net.Interfaces( )

  if err != nil {
    ErrorHandler( w, r, http.StatusInternalServerError, "Failed to get network interfaces data" )
    return
  } else {
    w.WriteHeader( http.StatusOK )
  }

  for _, f := range l {
    dev, err := GetDevice( f )

    if err != nil {
      ErrorHandler( w, r, http.StatusInternalServerError, "Failed to get network interfaces data" )
      return
    }

    if json_p == "true" {
      if verbose_p == "true" {
        json.NewEncoder( w ).Encode( dev )
      } else {
        json.NewEncoder( w ).Encode( dev.NetDev.Name )
      }
    } else {
      if verbose_p == "true" {
        fmt.Fprintln( w, dev.String( ) )
      } else {
        fmt.Fprintln( w, dev.NetDev.Name )
      }
    }
  }
}

func ErrorHandler( w http.ResponseWriter, r *http.Request, status int, text string ) {
  w.WriteHeader( status )
  fmt.Fprint( w, text )
}

func GetDevice( i net.Interface ) ( Device, error ) {
  var dev Device

  dev.NetDev = i
  dev.HardwareAddrStr = i.HardwareAddr.String( )
  dev.FlagsStr = i.Flags.String( )

  addrs, err := i.Addrs( )

  if err != nil {
    return dev, err
  } else {
    dev.Uniaddr = addrs
  }

  addrs, err = i.MulticastAddrs( )

  if err != nil {
    return dev, err
  } else {
    dev.Muladdr = addrs
  }

  return dev, err
}

// Stringers for Device
func( d Device ) String( ) string {
  var muladdr string
  var uniaddr string

  for _, f := range d.Uniaddr {
    uniaddr += fmt.Sprintf( "\n\t\t%s", f.String( ) )
  }

  for _, f := range d.Muladdr {
    muladdr += fmt.Sprintf( "\n\t\t%s", f.String( ) )
  }

  return fmt.Sprintf( "%s\n\tUnicast addresses: %s\n\tMulticast addresses: %s",
    InterfaceString( d.NetDev ), uniaddr, muladdr )
}

func InterfaceString( i net.Interface ) string {

  ret := fmt.Sprintf( "%d. %s\n\tMTU: %d\n\tHardwareAddr: %s\n\tFlags: %s",
  i.Index, i.Name, i.MTU, i.HardwareAddr.String( ), i.Flags.String( ) )

  return ret
}

