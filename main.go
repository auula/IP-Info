package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"strconv"

	//"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"io/ioutil"
	"net/http"
)

func main() {
	app := app.New()
	ip := widget.NewLabel("")
	isp := widget.NewLabel("")
	inputText := widget.NewEntry()
	location := widget.NewLabel("")
	inputText.SetPlaceHolder("请输入内容")
	//app.Settings().SetTheme(theme.LightTheme())
	w := app.NewWindow("IP归属地查询工具v1.0")
	w.Resize(fyne.NewSize(250,300))
	w.SetContent(widget.NewVBox(
		widget.NewForm(&widget.FormItem{Text: "IP地址:", Widget: ip},
			&widget.FormItem{Text: "地理位置:", Widget: location},
			&widget.FormItem{Text: "运营商名称:", Widget: isp}),
		inputText,
		widget.NewButton("查 询", func() {
			info := GetIpInfo(inputText.Text)
			var da  Artist
			json.Unmarshal([]byte(info),&da)
			fmt.Println(info)
			ip.SetText(inputText.Text)
			isp.SetText(da.Data.ISP)
			location.SetText(da.Data.Pos+" 经纬度:"+strconv.FormatFloat(da.Data.Location.Lat, 'E', -1, 32)+","+strconv.FormatInt(da.Data.Location.Lng,10))
		}),
	))
	w.ShowAndRun()
}
func GetIpInfo(ip string) string {
	if len(ip) == 0 {
		return ""
	}

	url := fmt.Sprintf("http://v1.alapi.cn/api/ip?ip=%s&format=json", ip)

	resp, err :=   http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	return string(body)
}
func UnmarshalArtist(data []byte) (Artist, error) {
	var r Artist
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Artist) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAlbum(data []byte) (Album, error) {
	var r Album
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Album) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTrack(data []byte) (Track, error) {
	var r Track
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Track) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Artist struct {
	Code    *int64   `json:"code,omitempty"`
	Msg     *string  `json:"msg,omitempty"`
	Data    *Data    `json:"data,omitempty"`
	Author  *Author  `json:"author,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Founded *int64   `json:"founded,omitempty"`
	Members []string `json:"members"`
}

type Author struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type Data struct {
	Beginip   string   `json:"beginip"`
	Endip     string   `json:"endip"`
	Pos       string   `json:"pos"`
	ISP       string   `json:"isp"`
	Location  Location `json:"location"`
	Rectangle string   `json:"rectangle"`
	AdInfo    AdInfo   `json:"ad_info"`
	IP        string   `json:"ip"`
}

type AdInfo struct {
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Adcode   int64  `json:"adcode"`
}

type Location struct {
	Lat float64 `json:"lat"`
	Lng int64   `json:"lng"`
}

type Album struct {
	Name   string      `json:"name"`
	Artist ArtistClass `json:"artist"`
	Tracks []Track     `json:"tracks"`
}

type ArtistClass struct {
	Name    string   `json:"name"`
	Founded int64    `json:"founded"`
	Members []string `json:"members"`
}

type Track struct {
	Name     string `json:"name"`
	Duration int64  `json:"duration"`
}