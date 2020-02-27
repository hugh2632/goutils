package crawler

import (
	"context"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/robertkrimen/otto"
	"goutils/model"
	"goutils/pkg/setting"
	"log"
	"strings"
	"time"
)

var browser context.Context

func init(){
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", setting.RunMode != "debug"),
		chromedp.Flag("ignore-certificate-errors", true),
	)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	browser, _ = chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	_ = chromedp.Run(browser)
}

func GetPdfBytes(url string) ([]byte, error){
	ch := make(chan struct{})
	var loaded = false
	var stopped = false

	taskCtx, cancel := chromedp.NewContext(browser)
	defer cancel()
	var err error
	go func() {
		chromedp.ListenTarget(taskCtx, func(ev interface{}) {
				//te :=  reflect.Indirect(reflect.ValueOf(ev)).Type()
				//name := te.String()
				////if strings.HasPrefix(name, "page.") || strings.HasPrefix(name, "network.") { //页面事件
				//if strings.HasPrefix(name, "page."){
				//	fmt.Println(name + "\t" + reflect.ValueOf(ev).Elem().String())
				//}
			switch ev.(type) {
			case *page.EventLoadEventFired://确保
				go func() {
					loaded= true
					if stopped{
						ch <- struct{}{}
					}
				}()
			case *page.EventFrameStoppedLoading://会多次触发。。 不知道原因
				go func() {
					stopped = true
					if loaded {
						ch <- struct{}{}
					}
				}()
			}
		})
	}()
	err = chromedp.Run(taskCtx,
		chromedp.Navigate(url))
	if err != nil {
		return nil, err
	}

	select {
	case <-time.After(setting.ReadTimeout):
		return nil, errors.New("加载超时")
	case <-ch:
		var pdfBuffer []byte
		err = chromedp.Run(taskCtx,
			chromedp.ActionFunc(func(ctx context.Context) error {
				pdfBuffer, _, err = page.PrintToPDF().WithPrintBackground(true).Do(ctx)
				return err
			}),
		)
		return pdfBuffer, err
	}
}

func GetData(task model.TaskInfo) error{
	ch := make(chan struct{})
	var loaded = false
	var stopped = false

	taskCtx, cancel := chromedp.NewContext(browser)
	defer cancel()
	var err error
	go func() {
		chromedp.ListenTarget(taskCtx, func(ev interface{}) {
			//te :=  reflect.Indirect(reflect.ValueOf(ev)).Type()
			//name := te.String()
			////if strings.HasPrefix(name, "page.") || strings.HasPrefix(name, "network.") { //页面事件
			//if strings.HasPrefix(name, "page."){
			//	fmt.Println(name + "\t" + reflect.ValueOf(ev).Elem().String())
			//}
			switch ev.(type) {
			case *page.EventLoadEventFired://确保
				go func() {
					loaded= true
					if stopped{
						ch <- struct{}{}
					}
				}()
			case *page.EventFrameStoppedLoading://会多次触发。。 不知道原因
				go func() {
					stopped = true
					if loaded {
						ch <- struct{}{}
					}
				}()
			}
		})
	}()
	err = chromedp.Run(taskCtx,
		chromedp.Navigate(task.Url))
	if err != nil {
		return err
	}

	select {
	case <-time.After(setting.ReadTimeout):
		return errors.New("加载超时")
	case <-ch:
		var datas []model.DataInfo
		dom, err := Gethtml(task.Url)
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(dom))
		var vm = otto.New()
		err = vm.Set("doc", doc)
		err = vm.Set("datas", datas)
		_, err = vm.Run(task.Rules)

		//获取data的文本内容
		for i, _ :=range datas{
			htmlStr, er := Gethtml(datas[i].Url)
			if er == nil {
				datas[i].Backup = htmlStr
			}
		}
		return err
	}
}

func Gethtml(url string) (string, error){
	ch := make(chan struct{})
	var loaded = false
	var stopped = false

	taskCtx, cancel := chromedp.NewContext(browser)
	defer cancel()
	var err error
	go func() {
		chromedp.ListenTarget(taskCtx, func(ev interface{}) {
			//te :=  reflect.Indirect(reflect.ValueOf(ev)).Type()
			//name := te.String()
			////if strings.HasPrefix(name, "page.") || strings.HasPrefix(name, "network.") { //页面事件
			//if strings.HasPrefix(name, "page."){
			//	fmt.Println(name + "\t" + reflect.ValueOf(ev).Elem().String())
			//}
			switch ev.(type) {
			case *page.EventLoadEventFired://确保
				go func() {
					loaded= true
					if stopped{
						ch <- struct{}{}
					}
				}()
			case *page.EventFrameStoppedLoading://会多次触发。。 不知道原因
				go func() {
					stopped = true
					if loaded {
						ch <- struct{}{}
					}
				}()
			}
		})
	}()
	err = chromedp.Run(taskCtx,
		chromedp.Navigate(url))
	if err != nil {
		return "", err
	}

	select {
	case <-time.After(setting.ReadTimeout):
		return "", errors.New("加载超时")
	case <-ch:
		var res string
		var ids []cdp.NodeID
		err = chromedp.Run(taskCtx,
			chromedp.NodeIDs(`document`, &ids, chromedp.ByJSPath),
			chromedp.ActionFunc(func(ctx context.Context) error {
				var err1 error
				if len(ids) > 0 {
					res, err1 =dom.GetOuterHTML().WithNodeID(ids[0]).Do(ctx)
					return  err1
				} else {
					panic(`获取HTML文档失败,长度为0`)
				}
				return err1
			}),
		)
		return res, err
	}
}