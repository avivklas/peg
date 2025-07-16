# peg
super simple app config library that uses struct-tags for declaration

## learn by example

```go
type AppConfig struct {
    Server Server `peg.name:"server"`
    DB     DB     `peg.name:"db"`
}

type Server struct {
    BindAddr string `peg.name:"bind-addr" peg.usage:"network address to listen to" peg.default:":8080"`
    TLS      string `peg.name:"tls"`
}

type DB struct {
    Host string `peg.name:"host" peg.usage:"hostname of the db"`
    Port int    `peg.name:"port" peg.usage:"port of the db" peg.default:"5432"`
    User string `peg.name:"user" peg.usage:"user of the db" peg.default:"user1"`
    Pass string `peg.name:"pass" peg.usage:"password of the db"`
}

type TLS struct {
    Enabled bool   `peg.name:"enabled" peg.usage:"weather to enable tls on the server" peg.default:"false"`
    Cert    string `peg.name:"cert" peg.usage:"path to the tls cert file"`
    Key     string `peg.name:"key" peg.usage:"path to the tls key file"` 
}

func func main() {
    var conf AppConfig
    peg.Bind(&conf).Read()

    // if app did run with -server.tls.enabled or env var SERVER_TLS_ENABLED=true
    if conf.TLS.Enabled {
        ...
    }
}
```


### tag names:
config fields are defined using the following tags:

`peg.name`  
the name of the config field

`peg.usage`  
usage to print in the help section

`peg.default`  
default value of the field

### sources:
config field values are read from the following sources.
the key of the field is composed of the full path of the field name according to the conventions of the source.

#### flags  
read values from the program arguments.   
in the above example, the flag name for whether to enable tls or not is `-server.tls.enabled`

#### env vars  
read values from environment variables or from `.env` file in the working dir.   
in the above example, the env var name for whether to enable tls or not is `SERVER_TLS_ENABLED`