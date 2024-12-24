package words

import (
    "context"

    "proxima/app/gateway/api/words/v1"
    words "proxima/app/word/api/words/v1"

    "github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {
    word, err := c.WordsClient.Get(ctx, &words.GetReq{
        Id: uint32(req.Id),
    })
    if err != nil {
        return nil, err
    }

    if word == nil {
        return nil, gerror.New("word not found")
    }
    return &v1.DetailRes{
        Id:         uint(word.Words.Id),
        Word:       word.Words.Word,
        Definition: word.Words.Definition,
    }, nil
}
