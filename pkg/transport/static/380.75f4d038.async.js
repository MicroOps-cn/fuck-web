(self.webpackChunkfuck_web_ui=self.webpackChunkfuck_web_ui||[]).push([[380],{69501:function(e,t,i){"use strict";i.d(t,{Z:function(){return X}});var n=i(15009),a=i.n(n),r=i(99289),l=i.n(r),o=i(97857),s=i.n(o),d=i(5574),c=i.n(d),m=i(42381),u=i(67294),h=i(93967),g=i.n(h),p=i(53124),f=i(35792),b=i(16777),v=i(14747),S=i(83559),$=i(87893);const y=e=>{const{componentCls:t,calc:i}=e;return{[t]:Object.assign(Object.assign({},(0,v.Wf)(e)),{margin:0,padding:0,listStyle:"none",[`${t}-item`]:{position:"relative",margin:0,paddingBottom:e.itemPaddingBottom,fontSize:e.fontSize,listStyle:"none","&-tail":{position:"absolute",insetBlockStart:e.itemHeadSize,insetInlineStart:i(i(e.itemHeadSize).sub(e.tailWidth)).div(2).equal(),height:`calc(100% - ${(0,b.unit)(e.itemHeadSize)})`,borderInlineStart:`${(0,b.unit)(e.tailWidth)} ${e.lineType} ${e.tailColor}`},"&-pending":{[`${t}-item-head`]:{fontSize:e.fontSizeSM,backgroundColor:"transparent"},[`${t}-item-tail`]:{display:"none"}},"&-head":{position:"absolute",width:e.itemHeadSize,height:e.itemHeadSize,backgroundColor:e.dotBg,border:`${(0,b.unit)(e.dotBorderWidth)} ${e.lineType} transparent`,borderRadius:"50%","&-blue":{color:e.colorPrimary,borderColor:e.colorPrimary},"&-red":{color:e.colorError,borderColor:e.colorError},"&-green":{color:e.colorSuccess,borderColor:e.colorSuccess},"&-gray":{color:e.colorTextDisabled,borderColor:e.colorTextDisabled}},"&-head-custom":{position:"absolute",insetBlockStart:i(e.itemHeadSize).div(2).equal(),insetInlineStart:i(e.itemHeadSize).div(2).equal(),width:"auto",height:"auto",marginBlockStart:0,paddingBlock:e.customHeadPaddingVertical,lineHeight:1,textAlign:"center",border:0,borderRadius:0,transform:"translate(-50%, -50%)"},"&-content":{position:"relative",insetBlockStart:i(i(e.fontSize).mul(e.lineHeight).sub(e.fontSize)).mul(-1).add(e.lineWidth).equal(),marginInlineStart:i(e.margin).add(e.itemHeadSize).equal(),marginInlineEnd:0,marginBlockStart:0,marginBlockEnd:0,wordBreak:"break-word"},"&-last":{[`> ${t}-item-tail`]:{display:"none"},[`> ${t}-item-content`]:{minHeight:i(e.controlHeightLG).mul(1.2).equal()}}},[`&${t}-alternate,\n        &${t}-right,\n        &${t}-label`]:{[`${t}-item`]:{"&-tail, &-head, &-head-custom":{insetInlineStart:"50%"},"&-head":{marginInlineStart:i(e.marginXXS).mul(-1).equal(),"&-custom":{marginInlineStart:i(e.tailWidth).div(2).equal()}},"&-left":{[`${t}-item-content`]:{insetInlineStart:`calc(50% - ${(0,b.unit)(e.marginXXS)})`,width:`calc(50% - ${(0,b.unit)(e.marginSM)})`,textAlign:"start"}},"&-right":{[`${t}-item-content`]:{width:`calc(50% - ${(0,b.unit)(e.marginSM)})`,margin:0,textAlign:"end"}}}},[`&${t}-right`]:{[`${t}-item-right`]:{[`${t}-item-tail,\n            ${t}-item-head,\n            ${t}-item-head-custom`]:{insetInlineStart:`calc(100% - ${(0,b.unit)(i(i(e.itemHeadSize).add(e.tailWidth)).div(2).equal())})`},[`${t}-item-content`]:{width:`calc(100% - ${(0,b.unit)(i(e.itemHeadSize).add(e.marginXS).equal())})`}}},[`&${t}-pending\n        ${t}-item-last\n        ${t}-item-tail`]:{display:"block",height:`calc(100% - ${(0,b.unit)(e.margin)})`,borderInlineStart:`${(0,b.unit)(e.tailWidth)} dotted ${e.tailColor}`},[`&${t}-reverse\n        ${t}-item-last\n        ${t}-item-tail`]:{display:"none"},[`&${t}-reverse ${t}-item-pending`]:{[`${t}-item-tail`]:{insetBlockStart:e.margin,display:"block",height:`calc(100% - ${(0,b.unit)(e.margin)})`,borderInlineStart:`${(0,b.unit)(e.tailWidth)} dotted ${e.tailColor}`},[`${t}-item-content`]:{minHeight:i(e.controlHeightLG).mul(1.2).equal()}},[`&${t}-label`]:{[`${t}-item-label`]:{position:"absolute",insetBlockStart:i(i(e.fontSize).mul(e.lineHeight).sub(e.fontSize)).mul(-1).add(e.tailWidth).equal(),width:`calc(50% - ${(0,b.unit)(e.marginSM)})`,textAlign:"end"},[`${t}-item-right`]:{[`${t}-item-label`]:{insetInlineStart:`calc(50% + ${(0,b.unit)(e.marginSM)})`,width:`calc(50% - ${(0,b.unit)(e.marginSM)})`,textAlign:"start"}}},"&-rtl":{direction:"rtl",[`${t}-item-head-custom`]:{transform:"translate(50%, -50%)"}}})}};var x=(0,S.I$)("Timeline",(e=>{const t=(0,$.mergeToken)(e,{itemHeadSize:10,customHeadPaddingVertical:e.paddingXXS,paddingInlineEnd:2});return[y(t)]}),(e=>({tailColor:e.colorSplit,tailWidth:e.lineWidthBold,dotBorderWidth:e.wireframe?e.lineWidthBold:3*e.lineWidth,dotBg:e.colorBgContainer,itemPaddingBottom:1.25*e.padding}))),I=function(e,t){var i={};for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&t.indexOf(n)<0&&(i[n]=e[n]);if(null!=e&&"function"==typeof Object.getOwnPropertySymbols){var a=0;for(n=Object.getOwnPropertySymbols(e);a<n.length;a++)t.indexOf(n[a])<0&&Object.prototype.propertyIsEnumerable.call(e,n[a])&&(i[n[a]]=e[n[a]])}return i};var k=e=>{var{prefixCls:t,className:i,color:n="blue",dot:a,pending:r=!1,position:l,label:o,children:s}=e,d=I(e,["prefixCls","className","color","dot","pending","position","label","children"]);const{getPrefixCls:c}=u.useContext(p.E_),m=c("timeline",t),h=g()(`${m}-item`,{[`${m}-item-pending`]:r},i),f=/blue|red|green|gray/.test(n||"")?void 0:n,b=g()(`${m}-item-head`,{[`${m}-item-head-custom`]:!!a,[`${m}-item-head-${n}`]:!f});return u.createElement("li",Object.assign({},d,{className:h}),o&&u.createElement("div",{className:`${m}-item-label`},o),u.createElement("div",{className:`${m}-item-tail`}),u.createElement("div",{className:b,style:{borderColor:f,color:f}},a),u.createElement("div",{className:`${m}-item-content`},s))},O=i(96641),w=i(79090),j=function(e,t){var i={};for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&t.indexOf(n)<0&&(i[n]=e[n]);if(null!=e&&"function"==typeof Object.getOwnPropertySymbols){var a=0;for(n=Object.getOwnPropertySymbols(e);a<n.length;a++)t.indexOf(n[a])<0&&Object.prototype.propertyIsEnumerable.call(e,n[a])&&(i[n[a]]=e[n[a]])}return i};var C=e=>{var{prefixCls:t,className:i,pending:n=!1,children:a,items:r,rootClassName:l,reverse:o=!1,direction:s,hashId:d,pendingDot:c,mode:m=""}=e,h=j(e,["prefixCls","className","pending","children","items","rootClassName","reverse","direction","hashId","pendingDot","mode"]);const p=(e,i)=>"alternate"===m?"right"===e?`${t}-item-right`:"left"===e||i%2==0?`${t}-item-left`:`${t}-item-right`:"left"===m?`${t}-item-left`:"right"===m||"right"===e?`${t}-item-right`:"",f=(0,O.Z)(r||[]),b="boolean"==typeof n?null:n;n&&f.push({pending:!!n,dot:c||u.createElement(w.Z,null),children:b}),o&&f.reverse();const v=f.length,S=`${t}-item-last`,$=f.filter((e=>!!e)).map(((e,t)=>{var i;const a=t===v-2?S:"",r=t===v-1?S:"",{className:l}=e,s=j(e,["className"]);return u.createElement(k,Object.assign({},s,{className:g()([l,!o&&n?a:r,p(null!==(i=null==e?void 0:e.position)&&void 0!==i?i:"",t)]),key:(null==e?void 0:e.key)||t}))})),y=f.some((e=>!!(null==e?void 0:e.label))),x=g()(t,{[`${t}-pending`]:!!n,[`${t}-reverse`]:!!o,[`${t}-${m}`]:!!m&&!y,[`${t}-label`]:y,[`${t}-rtl`]:"rtl"===s},i,l,d);return u.createElement("ul",Object.assign({},h,{className:x}),$)},P=i(37419);var T=function(e,t){return e&&Array.isArray(e)?e:(0,P.Z)(t).map((e=>{var t,i;return Object.assign({children:null!==(i=null===(t=null==e?void 0:e.props)||void 0===t?void 0:t.children)&&void 0!==i?i:""},e.props)}))},E=function(e,t){var i={};for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&t.indexOf(n)<0&&(i[n]=e[n]);if(null!=e&&"function"==typeof Object.getOwnPropertySymbols){var a=0;for(n=Object.getOwnPropertySymbols(e);a<n.length;a++)t.indexOf(n[a])<0&&Object.prototype.propertyIsEnumerable.call(e,n[a])&&(i[n[a]]=e[n[a]])}return i};const N=e=>{const{getPrefixCls:t,direction:i,timeline:n}=u.useContext(p.E_),{prefixCls:a,children:r,items:l,className:o,style:s}=e,d=E(e,["prefixCls","children","items","className","style"]),c=t("timeline",a);const m=(0,f.Z)(c),[h,b,v]=x(c,m),S=T(l,r);return h(u.createElement(C,Object.assign({},d,{className:g()(null==n?void 0:n.className,o,v,m),style:Object.assign(Object.assign({},null==n?void 0:n.style),s),prefixCls:c,direction:i,items:S,hashId:b})))};N.Item=k;var z=N,H=i(83062),B=i(76723),W=i(96486),M=i.n(W),R=i(30381),q=i.n(R),_=i(55431),D=i(99090),Y="logline___PPKuZ",Z="logtime___zBYgt",A=i(85893);q().locale((0,D.getLocale)());var L=m.Z.Paragraph,F=function(e){var t=e.eventId,i=e.expanded,n=e.request,a=(0,u.useState)(),r=c()(a,2),l=r[0],o=r[1];(0,u.useEffect)((function(){i&&n&&n({eventId:t}).then((function(e){e.data&&o(e.data)}))}),[i,t,n]);var s=function e(t){if(M().isString(t))try{return e(JSON.parse(t))}catch(e){return[t]}else{if(M().isArray(t))return t.map((function(t){return e(t)})).flat();if(M().isObject(t)){var i=[];for(var n in t)if(Object.prototype.hasOwnProperty.call(t,n)){var a=t[n];i.push((0,A.jsxs)("div",{className:Y,children:[(0,A.jsxs)("div",{className:"log-title",children:["[",n,"]"]}),(0,A.jsx)(L,{className:"log-content",ellipsis:{expandable:!0},copyable:!0,children:a})]},n))}return i}}return[t]};return(0,A.jsx)(z,{children:null==l?void 0:l.map((function(e){return(0,A.jsxs)(z.Item,{color:"green",children:[(0,A.jsx)("div",{children:q()(e.createTime).format("YYYY-MM-DD hh:mm:ss")}),s(e.log),(0,A.jsx)("div",{className:Z,children:q()(e.createTime).format("YYYY-MM-DD hh:mm:ss")})]},e.id)}))})},X=function(e){var t=e.parentIntl,i=e.logViewerProps,n=e.request;return(0,A.jsx)(_.ZP,{columns:[{title:t.t("title.eventTime","Event Time"),hideInSearch:!0,dataIndex:"createTime",width:280,render:function(e,t){if(null==t||!t.createTime)return"";var i=q()(t.createTime);return"".concat(i.format("YYYY-MM-DD HH:mm:ss")," (").concat(i.fromNow(),")")}},{title:t.t("title.username","Username"),dataIndex:"username",width:180,hideInSearch:!0,render:function(e,t){return t.username?t.userId?(0,A.jsx)(H.Z,{title:t.userId,children:t.username}):t.username:t.userId}},{title:t.t("title.clientIP","Client IP"),dataIndex:"clientIp",width:180,hideInSearch:!0},{title:t.t("title.location","Location"),dataIndex:"location",width:180,hideInSearch:!0},{title:t.t("title.action","Action"),dataIndex:"action",hideInSearch:!0},{title:t.t("title.status","Status"),hideInSearch:!0,dataIndex:"status",valueEnum:{0:{text:t.t("status.failed","Failed"),status:"Error"},1:{text:t.t("status.successd","Successd"),status:"Success"},false:{text:t.t("status.failed","Failed"),status:"Error"},true:{text:t.t("status.successd","Successd"),status:"Success"}},render:function(e,t){return!t.status&&t.message?(0,A.jsx)(H.Z,{title:t.message,children:(0,A.jsx)("span",{children:e})}):e}},{title:t.t("title.took","Took"),hideInSearch:!0,dataIndex:"took",render:function(e,i){if(!i.took)return"";var n=i.took/1e3/1e3;return n<10?t.t("took.millisecond","{took} milliseconds","",{took:n.toFixed(2)}):n<1e3?t.t("took.millisecond","{took} milliseconds","",{took:n.toFixed(0)}):n<1e4?t.t("took.second","{took} seconds","",{took:(n/1e3).toFixed(2)}):n<12e4?t.t("took.second","{took} seconds","",{took:(n/1e3).toFixed(0)}):q().duration(n).humanize()}},{title:t.t("title.keywords","Keywords"),hideInTable:!0,dataIndex:"keywords",hideInSearch:!1},{title:t.t("title.timeRange","Time Range"),hideInTable:!0,dataIndex:"timeRange",valueType:"dateTimeRange",hideInSearch:!1,search:{transform:function(e){return{startTime:q()(e[0]).format(),endTime:q()(e[1]).format()}}},initialValue:[q()().startOf("day"),q()().endOf("day")],formItemProps:{style:{width:500}},renderFormItem:function(){var e=[{text:t.t("timeRange.today","Today"),value:[q()().startOf("day"),q()().endOf("day")]},{text:t.t("timeRange.lastDay","Last day"),value:[q()().add(-1,"d"),q()()]},{text:t.t("timeRange.last3Days","Last 3 Days"),value:[q()().add(-3,"d"),q()()]},{text:t.t("timeRange.lastWeek","Last Week"),value:[q()().add(-1,"w"),q()()]},{text:t.t("timeRange.lastMonth","Last Month"),value:[q()().add(-1,"M"),q()()]},{text:t.t("timeRange.thisWeek","This week"),value:[q()().startOf("week"),q()().endOf("week")]},{text:t.t("timeRange.thisMonth","This month"),value:[q()().startOf("month"),q()().endOf("month")]}];return(0,A.jsx)(B.default.RangePicker,{showTime:!0,ranges:e.reduce((function(e,t){return e[t.text]=t.value,e}),{})})}}],request:function(){var e=l()(a()().mark((function e(t){return a()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",n(t));case 1:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),rowKey:"id",expandable:{expandedRowRender:function(e,t,n,a){return(0,A.jsx)(F,s()(s()({},i),{},{eventId:e.id,expanded:a}))}},pagination:{pageSize:10,showSizeChanger:!0,pageSizeOptions:[5,10,20,50,100]}})}},49677:function(e){e.exports=function(e){if(null==e)throw new TypeError("Cannot destructure "+e)},e.exports.__esModule=!0,e.exports.default=e.exports}}]);