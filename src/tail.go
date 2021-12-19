package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hpcloud/tail"
	"github.com/hpcloud/tail/ratelimiter"

	tpm "tailpipe/mail"
	"tailpipe/payload"
)

func tailf(conf tConfig) {
	t, err := tail.TailFile(conf.FileToWatch, tail.Config{
		Follow:      true,
		Location:    &tail.SeekInfo{Offset: 0, Whence: os.SEEK_END},
		RateLimiter: ratelimiter.NewLeakyBucket(10, time.Duration(5)*time.Second),
	})
	if err != nil {
		fmt.Printf("%q\n", err)
	}
	for line := range t.Lines {
		if rxMatch(conf.RegexFilter, line.Text) == true {
			tpm.SendMail(
				payload.NewPayload(line),
				conf.Mail,
			)
		}
	}
}
