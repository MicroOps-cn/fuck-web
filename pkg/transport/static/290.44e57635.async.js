(self.webpackChunkfuck_web_ui=self.webpackChunkfuck_web_ui||[]).push([[290],{5966:function(e,t,n){"use strict";var r=n(97685),o=n(1413),s=n(91),i=n(21770),l=n(54120),a=n(55241),u=n(97435),h=n(67294),c=n(97658),d=n(85893),f=["fieldProps","proFieldProps"],g=["fieldProps","proFieldProps"],m="text",p=function(e){var t=(0,i.Z)(e.open||!1,{value:e.open,onChange:e.onOpenChange}),n=(0,r.Z)(t,2),s=n[0],u=n[1];return(0,d.jsx)(l.Z.Item,{shouldUpdate:!0,noStyle:!0,children:function(t){var n,r=t.getFieldValue(e.name||[]);return(0,d.jsx)(a.Z,(0,o.Z)((0,o.Z)({getPopupContainer:function(e){return e&&e.parentNode?e.parentNode:e},onOpenChange:function(e){return u(e)},content:(0,d.jsxs)("div",{style:{padding:"4px 0"},children:[null===(n=e.statusRender)||void 0===n?void 0:n.call(e,r),e.strengthText?(0,d.jsx)("div",{style:{marginTop:10},children:(0,d.jsx)("span",{children:e.strengthText})}):null]}),overlayStyle:{width:240},placement:"rightTop"},e.popoverProps),{},{open:s,children:e.children}))}})},E=function(e){var t=e.fieldProps,n=e.proFieldProps,r=(0,s.Z)(e,f);return(0,d.jsx)(c.Z,(0,o.Z)({valueType:m,fieldProps:t,filedConfig:{valueType:m},proFieldProps:n},r))};E.Password=function(e){var t=e.fieldProps,n=e.proFieldProps,i=(0,s.Z)(e,g),l=(0,h.useState)(!1),a=(0,r.Z)(l,2),f=a[0],E=a[1];return null!=t&&t.statusRender&&i.name?(0,d.jsx)(p,{name:i.name,statusRender:null==t?void 0:t.statusRender,popoverProps:null==t?void 0:t.popoverProps,strengthText:null==t?void 0:t.strengthText,open:f,onOpenChange:E,children:(0,d.jsx)("div",{children:(0,d.jsx)(c.Z,(0,o.Z)({valueType:"password",fieldProps:(0,o.Z)((0,o.Z)({},(0,u.Z)(t,["statusRender","popoverProps","strengthText"])),{},{onBlur:function(e){var n;null==t||null===(n=t.onBlur)||void 0===n||n.call(t,e),E(!1)},onClick:function(e){var n;null==t||null===(n=t.onClick)||void 0===n||n.call(t,e),E(!0)}}),proFieldProps:n,filedConfig:{valueType:m}},i))})}):(0,d.jsx)(c.Z,(0,o.Z)({valueType:"password",fieldProps:t,proFieldProps:n,filedConfig:{valueType:m}},i))},E.displayName="ProFormComponent",t.Z=E},84059:function(e,t,n){"use strict";n.d(t,{Qd:function(){return v}});var r,o=n(67294),s=Object.defineProperty,i=Object.getOwnPropertySymbols,l=Object.prototype.hasOwnProperty,a=Object.prototype.propertyIsEnumerable,u=(e,t,n)=>t in e?s(e,t,{enumerable:!0,configurable:!0,writable:!0,value:n}):e[t]=n,h=(e,t)=>{for(var n in t||(t={}))l.call(t,n)&&u(e,n,t[n]);if(i)for(var n of i(t))a.call(t,n)&&u(e,n,t[n]);return e},c=(e,t)=>{var n={};for(var r in e)l.call(e,r)&&t.indexOf(r)<0&&(n[r]=e[r]);if(null!=e&&i)for(var r of i(e))t.indexOf(r)<0&&a.call(e,r)&&(n[r]=e[r]);return n};(e=>{const t=class{constructor(e,n,r,o){if(this.version=e,this.errorCorrectionLevel=n,this.modules=[],this.isFunction=[],e<t.MIN_VERSION||e>t.MAX_VERSION)throw new RangeError("Version value out of range");if(o<-1||o>7)throw new RangeError("Mask value out of range");this.size=4*e+17;let i=[];for(let e=0;e<this.size;e++)i.push(!1);for(let e=0;e<this.size;e++)this.modules.push(i.slice()),this.isFunction.push(i.slice());this.drawFunctionPatterns();const l=this.addEccAndInterleave(r);if(this.drawCodewords(l),-1==o){let e=1e9;for(let t=0;t<8;t++){this.applyMask(t),this.drawFormatBits(t);const n=this.getPenaltyScore();n<e&&(o=t,e=n),this.applyMask(t)}}s(0<=o&&o<=7),this.mask=o,this.applyMask(o),this.drawFormatBits(o),this.isFunction=[]}static encodeText(n,r){const o=e.QrSegment.makeSegments(n);return t.encodeSegments(o,r)}static encodeBinary(n,r){const o=e.QrSegment.makeBytes(n);return t.encodeSegments([o],r)}static encodeSegments(e,n,o=1,i=40,a=-1,u=!0){if(!(t.MIN_VERSION<=o&&o<=i&&i<=t.MAX_VERSION)||a<-1||a>7)throw new RangeError("Invalid value");let h,c;for(h=o;;h++){const r=8*t.getNumDataCodewords(h,n),o=l.getTotalBits(e,h);if(o<=r){c=o;break}if(h>=i)throw new RangeError("Data too long")}for(const e of[t.Ecc.MEDIUM,t.Ecc.QUARTILE,t.Ecc.HIGH])u&&c<=8*t.getNumDataCodewords(h,e)&&(n=e);let d=[];for(const t of e){r(t.mode.modeBits,4,d),r(t.numChars,t.mode.numCharCountBits(h),d);for(const e of t.getData())d.push(e)}s(d.length==c);const f=8*t.getNumDataCodewords(h,n);s(d.length<=f),r(0,Math.min(4,f-d.length),d),r(0,(8-d.length%8)%8,d),s(d.length%8==0);for(let e=236;d.length<f;e^=253)r(e,8,d);let g=[];for(;8*g.length<d.length;)g.push(0);return d.forEach(((e,t)=>g[t>>>3]|=e<<7-(7&t))),new t(h,n,g,a)}getModule(e,t){return 0<=e&&e<this.size&&0<=t&&t<this.size&&this.modules[t][e]}getModules(){return this.modules}drawFunctionPatterns(){for(let e=0;e<this.size;e++)this.setFunctionModule(6,e,e%2==0),this.setFunctionModule(e,6,e%2==0);this.drawFinderPattern(3,3),this.drawFinderPattern(this.size-4,3),this.drawFinderPattern(3,this.size-4);const e=this.getAlignmentPatternPositions(),t=e.length;for(let n=0;n<t;n++)for(let r=0;r<t;r++)0==n&&0==r||0==n&&r==t-1||n==t-1&&0==r||this.drawAlignmentPattern(e[n],e[r]);this.drawFormatBits(0),this.drawVersion()}drawFormatBits(e){const t=this.errorCorrectionLevel.formatBits<<3|e;let n=t;for(let e=0;e<10;e++)n=n<<1^1335*(n>>>9);const r=21522^(t<<10|n);s(r>>>15==0);for(let e=0;e<=5;e++)this.setFunctionModule(8,e,o(r,e));this.setFunctionModule(8,7,o(r,6)),this.setFunctionModule(8,8,o(r,7)),this.setFunctionModule(7,8,o(r,8));for(let e=9;e<15;e++)this.setFunctionModule(14-e,8,o(r,e));for(let e=0;e<8;e++)this.setFunctionModule(this.size-1-e,8,o(r,e));for(let e=8;e<15;e++)this.setFunctionModule(8,this.size-15+e,o(r,e));this.setFunctionModule(8,this.size-8,!0)}drawVersion(){if(this.version<7)return;let e=this.version;for(let t=0;t<12;t++)e=e<<1^7973*(e>>>11);const t=this.version<<12|e;s(t>>>18==0);for(let e=0;e<18;e++){const n=o(t,e),r=this.size-11+e%3,s=Math.floor(e/3);this.setFunctionModule(r,s,n),this.setFunctionModule(s,r,n)}}drawFinderPattern(e,t){for(let n=-4;n<=4;n++)for(let r=-4;r<=4;r++){const o=Math.max(Math.abs(r),Math.abs(n)),s=e+r,i=t+n;0<=s&&s<this.size&&0<=i&&i<this.size&&this.setFunctionModule(s,i,2!=o&&4!=o)}}drawAlignmentPattern(e,t){for(let n=-2;n<=2;n++)for(let r=-2;r<=2;r++)this.setFunctionModule(e+r,t+n,1!=Math.max(Math.abs(r),Math.abs(n)))}setFunctionModule(e,t,n){this.modules[t][e]=n,this.isFunction[t][e]=!0}addEccAndInterleave(e){const n=this.version,r=this.errorCorrectionLevel;if(e.length!=t.getNumDataCodewords(n,r))throw new RangeError("Invalid argument");const o=t.NUM_ERROR_CORRECTION_BLOCKS[r.ordinal][n],i=t.ECC_CODEWORDS_PER_BLOCK[r.ordinal][n],l=Math.floor(t.getNumRawDataModules(n)/8),a=o-l%o,u=Math.floor(l/o);let h=[];const c=t.reedSolomonComputeDivisor(i);for(let n=0,r=0;n<o;n++){let o=e.slice(r,r+u-i+(n<a?0:1));r+=o.length;const s=t.reedSolomonComputeRemainder(o,c);n<a&&o.push(0),h.push(o.concat(s))}let d=[];for(let e=0;e<h[0].length;e++)h.forEach(((t,n)=>{(e!=u-i||n>=a)&&d.push(t[e])}));return s(d.length==l),d}drawCodewords(e){if(e.length!=Math.floor(t.getNumRawDataModules(this.version)/8))throw new RangeError("Invalid argument");let n=0;for(let t=this.size-1;t>=1;t-=2){6==t&&(t=5);for(let r=0;r<this.size;r++)for(let s=0;s<2;s++){const i=t-s,l=0==(t+1&2)?this.size-1-r:r;!this.isFunction[l][i]&&n<8*e.length&&(this.modules[l][i]=o(e[n>>>3],7-(7&n)),n++)}}s(n==8*e.length)}applyMask(e){if(e<0||e>7)throw new RangeError("Mask value out of range");for(let t=0;t<this.size;t++)for(let n=0;n<this.size;n++){let r;switch(e){case 0:r=(n+t)%2==0;break;case 1:r=t%2==0;break;case 2:r=n%3==0;break;case 3:r=(n+t)%3==0;break;case 4:r=(Math.floor(n/3)+Math.floor(t/2))%2==0;break;case 5:r=n*t%2+n*t%3==0;break;case 6:r=(n*t%2+n*t%3)%2==0;break;case 7:r=((n+t)%2+n*t%3)%2==0;break;default:throw new Error("Unreachable")}!this.isFunction[t][n]&&r&&(this.modules[t][n]=!this.modules[t][n])}}getPenaltyScore(){let e=0;for(let n=0;n<this.size;n++){let r=!1,o=0,s=[0,0,0,0,0,0,0];for(let i=0;i<this.size;i++)this.modules[n][i]==r?(o++,5==o?e+=t.PENALTY_N1:o>5&&e++):(this.finderPenaltyAddHistory(o,s),r||(e+=this.finderPenaltyCountPatterns(s)*t.PENALTY_N3),r=this.modules[n][i],o=1);e+=this.finderPenaltyTerminateAndCount(r,o,s)*t.PENALTY_N3}for(let n=0;n<this.size;n++){let r=!1,o=0,s=[0,0,0,0,0,0,0];for(let i=0;i<this.size;i++)this.modules[i][n]==r?(o++,5==o?e+=t.PENALTY_N1:o>5&&e++):(this.finderPenaltyAddHistory(o,s),r||(e+=this.finderPenaltyCountPatterns(s)*t.PENALTY_N3),r=this.modules[i][n],o=1);e+=this.finderPenaltyTerminateAndCount(r,o,s)*t.PENALTY_N3}for(let n=0;n<this.size-1;n++)for(let r=0;r<this.size-1;r++){const o=this.modules[n][r];o==this.modules[n][r+1]&&o==this.modules[n+1][r]&&o==this.modules[n+1][r+1]&&(e+=t.PENALTY_N2)}let n=0;for(const e of this.modules)n=e.reduce(((e,t)=>e+(t?1:0)),n);const r=this.size*this.size,o=Math.ceil(Math.abs(20*n-10*r)/r)-1;return s(0<=o&&o<=9),e+=o*t.PENALTY_N4,s(0<=e&&e<=2568888),e}getAlignmentPatternPositions(){if(1==this.version)return[];{const e=Math.floor(this.version/7)+2,t=32==this.version?26:2*Math.ceil((4*this.version+4)/(2*e-2));let n=[6];for(let r=this.size-7;n.length<e;r-=t)n.splice(1,0,r);return n}}static getNumRawDataModules(e){if(e<t.MIN_VERSION||e>t.MAX_VERSION)throw new RangeError("Version number out of range");let n=(16*e+128)*e+64;if(e>=2){const t=Math.floor(e/7)+2;n-=(25*t-10)*t-55,e>=7&&(n-=36)}return s(208<=n&&n<=29648),n}static getNumDataCodewords(e,n){return Math.floor(t.getNumRawDataModules(e)/8)-t.ECC_CODEWORDS_PER_BLOCK[n.ordinal][e]*t.NUM_ERROR_CORRECTION_BLOCKS[n.ordinal][e]}static reedSolomonComputeDivisor(e){if(e<1||e>255)throw new RangeError("Degree out of range");let n=[];for(let t=0;t<e-1;t++)n.push(0);n.push(1);let r=1;for(let o=0;o<e;o++){for(let e=0;e<n.length;e++)n[e]=t.reedSolomonMultiply(n[e],r),e+1<n.length&&(n[e]^=n[e+1]);r=t.reedSolomonMultiply(r,2)}return n}static reedSolomonComputeRemainder(e,n){let r=n.map((e=>0));for(const o of e){const e=o^r.shift();r.push(0),n.forEach(((n,o)=>r[o]^=t.reedSolomonMultiply(n,e)))}return r}static reedSolomonMultiply(e,t){if(e>>>8!=0||t>>>8!=0)throw new RangeError("Byte out of range");let n=0;for(let r=7;r>=0;r--)n=n<<1^285*(n>>>7),n^=(t>>>r&1)*e;return s(n>>>8==0),n}finderPenaltyCountPatterns(e){const t=e[1];s(t<=3*this.size);const n=t>0&&e[2]==t&&e[3]==3*t&&e[4]==t&&e[5]==t;return(n&&e[0]>=4*t&&e[6]>=t?1:0)+(n&&e[6]>=4*t&&e[0]>=t?1:0)}finderPenaltyTerminateAndCount(e,t,n){return e&&(this.finderPenaltyAddHistory(t,n),t=0),t+=this.size,this.finderPenaltyAddHistory(t,n),this.finderPenaltyCountPatterns(n)}finderPenaltyAddHistory(e,t){0==t[0]&&(e+=this.size),t.pop(),t.unshift(e)}};let n=t;function r(e,t,n){if(t<0||t>31||e>>>t!=0)throw new RangeError("Value out of range");for(let r=t-1;r>=0;r--)n.push(e>>>r&1)}function o(e,t){return 0!=(e>>>t&1)}function s(e){if(!e)throw new Error("Assertion error")}n.MIN_VERSION=1,n.MAX_VERSION=40,n.PENALTY_N1=3,n.PENALTY_N2=3,n.PENALTY_N3=40,n.PENALTY_N4=10,n.ECC_CODEWORDS_PER_BLOCK=[[-1,7,10,15,20,26,18,20,24,30,18,20,24,26,30,22,24,28,30,28,28,28,28,30,30,26,28,30,30,30,30,30,30,30,30,30,30,30,30,30,30],[-1,10,16,26,18,24,16,18,22,22,26,30,22,22,24,24,28,28,26,26,26,26,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28,28],[-1,13,22,18,26,18,24,18,22,20,24,28,26,24,20,30,24,28,28,26,30,28,30,30,30,30,28,30,30,30,30,30,30,30,30,30,30,30,30,30,30],[-1,17,28,22,16,22,28,26,26,24,28,24,28,22,24,24,30,28,28,26,28,30,24,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30,30]],n.NUM_ERROR_CORRECTION_BLOCKS=[[-1,1,1,1,1,1,2,2,2,2,4,4,4,4,4,6,6,6,6,7,8,8,9,9,10,12,12,12,13,14,15,16,17,18,19,19,20,21,22,24,25],[-1,1,1,1,2,2,4,4,4,5,5,5,8,9,9,10,10,11,13,14,16,17,17,18,20,21,23,25,26,28,29,31,33,35,37,38,40,43,45,47,49],[-1,1,1,2,2,4,4,6,6,8,8,8,10,12,16,12,17,16,18,21,20,23,23,25,27,29,34,34,35,38,40,43,45,48,51,53,56,59,62,65,68],[-1,1,1,2,4,4,4,5,6,8,8,11,11,16,16,18,16,19,21,25,25,25,34,30,32,35,37,40,42,45,48,51,54,57,60,63,66,70,74,77,81]],e.QrCode=n;const i=class{constructor(e,t,n){if(this.mode=e,this.numChars=t,this.bitData=n,t<0)throw new RangeError("Invalid argument");this.bitData=n.slice()}static makeBytes(e){let t=[];for(const n of e)r(n,8,t);return new i(i.Mode.BYTE,e.length,t)}static makeNumeric(e){if(!i.isNumeric(e))throw new RangeError("String contains non-numeric characters");let t=[];for(let n=0;n<e.length;){const o=Math.min(e.length-n,3);r(parseInt(e.substr(n,o),10),3*o+1,t),n+=o}return new i(i.Mode.NUMERIC,e.length,t)}static makeAlphanumeric(e){if(!i.isAlphanumeric(e))throw new RangeError("String contains unencodable characters in alphanumeric mode");let t,n=[];for(t=0;t+2<=e.length;t+=2){let o=45*i.ALPHANUMERIC_CHARSET.indexOf(e.charAt(t));o+=i.ALPHANUMERIC_CHARSET.indexOf(e.charAt(t+1)),r(o,11,n)}return t<e.length&&r(i.ALPHANUMERIC_CHARSET.indexOf(e.charAt(t)),6,n),new i(i.Mode.ALPHANUMERIC,e.length,n)}static makeSegments(e){return""==e?[]:i.isNumeric(e)?[i.makeNumeric(e)]:i.isAlphanumeric(e)?[i.makeAlphanumeric(e)]:[i.makeBytes(i.toUtf8ByteArray(e))]}static makeEci(e){let t=[];if(e<0)throw new RangeError("ECI assignment value out of range");if(e<128)r(e,8,t);else if(e<16384)r(2,2,t),r(e,14,t);else{if(!(e<1e6))throw new RangeError("ECI assignment value out of range");r(6,3,t),r(e,21,t)}return new i(i.Mode.ECI,0,t)}static isNumeric(e){return i.NUMERIC_REGEX.test(e)}static isAlphanumeric(e){return i.ALPHANUMERIC_REGEX.test(e)}getData(){return this.bitData.slice()}static getTotalBits(e,t){let n=0;for(const r of e){const e=r.mode.numCharCountBits(t);if(r.numChars>=1<<e)return 1/0;n+=4+e+r.bitData.length}return n}static toUtf8ByteArray(e){e=encodeURI(e);let t=[];for(let n=0;n<e.length;n++)"%"!=e.charAt(n)?t.push(e.charCodeAt(n)):(t.push(parseInt(e.substr(n+1,2),16)),n+=2);return t}};let l=i;l.NUMERIC_REGEX=/^[0-9]*$/,l.ALPHANUMERIC_REGEX=/^[A-Z0-9 $%*+.\/:-]*$/,l.ALPHANUMERIC_CHARSET="0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:",e.QrSegment=l})(r||(r={})),(e=>{let t;(e=>{const t=class{constructor(e,t){this.ordinal=e,this.formatBits=t}};let n=t;n.LOW=new t(0,1),n.MEDIUM=new t(1,0),n.QUARTILE=new t(2,3),n.HIGH=new t(3,2),e.Ecc=n})(t=e.QrCode||(e.QrCode={}))})(r||(r={})),(e=>{let t;(e=>{const t=class{constructor(e,t){this.modeBits=e,this.numBitsCharCount=t}numCharCountBits(e){return this.numBitsCharCount[Math.floor((e+7)/17)]}};let n=t;n.NUMERIC=new t(1,[10,12,14]),n.ALPHANUMERIC=new t(2,[9,11,13]),n.BYTE=new t(4,[8,16,16]),n.KANJI=new t(8,[8,10,12]),n.ECI=new t(7,[0,0,0]),e.Mode=n})(t=e.QrSegment||(e.QrSegment={}))})(r||(r={}));var d=r,f={L:d.QrCode.Ecc.LOW,M:d.QrCode.Ecc.MEDIUM,Q:d.QrCode.Ecc.QUARTILE,H:d.QrCode.Ecc.HIGH},g=128,m="L",p="#FFFFFF",E="#000000",w=!1;function C(e,t=0){const n=[];return e.forEach((function(e,r){let o=null;e.forEach((function(s,i){if(!s&&null!==o)return n.push(`M${o+t} ${r+t}h${i-o}v1H${o+t}z`),void(o=null);if(i!==e.length-1)s&&null===o&&(o=i);else{if(!s)return;null===o?n.push(`M${i+t},${r+t} h1v1H${i+t}z`):n.push(`M${o+t},${r+t} h${i+1-o}v1H${o+t}z`)}}))})),n.join("")}function M(e,t){return e.slice().map(((e,n)=>n<t.y||n>=t.y+t.h?e:e.map(((e,n)=>(n<t.x||n>=t.x+t.w)&&e))))}function R(e,t,n,r){if(null==r)return null;const o=n?4:0,s=e.length+2*o,i=Math.floor(.1*t),l=s/t,a=(r.width||i)*l,u=(r.height||i)*l,h=null==r.x?e.length/2-a/2:r.x*l,c=null==r.y?e.length/2-u/2:r.y*l;let d=null;if(r.excavate){let e=Math.floor(h),t=Math.floor(c);d={x:e,y:t,w:Math.ceil(a+h-e),h:Math.ceil(u+c-t)}}return{x:h,y:c,h:u,w:a,excavation:d}}var P=function(){try{(new Path2D).addPath(new Path2D)}catch(e){return!1}return!0}();function v(e){const t=e,{value:n,size:r=g,level:s=m,bgColor:i=p,fgColor:l=E,includeMargin:a=w,style:u,imageSettings:v}=t,N=c(t,["value","size","level","bgColor","fgColor","includeMargin","style","imageSettings"]),y=null==v?void 0:v.src,A=(0,o.useRef)(null),I=(0,o.useRef)(null),[_,S]=(0,o.useState)(!1);(0,o.useEffect)((()=>{if(null!=A.current){const e=A.current,t=e.getContext("2d");if(!t)return;let o=d.QrCode.encodeText(n,f[s]).getModules();const u=a?4:0,h=o.length+2*u,c=R(o,r,a,v),g=I.current,m=null!=c&&null!==g&&g.complete&&0!==g.naturalHeight&&0!==g.naturalWidth;m&&null!=c.excavation&&(o=M(o,c.excavation));const p=window.devicePixelRatio||1;e.height=e.width=r*p;const E=r/h*p;t.scale(E,E),t.fillStyle=i,t.fillRect(0,0,h,h),t.fillStyle=l,P?t.fill(new Path2D(C(o,u))):o.forEach((function(e,n){e.forEach((function(e,r){e&&t.fillRect(r+u,n+u,1,1)}))})),m&&t.drawImage(g,c.x+u,c.y+u,c.w,c.h)}})),(0,o.useEffect)((()=>{S(!1)}),[y]);const F=h({height:r,width:r},u);let O=null;return null!=y&&(O=o.createElement("img",{src:y,key:y,style:{display:"none"},onLoad:()=>{S(!0)},ref:I})),o.createElement(o.Fragment,null,o.createElement("canvas",h({style:F,height:r,width:r,ref:A},N)),O)}},49677:function(e){e.exports=function(e){if(null==e)throw new TypeError("Cannot destructure "+e)},e.exports.__esModule=!0,e.exports.default=e.exports}}]);