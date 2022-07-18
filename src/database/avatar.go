package database

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"log"
	"strconv"
	"time"
)

var encoder = png.Encoder{
	CompressionLevel: png.BestCompression,
}

func (db *Database) GetAvatar(avatarURL string, size int) (image.Image, error) {
	ctx, finishCtx := RedisContext()

	res, err := db.Client.Get(ctx, avatarURL+"-"+strconv.Itoa(size)).Result()
	finishCtx()

	if err != nil {
		return nil, err
	}

	decoded, err := base64.StdEncoding.DecodeString(res)

	if err != nil {
		return nil, err
	}

	image, _, err := image.Decode(bytes.NewReader(decoded))

	if err != nil {
		return nil, err
	}

	return image, nil
}

func (db *Database) SetAvatar(avatarURL string, size int, image image.Image) error {
	buff := new(bytes.Buffer)
	err := encoder.Encode(buff, image)

	if err != nil {
		log.Print(err)
		return err
	}

	ctx, finishCtx := RedisContext()

	err = db.Client.Set(ctx, avatarURL+"-"+strconv.Itoa(size), base64.StdEncoding.EncodeToString(buff.Bytes()), time.Hour).Err()
	finishCtx()

	return err
}
