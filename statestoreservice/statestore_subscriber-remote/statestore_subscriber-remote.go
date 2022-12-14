// Code generated by Thrift Compiler (0.17.0). DO NOT EDIT.

package main

import (
	"context"
	"flag"
	"fmt"
	thrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/believems/impala-thrift/statestoreservice"
	"github.com/believems/impala-thrift/status"
	"github.com/believems/impala-thrift/types"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

var _ = status.GoUnusedProtection__
var _ = types.GoUnusedProtection__
var _ = statestoreservice.GoUnusedProtection__

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  TUpdateStateResponse UpdateState(TUpdateStateRequest params)")
	fmt.Fprintln(os.Stderr, "  THeartbeatResponse Heartbeat(THeartbeatRequest params)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
	var m map[string]string = h
	return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
	parts := strings.Split(value, ": ")
	if len(parts) != 2 {
		return fmt.Errorf("header should be of format 'Key: Value'")
	}
	h[parts[0]] = parts[1]
	return nil
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	headers := make(httpHeaders)
	var parsedUrl *url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
	flag.Parse()

	if len(urlString) > 0 {
		var err error
		parsedUrl, err = url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	var cfg *thrift.TConfiguration = nil
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
		if len(headers) > 0 {
			httptrans := trans.(*thrift.THttpClient)
			for key, value := range headers {
				httptrans.SetHeader(key, value)
			}
		}
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans = thrift.NewTSocketConf(net.JoinHostPort(host, portStr), cfg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransportConf(trans, cfg)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactoryConf(cfg)
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactoryConf(cfg)
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryConf(cfg)
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client := statestoreservice.NewStatestoreSubscriberClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "UpdateState":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "UpdateState requires 1 args")
			flag.Usage()
		}
		arg41 := flag.Arg(1)
		mbTrans42 := thrift.NewTMemoryBufferLen(len(arg41))
		defer mbTrans42.Close()
		_, err43 := mbTrans42.WriteString(arg41)
		if err43 != nil {
			Usage()
			return
		}
		factory44 := thrift.NewTJSONProtocolFactory()
		jsProt45 := factory44.GetProtocol(mbTrans42)
		argvalue0 := statestoreservice.NewTUpdateStateRequest()
		err46 := argvalue0.Read(context.Background(), jsProt45)
		if err46 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.UpdateState(context.Background(), value0))
		fmt.Print("\n")
		break
	case "Heartbeat":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Heartbeat requires 1 args")
			flag.Usage()
		}
		arg47 := flag.Arg(1)
		mbTrans48 := thrift.NewTMemoryBufferLen(len(arg47))
		defer mbTrans48.Close()
		_, err49 := mbTrans48.WriteString(arg47)
		if err49 != nil {
			Usage()
			return
		}
		factory50 := thrift.NewTJSONProtocolFactory()
		jsProt51 := factory50.GetProtocol(mbTrans48)
		argvalue0 := statestoreservice.NewTHeartbeatRequest()
		err52 := argvalue0.Read(context.Background(), jsProt51)
		if err52 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Heartbeat(context.Background(), value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
