package torrent

import (
	_ "embed"
)

//go:embed release.txt
var LatestTorrent string
