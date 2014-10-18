package libcloud
import (
	"io/ioutil"
	"fmt"
	"github.com/BurntSushi/toml"
)


type PeerInfo struct {
	addr string
	bind_addr string
} 
type EtcdConfig struct {
	addr string
	bind_addr string
	peer PeerInfo
}

func main(){
  raw_data, err := ioutil.ReadFile("input.txt")
  if err != nil {
      panic(err)
  }
  data := string(raw_data)
  var conf EtcdConfig
  if _, err := toml.Decode(data, &conf); err != nil {
  // handle error
  }
  fmt.Printf("%s (%s)\n",conf.addr)
  fmt.Printf("%s (%s)\n",conf.bind_addr)
}

