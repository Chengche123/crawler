package pmodel

import "google.golang.org/protobuf/proto"

func UnmashalComicDetailResponse(bs []byte) (*ComicDetailResponse, error) {
	var res ComicDetailResponse

	err := proto.Unmarshal(bs, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
