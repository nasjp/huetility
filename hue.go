package main

import (
	"fmt"
)

type Hue map[HueLightID]*HueLight

type HueLightID string

type HueLight struct {
	State            *HueState        `json:"state"`
	Swupdate         *HueSwupdate     `json:"swupdate"`
	Type             string           `json:"type"`
	Name             string           `json:"name"`
	Modelid          string           `json:"modelid"`
	Manufacturername string           `json:"manufacturername"`
	Productname      string           `json:"productname"`
	Capabilities     *HueCapabilities `json:"capabilities"`
	Config           *HueConfig       `json:"config"`
	Uniqueid         string           `json:"uniqueid"`
	Swversion        string           `json:"swversion"`
}

type HueState struct {
	Bri       *int   `json:"bri,omitempty"`
	Ct        *int   `json:"ct,omitempty"`
	Alert     string `json:"alert,omitempty"`
	Colormode string `json:"colormode,omitempty"`
	Mode      string `json:"mode,omitempty"`
	Reachable bool   `json:"reachable,omitempty"`
	On        bool   `json:"on"`
}

type HueSwupdate struct {
	State       string `json:"state"`
	Lastinstall string `json:"lastinstall"`
}

type HueCapabilities struct {
	Control   *HueControl   `json:"control"`
	Streaming *HueStreaming `json:"streaming"`
	Certified bool          `json:"certified"`
}

type HueControl struct {
	Mindimlevel int    `json:"mindimlevel"`
	Maxlumen    int    `json:"maxlumen"`
	Ct          *HueCt `json:"ct"`
}

type HueCt struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type HueStreaming struct {
	Renderer bool `json:"renderer"`
	Proxy    bool `json:"proxy"`
}

type HueConfig struct {
	Archetype string      `json:"archetype"`
	Function  string      `json:"function"`
	Direction string      `json:"direction"`
	Startup   *HueStartup `json:"startup"`
}

type HueStartup struct {
	Mode       string `json:"mode"`
	Configured bool   `json:"configured"`
}

var (
	_ fmt.Stringer = (Hue)(nil)
	_ fmt.Stringer = (*HueLight)(nil)
	_ fmt.Stringer = (*HueState)(nil)
)

func (h Hue) String() string {
	var str string
	for id, light := range h {
		str += fmt.Sprintf(`id: %s
%s`, id, light)
	}

	return str
}

func (l *HueLight) String() string {
	return fmt.Sprintf(
		`  name: %s
%s`,
		l.Name,
		l.State,
	)
}

func (s *HueState) String() string {
	return fmt.Sprintf(
		`  state:
    bri       : %d
    ct        : %d
    alert     : "%s"
    colormode : "%s"
    mode      : "%s"
    reachable : %t
    on        : %t`,
		*s.Bri,
		*s.Ct,
		s.Alert,
		s.Colormode,
		s.Mode,
		s.Reachable,
		s.On,
	)
}

func hueEndPoint(cfg *Config) string {
	return fmt.Sprintf("http://%s/api/%s", cfg.IPAddress, cfg.UserName)
}
