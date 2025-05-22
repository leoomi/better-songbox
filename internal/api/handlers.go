package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/leoomi/better-songbox/internal/api/parser"
)

const test = "<div class='container'><table class='table table-hover table-striped '><tr><th>CANTOR</th><th>MÚSICA</th><th>CÓDIGO</th><th>CLIPE</th><th>AT</th></tr><tr class='SearchRow'><td>Tequendama De Oro</td><td>Juana La Cubana</td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>13-008-10</a></td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>CDG</a></td><td>G02</td></tr><tr class='SearchRow'><td>Verve</td><td>Bitter Sweet Symphony</td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>68-452-15</a></td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>CDG</a></td><td>G05</td></tr><tr class='SearchRow'><td>Verve</td><td>The Drugs Don't Work</td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>72-114-15</a></td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>CDG</a></td><td>G07</td></tr><tr class='SearchRow'><td>Verve</td><td>Lucky Man</td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>72-117-01</a></td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>CDG</a></td><td>G06</td></tr><tr class='SearchRow'><td>Verve</td><td>Love Is Noise</td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>72-270-02</a></td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>CDG</a></td><td>G03</td></tr><tr class='SearchRow'><td>Verve</td><td>Sonnet</td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>74-809-10</a></td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>CDG</a></td><td>G01</td></tr><tr class='SearchRow'><td>Verve</td><td>Rather Be</td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>87-051-15</a></td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>CDG</a></td><td>G04</td></tr><tr class='SearchRow'><td>Verve Pipe</td><td>The Freshman</td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>59-056-03</a></td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>CDG</a></td><td>G05</td></tr><tr class='SearchRow'><td>Verve Pipe</td><td>Photograph</td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>68-308-06</a></td><td class='NumCel'><a href='#' id='NumCel' onclick='AdicionarNum(this)' rel='popover' data-toggle='popover' data-content='Adicionado!' data-placement='left'>CDG</a></td><td>G02</td></tr></table></div>"

func SearchHandler(w http.ResponseWriter, req *http.Request) {
	resp, err := http.PostForm("https://www.songbox.info/Search.php", url.Values{
		"CANT":  {"teste"},
		"MUS":   {""},
		"POSIC": {"comeco"},
		"OPT2":  {"karaoke"},
	})
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	parser.ParseKaraokeSearch(strings.NewReader(test))
}
