var rt=Object.defineProperty,st=Object.defineProperties;var ct=Object.getOwnPropertyDescriptors;var F=Object.getOwnPropertySymbols;var lt=Object.prototype.hasOwnProperty,ft=Object.prototype.propertyIsEnumerable;var B=(t,n,e)=>n in t?rt(t,n,{enumerable:!0,configurable:!0,writable:!0,value:e}):t[n]=e,p=(t,n)=>{for(var e in n||(n={}))lt.call(n,e)&&B(t,e,n[e]);if(F)for(var e of F(n))ft.call(n,e)&&B(t,e,n[e]);return t},C=(t,n)=>st(t,ct(n));function $(t){return t.split("-")[0]}function q(t){return t.split("-")[1]}function S(t){return["top","bottom"].includes($(t))?"x":"y"}function G(t){return t==="y"?"height":"width"}function I(t,n,e){let{reference:o,floating:i}=t;const l=o.x+o.width/2-i.width/2,c=o.y+o.height/2-i.height/2,r=S(n),s=G(r),f=o[s]/2-i[s]/2,d=$(n),u=r==="x";let a;switch(d){case"top":a={x:l,y:o.y-i.height};break;case"bottom":a={x:l,y:o.y+o.height};break;case"right":a={x:o.x+o.width,y:c};break;case"left":a={x:o.x-i.width,y:c};break;default:a={x:o.x,y:o.y}}switch(q(n)){case"start":a[r]-=f*(e&&u?-1:1);break;case"end":a[r]+=f*(e&&u?-1:1);break}return a}const ut=async(t,n,e)=>{const{placement:o="bottom",strategy:i="absolute",middleware:l=[],platform:c}=e,r=await(c.isRTL==null?void 0:c.isRTL(n));if(c==null&&console.error(["Floating UI: `platform` property was not passed to config. If you","want to use Floating UI on the web, install @floating-ui/dom","instead of the /core package. Otherwise, you can create your own","`platform`: https://floating-ui.com/docs/platform"].join(" ")),l.filter(g=>{let{name:w}=g;return w==="autoPlacement"||w==="flip"}).length>1)throw new Error(["Floating UI: duplicate `flip` and/or `autoPlacement`","middleware detected. This will lead to an infinite loop. Ensure only","one of either has been passed to the `middleware` array."].join(" "));let s=await c.getElementRects({reference:t,floating:n,strategy:i}),{x:f,y:d}=I(s,o,r),u=o,a={},h=0;for(let g=0;g<l.length;g++){if(h++,h>100)throw new Error(["Floating UI: The middleware lifecycle appears to be","running in an infinite loop. This is usually caused by a `reset`","continually being returned without a break condition."].join(" "));const{name:w,fn:P}=l[g],{x:O,y:T,data:V,reset:y}=await P({x:f,y:d,initialPlacement:o,placement:u,strategy:i,middlewareData:a,rects:s,platform:c,elements:{reference:t,floating:n}});if(f=O!=null?O:f,d=T!=null?T:d,a=C(p({},a),{[w]:p(p({},a[w]),V)}),y){typeof y=="object"&&(y.placement&&(u=y.placement),y.rects&&(s=y.rects===!0?await c.getElementRects({reference:t,floating:n,strategy:i}):y.rects),{x:f,y:d}=I(s,u,r)),g=-1;continue}}return{x:f,y:d,placement:u,strategy:i,middlewareData:a}};function at(t){return p({top:0,right:0,bottom:0,left:0},t)}function dt(t){return typeof t!="number"?at(t):{top:t,right:t,bottom:t,left:t}}function j(t){return C(p({},t),{top:t.y,left:t.x,right:t.x+t.width,bottom:t.y+t.height})}const ht=Math.min,gt=Math.max;function pt(t,n,e){return gt(t,ht(n,e))}const Wt=t=>({name:"arrow",options:t,async fn(n){const{element:e,padding:o=0}=t!=null?t:{},{x:i,y:l,placement:c,rects:r,platform:s}=n;if(e==null)return console.warn("Floating UI: No `element` was passed to the `arrow` middleware."),{};const f=dt(o),d={x:i,y:l},u=S(c),a=G(u),h=await s.getDimensions(e),g=u==="y"?"top":"left",w=u==="y"?"bottom":"right",P=r.reference[a]+r.reference[u]-d[u]-r.floating[a],O=d[u]-r.reference[u],T=await(s.getOffsetParent==null?void 0:s.getOffsetParent(e)),V=T?u==="y"?T.clientHeight||0:T.clientWidth||0:0,y=P/2-O/2,ot=f[g],it=V-h[a]-f[w],M=V/2-h[a]/2+y,_=pt(ot,M,it);return{data:{[u]:_,centerOffset:M-_}}}});function wt(t,n,e,o){o===void 0&&(o=!1);const i=$(t),l=q(t),c=S(t)==="x",r=["left","top"].includes(i)?-1:1,s=o&&c?-1:1,f=typeof e=="function"?e(C(p({},n),{placement:t})):e,d=typeof f=="number";let{mainAxis:u,crossAxis:a,alignmentAxis:h}=d?{mainAxis:f,crossAxis:0,alignmentAxis:null}:p({mainAxis:0,crossAxis:0,alignmentAxis:null},f);return l&&typeof h=="number"&&(a=l==="end"?h*-1:h),c?{x:a*s,y:u*r}:{x:u*r,y:a*s}}const Dt=function(t){return t===void 0&&(t=0),{name:"offset",options:t,async fn(n){const{x:e,y:o,placement:i,rects:l,platform:c,elements:r}=n,s=wt(i,l,t,await(c.isRTL==null?void 0:c.isRTL(r.floating)));return{x:e+s.x,y:o+s.y,data:s}}}};function J(t){return t&&t.document&&t.location&&t.alert&&t.setInterval}function b(t){if(t==null)return window;if(!J(t)){const n=t.ownerDocument;return n&&n.defaultView||window}return t}function E(t){return b(t).getComputedStyle(t)}function x(t){return J(t)?"":t?(t.nodeName||"").toLowerCase():""}function m(t){return t instanceof b(t).HTMLElement}function R(t){return t instanceof b(t).Element}function mt(t){return t instanceof b(t).Node}function k(t){const n=b(t).ShadowRoot;return t instanceof n||t instanceof ShadowRoot}function W(t){const{overflow:n,overflowX:e,overflowY:o}=E(t);return/auto|scroll|overlay|hidden/.test(n+o+e)}function yt(t){return["table","td","th"].includes(x(t))}function K(t){const n=navigator.userAgent.toLowerCase().includes("firefox"),e=E(t);return e.transform!=="none"||e.perspective!=="none"||e.contain==="paint"||["transform","perspective"].includes(e.willChange)||n&&e.willChange==="filter"||n&&(e.filter?e.filter!=="none":!1)}function Q(){return!/^((?!chrome|android).)*safari/i.test(navigator.userAgent)}const X=Math.min,L=Math.max,N=Math.round;function A(t,n,e){var o,i,l,c;n===void 0&&(n=!1),e===void 0&&(e=!1);const r=t.getBoundingClientRect();let s=1,f=1;n&&m(t)&&(s=t.offsetWidth>0&&N(r.width)/t.offsetWidth||1,f=t.offsetHeight>0&&N(r.height)/t.offsetHeight||1);const d=R(t)?b(t):window,u=!Q()&&e,a=(r.left+(u&&(o=(i=d.visualViewport)==null?void 0:i.offsetLeft)!=null?o:0))/s,h=(r.top+(u&&(l=(c=d.visualViewport)==null?void 0:c.offsetTop)!=null?l:0))/f,g=r.width/s,w=r.height/f;return{width:g,height:w,top:h,right:a+g,bottom:h+w,left:a,x:a,y:h}}function v(t){return((mt(t)?t.ownerDocument:t.document)||window.document).documentElement}function D(t){return R(t)?{scrollLeft:t.scrollLeft,scrollTop:t.scrollTop}:{scrollLeft:t.pageXOffset,scrollTop:t.pageYOffset}}function Z(t){return A(v(t)).left+D(t).scrollLeft}function xt(t){const n=A(t);return N(n.width)!==t.offsetWidth||N(n.height)!==t.offsetHeight}function bt(t,n,e){const o=m(n),i=v(n),l=A(t,o&&xt(n),e==="fixed");let c={scrollLeft:0,scrollTop:0};const r={x:0,y:0};if(o||!o&&e!=="fixed")if((x(n)!=="body"||W(i))&&(c=D(n)),m(n)){const s=A(n,!0);r.x=s.x+n.clientLeft,r.y=s.y+n.clientTop}else i&&(r.x=Z(i));return{x:l.left+c.scrollLeft-r.x,y:l.top+c.scrollTop-r.y,width:l.width,height:l.height}}function tt(t){return x(t)==="html"?t:t.assignedSlot||t.parentNode||(k(t)?t.host:null)||v(t)}function U(t){return!m(t)||getComputedStyle(t).position==="fixed"?null:t.offsetParent}function vt(t){let n=tt(t);for(k(n)&&(n=n.host);m(n)&&!["html","body"].includes(x(n));){if(K(n))return n;n=n.parentNode}return null}function H(t){const n=b(t);let e=U(t);for(;e&&yt(e)&&getComputedStyle(e).position==="static";)e=U(e);return e&&(x(e)==="html"||x(e)==="body"&&getComputedStyle(e).position==="static"&&!K(e))?n:e||vt(t)||n}function Y(t){if(m(t))return{width:t.offsetWidth,height:t.offsetHeight};const n=A(t);return{width:n.width,height:n.height}}function At(t){let{rect:n,offsetParent:e,strategy:o}=t;const i=m(e),l=v(e);if(e===l)return n;let c={scrollLeft:0,scrollTop:0};const r={x:0,y:0};if((i||!i&&o!=="fixed")&&((x(e)!=="body"||W(l))&&(c=D(e)),m(e))){const s=A(e,!0);r.x=s.x+e.clientLeft,r.y=s.y+e.clientTop}return C(p({},n),{x:n.x-c.scrollLeft+r.x,y:n.y-c.scrollTop+r.y})}function Tt(t,n){const e=b(t),o=v(t),i=e.visualViewport;let l=o.clientWidth,c=o.clientHeight,r=0,s=0;if(i){l=i.width,c=i.height;const f=Q();(f||!f&&n==="fixed")&&(r=i.offsetLeft,s=i.offsetTop)}return{width:l,height:c,x:r,y:s}}function Ct(t){var n;const e=v(t),o=D(t),i=(n=t.ownerDocument)==null?void 0:n.body,l=L(e.scrollWidth,e.clientWidth,i?i.scrollWidth:0,i?i.clientWidth:0),c=L(e.scrollHeight,e.clientHeight,i?i.scrollHeight:0,i?i.clientHeight:0);let r=-o.scrollLeft+Z(t);const s=-o.scrollTop;return E(i||e).direction==="rtl"&&(r+=L(e.clientWidth,i?i.clientWidth:0)-l),{width:l,height:c,x:r,y:s}}function nt(t){const n=tt(t);return["html","body","#document"].includes(x(n))?t.ownerDocument.body:m(n)&&W(n)?n:nt(n)}function et(t,n){var e;n===void 0&&(n=[]);const o=nt(t),i=o===((e=t.ownerDocument)==null?void 0:e.body),l=b(o),c=i?[l].concat(l.visualViewport||[],W(o)?o:[]):o,r=n.concat(c);return i?r:r.concat(et(c))}function Rt(t,n){const e=n==null||n.getRootNode==null?void 0:n.getRootNode();if(t!=null&&t.contains(n))return!0;if(e&&k(e)){let o=n;do{if(o&&t===o)return!0;o=o.parentNode||o.host}while(o)}return!1}function Lt(t,n){const e=A(t,!1,n==="fixed"),o=e.top+t.clientTop,i=e.left+t.clientLeft;return{top:o,left:i,x:i,y:o,right:i+t.clientWidth,bottom:o+t.clientHeight,width:t.clientWidth,height:t.clientHeight}}function z(t,n,e){return n==="viewport"?j(Tt(t,e)):R(n)?Lt(n,e):j(Ct(v(t)))}function Et(t){const n=et(t),o=["absolute","fixed"].includes(E(t).position)&&m(t)?H(t):t;return R(o)?n.filter(i=>R(i)&&Rt(i,o)&&x(i)!=="body"):[]}function Ot(t){let{element:n,boundary:e,rootBoundary:o,strategy:i}=t;const c=[...e==="clippingAncestors"?Et(n):[].concat(e),o],r=c[0],s=c.reduce((f,d)=>{const u=z(n,d,i);return f.top=L(u.top,f.top),f.right=X(u.right,f.right),f.bottom=X(u.bottom,f.bottom),f.left=L(u.left,f.left),f},z(n,r,i));return{width:s.right-s.left,height:s.bottom-s.top,x:s.left,y:s.top}}const Vt={getClippingRect:Ot,convertOffsetParentRelativeRectToViewportRelativeRect:At,isElement:R,getDimensions:Y,getOffsetParent:H,getDocumentElement:v,getElementRects:t=>{let{reference:n,floating:e,strategy:o}=t;return{reference:bt(n,H(e),o),floating:C(p({},Y(e)),{x:0,y:0})}},getClientRects:t=>Array.from(t.getClientRects()),isRTL:t=>E(t).direction==="rtl"},Pt=(t,n,e)=>ut(t,n,p({platform:Vt},e));export{Wt as a,Pt as c,Dt as o};
//# sourceMappingURL=_@floating-ui-5b7f0efd.js.map