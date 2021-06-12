package picker

import (
	"astropetal/notify"
	"astropetal/timing"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"git.sr.ht/~adnano/go-gemini"
	"log"
	"net/url"
	"regexp"
	"time"
)

const (
	gardenUrl = "/app/visit"
	pickUrl   = "/search"
)

var floweringRe, pickedRe *regexp.Regexp

func init() {
	floweringRe = regexp.MustCompile(`.+ flowering *`)
	pickedRe = regexp.MustCompile(`You spot a (.+) petal lying on the ground nearby`)
}

type Picker struct {
	baseUrl string
	tlsCert *tls.Certificate
	client  gemini.Client
}

func NewPicker(baseUrl string, tlsCert *tls.Certificate) *Picker {
	return &Picker{baseUrl, tlsCert, gemini.Client{}}
}

func (p *Picker) Pick() *notify.Report {
	report := notify.NewReport()
	msg := "Let's pick some petals"
	log.Print(msg)
	report.Push(notify.StatusOk, msg)

	urls, err := p.getFloweringPlants()
	if err != nil {
		log.Print(err)
		report.Push(notify.StatusErr, err.Error())
		return report
	}

	msg = fmt.Sprintf("Found %d flowering plants", len(urls))
	log.Print(msg)
	report.Push(notify.StatusInfo, msg)

	count := 0
	for _, u := range urls {
		time.Sleep(timing.Approx(10*time.Second, 5*time.Second))
		picked, err := p.pickPetal(u)
		if err != nil {
			log.Print(err)
			report.Push(notify.StatusErr, err.Error())
			continue
		}
		if picked {
			count++
		}
	}

	msg = fmt.Sprintf("Total %d petals picked", count)
	log.Print(msg)
	report.Push(notify.StatusInfo, msg)
	return report
}

func (p *Picker) getFloweringPlants() ([]string, error) {
	res, err := p.getPage(p.baseUrl + gardenUrl)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	var urls []string
	err = gemini.ParseLines(res.Body, func(line gemini.Line) {
		if line, ok := line.(gemini.LineLink); ok {
			if floweringRe.MatchString(line.Name) {
				urls = append(urls, line.URL)
			}
		}
	})
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func (p *Picker) pickPetal(plantUrl string) (bool, error) {
	res, err := p.getPage(p.baseUrl + plantUrl + pickUrl)
	if err != nil {
		return false, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer res.Body.Close()

	picked := false
	err = gemini.ParseLines(res.Body, func(line gemini.Line) {
		if pickedRe.MatchString(line.String()) {
			picked = true
		}
	})
	if err != nil {
		return false, err
	}
	return picked, nil
}

func (p *Picker) getPage(u string) (*gemini.Response, error) {
	var doRequest func(*gemini.Request, int) (*gemini.Response, error)
	doRequest = func(req *gemini.Request, redirects int) (*gemini.Response, error) {
		if redirects > 5 {
			return nil, errors.New("to many redirects")
		}
		ctx := context.Background()
		res, err := p.client.Do(ctx, req)
		if err != nil {
			return nil, err
		}
		if res.Status.Class() == gemini.StatusRedirect {
			target, err := url.Parse(res.Meta)
			if err != nil {
				return nil, err
			}
			target = req.URL.ResolveReference(target)
			redirect := *req
			redirect.URL = target
			return doRequest(&redirect, redirects+1)
		}
		return res, nil
	}

	req, err := gemini.NewRequest(u)
	if err != nil {
		return nil, err
	}
	req.Certificate = p.tlsCert
	return doRequest(req, 0)
}
