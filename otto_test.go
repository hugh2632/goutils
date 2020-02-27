package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
	"strings"
	"testing"
)

func Test_Otto(t *testing.T){
	//var ctx crawler.Context
	//var datas []model.DataInfo
	var dom = `<!DOCTYPE html><html lang="en"><head>
    <meta charset="UTF-8">
<meta name="keywords" content="工业互联网安全应急响应中心，工控，互联网安全，应急响应中心，ics-cert">
<meta name="description" content="工业互联网安全应急响应中心，工控，互联网安全，应急响应中心，ics-cert">
<title>工业互联网安全应急响应中心</title>
<link href="/portal/pub/1/images/cms/site/2018/02/favicon.ico" rel="icon" type="baseImages/x-icon">
<link href="https://www.ics-cert.org.cn/portal/css/bootstrap/bootstrap.min.css" rel="stylesheet" type="text/css">
<link href="https://www.ics-cert.org.cn/portal/css/pages/header.css?t=0339285032" rel="stylesheet" type="text/css">
<link href="https://www.ics-cert.org.cn/portal/css/pages/footer.css?t=0339285032" rel="stylesheet" type="text/css">
<link href="https://www.ics-cert.org.cn/portal/css/pages/common.css?t=0339285032" rel="stylesheet" type="text/css">    <link href="../../css/pages/listPage.css?t=0339285032" rel="stylesheet" type="text/css">
</head>
<body>
<header class="container-fluid">
	<input id="siteDomain" type="hidden" value="https://www.ics-cert.org.cn/portal">
    <div class="center-block mid">
        <div class="nav-top">
            <div class="logo">
                <img src="https://www.ics-cert.org.cn/portal/baseImages/logo.png">
            </div>
            <div class="right">
               <div class="right-box">
                    <div class="tog">
                        <a href="#" id="lg_link">
                            <span class="img-circle img-box"><img class="" src="https://www.ics-cert.org.cn/portal/baseImages/user_default_head.png" alt="user"></span>
                            <span class="name">用户123</span>
                            <span class="ic-row"></span>
                        </a>
                        <ul id="lg_box" class="nav">
                            <li><a href="https://www.ics-cert.org.cn/portal/toMemberCentre" target="_blank"><span class="ic ic-1"></span>会员中心</a></li>
                            <li><a id="exit" href="https://www.ics-cert.org.cn/portal/logout"><span class="ic ic-2"></span>退出</a> </li>
                        </ul>
                    </div>
                </div>
                <button id="lg" type="button" class="btn btn-lg">登录</button>
                <button id="lg-register-btn" type="button" class="btn btn-lg">注册</button>
                <div class="search">
                    <form action="https://www.ics-cert.org.cn/portal/index.html" name="getSearchResultForm" method="post">
                   		<input class="box" type="text"><a class="img" type="btn" onclick="getSearchResultForm.submit()"><img src="https://www.ics-cert.org.cn/portal/baseImages/serch.png"></a>
					</form>
                </div>
            </div>
        </div>
        <div class="nav-bottom">
            <ul class="navbar navbar-nav">
                <li><a href="https://www.ics-cert.org.cn/portal/index.html">首页</a> </li>
	                 	<li class="active"><a href="#">威胁预警</a>
		                 	<ul id="sub-list">
							       		<li><a href="https://www.ics-cert.org.cn/portal/page/111/index_1.html" target="_blank">漏洞公告</a></li>
							       		<li><a href="https://www.ics-cert.org.cn/portal/page/112/index_1.html" target="_blank">恶意代码公告</a></li>
		 					</ul>
	 					</li>
	                 	<li><a href="#">态势感知</a>
		                 	<ul id="sub-list">
							       		<li><a href="https://www.ics-cert.org.cn/portal/page/121/index_1.html" target="_blank">安全态势</a></li>
							       		<li><a href="https://www.ics-cert.org.cn/portal/page/122/index_1.html" target="_blank">漏洞态势</a></li>
		 					</ul>
	 					</li>
	                 	<li><a href="#">新闻动态</a>
		                 	<ul id="sub-list">
							       		<li><a href="https://www.ics-cert.org.cn/portal/page/131/index_1.html" target="_blank">CII-SRC资讯</a></li>
							       		<li><a href="https://www.ics-cert.org.cn/portal/page/132/index_1.html" target="_blank">国内资讯</a></li>
							       		<li><a href="https://www.ics-cert.org.cn/portal/page/133/index_1.html" target="_blank">国外资讯</a></li>
		 					</ul>
	 					</li>
	                 	<li><a href="#">检测评估</a>
		                 	<ul id="sub-list">
							       		<li><a href="http://www.ics-cert.org.cn/portal/jcrz.html" target="_blank">检测认证</a></li>
							       		<li><a href="https://www.ics-cert.org.cn/portal/page/142/index_1.html" target="_blank">风险评估</a></li>
		 					</ul>
	 					</li>
	                 	<li><a href="#">标准指南</a>
		                 	<ul id="sub-list">
							       		<li><a href="https://www.ics-cert.org.cn/portal/page/151/index_1.html" target="_blank">标准指南</a></li>
		 					</ul>
	 					</li>
	                 	<li><a href="#">合作体系</a>
		                 	<ul id="sub-list">
							       		<li><a href="http://www.ics-cert.org.cn" target="_blank">合作体系</a></li>
		 					</ul>
	 					</li>
	                 	<li><a href="#">CII-SRC在线</a>
		                 	<ul id="sub-list">
							       		<li><a href="http://www.ics-cert.org.cn/portal/leak/reported/vulnReport" target="_blank">漏洞上报</a></li>
							       		<li><a href="http://www.ics-cert.org.cn" target="_blank">事件受理</a></li>
							       		<li><a href="http://www.ics-cert.org.cn" target="_blank">工具下载</a></li>
		 					</ul>
	 					</li>
            </ul>
        </div>
    </div>
</header><div class="main-content">
    <div class="title">
        <img src="../../baseImages/home.png">
        <span class="mine">我的位置</span>
	        	<a href="../111/index_1.html">威胁预警</a><span class="s-line">&nbsp;&nbsp;&gt;&nbsp;</span>
		        <a href="../111/index_1.html">漏洞公告</a>
    </div>
    <div class="sidebar">
        <a class="top" href="javascript:;">威胁预警</a>
		      	    <a class="list sidebar-list-active" href="../111/index_1.html">漏洞公告</a>	
		          	<a class="list" href="../112/index_1.html">恶意代码公告</a>
        <a class="icon first-icon" href="https://www.ics-cert.org.cn/portal/toZhongCe" target="_blank"><img src="../../baseImages/list-ics-cert-test.png"></a>
        <a class="icon" href="https://openlab.ics-cert.org.cn" target="_blank"><img src="../../baseImages/list-ics-cert-research.png"></a>
        <a class="icon" href="https://seclab.ics-cert.org.cn" target="_blank"><img src="../../baseImages/list-ics-cert-lab.png"></a>
        <a class="icon" href="http://secrank.ics-cert.org.cn:8084/search" target="_blank"><img src="../../baseImages/wlaqnlpm-list.png"></a>
    </div>
    <div class="info">
        <h3>漏洞公告</h3>
        <ul>
        	<ul>
					<li class="top">
						<img src="../../baseImages/right.png">
							<a href="./f37fc48997a9437e9c15ed28fbe327e5.html" target="_blank">联网医疗设备的网络安全和漏洞披露面临新挑战</a>
						<span>2020/02/27</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./cedd1fbe2a654bf58b8051c8647db4fb.html" target="_blank">原创｜GE医疗集团的患者监护设备中发现安全漏洞</a>
						<span>2020/02/25</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./d94a036027f6461ea78ca066c8799e61.html" target="_blank">新的勒索软件EKANS 针对工业控制系统</a>
						<span>2020/02/14</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./032d80a188fc4f81b31face3be8e3586.html" target="_blank">思科底层协议爆5个零日漏洞，旗下“几乎所有”设备都中招</a>
						<span>2020/02/11</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./8b7096add9794b93a90ae1ed955236ea.html" target="_blank">威胁情报：网络安全的下一个引爆点</a>
						<span>2020/01/09</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./1a31d9798e6d4e0d8801d88b86f7a644.html" target="_blank">西门子SPPA-T3000工控系统爆出致命漏洞且未完全修复，全球电厂或再遭劫难！</a>
						<span>2019/12/17</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./054d25bd77af4a1cb69a082396046867.html" target="_blank">国家级黑客组织“海莲花” 攻陷宝马，汽车工业成APT新目标</a>
						<span>2019/12/10</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./278a3dc9170a47af86be1901a1bf101e.html" target="_blank">英国核发电厂遭受网络攻击，疑似法国电力公司受影响</a>
						<span>2019/12/04</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./e0c2891ce6b948f5b291d68d2e3ed83d.html" target="_blank">利用震网三代和某PLC漏洞组合攻击工控系统</a>
						<span>2019/11/29</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./100255baac844f4a991502792e33644a.html" target="_blank">工控系统再迎大波澜，伊朗APT组织将其作为重点攻击目标</a>
						<span>2019/11/27</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./5ac7447ded5b4b41b81762e8904d20de.html" target="_blank">入侵俄罗斯铁路信息系统只需要20分钟？</a>
						<span>2019/11/22</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./7b4dd01cae534c4e9e3acb4ccb294932.html" target="_blank">Attack Surface | 描绘机场信息系统攻击面</a>
						<span>2019/11/19</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./5c9878b2b793481f88b1bdadd23fec62.html" target="_blank">亚马逊Ring智能门锁漏洞可窃取WiFi密码</a>
						<span>2019/11/14</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./6c90cc6926a04e19936e8ba66748503e.html" target="_blank">核心工业系统陷入危机？印度核电厂遭受网络攻击事件梳理与分析</a>
						<span>2019/10/31</span>
					</li>
					<li>
						<img src="../../baseImages/right.png">
							<a href="./a3530b7e1edc4ee9abee4d0e1a1a287c.html" target="_blank">首次针对继电保护装置的网络攻击——2016年乌克兰电力事件再分析</a>
						<span>2019/10/15</span>
					</li>
		    </ul>
        </ul>
    </div>
    <div class="page">
        <ul>
			             <li class="active"><a href="javascript:;">1</a></li>
			             <li><a href="./index_2.html">2</a></li>
			             <li><a href="./index_3.html">3</a></li>
			             <li><a href="./index_4.html">4</a></li>
			             <li><a href="./index_5.html">5</a></li>
			             <li><a href="./index_6.html">6</a></li>
			             <li><a href="./index_7.html">7</a></li>
        			<li><a href="./index_2.html" title="下一页">»</a></li>
                    <li><a href="./index_7.html">末页</a></li>
        </ul>
    </div>
</div>
<footer id="footer">
    <div class="f-top">
        <a href="#">关于我们</a><span class="v-line">|</span>
        <a href="#">网站声明</a><span class="v-line">|</span>
        <a href="https://www.ics-cert.org.cn/portal/map.html" target="_blank">网站地图</a><span class="v-line">|</span>
        <a href="#">联系我们</a><span class="v-line">|</span>
        <a href="#">版权声明</a>
    </div>
    <div class="f-bottom">
        <p>国家计算机网络应急技术处理协调中心版权所有</p>
        <p>Email:ics-cert@cert.org.cn&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;京ICP备10012421号-3</p>
    </div>
</footer>
<script src="https://www.ics-cert.org.cn/portal/js/jQuery/jQuery-2.1.4.min.js"></script>
<script src="https://www.ics-cert.org.cn/portal/js/bootstrap/bootstrap.min.js"></script>
<script src="https://www.ics-cert.org.cn/portal/js/common.js?t=0339286032string[" hhmmsssss"]}"=""></script>
</body></html>`
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(dom))

	//doc.Find(".info li").Each(func(i int, s *goquery.Selection) {
	//	var tmp = model.DataInfo{};
	//	tmp.Url, _= s.Find("a").Attr("href");
	//	tmp.Title = s.Find("a").Text();
	//	tmp.Date =  s.Find("span").Text();
	//	datas = append(datas, tmp);
	//})
	var vm = otto.New()
	//vm.Set("ctx", ctx)
	_ = vm.Set("doc", doc)
	//_ = vm.Set("datas", datas)
	_, err := vm.Run(`
		var s = doc.Find(".info li");
			console.log(doc.Length());
`)
	if err != nil {
		t.Log(err.Error())
	}
}

type TT struct{
	Name string
	Age int
}
