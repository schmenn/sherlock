package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/Schmenn/sherlock/requests"
	"github.com/fatih/color"
)

// SitesStruct is what's seen in the sites.json file
type SitesStruct struct {
	URL      string      `json:"url"`
	PrintURL string      `json:"printURL"`
	Method   string      `json:"method"`
	Check    string      `json:"check"`
	Body     string      `json:"body"`
	Match    interface{} `json:"match"`
	Headers  [][]string  `json:"headers"`
}

func main() {

	if len(os.Args) < 2 {
		color.HiRed("Please specify a username")
		os.Exit(1)
	}
	name := os.Args[1]
	f, err := os.ReadFile("sites.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	var sites []SitesStruct
	err = json.Unmarshal(f, &sites)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	var wg sync.WaitGroup
	for _, site := range sites {
		wg.Add(1)
		go func(site SitesStruct) {
			defer wg.Done()
			req := requests.NewRequest(site.Method, site.URL, name, site.Body, site.Headers)
			if site.PrintURL != "" {
				site.URL = site.PrintURL
			}
			if req == nil {
				color.HiRed("[ ] " + strings.Split(strings.Split(site.URL, "//")[1], "/")[0])
				return
			}
			resp := requests.MakeRequest(req, &http.Client{})
			if resp == nil {
				color.HiRed("[ ] " + strings.Split(strings.Split(site.URL, "//")[1], "/")[0])
				return
			}
			if requests.CheckValidity(resp, site.Check, site.Match) {
				color.HiGreen(fmt.Sprintf("[X] %s", strings.Replace(site.URL, "{}", name, 1)))
			} else {
				color.HiRed("[ ] " + strings.Split(strings.Split(site.URL, "//")[1], "/")[0])
			}
		}(site)
	}
	wg.Wait()
}

/*
9GAG
About.me
AllTrails
Archive.org
Asciinema
Ask Fedora
AskFM
BinarySearch
BitBucket
BitCoinForum
Blogger
BodyBuilding
Bookcrossing
BuyMeACoffee
BuzzFeed
CNET
CapFriendly
Carbonmade
Career.habr
Cent
Championat
Chatujme.cz
Chess
Cloob
CloudflareCommunity
Clozemaster
Codecademy
Codechef
Codepen
Codewars
ColourLovers
Contently
Coroflot
Countable
Cracked
Crevado
DEV Community
DailyMotion
Designspiration
DeviantART
Discogs
Discuss.Elastic.co
Disqus
Docker Hub
Dribbble
Duolingo
Ello
Etsy
Euw
EyeEm
F3.cool
Facebook
Facenama
Fandom
Flickr
Flightradar24
Flipboard
Football
FortniteTracker
Freelance.habr
Freelancer.com
Freesound
GDProfiles
Gamespot
GetMyUni
Giphy
GitHub Support Community
GitLab
Gitee
GoodReads
Gravatar
Gumroad
GunsAndAmmo
GuruShots
HackTheBox
Hackaday
HackerNews
HackerOne
HackerRank
House-Mixes.com
Houzz
HubPages
Hubski
ICQ
IFTTT
ImgUp.cz
Imgur
Instructables
Issuu
Itch.io
Jimdo
Kaggle
Kali community
Keybase
Kik
Kongregate
LOR
Launchpad
LeetCode
Letterboxd
Lichess
LiveJournal
LiveLeak
Lobsters
Lolchess
Medium
Memrise
MixCloud
Munzee
MyAnimeList
MyMiniFactory
Myspace
NICommunityForum
NameMC (Minecraft.net skins)
NationStates Nation
NationStates Region
Naver
Newgrounds
Nightbot
NotABug.org
OK
OpenStreetMap
Opensource
Oracle Community
Otzovik
OurDJTalk
PCGamer
PCPartPicker
PSNProfiles.com
Packagist
Pastebin
Patreon
Periscope
Pinkbike
Pinterest
PlayStore
Plug.DJ
Pokemon Showdown
Polarsteps
Polygon
ProductHunt
PromoDJ
PyPi
Quizlet
Quora
Raidforums
Rajce.net
Rate Your Music
Redbubble
Reddit
Repl.it
ResearchGate
ReverbNation
Roblox
RubyGems
Sbazar.cz
Scratch
Scribd
ShitpostBot5000
Signal
Slack
Slashdot
SlideShare
Smashcast
Smule
SoundCloud
SourceForge
SoylentNews
SparkPeople
Speedrun.com
Splits.io
Sporcle
SportsRU
Spotify
Star Citizen
Steam
SteamGroup
Steamid
Strava
SublimeForum
TETR.IO
Telegram
Tellonym.me
TikTok
Tinder
TrackmaniaLadder
TradingView
Trakt
TrashboxRU
Trello
TripAdvisor
TryHackMe
Twitch
Twitter
Typeracer
Ultimate-Guitar
Unsplash
VK
VSCO
Velomania
Venmo
Vero
Vimeo
Virgool
VirusTotal
Warrior Forum
Wattpad
We Heart It
WebNode
Whonix Forum
Wikidot
Wikipedia
Windy
Wix
WordPress
WordPressOrg
Xbox Gamertag
YouNow
YouPic
YouTube
Zhihu
akniga
allmylinks
aminoapp
authorSTREAM
babyRU
babyblogRU
chaos.social
couchsurfing
d3RU
dailykos
datingRU
devRant
drive2
eGPU
eintracht
fixya
fl
forum_guns
forumhouseRU
geocaching
gfycat
habr
hackster
hunting
iMGSRC.RU
igromania
interpals
irecommend
jbzd.com.pl
jeuxvideo
kofi
kwork
labpentestit
last.fm
leasehackr
livelib
mastodon.cloud
mastodon.social
mastodon.technology
mastodon.xyz
mercadolivre
metacritic
moikrug
mstdn.io
nairaland.com
nnRU
note
npm
opennet
osu!
phpRU
pikabu
pr0gramm
prog.hu
radio_echo_msk
satsisRU
social.tchncs.de
spletnik
svidbook
toster
uid
*/
