package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type music struct {
	Data []struct {
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
		Artist                struct {
			ID            int    `json:"id"`
			Name          string `json:"name"`
			Link          string `json:"link"`
			Picture       string `json:"picture"`
			PictureSmall  string `json:"picture_small"`
			PictureMedium string `json:"picture_medium"`
			PictureBig    string `json:"picture_big"`
			PictureXl     string `json:"picture_xl"`
			Tracklist     string `json:"tracklist"`
			Type          string `json:"type"`
		} `json:"artist"`
		Album struct {
			ID          int    `json:"id"`
			Title       string `json:"title"`
			Cover       string `json:"cover"`
			CoverSmall  string `json:"cover_small"`
			CoverMedium string `json:"cover_medium"`
			CoverBig    string `json:"cover_big"`
			CoverXl     string `json:"cover_xl"`
			Tracklist   string `json:"tracklist"`
			Type        string `json:"type"`
		} `json:"album"`
		Type string `json:"type"`
	} `json:"data"`
	Total int    `json:"total"`
	Next  string `json:"next"`
}

type news struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Source struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"source"`
		Author      string    `json:"author"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		URL         string    `json:"url"`
		URLToImage  string    `json:"urlToImage"`
		PublishedAt time.Time `json:"publishedAt"`
		Content     string    `json:"content"`
	} `json:"articles"`
}
type chuck struct {
	Categories []interface{} `json:"categories"`
	CreatedAt  string        `json:"created_at"`
	IconURL    string        `json:"icon_url"`
	ID         string        `json:"id"`
	UpdatedAt  string        `json:"updated_at"`
	URL        string        `json:"url"`
	Value      string        `json:"value"`
}
type unsplash struct {
	ID             string      `json:"id"`
	CreatedAt      string      `json:"created_at"`
	UpdatedAt      string      `json:"updated_at"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	Color          string      `json:"color"`
	Description    interface{} `json:"description"`
	AltDescription interface{} `json:"alt_description"`
	Urls           struct {
		Raw     string `json:"raw"`
		Full    string `json:"full"`
		Regular string `json:"regular"`
		Small   string `json:"small"`
		Thumb   string `json:"thumb"`
	} `json:"urls"`
	Links struct {
		Self             string `json:"self"`
		HTML             string `json:"html"`
		Download         string `json:"download"`
		DownloadLocation string `json:"download_location"`
	} `json:"links"`
	Categories             []interface{} `json:"categories"`
	Sponsored              bool          `json:"sponsored"`
	SponsoredBy            interface{}   `json:"sponsored_by"`
	SponsoredImpressionsID interface{}   `json:"sponsored_impressions_id"`
	Likes                  int           `json:"likes"`
	LikedByUser            bool          `json:"liked_by_user"`
	CurrentUserCollections []interface{} `json:"current_user_collections"`
	User                   struct {
		ID              string      `json:"id"`
		UpdatedAt       string      `json:"updated_at"`
		Username        string      `json:"username"`
		Name            string      `json:"name"`
		FirstName       string      `json:"first_name"`
		LastName        string      `json:"last_name"`
		TwitterUsername interface{} `json:"twitter_username"`
		PortfolioURL    interface{} `json:"portfolio_url"`
		Bio             string      `json:"bio"`
		Location        string      `json:"location"`
		Links           struct {
			Self      string `json:"self"`
			HTML      string `json:"html"`
			Photos    string `json:"photos"`
			Likes     string `json:"likes"`
			Portfolio string `json:"portfolio"`
			Following string `json:"following"`
			Followers string `json:"followers"`
		} `json:"links"`
		ProfileImage struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
			Large  string `json:"large"`
		} `json:"profile_image"`
		InstagramUsername string `json:"instagram_username"`
		TotalCollections  int    `json:"total_collections"`
		TotalLikes        int    `json:"total_likes"`
		TotalPhotos       int    `json:"total_photos"`
		AcceptedTos       bool   `json:"accepted_tos"`
	} `json:"user"`
	Exif struct {
		Make         string `json:"make"`
		Model        string `json:"model"`
		ExposureTime string `json:"exposure_time"`
		Aperture     string `json:"aperture"`
		FocalLength  string `json:"focal_length"`
		Iso          int    `json:"iso"`
	} `json:"exif"`
	Views     int `json:"views"`
	Downloads int `json:"downloads"`
}
type kek struct {
	Ok     bool `json:"ok"`
	Result struct {
		User struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"user"`
		Status string `json:"status"`
	} `json:"result"`
}

var bh [100]string

func wordCheck(text string, userID int, chatID int64) {
	for j := 0; j < len(bh); j++ {
		if strings.Contains(text, bh[j]) {
			kick(userID, chatID)
		}
	}
}

func main() {
	var (
		unsplashResponse = "https://api.unsplash.com/photos/random?client_id=1435c8eaadfbeacd502ec854e73123059456f3a601722e790c009bd40fdfe15b"
		chuckResponse    = "https://api.chucknorris.io/jokes/random"
		newsResponse     = "https://newsapi.org/v2/top-headlines?country=ru&apiKey=4ae2630c606c46bb99756be01d9bb174"
	)
	var (
		text string
	)
	var (
		random    int
		counter   int
		strrandom string
	)
	var (
		replyID int
	)
	var (
		chatID    int64
		strchatID string
	)
	var (
		userID    int
		struserID string
	)
	var (
		stableID    int
		strstableID string
	)
	var token = "858109721:AAFLuKz-S0XA-6Yv5IhMOa6jbePlCGbyhYE"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Connection complete %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if !update.Message.IsCommand() {
			if update.Message.Text != "" {
				userID = update.Message.From.ID
				chatID = update.Message.Chat.ID
				go wordCheck(update.Message.Text, userID, chatID)
			}

		}
		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "suicide":
				chatID = update.Message.Chat.ID
				userID = update.Message.From.ID
				struserID = strconv.Itoa(userID)
				strchatID = strconv.FormatInt(chatID, 10)
				http.Get("https://api.telegram.org/bot" + token + "/kickChatMember?chat_id=" + strchatID + "&user_id=" + struserID)
			case "pyhton":
				msg.Text = "I hate this"
			case "help":
				msg.Text = "try /ban and /f"
			case "ban":
				if update.Message.ReplyToMessage != nil {
					chatID = update.Message.Chat.ID
					userID = update.Message.From.ID
					if checkAdmin(&strchatID, &struserID, &chatID, &userID, &replyID) {
						kick(userID, chatID)
					} else {
						msg.Text = "I fucked up"
					}
				} else {
					msg.Text = "No no no, please reply, and only after - try ban"
				}
			case "bhadd":
				for i := 0; i < len(bh); i++ {
					s := strings.Split(update.Message.Text, " ")
					if bh[i] == "" {
						if len(s) == 2 {
							bh[i] = s[1]
							break
						}
					}
				}
			case "bhlist":
				for i := 0; i < len(bh); i++ {
					msg.Text = msg.Text + "\n" + bh[i]
				}
			case "ping":
				message := tgbotapi.NewMessage(update.Message.Chat.ID, "Fuck you")
				bot.Send(message)
			case "savestab":
				stableID = update.Message.MessageID
				chatID = update.Message.Chat.ID
			case "f":
				chatID = update.Message.Chat.ID
				strchatID = strconv.FormatInt(chatID, 10)
				http.Get("https://api.telegram.org/bot" + token + "/sendSticker?chat_id=" + strchatID + "&sticker=CAADAgADsgADTptkAm1WnTBWvUfiAg")
			case "gay":
				random = rand.Intn(100)
				if counter == 5 {
					msg.Text = "Wait"
					bot.Send(msg)
					time.Sleep(10 * time.Second)
					counter = 0
				} else {
					msg.Text = "You are gay with chance:" + strconv.Itoa(random) + "%"
					counter++
				}
			case "8":
				random = rand.Intn(6)
				switch random {
				case 0:
					msg.Text = "Мой ответ - 'да'"
				case 1:
					msg.Text = "Скорее всего да"
				case 2:
					msg.Text = "хз"
				case 3:
					msg.Text = "Скорее всего нет"
				case 4:
					msg.Text = "Давай ещё раз"
				case 5:
					msg.Text = "Мой ответ-'нет'"
				}
			case "unsplash":
				userID = update.Message.From.ID
				if userID == 847529348 {
					msg.Text = "Suck"
				} else {

					httpGet, err := http.Get(unsplashResponse)
					errcheck(&err)
					var photos = unsplash{}
					json.NewDecoder(httpGet.Body).Decode(&photos)
					text = photos.Links.Download
					msg.Text = text
				}
			case "suck":
				msg.Text = "Suck"
			case "music":
				random = rand.Intn(300000-100000) + 100000
				httpGet, err := http.Get("https://api.deezer.com/search?q=queen")
				mp3, er := os.Create("music.mp3")
				errcheck(&er)
				errcheck(&err)
				var mus = music{}
				json.NewDecoder(httpGet.Body).Decode(&mus)
				log.Print(mus.Total)
				out, ver := http.Get(mus.Data[1].Link)
				errcheck(&ver)
				io.Copy(mp3, out.Body)
				chatID = update.Message.Chat.ID
				strchatID = strconv.FormatInt(chatID, 10)
				http.Get("https://api.telegram.org/bot" + token + "/sendAudio?chat_id=404334300&audio=music.mp3")
			case "flex":
				chatID = update.Message.Chat.ID
				strchatID = strconv.FormatInt(chatID, 10)
				http.Get("https://api.telegram.org/bot" + token + "/sendAnimation?chat_id=" + strchatID + "&animation=CgADAgADLQMAAn-E6UlWs6GdWI1ZvgI")
			case "shrug":
				msg.Text = "¯\\_(ツ)_/¯"
			case "fix":
				if update.Message.ReplyToMessage != nil {
					replyID = update.Message.ReplyToMessage.From.ID
					msg.Text = strconv.Itoa(replyID)
				}
			case "Foxed":
				msg.Text = "http://qiwi.me/f0x1d"
			case "chuck":
				httpGet, err := http.Get(chuckResponse)
				errcheck(&err)
				var Chuck = chuck{}
				json.NewDecoder(httpGet.Body).Decode(&Chuck)
				text = Chuck.Value
				URLText := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				URLText.Text = Chuck.IconURL
				bot.Send(URLText)
				msg.Text = text
			case "news":
				random = rand.Intn(7)
				httpGet, err := http.Get(newsResponse)
				errcheck(&err)
				var news = news{}
				json.NewDecoder(httpGet.Body).Decode(&news)
				msg.Text = news.Articles[random].Title + "\n" + "\n" + news.Articles[random].URL + "\n"
			case "random":
				random = rand.Intn(32769)
				strrandom = strconv.Itoa(random)
				msg.Text = strrandom
			case "stable":
				strstableID = strconv.Itoa(stableID)
				chatID = update.Message.Chat.ID
				strchatID = strconv.FormatInt(chatID, 10)
				chatID2 := update.Message.Chat.ID
				strchatID2 := strconv.FormatInt(chatID2, 10)
				http.Get("https://api.telegram.org/bot" + token + "/forwardMessage?chat_id=" + strchatID2 + "&from_chat_id=" + strchatID + "&message_id=" + strstableID)
			case "productplacement":
				msg.Text = "Привет, сегодня днем тут в чате у меня спрашивали про инстересную тему, которую я нашел, вот ссылка на нее - @Kernux(ссылка в ЛС)\n\nГлавное понять правильно как использовать выгодно инфу что там есть, у меня получилось ну очень прибыльно!)"
			case "info":
				msg.Text = "Author:@Kernux\nHello World:Hello,World!"
			case "changelog":
				msg.Text = "[+]changelog\n [-] old pyhton:))\n [+]banhammer(I do not know why this work(лютый *****код))\n [=]Небольшой рефакторинг(нужно было для введения банахммера)\n [=]Разграничение каждого кейса на потоки + забивание новыми :DDDDD\n[=]Trash remooved\n[=]Old git commit :|"
			default:
				//msg.Text = "Correct your command please"
			}
			bot.Send(msg)
		}
	}
}
