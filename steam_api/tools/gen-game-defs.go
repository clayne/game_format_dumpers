package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type Game struct {
	Name      string
	Defs      string
	ExtraDefs string
	Versions  map[string]int
}

type Version struct {
	id    string
	key   string
	value string
}

func main() {
	log.SetPrefix("gen-games-defs: ")
	log.SetFlags(0)

	var maketype string
	flag.StringVar(&maketype, "t", "unix", "specify makefile type")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() < 1 {
		usage()
	}

	data, err := os.ReadFile(flag.Arg(0))
	check(err)

	var games []Game
	err = json.Unmarshal(data, &games)
	check(err)

	def := findtitle("DEFAULT", games)
	if def == nil {
		log.Fatal("No default definition specified!")
	}
	for i := range games {
		g := &games[i]
		inherit(def, g)
	}

	switch maketype {
	case "unix":
		makeunix(games)
	case "windows":
		makewindows(games)
	default:
		log.Fatal("Unknown makefile type")
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: [options] games.json")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "Supported makefile types: unix windows")
	os.Exit(2)
}

func findtitle(name string, games []Game) *Game {
	for _, g := range games {
		if name == g.Name {
			return &g
		}
	}
	return nil
}

func inherit(parent, child *Game) {
	if child.Defs == "" {
		child.Defs = parent.Defs
	}
	if child.ExtraDefs == "" {
		child.ExtraDefs = parent.ExtraDefs
	}
	if child.Versions == nil {
		child.Versions = make(map[string]int)
	}
	for ident, version := range parent.Versions {
		_, exist := child.Versions[ident]
		if !exist {
			child.Versions[ident] = version
		}
	}
}

func makeunix(games []Game) {
	fmt.Printf("# AUTOGENERATED FILE\n\n")
	for _, g := range games {
		fmt.Printf("ifeq ($(GAME), %s)\n", g.Name)
		fmt.Printf("\tDEFS = %s\n", g.Defs)
		if g.ExtraDefs != "" {
			fmt.Printf("\tDEFS += %s\n", g.ExtraDefs)
		}

		for _, v := range VERSIONS {
			steamver, exist := g.Versions[v.id]
			if !exist {
				log.Fatalf("Missing steam version for %s: %s", g.Name, v.id)
			}

			fmt.Printf("\tDEFS += -D %s='\"%s%03d\"'\n", v.key, v.value, steamver)
		}
		fmt.Printf("endif\n")
		fmt.Println()
	}
}

func makewindows(games []Game) {
	fmt.Printf("# AUTOGENERATED FILE\n\n")
	for _, g := range games {
		fmt.Printf("!if \"$(GAME)\" == %q\n", g.Name)
		fmt.Printf("DEFS = %s\n", g.Defs)
		if g.ExtraDefs != "" {
			fmt.Printf("DEFS = $(DEFS) %s\n", g.ExtraDefs)
		}

		for _, v := range VERSIONS {
			steamver, exist := g.Versions[v.id]
			if !exist {
				log.Fatalf("Missing steam version for %s: %s", g.Name, v.id)
			}

			fmt.Printf("DEFS = $(DEFS) -D %s=\"\\\"%s%03d\\\"\"\n", v.key, v.value, steamver)
		}
		fmt.Printf("!endif\n")
		fmt.Println()
	}
}

var VERSIONS = []Version{
	{"applist", "STEAM_APPLIST_VERSION", "STEAMAPPLIST_INTERFACE_VERSION"},
	{"appticket", "STEAM_APP_TICKET_VERSION", "STEAMAPPTICKET_INTERFACE_VERSION"},
	{"apps", "STEAM_APPS_VERSION", "STEAMAPPS_INTERFACE_VERSION"},
	{"client", "STEAM_CLIENT_VERSION", "SteamClient"},
	{"controller", "STEAM_CONTROLLER_VERSION", "SteamController"},
	{"friends", "STEAM_FRIENDS_VERSION", "SteamFriends"},
	{"gamesearch", "STEAMGAMESEARCH_VERSION", "SteamMatchGameSearch"},
	{"gameserver", "STEAM_GAMESERVER_VERSION", "SteamGameServer"},
	{"gameserverapps", "STEAM_GAMESERVERAPPS_VERSION", "SteamGameServerApps"},
	{"htmlsurface", "STEAM_HTMLSURFACE_VERSION", "STEAMHTMLSURFACE_INTERFACE_VERSION_"},
	{"http", "STEAM_HTTP_VERSION", "STEAMHTTP_INTERFACE_VERSION"},
	{"input", "STEAM_INPUT_VERSION", "SteamInput"},
	{"matchmaking", "STEAM_MATCHMAKING_VERSION", "SteamMatchMaking"},
	{"matchmakingservers", "STEAM_MATCHMAKINGSERVERS_VERSION", "SteamMatchMakingServers"},
	{"music", "STEAM_MUSIC_VERSION", "STEAMMUSIC_INTERFACE_VERSION"},
	{"musicremote", "STEAM_MUSICREMOTE_VERSION", "STEAMMUSICREMOTE_INTERFACE_VERSION"},
	{"networking", "STEAM_NETWORKING_VERSION", "SteamNetworking"},
	{"parentalsettings", "STEAM_PARENTALSETTINGS_VERSION", "STEAMPARENTALSETTINGS_INTERFACE_VERSION"},
	{"parties", "STEAM_PARTIES_VERSION", "SteamParties"},
	{"remoteplay", "STEAM_REMOTEPLAY_VERSION", "STEAMREMOTEPLAY_INTERFACE_VERSION"},
	{"remotestorage", "STEAM_REMOTESTORAGE_VERSION", "STEAMREMOTESTORAGE_INTERFACE_VERSION"},
	{"screenshots", "STEAM_SCREENSHOTS_VERSION", "STEAMSCREENSHOTS_INTERFACE_VERSION"},
	{"ugc", "STEAM_UGC_VERSION", "STEAMUGC_INTERFACE_VERSION"},
	{"unifiedmessages", "STEAM_UNIFIEDMESSAGES_VERSION", "STEAMUNIFIEDMESSAGES_INTERFACE_VERSION"},
	{"user", "STEAM_USER_VERSION", "SteamUser"},
	{"userstats", "STEAM_USERSTATS_VERSION", "STEAMUSERSTATS_INTERFACE_VERSION"},
	{"utils", "STEAM_UTILS_VERSION", "SteamUtils"},
	{"video", "STEAM_VIDEO_VERSION", "STEAMVIDEO_INTERFACE_V"},
	{"tv", "STEAM_TV_VERSION", "STEAMTV_INTERFACE_V"},
}
