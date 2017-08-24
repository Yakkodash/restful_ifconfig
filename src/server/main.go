package main

import (
    "fmt"
    "log"
    "net/http"
    "net"
    "encoding/json"
)

type Device struct {
  NetDev        net.Interface
  HardwareAddr  net.HardwareAddr
  HardwareAddrStr string
  Flags         net.Flags
  FlagsStr      string
  Uniaddr       []net.Addr
  Muladdr       []net.Addr
}

func main( ) {
  log.Printf("Server started")

  http.HandleFunc("/help", Help)
  http.HandleFunc("/list", List)

  log.Fatal(http.ListenAndServe(":8080", nil))
}

func Help( w http.ResponseWriter, r *http.Request ) {
  fmt.Fprintf( w, "#Description\n" +
                  "\tRESTful API to list network interfaces of the host machine.\n" +
                  "#Synopsis\n" +
                  "\tapi_address/list[?json=true|verbose=true]\n" +
                  "\tapi_address/help\n" )

  log.Println( "Called help, method: " + r.Method )
}

func List( w http.ResponseWriter, r *http.Request ) {
  r.ParseForm( )

  verbose_p := r.Form.Get( "verbose" )
  json_p := r.Form.Get( "json" )

  log.Printf( "Called list, method: %s, json: %s, verbose: %s\n", r.Method, json_p, verbose_p )

  l, err := net.Interfaces( )

  if err != nil {
    panic(err)
  }

  for _, f := range l {
    dev := GetDevice( f )
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

func GetDevice( i net.Interface ) ( Device ) {
  var dev Device
  dev.NetDev = i
  dev.HardwareAddrStr = i.HardwareAddr.String( )
  dev.FlagsStr = i.Flags.String( )
  dev.Uniaddr, _ = i.Addrs( )
  dev.Muladdr, _ = i.MulticastAddrs( )

  return dev
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
    InterfaceString( d.NetDev ), muladdr, uniaddr )
}

func InterfaceString( i net.Interface ) string {

  ret := fmt.Sprintf( "%d. %s\n\tMTU: %d\n\tHardwareAddr: %s\n\tFlags: %s",
  i.Index, i.Name, i.MTU, i.HardwareAddr.String( ), i.Flags.String( ) )

  return ret
}

