# gpd

An attempt to build a tiny cmd presenter in go. In the end it should fit on a floppy disk together with an md presentation. Inspired by some tweets https://twitter.com/MaartjeME/status/1356144138748637184

Since it's should be compatible with mpd, but in go and I'm not that creative I'm gonna name it gpd.



## build

To get a tiny binary follow this tutoarial https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/

basically `go build -ldflags "-s -w"` and `upx gpd`
