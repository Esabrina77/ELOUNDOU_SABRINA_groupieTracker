package manager

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// structure de sauvegarde du login de chaque  user
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Pseudo   string `json:"pseudo"`
}

var ListUser []LoginUser

// STRUCTURES POUR LES ARTISTS
type ArtistResponse struct {
	Data  []Artist `json:"data"`
	Total int      `json:"total"`
	Next  string   `json:"next"`
}
type Artist struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	NbAlbum       int    `json:"nb_album"`
	NbFan         int    `json:"nb_fan"`
	Radio         bool   `json:"radio"`
	Tracklist     string `json:"tracklist"`
	Type          string `json:"type"`
}

// STRUCTURES POUR LES OPTIONS
type Options struct {
	Streaming         bool   `json:"streaming"`
	StreamingDuration int    `json:"streaming_duration"`
	Offline           bool   `json:"offline"`
	Hq                bool   `json:"hq"`
	AdsDisplay        bool   `json:"ads_display"`
	AdsAudio          bool   `json:"ads_audio"`
	TooManyDevices    bool   `json:"too_many_devices"`
	CanSubscribe      bool   `json:"can_subscribe"`
	RadioSkips        int    `json:"radio_skips"`
	Lossless          bool   `json:"lossless"`
	Preview           bool   `json:"preview"`
	Radio             bool   `json:"radio"`
	Type              string `json:"type"`
}

// STRUCTURES POUR LES RADIOS
type RadioResponse struct {
	Data []Radio `json:"data"`
}
type Radio struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	Tracklist     string `json:"tracklist"`
	Md5Image      string `json:"md5_image"`
	Type          string `json:"type"`
}

// STRUCTURES POUR LES EDITORIAUX
type EditorialResponse struct {
	Data  []Editorial `json:"data"`
	Total int         `json:"total"`
	Next  string      `json:"next"`
}
type Editorial struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	Type          string `json:"type"`
}

type Genres struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	PictureSmall  string `json:"picture_small"`
	PictureMedium string `json:"picture_medium"`
	PictureBig    string `json:"picture_big"`
	PictureXl     string `json:"picture_xl"`
	Type          string `json:"type"`
}

// STRUCTURE DES PLAYLISTS
type Playlist struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Duration      int     `json:"duration"`
	Public        bool    `json:"public"`
	IsLovedTrack  bool    `json:"is_loved_track"`
	Collaborative bool    `json:"collaborative"`
	NbTracks      int     `json:"nb_tracks"`
	Fans          int     `json:"fans"`
	Link          string  `json:"link"`
	Share         string  `json:"share"`
	Picture       string  `json:"picture"`
	PictureSmall  string  `json:"picture_small"`
	PictureMedium string  `json:"picture_medium"`
	PictureBig    string  `json:"picture_big"`
	PictureXl     string  `json:"picture_xl"`
	Checksum      string  `json:"checksum"`
	Tracklist     string  `json:"tracklist"`
	CreationDate  string  `json:"creation_date"`
	Md5Image      string  `json:"md5_image"`
	PictureType   string  `json:"picture_type"`
	Creator       Creator `json:"creator"`
	Type          string  `json:"type"`
	Tracks        Tracks  `json:"tracks"`
}
type Creator struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Tracklist string `json:"tracklist"`
	Type      string `json:"type"`
}
type PlaylistResponse struct {
	ID                    int    `json:"id"`
	Readable              bool   `json:"readable"`
	Title                 string `json:"title"`
	TitleShort            string `json:"title_short"`
	TitleVersion          string `json:"title_version,omitempty"`
	Link                  string `json:"link"`
	Duration              int    `json:"duration"`
	Rank                  int    `json:"rank"`
	ExplicitLyrics        bool   `json:"explicit_lyrics"`
	ExplicitContentLyrics int    `json:"explicit_content_lyrics"`
	ExplicitContentCover  int    `json:"explicit_content_cover"`
	Preview               string `json:"preview"`
	Md5Image              string `json:"md5_image"`
	TimeAdd               int    `json:"time_add"`
	Artist                Artist `json:"artist"`
	Album                 Album  `json:"album"`
	Type                  string `json:"type"`
}
type Tracks struct {
	Data     []Playlist `json:"data"`
	Checksum string     `json:"checksum"`
}

// STRUCTURE DES ALBUMS
type Album struct {
	ID                    int    `json:"id"`
	Title                 string `json:"title"`
	Upc                   string `json:"upc"`
	Link                  string `json:"link"`
	Share                 string `json:"share"`
	Cover                 string `json:"cover"`
	CoverSmall            string `json:"cover_small"`
	CoverMedium           string `json:"cover_medium"`
	CoverBig              string `json:"cover_big"`
	CoverXl               string `json:"cover_xl"`
	Md5Image              string `json:"md5_image"`
	GenreID               int    `json:"genre_id"`
	Genres                Genres `json:"genres"`
	Label                 string `json:"label"`
	NbTracks              int    `json:"nb_tracks"`
	Duration              int    `json:"duration"`
	Fans                  int    `json:"fans"`
	ReleaseDate           string `json:"release_date"`
	RecordType            string `json:"record_type"`
	Available             bool   `json:"available"`
	Tracklist             string `json:"tracklist"`
	ExplicitLyrics        bool   `json:"explicit_lyrics"`
	ExplicitContentLyrics int    `json:"explicit_content_lyrics"`
	ExplicitContentCover  int    `json:"explicit_content_cover"`
	Artist                Artist `json:"artist"`
	Type                  string `json:"type"`
	Tracks                Tracks `json:"tracks"`
}
type DataAlbums struct {
	ID                    int    `json:"id"`
	Readable              bool   `json:"readable"`
	Title                 string `json:"title"`
	TitleShort            string `json:"title_short"`
	TitleVersion          string `json:"title_version"`
	Link                  string `json:"link"`
	Duration              int    `json:"duration"`
	Rank                  int    `json:"rank"`
	ExplicitLyrics        bool   `json:"explicit_lyrics"`
	ExplicitContentLyrics int    `json:"explicit_content_lyrics"`
	ExplicitContentCover  int    `json:"explicit_content_cover"`
	Preview               string `json:"preview"`
	Md5Image              string `json:"md5_image"`
	Artist                Artist `json:"artist"`
	Album                 Album  `json:"album"`
	Type                  string `json:"type"`
}
type TracksAlbums struct {
	Data []DataAlbums `json:"data"`
}

func RetrieveUser() []LoginUser {
	data, err := os.ReadFile("login.json")

	if err != nil {
		fmt.Printf("Erreur lors de la lecture du fichier:%v", err)
		return nil
	}
	var Users []LoginUser
	err = json.Unmarshal(data, &Users)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("list des users : %#v\n", Users)
	return Users
}

// Marquer ( enregistrer) les nouveaux users dans le fichiers De login
func MarkLogin(email string, password string, pseudo string) {
	var newLogin = LoginUser{
		Email:    email,
		Password: password,
		Pseudo:   pseudo,
	}
	users := RetrieveUser()
	users = append(users, newLogin)

	//Convertir lelogin en JSON
	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	//Ecrire les données JSON dans le fichier
	err = os.WriteFile("login.json", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("list des users : %#v\n", users)

}
func PrintColorResult(color string, message string) {
	colorCode := ""
	switch color {
	case "red":
		colorCode = "\033[31m"
	case "green":
		colorCode = "\033[32m"
	case "yellow":
		colorCode = "\033[33m"
	case "blue":
		colorCode = "\033[34m"
	case "purple":
		colorCode = "\033[35m"

	default: //REMETTRE LA COULEUR INITIALE (blanc)
		colorCode = "\033[0m"
	}
	fmt.Printf("%s%s\033[0m", colorCode, message)
}
