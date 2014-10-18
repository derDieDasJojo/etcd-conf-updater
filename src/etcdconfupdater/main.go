package etcdconfupdater

import (
	"os"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jessevdk/go-flags"
	"strings"
	"bufio"
	"os/exec"
)

type peerInfo struct {
	Addr string
	Bind_addr string
} 
type etcdConfig struct {
	Addr string
	Bind_addr string
	Peer peerInfo
}


type Options struct {
// Example of verbosity with level
Verbose bool `short:"v" long:"verbose" description:"Verbose output"`
// Example of optional value
Peers []string `short:"p" long:"peers" description:"IP adresses of other etcds" required:"true"`
Debug bool `short:"d" long:"debug" description:"Debug options"`
// Example of map with multiple default values
//Users map[string]string `long:"users" description:"User e-mail map" default:"system:system@example.org" default:"admin:admin@example.org"`
// Example of option group
//Editor EditorOptions `group:"Editor Options"`
// Example of custom type Marshal/Unmarshal
//Point Point `long:"point" description:"A x,y point" default:"1,2"`
}



func main() {
	//parse command line opts
	var opts Options
	args := os.Args
	args, err := flags.ParseArgs(&opts, args)
	if err != nil {
	    panic(err)
	    os.Exit(1)
	}
	if (opts.Verbose){
		fmt.Print("opts:")
		fmt.Println(opts.Peers)
	}
		
	//read config
	var config etcdConfig
	if _, err := toml.DecodeFile("etcd.conf", &config); err != nil {
		fmt.Println(err)
		return
	}
	//append port
	for id, peer := range opts.Peers{
		//TODO: check if peer is not ending with a port
		opts.Peers[id] = peer + ":7001"
	}
	config.Peer.Addr=strings.Join(opts.Peers,", ")

	//write config
	var filename string;
	if opts.Debug{
		filename="etcd2.conf"
	}else{
		filename="/etc/etcd/etcd.conf"
	}
	
	file, err := os.Create(filename) // For read access.
	if err != nil {
		fmt.Println(err)
		return
	}
	w := bufio.NewWriter(file)
	if err := toml.NewEncoder(w).Encode(config); err != nil {
		fmt.Println(err)
		return
	}
	
	//print some relevant config params
	if (opts.Verbose){
		fmt.Printf("Bind_addr: %s\n",config.Bind_addr)
		fmt.Printf("Addr: %s\n",config.Addr)
		fmt.Printf("Peer.Bind_addr: %s\n",config.Peer.Bind_addr)
		fmt.Printf("Peer.Addr: %s\n",config.Peer.Addr)
	}
	//do something
	cmd := exec.Command("echo","-n","hello world")
	//stdout, err := cmd.StdoutPipe()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	
	if err := cmd.Start(); err != nil{
		fmt.Println(err)
		return
	}
	//out,err := stdout.Read()
	//fmt.Println(stdout)
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		return
	}


	
}
