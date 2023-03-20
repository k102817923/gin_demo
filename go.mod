module gin_demo

go 1.20

require (
	github.com/astaxie/beego v1.12.3
	github.com/gin-gonic/gin v1.9.0
	github.com/go-ini/ini v1.67.0
	github.com/jinzhu/gorm v1.9.16
	github.com/unknwon/com v1.0.1
)

require (
	github.com/bytedance/sonic v1.8.5 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/cilium/ebpf v0.10.0 // indirect
	github.com/cosiner/argv v0.1.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/derekparker/trie v0.0.0-20221221181808-1424fce0c981 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-delve/delve v1.20.1 // indirect
	github.com/go-delve/liner v1.2.3-0.20220127212407-d32d89dd2a5d // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.11.2 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/goccy/go-json v0.10.1 // indirect
	github.com/google/go-dap v0.7.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/leodido/go-urn v1.2.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.7 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/cobra v1.6.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	go.starlark.net v0.0.0-20230302034142-4b1e35fe2254 // indirect
	golang.org/x/arch v0.3.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/EDDYCJY/go-gin-example/conf => /pkg/conf
	github.com/EDDYCJY/go-gin-example/middleware => ./middleware
	github.com/EDDYCJY/go-gin-example/models => /models
	github.com/EDDYCJY/go-gin-example/pkg/e => /pkg/e
	github.com/EDDYCJY/go-gin-example/pkg/setting => /pkg/setting
	github.com/EDDYCJY/go-gin-example/pkg/util => /pkg/util
	github.com/EDDYCJY/go-gin-example/routers => /routers
	github.com/EDDYCJY/go-gin-example/routers/api => /routers/api
	github.com/EDDYCJY/go-gin-example/routers/api/v1/ => /routers/api/v1
)
