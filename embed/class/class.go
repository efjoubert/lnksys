package class

import (
	"io"
	"strings"
)

const classjs string=`/*!
* @author      Angelo Dini
* @version     1.0
* @copyright   Distributed under the BSD License.
*/
(function(){var d="1.0";var c=window.Class;var b=window.Class=function(n){n=n||{};var m=function(){return(this.initialize)?this.initialize.apply(this,arguments):j};if(n.implement){var j=window===this?g(m.prototype):this;var l=n.implement;a(n,"implement");n=f(n,e(l))}m.prototype=g(n);m.constructor=m;m._parent=g(n);for(var k=0,h=["extend","implement","getOptions","setOptions"];k<h.length;k++){m[h[k]]=b[h[k]]}return m};b.extend=function(j){var h=this;if(j.implement){this.prototype=f(this.prototype,e(j.implement));a(j,"implement")}for(var i in j){j[i]=typeof j[i]==="function"&&/parent/.test(j[i].toString())?(function(l,k){return function(){this.parent=h._parent[k];return l.apply(this,arguments)}})(j[i],i):j[i]}this._parent=f(this._parent,j,true);this.prototype=f(this.prototype,j);return this};b.implement=function(h){return this.prototype=f(this.prototype,e(h))};b.getOptions=function(){return this.prototype.options||{}};b.setOptions=function(h){return this.prototype.options=f(this.prototype.options,h)};b.noConflict=function(){window.Class=c;return b};b.version=d;function g(i){var h=function(){};h.prototype=i.prototype||i;return new h()}function a(l,i,k){if(k){var h={};for(var j in l){if(j!==i){h[j]=l[j]}}}else{delete l[i]}return h||l}function f(h,i,k){if(!h||!i){return h||i||{}}h=g(h);i=g(i);for(var j in i){if(Object.prototype.toString.call(i[j])==="[object Object]"){f(h[j],i[j])}else{h[j]=(k&&h[j])?h[j]:i[j]}}return h}function e(l){var k={};for(var h=0;h<l.length;h++){if(typeof(l[h])==="function"){l[h]=l[h].prototype}var j=a(l[h],"initialize",true);if(j.implement){k=e(j.implement)}else{k=f(k,j)}}return k}})();`

func ClassJS() io.Reader {
	return strings.NewReader(classjs);
}

func ClassFindJS(findclassjs string) io.Reader {
	if strings.LastIndex(findclassjs,"/")>-1 {
		findclassjs=findclassjs[strings.LastIndex(findclassjs,"/")+1:]
	}
	if findclassjs=="class.js" {
		return ClassJS()
	}
	return nil
}