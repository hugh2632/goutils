package htmlUtil

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/url"
	"strings"
)

func GetSelfNodeStr(node *html.Node) (res string, err error) {
	switch node.Type {
	case html.ErrorNode:
		return "", errors.New("不能转换error节点")
	case html.TextNode:
		return html.EscapeString(node.Data), nil
	case html.DocumentNode:
		return "", nil
	case html.ElementNode:
		// No-op.
	case html.CommentNode:
		res += "<!--" + node.Data + "-->"
		return res, nil
	case html.DoctypeNode:
		return "", nil
	default:
		return "", errors.New("未知的html节点")
	}
	res += "<" + node.Data
	for _, a := range node.Attr {
		res += " "
		if a.Namespace != "" {
			res += a.Namespace
			res += ":"
		}
		res += a.Key + "=\"" + html.EscapeString(a.Val) + "\""
	}
	if voidElements[node.Data] {
		if node.FirstChild != nil {
			err = fmt.Errorf("不支持子节点的类型却有了子节点", node.Data)
		}
		res += "/>"
	} else {
		res += "></" + node.Data + ">"
	}
	return res, err
}

// Section 12.1.2, "Elements", gives this list of void elements. Void elements
// are those that can't have any contents.
var voidElements = map[string]bool{
	"area":    true,
	"base":    true,
	"br":      true,
	"col":     true,
	"command": true,
	"embed":   true,
	"hr":      true,
	"img":     true,
	"input":   true,
	"keygen":  true,
	"link":    true,
	"meta":    true,
	"param":   true,
	"source":  true,
	"track":   true,
	"wbr":     true,
}

func SaveDocument(doc *html.Node, fileName string) error{
	var buf = bytes.NewBuffer([]byte{})
	err := html.Render(buf, doc)
	err = ioutil.WriteFile( fileName, buf.Bytes(), 0644)
	return err
}

func GetUrl(node *html.Node, root *url.URL) *url.URL{
	switch node.Type {
	case html.ElementNode: //todo url发现和记录
		if node.Data == "use"{
			return nil
		}
		for _, a := range node.Attr {
			if a.Key == "href" || a.Key == "src" {
				u, err := url.Parse(strings.Trim(a.Val, " "))
				if err != nil {
					fmt.Println( "转换URL失败" + a.Val + err.Error())
					return nil
				}
				switch u.Scheme {
				case "javascript", "data":
					return nil
				case "https","http":
					return u
				default:
					tmpUrl := a.Val
					if strings.HasPrefix(a.Val, "//") { //双斜杠开始的
						tmpUrl = root.Scheme + ":" + tmpUrl
					} else if strings.HasPrefix(a.Val, "./"){
						if u.Path != ""{
							fileN := GetFileNameFromRaw(a.Val)
							direc := GetAbsDirectory(root)
							tmpUrl = direc + "/" + fileN
						}
					} else if strings.HasPrefix(a.Val, "../"){
						var direc string
						list := strings.Split(root.Path, "/")
						for range list{
							direc = GetAbsDirectory(root)
						}
						tmpUrl = direc + "/" + GetFileNameFromRaw(a.Val)
					}
					tu , err := url.Parse(tmpUrl)
					if err != nil {
						fmt.Println( "转换URL失败" + tmpUrl + err.Error())
						return nil
					}
					return tu
				}
			}
		}
	}
	return nil
}

func GetFileNameFromRaw(u string) string{
	lashSlashIndex := strings.LastIndex(u, "/")
	if lashSlashIndex > -1 {
		return u[lashSlashIndex + 1:]
	}
	return ""
}

func GetAbsDirectory(u *url.URL) string{
	if u.Path == "" || u.Path == "/" {
		return u.Scheme + "://" + u.Host
	} else {
		lastSlashIndex := strings.LastIndex(u.Path , "/")
		return u.Scheme + "://" + u.Host +  u.Path[0: lastSlashIndex]
	}
}