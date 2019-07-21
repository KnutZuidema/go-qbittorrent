package sync

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/KnutZuidema/go-qbittorrent/pkg"
	"github.com/KnutZuidema/go-qbittorrent/pkg/model"
)

type Client struct {
	BaseUrl string
	Client  *http.Client
	Logger  logrus.FieldLogger
}

func (c Client) GetMainData(rid int) (*model.SyncMainData, error) {
	params := url.Values{}
	params.Add("rid", strconv.Itoa(rid))
	endpoint := c.BaseUrl + "/maindata?" + params.Encode()
	var res model.SyncMainData
	if err := pkg.GetInto(c.Client, &res, endpoint, nil); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c Client) GetTorrentPeersData(hash string, rid int) (*model.SyncPeersData, error) {
	params := url.Values{}
	params.Add("hash", hash)
	params.Add("rid", strconv.Itoa(rid))
	endpoint := c.BaseUrl + "/torrentPeers?" + params.Encode()
	var res model.SyncPeersData
	if err := pkg.GetInto(c.Client, &res, endpoint, nil); err != nil {
		return nil, err
	}
	return &res, nil
}
