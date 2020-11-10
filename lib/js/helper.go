// generated by "lib/js/generate"

package js

var Element = &Function{
	Name:         "element",
	Definition:   `function(e){return functions.selectable(this).querySelector(e)}`,
	Dependencies: []*Function{Selectable},
}

var Elements = &Function{
	Name:         "elements",
	Definition:   `function(e){return functions.selectable(this).querySelectorAll(e)}`,
	Dependencies: []*Function{Selectable},
}

var ElementX = &Function{
	Name:         "elementX",
	Definition:   `function(e){const t=functions.selectable(this);return document.evaluate(e,t,null,XPathResult.FIRST_ORDERED_NODE_TYPE).singleNodeValue}`,
	Dependencies: []*Function{Selectable},
}

var ElementsX = &Function{
	Name:         "elementsX",
	Definition:   `function(e){const t=functions.selectable(this),n=document.evaluate(e,t,null,XPathResult.ORDERED_NODE_ITERATOR_TYPE),i=[];let s;for(;s=n.iterateNext();)i.push(s);return i}`,
	Dependencies: []*Function{Selectable},
}

var ElementR = &Function{
	Name:         "elementR",
	Definition:   `function(e,t){const n=new RegExp(t),i=Array.from((this.document||this).querySelectorAll(e)).find(e=>n.test(functions.text.call(e)));return i||null}`,
	Dependencies: []*Function{Text},
}

var Parents = &Function{
	Name:         "parents",
	Definition:   `function(e){let t=this.parentElement;const n=[];for(;t;)t.matches(e)&&n.push(t),t=t.parentElement;return n}`,
	Dependencies: []*Function{},
}

var ContainsElement = &Function{
	Name:         "containsElement",
	Definition:   `function(e){for(var t=e;null!=t;){if(t===this)return!0;t=t.parentElement}return!1}`,
	Dependencies: []*Function{},
}

var InitMouseTracer = &Function{
	Name:         "initMouseTracer",
	Definition:   `async function(e,t){if(await functions.waitLoad(),document.getElementById(e))return;const n=document.createElement("div");n.innerHTML=t;const i=n.lastChild;i.id=e,i.style="position: absolute; z-index: 2147483647; width: 17px; pointer-events: none;",i.removeAttribute("width"),i.removeAttribute("height"),document.body.appendChild(i)}`,
	Dependencies: []*Function{WaitLoad},
}

var UpdateMouseTracer = &Function{
	Name:         "updateMouseTracer",
	Definition:   `function(e,t,n){const i=document.getElementById(e);return!!i&&(i.style.left=t-2+"px",i.style.top=n-3+"px",!0)}`,
	Dependencies: []*Function{},
}

var Rect = &Function{
	Name:         "rect",
	Definition:   `function(){const e=functions.tag(this).getBoundingClientRect();return{x:e.x,y:e.y,width:e.width,height:e.height}}`,
	Dependencies: []*Function{Tag},
}

var Overlay = &Function{
	Name:         "overlay",
	Definition:   `async function(e,t,n,i,s,o){await functions.waitLoad();const r=document.createElement("div");if(r.id=e,r.style=` + "`" + `position: fixed; z-index:2147483647; border: 2px dashed red;\n        border-radius: 3px; box-shadow: #5f3232 0 0 3px; pointer-events: none;\n        box-sizing: border-box;\n        left: ${t}px;\n        top: ${n}px;\n        height: ${s}px;\n        width: ${i}px;` + "`" + `,i*s==0&&(r.style.border="none"),!o)return void document.body.appendChild(r);const l=document.createElement("div");l.style=` + "`" + `position: absolute; color: #cc26d6; font-size: 12px; background: #ffffffeb;\n        box-shadow: #333 0 0 3px; padding: 2px 5px; border-radius: 3px; white-space: nowrap;\n        top: ${s}px;` + "`" + `,l.innerHTML=o,r.appendChild(l),document.body.appendChild(r),window.innerHeight<l.offsetHeight+n+s&&(l.style.top=-l.offsetHeight-2+"px"),window.innerWidth<l.offsetWidth+t&&(l.style.left=window.innerWidth-l.offsetWidth-t+"px")}`,
	Dependencies: []*Function{WaitLoad},
}

var ElementOverlay = &Function{
	Name:         "elementOverlay",
	Definition:   `async function(e,t){const n=functions.tag(this);let i=n.getBoundingClientRect();await functions.overlay(e,i.left,i.top,i.width,i.height,t);const s=()=>{const t=document.getElementById(e);if(null===t)return;const o=n.getBoundingClientRect();i.left!==o.left||i.top!==o.top||i.width!==o.width||i.height!==o.height?(t.style.left=o.left+"px",t.style.top=o.top+"px",t.style.width=o.width+"px",t.style.height=o.height+"px",i=o,setTimeout(s,100)):setTimeout(s,100)};setTimeout(s,100)}`,
	Dependencies: []*Function{Tag, Overlay},
}

var RemoveOverlay = &Function{
	Name:         "removeOverlay",
	Definition:   `function(e){const t=document.getElementById(e);t&&t.remove()}`,
	Dependencies: []*Function{},
}

var WaitIdle = &Function{
	Name:         "waitIdle",
	Definition:   `e=>new Promise(t=>{window.requestIdleCallback(t,{timeout:e})})`,
	Dependencies: []*Function{},
}

var WaitLoad = &Function{
	Name:         "waitLoad",
	Definition:   `function(){const e=this===window;return new Promise((t,n)=>{if(e){if("complete"===document.readyState)return t();window.addEventListener("load",t)}else void 0===this.complete||this.complete?t():(this.addEventListener("load",t),this.addEventListener("error",n))})}`,
	Dependencies: []*Function{},
}

var InputEvent = &Function{
	Name:         "inputEvent",
	Definition:   `function(){this.dispatchEvent(new Event("input",{bubbles:!0})),this.dispatchEvent(new Event("change",{bubbles:!0}))}`,
	Dependencies: []*Function{},
}

var InputTime = &Function{
	Name:         "inputTime",
	Definition:   `function(e){const t=new Date(e),n=e=>e.toString().padStart(2,"0"),i=t.getFullYear(),s=n(t.getMonth()+1),o=n(t.getDate()),r=n(t.getHours()),l=n(t.getMinutes());switch(this.type){case"date":this.value=` + "`" + `${i}-${s}-${o}` + "`" + `;break;case"datetime-local":this.value=` + "`" + `${i}-${s}-${o}T${r}:${l}` + "`" + `;break;case"month":this.value=s;break;case"time":this.value=` + "`" + `${r}:${l}` + "`" + `}functions.inputEvent.call(this)}`,
	Dependencies: []*Function{InputEvent},
}

var SelectText = &Function{
	Name:         "selectText",
	Definition:   `function(e){const t=this.value.match(new RegExp(e));t&&this.setSelectionRange(t.index,t.index+t[0].length)}`,
	Dependencies: []*Function{},
}

var SelectAllText = &Function{
	Name:         "selectAllText",
	Definition:   `function(){this.select()}`,
	Dependencies: []*Function{},
}

var Select = &Function{
	Name:         "select",
	Definition:   `function(e,t,n){let i;switch(n){case"regex":i=e.map(e=>{const t=new RegExp(e);return e=>t.test(e.innerText)});break;case"css-selector":i=e.map(e=>t=>t.matches(e));break;default:i=e.map(e=>t=>t.innerText.includes(e))}const s=Array.from(this.options);i.forEach(e=>{const n=s.find(e);n&&(n.selected=t)}),this.dispatchEvent(new Event("input",{bubbles:!0})),this.dispatchEvent(new Event("change",{bubbles:!0}))}`,
	Dependencies: []*Function{},
}

var Visible = &Function{
	Name:         "visible",
	Definition:   `function(){const e=functions.tag(this),t=e.getBoundingClientRect(),n=window.getComputedStyle(e);return"none"!==n.display&&"hidden"!==n.visibility&&!!(t.top||t.bottom||t.width||t.height)}`,
	Dependencies: []*Function{Tag},
}

var Invisible = &Function{
	Name:         "invisible",
	Definition:   `function(){return!functions.visible.apply(this)}`,
	Dependencies: []*Function{Visible},
}

var Text = &Function{
	Name:         "text",
	Definition:   `function(){switch(this.tagName){case"INPUT":case"TEXTAREA":return this.value||this.placeholder;case"SELECT":return Array.from(this.selectedOptions).map(e=>e.innerText).join();case void 0:return this.textContent;default:return this.innerText}}`,
	Dependencies: []*Function{},
}

var Resource = &Function{
	Name:         "resource",
	Definition:   `function(){return new Promise((e,t)=>{if(this.complete)return e(this.currentSrc);this.addEventListener("load",()=>e(this.currentSrc)),this.addEventListener("error",e=>t(e))})}`,
	Dependencies: []*Function{},
}

var AddScriptTag = &Function{
	Name:         "addScriptTag",
	Definition:   `function(e,t,n){if(!document.getElementById(e))return new Promise((i,s)=>{var o=document.createElement("script");t?(o.src=t,o.onload=i):(o.type="text/javascript",o.text=n,i()),o.id=e,o.onerror=s,document.head.appendChild(o)})}`,
	Dependencies: []*Function{},
}

var AddStyleTag = &Function{
	Name:         "addStyleTag",
	Definition:   `function(e,t,n){if(!document.getElementById(e))return new Promise((i,s)=>{var o;t?((o=document.createElement("link")).rel="stylesheet",o.href=t):((o=document.createElement("style")).type="text/css",o.appendChild(document.createTextNode(n)),i()),o.id=e,o.onload=i,o.onerror=s,document.head.appendChild(o)})}`,
	Dependencies: []*Function{},
}

var FetchAsDataURL = &Function{
	Name:         "fetchAsDataURL",
	Definition:   `e=>fetch(e).then(e=>e.blob()).then(e=>new Promise((t,n)=>{var i=new FileReader;i.onload=(()=>t(i.result)),i.onerror=(()=>n(i.error)),i.readAsDataURL(e)}))`,
	Dependencies: []*Function{},
}

var Selectable = &Function{
	Name:         "selectable",
	Definition:   `e=>e===window?e.document:e`,
	Dependencies: []*Function{},
}

var Tag = &Function{
	Name:         "tag",
	Definition:   `e=>e.tagName?e:e.parentElement`,
	Dependencies: []*Function{},
}

var ExposeFunc = &Function{
	Name:         "exposeFunc",
	Definition:   `function(e,t){let n=0;window[e]=(e=>new Promise((i,s)=>{const o=t+"_cb"+n++;window[o]=((e,t)=>{delete window[o],t?s(t):i(e)}),window[t](JSON.stringify({req:e,cb:o}))}))}`,
	Dependencies: []*Function{},
}
