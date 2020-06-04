package assets

// Template provides as simple html page
const Template = `
<html>
<head>
<style type="text/css">
body{
	font-family: "Droid Sans", Arial, sans-serif;
	padding: 0;
	margin: 0;
}
.card{
	border: 1px solid #ccc;
	border-radius: 2px;
	min-width: 153px;
	max-width: 90vw;
	overflow-x: auto;
}

.card-body{
    min-width: 250px;
}
.actions{
	white-space: nowrap;
	position: absolute;
	top: -30px;
	margin-left: -1px;
	
}


.btn{
	text-decoration: none;
	color: #666;
	border: 1px solid #ccc;
	border-radius: 2px;
	padding: 5px;
	font-size: 12px;
	cursor: pointer;
	background-color: transparent;

}
#content{
	color: #555;
	padding: 10px;
	margin: 0;
}

.header{
	padding: 0;
	border-bottom: 2px solid #CCC;
	margin-bottom: 40px;
	text-align: center;
	min-height: 35px;
    font-size: 65px;
    color: #777;
}
.header img{
	max-height: 75px;
}
.header-image{
	display: inline-block;
	overflow-y: hidden;
}

.header-text{
	vertical-align: top;
}
.github{
	position: absolute;
    left: 10px;
	top: 10px;
}
.github img{
    height: 20px;
	opacity: 0.5;
}


</style>
</head>
<body>

<div class="github">
	<a target="_blank" href="https://github.com/itsy-sh/is" rel="nofollow">
		<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyRpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuMy1jMDExIDY2LjE0NTY2MSwgMjAxMi8wMi8wNi0xNDo1NjoyNyAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENTNiAoTWFjaW50b3NoKSIgeG1wTU06SW5zdGFuY2VJRD0ieG1wLmlpZDpFNTE3OEEyRTk5QTAxMUUyOUExNUJDMTA0NkE4OTA0RCIgeG1wTU06RG9jdW1lbnRJRD0ieG1wLmRpZDpFNTE3OEEyRjk5QTAxMUUyOUExNUJDMTA0NkE4OTA0RCI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAuaWlkOkU1MTc4QTJDOTlBMDExRTI5QTE1QkMxMDQ2QTg5MDREIiBzdFJlZjpkb2N1bWVudElEPSJ4bXAuZGlkOkU1MTc4QTJEOTlBMDExRTI5QTE1QkMxMDQ2QTg5MDREIi8+IDwvcmRmOkRlc2NyaXB0aW9uPiA8L3JkZjpSREY+IDwveDp4bXBtZXRhPiA8P3hwYWNrZXQgZW5kPSJyIj8+FYrpWAAABrNJREFUeNrkW2lsVFUUvjMWirYUkS5BXApUa2vd6gL+wAWjoP5RiW2EUBajAiqSuPADQ0w1UUQTrcFAUUSJEKriEuMWFKuJIElFSS24YNpQK6WoBbuAktbva880M8O8vnfevJm+CSf5cme599xzvnfffffce17AJFjycnLzUVwDXAgUAucBY4BMIEOqdQIdwJ/Az4J64OvWtoONibQvkACHgyiuBe4CbgLOjVNlE/AZsAmoBSE9viQAjueieBCYC5yVoAvWDKwHqkBEmy8IgON09lHgXmCESY4cBaqBlSCieUgIgOPDUCwBngBOM0MjXdL/CyDiv6QRAOcvR7EBKDL+kD3AbJBQl1AC4DjrLwaeBYYbf8m/ciu+BCJ6PScAzp+K4nXgTuNveQuYAxK6PSMAzo9C8TFwtUkN2Q7cDBIOx02AOP8FUGpSSzgf3GBHQsDGec7unwOTTWrKDiGhS02ATHjvALeb1JZ3gRlWE+MpVq0yMzIekRk/1YWP6o7Ors5vHI8AXH1Odl8BaTbKrwd4j10MTAduS8JqkKvA94BPgN0A56htNm2OMyDDKNhuSwCcT5dIrMBG6S4oLI1qezqKBcBjwGiPHW8HVgCr0W97VL/fobjMpv2vQAnaHgv/MdYVXurAeSNPhggRw56BQatRVgL3A0H5+xDwI8Dw9g/5Hlq+clmdDYwF8iV0zpb/GP2tApZHOx4m2xwQUCC+VVqOABg+AUUDkO6AgHkwaL2DJXORxPVNylUnw+gpXObaLXFRlxHoaw7U8uoXQ99vViNgqUPnKQfsKojhdW7GuxDW5JUtIuni432hH4JhLJ7Dq6qwcZiPZnpNXDJPfI0kQEJbjVM5PiIgW3nhlkQQILH9LGWnV/iIAK0ts8TngREwDchVKrnKRwRobckVnwcIKFcq4ONrkY8IWBT2SHUq5eEE3Khs/CRm6Z1+8V5sqVQ26/M5gHuhSJ79TqUFmIhOj/ppwQ8/Rshqb5yiWXFQFhsaWeU352UU0KaXlc2mBI1+Y3OzjyO/Gm2kSAIKFQ2awfQ+v3oP23gL/K5oUhh0GPiEZG8KxP97FHULgsqwtTUFCDioqHsGCRipaHA8BQjQrAcyg4roj5KVAgSMUtRNDyqVj0wBAlQ2koBuRf3xKUBAvqJuN1eCrYpAiHNAltNjpyFYDfL47oix38wdmDA5AvYr+kjzWRgcLVcqnKfsJwGNyk5u9TEBtyjrNwaVgRClTPKA/Db8aVOZslkDG2nD2vEuOkqGlLmYpHcGJLlJu8LjtvJFgx06Jvnq8xC33gUBeUE4waWjduua5wdVPrr6VS6cr6PvoXv5Ixed3g3mH/fB1V9OW1w07fM5IEouUEZR4bIWWJzsTRJ55r8I3ONSRRFs3hsIU8hkgkkulf0CPAx8qElQcuk4beYp9Epgoks138LOvqSPgfyAzIwMZlnFSobgIegc4H3gH6AkxmKDub9Mjb0DeoYDrZ1dne0eO14AvfPx8RXgAYaycahbBvt+GLgFpIM0md3PjqrMTMxpYKxB6p1v+s/n7bbSuMCqldmZyc+fRh9ND+IsAxrmG3C3qtj0J1uP84hLrnwnwJbjEQRIxzw0XB2jER93C9Bog9TjsRgzLpzuJr0BzHV6e8gwf9XoziqdCv1YE/oSTQBHwfem/3w+5syPxuukLtfdO0zk+WIs+YuPKLQ7ohzyWTIix3joPPMTLg1d/Yg5gIL7ogf32U/4WGGhYDr+34J6bUALPpPA62w6XYMOP9BaCv3HoD/PeJubODN6U/eEq4cKTIurttpBAZ4L+87TmKdtOt0ah8FbPXS+WnyLEKskqUy5FaweM5dA2e6w+pNkZuajhfMD3/zYBfDKb3Y6+cWwgytOL7bh98nQ73BEgHReIvd4Roy/a6Cs3CRYJOnq7zjV8HWcybC33mpLLKZIA84FPRYhcSokUNL2Civnjd0MjoZbUCy0+PtNkDDD5wQsFB8sxWm2+GJZd8eSt4HnZXnZ66Nb4CHYYxuxat4XmI1inbHeczskq77DMrK4z8AgK3+Q/L5EEMBn/PzQos0zAsQgvg5XY3TpNKOTSAD3NsrQX63TBqq9PVHM9NgvfXi/06ZSjfNqAoQEHj9Pled+pw8cpw2co6aKbSoJxDlJnYniKdP/sqSVrrEw7IBL/TnG+rSXEy7fYVoG/S1uffDkzVEYypB1qewJRCdb5rp9yxN6mQDZFmOS2wisCIXo8Yin7w7LiKiQEcFYfhOMnBmnzo1CLIO09Qyt47niJxDQ29trTmY56Qn4X4ABAFR7IoDmVT5NAAAAAElFTkSuQmCC">
	</a>
</div>
<div class="header">
	<span class="header-text">
		Itsy Share
	</span>
	<div class="header-image">
	<img src="data:image/png;base64,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">
	</div>
</div>

	<div style="text-align: center; position: relative">
		
		<div class="card" style="display: inline-block; text-align: left; ">
		<div class="actions">
			{{if index . "ascii"}}
				<button id="copy"  
				class="btn" 
				data-clipboard-action="copy" 
				data-clipboard-target="#content">
					Copy
				</button>
				
			{{end}}
			<a href="/download?key={{index . "key"}}" class="btn">Download</a>
			
		</div>
		  <div class="card-body">
{{$key := index . "key"}}
{{$filesizes := index . "filesizes"}}
{{if index . "ascii"}}
<pre id="content">{{ index . "content"}}</pre>
{{else}}
<pre id="content">
{{range $i, $path := index . "filepaths"}}<div style="width: 100px; display: inline-block">({{index $filesizes $i}})</div><a href="/download?key={{$key}}&file={{$path}}">{{filename $path}}</a>   
{{end}}
</pre>
{{end}}
		  </div>
		</div>
    </div>
</div>



</body>
<script>
	{{index . "clipboard"}}
</script>
<script>
	new ClipboardJS('#copy');
</script>

</html>
`

// JSClipboard is a js lib for copying text from browser to clipboard
const JSClipboard = `/*!
 * clipboard.js v2.0.4
 * https://zenorocha.github.io/clipboard.js
 * 
 * Licensed MIT © Zeno Rocha
 */
!function(t,e){"object"==typeof exports&&"object"==typeof module?module.exports=e():"function"==typeof define&&define.amd?define([],e):"object"==typeof exports?exports.ClipboardJS=e():t.ClipboardJS=e()}(this,function(){return function(n){var o={};function r(t){if(o[t])return o[t].exports;var e=o[t]={i:t,l:!1,exports:{}};return n[t].call(e.exports,e,e.exports,r),e.l=!0,e.exports}return r.m=n,r.c=o,r.d=function(t,e,n){r.o(t,e)||Object.defineProperty(t,e,{enumerable:!0,get:n})},r.r=function(t){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(t,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(t,"__esModule",{value:!0})},r.t=function(e,t){if(1&t&&(e=r(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(r.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var o in e)r.d(n,o,function(t){return e[t]}.bind(null,o));return n},r.n=function(t){var e=t&&t.__esModule?function(){return t.default}:function(){return t};return r.d(e,"a",e),e},r.o=function(t,e){return Object.prototype.hasOwnProperty.call(t,e)},r.p="",r(r.s=0)}([function(t,e,n){"use strict";var r="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t},i=function(){function o(t,e){for(var n=0;n<e.length;n++){var o=e[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(t,o.key,o)}}return function(t,e,n){return e&&o(t.prototype,e),n&&o(t,n),t}}(),a=o(n(1)),c=o(n(3)),u=o(n(4));function o(t){return t&&t.__esModule?t:{default:t}}var l=function(t){function o(t,e){!function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}(this,o);var n=function(t,e){if(!t)throw new ReferenceError("this hasn't been initialised - super() hasn't been called");return!e||"object"!=typeof e&&"function"!=typeof e?t:e}(this,(o.__proto__||Object.getPrototypeOf(o)).call(this));return n.resolveOptions(e),n.listenClick(t),n}return function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Super expression must either be null or a function, not "+typeof e);t.prototype=Object.create(e&&e.prototype,{constructor:{value:t,enumerable:!1,writable:!0,configurable:!0}}),e&&(Object.setPrototypeOf?Object.setPrototypeOf(t,e):t.__proto__=e)}(o,c.default),i(o,[{key:"resolveOptions",value:function(){var t=0<arguments.length&&void 0!==arguments[0]?arguments[0]:{};this.action="function"==typeof t.action?t.action:this.defaultAction,this.target="function"==typeof t.target?t.target:this.defaultTarget,this.text="function"==typeof t.text?t.text:this.defaultText,this.container="object"===r(t.container)?t.container:document.body}},{key:"listenClick",value:function(t){var e=this;this.listener=(0,u.default)(t,"click",function(t){return e.onClick(t)})}},{key:"onClick",value:function(t){var e=t.delegateTarget||t.currentTarget;this.clipboardAction&&(this.clipboardAction=null),this.clipboardAction=new a.default({action:this.action(e),target:this.target(e),text:this.text(e),container:this.container,trigger:e,emitter:this})}},{key:"defaultAction",value:function(t){return s("action",t)}},{key:"defaultTarget",value:function(t){var e=s("target",t);if(e)return document.querySelector(e)}},{key:"defaultText",value:function(t){return s("text",t)}},{key:"destroy",value:function(){this.listener.destroy(),this.clipboardAction&&(this.clipboardAction.destroy(),this.clipboardAction=null)}}],[{key:"isSupported",value:function(){var t=0<arguments.length&&void 0!==arguments[0]?arguments[0]:["copy","cut"],e="string"==typeof t?[t]:t,n=!!document.queryCommandSupported;return e.forEach(function(t){n=n&&!!document.queryCommandSupported(t)}),n}}]),o}();function s(t,e){var n="data-clipboard-"+t;if(e.hasAttribute(n))return e.getAttribute(n)}t.exports=l},function(t,e,n){"use strict";var o,r="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t},i=function(){function o(t,e){for(var n=0;n<e.length;n++){var o=e[n];o.enumerable=o.enumerable||!1,o.configurable=!0,"value"in o&&(o.writable=!0),Object.defineProperty(t,o.key,o)}}return function(t,e,n){return e&&o(t.prototype,e),n&&o(t,n),t}}(),a=n(2),c=(o=a)&&o.__esModule?o:{default:o};var u=function(){function e(t){!function(t,e){if(!(t instanceof e))throw new TypeError("Cannot call a class as a function")}(this,e),this.resolveOptions(t),this.initSelection()}return i(e,[{key:"resolveOptions",value:function(){var t=0<arguments.length&&void 0!==arguments[0]?arguments[0]:{};this.action=t.action,this.container=t.container,this.emitter=t.emitter,this.target=t.target,this.text=t.text,this.trigger=t.trigger,this.selectedText=""}},{key:"initSelection",value:function(){this.text?this.selectFake():this.target&&this.selectTarget()}},{key:"selectFake",value:function(){var t=this,e="rtl"==document.documentElement.getAttribute("dir");this.removeFake(),this.fakeHandlerCallback=function(){return t.removeFake()},this.fakeHandler=this.container.addEventListener("click",this.fakeHandlerCallback)||!0,this.fakeElem=document.createElement("textarea"),this.fakeElem.style.fontSize="12pt",this.fakeElem.style.border="0",this.fakeElem.style.padding="0",this.fakeElem.style.margin="0",this.fakeElem.style.position="absolute",this.fakeElem.style[e?"right":"left"]="-9999px";var n=window.pageYOffset||document.documentElement.scrollTop;this.fakeElem.style.top=n+"px",this.fakeElem.setAttribute("readonly",""),this.fakeElem.value=this.text,this.container.appendChild(this.fakeElem),this.selectedText=(0,c.default)(this.fakeElem),this.copyText()}},{key:"removeFake",value:function(){this.fakeHandler&&(this.container.removeEventListener("click",this.fakeHandlerCallback),this.fakeHandler=null,this.fakeHandlerCallback=null),this.fakeElem&&(this.container.removeChild(this.fakeElem),this.fakeElem=null)}},{key:"selectTarget",value:function(){this.selectedText=(0,c.default)(this.target),this.copyText()}},{key:"copyText",value:function(){var e=void 0;try{e=document.execCommand(this.action)}catch(t){e=!1}this.handleResult(e)}},{key:"handleResult",value:function(t){this.emitter.emit(t?"success":"error",{action:this.action,text:this.selectedText,trigger:this.trigger,clearSelection:this.clearSelection.bind(this)})}},{key:"clearSelection",value:function(){this.trigger&&this.trigger.focus(),window.getSelection().removeAllRanges()}},{key:"destroy",value:function(){this.removeFake()}},{key:"action",set:function(){var t=0<arguments.length&&void 0!==arguments[0]?arguments[0]:"copy";if(this._action=t,"copy"!==this._action&&"cut"!==this._action)throw new Error('Invalid "action" value, use either "copy" or "cut"')},get:function(){return this._action}},{key:"target",set:function(t){if(void 0!==t){if(!t||"object"!==(void 0===t?"undefined":r(t))||1!==t.nodeType)throw new Error('Invalid "target" value, use a valid Element');if("copy"===this.action&&t.hasAttribute("disabled"))throw new Error('Invalid "target" attribute. Please use "readonly" instead of "disabled" attribute');if("cut"===this.action&&(t.hasAttribute("readonly")||t.hasAttribute("disabled")))throw new Error('Invalid "target" attribute. You can\'t cut text from elements with "readonly" or "disabled" attributes');this._target=t}},get:function(){return this._target}}]),e}();t.exports=u},function(t,e){t.exports=function(t){var e;if("SELECT"===t.nodeName)t.focus(),e=t.value;else if("INPUT"===t.nodeName||"TEXTAREA"===t.nodeName){var n=t.hasAttribute("readonly");n||t.setAttribute("readonly",""),t.select(),t.setSelectionRange(0,t.value.length),n||t.removeAttribute("readonly"),e=t.value}else{t.hasAttribute("contenteditable")&&t.focus();var o=window.getSelection(),r=document.createRange();r.selectNodeContents(t),o.removeAllRanges(),o.addRange(r),e=o.toString()}return e}},function(t,e){function n(){}n.prototype={on:function(t,e,n){var o=this.e||(this.e={});return(o[t]||(o[t]=[])).push({fn:e,ctx:n}),this},once:function(t,e,n){var o=this;function r(){o.off(t,r),e.apply(n,arguments)}return r._=e,this.on(t,r,n)},emit:function(t){for(var e=[].slice.call(arguments,1),n=((this.e||(this.e={}))[t]||[]).slice(),o=0,r=n.length;o<r;o++)n[o].fn.apply(n[o].ctx,e);return this},off:function(t,e){var n=this.e||(this.e={}),o=n[t],r=[];if(o&&e)for(var i=0,a=o.length;i<a;i++)o[i].fn!==e&&o[i].fn._!==e&&r.push(o[i]);return r.length?n[t]=r:delete n[t],this}},t.exports=n},function(t,e,n){var d=n(5),h=n(6);t.exports=function(t,e,n){if(!t&&!e&&!n)throw new Error("Missing required arguments");if(!d.string(e))throw new TypeError("Second argument must be a String");if(!d.fn(n))throw new TypeError("Third argument must be a Function");if(d.node(t))return s=e,f=n,(l=t).addEventListener(s,f),{destroy:function(){l.removeEventListener(s,f)}};if(d.nodeList(t))return a=t,c=e,u=n,Array.prototype.forEach.call(a,function(t){t.addEventListener(c,u)}),{destroy:function(){Array.prototype.forEach.call(a,function(t){t.removeEventListener(c,u)})}};if(d.string(t))return o=t,r=e,i=n,h(document.body,o,r,i);throw new TypeError("First argument must be a String, HTMLElement, HTMLCollection, or NodeList");var o,r,i,a,c,u,l,s,f}},function(t,n){n.node=function(t){return void 0!==t&&t instanceof HTMLElement&&1===t.nodeType},n.nodeList=function(t){var e=Object.prototype.toString.call(t);return void 0!==t&&("[object NodeList]"===e||"[object HTMLCollection]"===e)&&"length"in t&&(0===t.length||n.node(t[0]))},n.string=function(t){return"string"==typeof t||t instanceof String},n.fn=function(t){return"[object Function]"===Object.prototype.toString.call(t)}},function(t,e,n){var a=n(7);function i(t,e,n,o,r){var i=function(e,n,t,o){return function(t){t.delegateTarget=a(t.target,n),t.delegateTarget&&o.call(e,t)}}.apply(this,arguments);return t.addEventListener(n,i,r),{destroy:function(){t.removeEventListener(n,i,r)}}}t.exports=function(t,e,n,o,r){return"function"==typeof t.addEventListener?i.apply(null,arguments):"function"==typeof n?i.bind(null,document).apply(null,arguments):("string"==typeof t&&(t=document.querySelectorAll(t)),Array.prototype.map.call(t,function(t){return i(t,e,n,o,r)}))}},function(t,e){if("undefined"!=typeof Element&&!Element.prototype.matches){var n=Element.prototype;n.matches=n.matchesSelector||n.mozMatchesSelector||n.msMatchesSelector||n.oMatchesSelector||n.webkitMatchesSelector}t.exports=function(t,e){for(;t&&9!==t.nodeType;){if("function"==typeof t.matches&&t.matches(e))return t;t=t.parentNode}}}])});`
