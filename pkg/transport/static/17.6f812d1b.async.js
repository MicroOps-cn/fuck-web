(self.webpackChunkfuck_web_ui=self.webpackChunkfuck_web_ui||[]).push([[17],{84017:function(e,n,t){"use strict";t.d(n,{_z:function(){return ne}});var r=t(4942),o=t(91),a=t(1413),i=t(71002),l=t(61923),c=t(53220),u=t(67159),d=t(28459),s=t(64218),f=t(93967),p=t.n(f),g=t(67294),m=t(76509),y=t(12044),b=t(97435),v=t(73935),h=t(7563),C=function(e){return(0,r.Z)({},e.componentCls,{position:"fixed",insetInlineEnd:0,bottom:0,zIndex:99,display:"flex",alignItems:"center",width:"100%",paddingInline:24,paddingBlock:0,boxSizing:"border-box",lineHeight:"64px",backgroundColor:(0,h.uK)(e.colorBgElevated,.6),borderBlockStart:"1px solid ".concat(e.colorSplit),"-webkit-backdrop-filter":"blur(8px)",backdropFilter:"blur(8px)",color:e.colorText,transition:"all 0.2s ease 0s","&-left":{flex:1,color:e.colorText},"&-right":{color:e.colorText,"> *":{marginInlineEnd:8,"&:last-child":{marginBlock:0,marginInline:0}}}})};var O=t(85893),j=["children","className","extra","portalDom","style","renderContent"],x=function(e){var n=e.children,t=e.className,l=e.extra,c=e.portalDom,u=void 0===c||c,s=e.style,f=e.renderContent,x=(0,o.Z)(e,j),w=(0,g.useContext)(d.ZP.ConfigContext),P=w.getPrefixCls,Z=w.getTargetContainer,S=e.prefixCls||P("pro"),k="".concat(S,"-footer-bar"),I=function(e){return(0,h.Xj)("ProLayoutFooterToolbar",(function(n){var t=(0,a.Z)((0,a.Z)({},n),{},{componentCls:".".concat(e)});return[C(t)]}))}(k),_=I.wrapSSR,M=I.hashId,T=(0,g.useContext)(m.X),N=(0,g.useMemo)((function(){var e=T.hasSiderMenu,n=T.isMobile,t=T.siderWidth;if(e)return t?n?"100%":"calc(100% - ".concat(t,"px)"):"100%"}),[T.collapsed,T.hasSiderMenu,T.isMobile,T.siderWidth]),E=(0,g.useMemo)((function(){return void 0===("undefined"==typeof window?"undefined":(0,i.Z)(window))||void 0===("undefined"==typeof document?"undefined":(0,i.Z)(document))?null:(null==Z?void 0:Z())||document.body}),[]),B=function(e,n){var t=n.stylish;return(0,h.Xj)("ProLayoutFooterToolbarStylish",(function(n){var o=(0,a.Z)((0,a.Z)({},n),{},{componentCls:".".concat(e)});return t?[(0,r.Z)({},"".concat(o.componentCls),null==t?void 0:t(o))]:[]}))}("".concat(k,".").concat(k,"-stylish"),{stylish:e.stylish}),H=(0,O.jsxs)(O.Fragment,{children:[(0,O.jsx)("div",{className:"".concat(k,"-left ").concat(M).trim(),children:l}),(0,O.jsx)("div",{className:"".concat(k,"-right ").concat(M).trim(),children:n})]});(0,g.useEffect)((function(){return T&&null!=T&&T.setHasFooterToolbar?(null==T||T.setHasFooterToolbar(!0),function(){var e;null==T||null===(e=T.setHasFooterToolbar)||void 0===e||e.call(T,!1)}):function(){}}),[]);var D=(0,O.jsx)("div",(0,a.Z)((0,a.Z)({className:p()(t,M,k,(0,r.Z)({},"".concat(k,"-stylish"),!!e.stylish)),style:(0,a.Z)({width:N},s)},(0,b.Z)(x,["prefixCls"])),{},{children:f?f((0,a.Z)((0,a.Z)((0,a.Z)({},e),T),{},{leftWidth:N}),H):H})),R=(0,y.j)()&&u&&E?(0,v.createPortal)(D,E,k):D;return B.wrapSSR(_((0,O.jsx)(g.Fragment,{children:R},k)))},w=t(88372),P=t(97685),Z=t(3770),S=t.n(Z),k=t(77059),I=t.n(k),_=t(85673),M=t(7134),T=t(78957),N=t(9220),E=t(30967),B=function(e){var n;return(0,r.Z)({},e.componentCls,(0,a.Z)((0,a.Z)({},null===E.Wf||void 0===E.Wf?void 0:(0,E.Wf)(e)),{},(0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)({position:"relative",backgroundColor:e.pageHeaderBgGhost,paddingBlock:e.pageHeaderPaddingVertical+2,paddingInline:e.pageHeaderPadding,"&-no-children":{height:null===(n=e.layout)||void 0===n||null===(n=n.pageContainer)||void 0===n?void 0:n.paddingBlockPageContainerContent},"& &-has-breadcrumb":{paddingBlockStart:e.pageHeaderPaddingBreadCrumb},"& &-has-footer":{paddingBlockEnd:0},"& &-back":(0,r.Z)({marginInlineEnd:e.margin,fontSize:16,lineHeight:1,"&-button":(0,a.Z)((0,a.Z)({fontSize:16},null===E.Nd||void 0===E.Nd?void 0:(0,E.Nd)(e)),{},{color:e.pageHeaderColorBack,cursor:"pointer"})},"".concat(e.componentCls,"-rlt &"),{float:"right",marginInlineEnd:0,marginInlineStart:0})},"& ".concat("ant","-divider-vertical"),{height:14,marginBlock:0,marginInline:e.marginSM,verticalAlign:"middle"}),"& &-breadcrumb + &-heading",{marginBlockStart:e.marginXS}),"& &-heading",{display:"flex",justifyContent:"space-between","&-left":{display:"flex",alignItems:"center",marginBlock:e.marginXS/2,marginInlineEnd:0,marginInlineStart:0,overflow:"hidden"},"&-title":(0,a.Z)((0,a.Z)({marginInlineEnd:e.marginSM,marginBlockEnd:0,color:e.colorTextHeading,fontWeight:600,fontSize:e.pageHeaderFontSizeHeaderTitle,lineHeight:e.controlHeight+"px"},{overflow:"hidden",whiteSpace:"nowrap",textOverflow:"ellipsis"}),{},(0,r.Z)({},"".concat(e.componentCls,"-rlt &"),{marginInlineEnd:0,marginInlineStart:e.marginSM})),"&-avatar":(0,r.Z)({marginInlineEnd:e.marginSM},"".concat(e.componentCls,"-rlt &"),{float:"right",marginInlineEnd:0,marginInlineStart:e.marginSM}),"&-tags":(0,r.Z)({},"".concat(e.componentCls,"-rlt &"),{float:"right"}),"&-sub-title":(0,a.Z)((0,a.Z)({marginInlineEnd:e.marginSM,color:e.colorTextSecondary,fontSize:e.pageHeaderFontSizeHeaderSubTitle,lineHeight:e.lineHeight},{overflow:"hidden",whiteSpace:"nowrap",textOverflow:"ellipsis"}),{},(0,r.Z)({},"".concat(e.componentCls,"-rlt &"),{float:"right",marginInlineEnd:0,marginInlineStart:12})),"&-extra":(0,r.Z)((0,r.Z)({marginBlock:e.marginXS/2,marginInlineEnd:0,marginInlineStart:0,whiteSpace:"nowrap","> *":(0,r.Z)({"white-space":"unset"},"".concat(e.componentCls,"-rlt &"),{marginInlineEnd:e.marginSM,marginInlineStart:0})},"".concat(e.componentCls,"-rlt &"),{float:"left"}),"*:first-child",(0,r.Z)({},"".concat(e.componentCls,"-rlt &"),{marginInlineEnd:0}))}),"&-content",{paddingBlockStart:e.pageHeaderPaddingContentPadding}),"&-footer",{marginBlockStart:e.margin}),"&-compact &-heading",{flexWrap:"wrap"}),"&-wide",{maxWidth:1152,margin:"0 auto"}),"&-rtl",{direction:"rtl"})))};var H=function(e,n,t,r){return t&&r?(0,O.jsx)("div",{className:"".concat(e,"-back ").concat(n).trim(),children:(0,O.jsx)("div",{role:"button",onClick:function(e){null==r||r(e)},className:"".concat(e,"-back-button ").concat(n).trim(),"aria-label":"back",children:t})}):null},D=function(e){var n=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"ltr";return void 0!==e.backIcon?e.backIcon:"rtl"===n?(0,O.jsx)(I(),{}):(0,O.jsx)(S(),{})},R=function(e){var n,t=g.useState(!1),o=(0,P.Z)(t,2),i=o[0],l=o[1],c=g.useContext(d.ZP.ConfigContext),u=c.getPrefixCls,s=c.direction,f=e.prefixCls,m=e.style,y=e.footer,b=e.children,v=e.breadcrumb,h=e.breadcrumbRender,C=e.className,j=e.contentWidth,x=e.layout,w=u("page-header",f),Z=function(e){return(0,E.Xj)("ProLayoutPageHeader",(function(n){var t=(0,a.Z)((0,a.Z)({},n),{},{componentCls:".".concat(e),pageHeaderBgGhost:"transparent",pageHeaderPadding:16,pageHeaderPaddingVertical:4,pageHeaderPaddingBreadCrumb:n.paddingSM,pageHeaderColorBack:n.colorTextHeading,pageHeaderFontSizeHeaderTitle:n.fontSizeHeading4,pageHeaderFontSizeHeaderSubTitle:14,pageHeaderPaddingContentPadding:n.paddingSM});return[B(t)]}))}(w),S=Z.wrapSSR,k=Z.hashId,I=(!v||null!=v&&v.items||null==v||!v.routes||(v.items=v.routes),null!=v&&v.items?function(e,n){var t;return null!==(t=e.items)&&void 0!==t&&t.length?(0,O.jsx)(_.Z,(0,a.Z)((0,a.Z)({},e),{},{className:p()("".concat(n,"-breadcrumb"),e.className)})):null}(v,w):null),R=v&&"props"in v,W=null!==(n=null==h?void 0:h((0,a.Z)((0,a.Z)({},e),{},{prefixCls:w}),I))&&void 0!==n?n:I,A=R?v:W,z=p()(w,k,C,(0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)({},"".concat(w,"-has-breadcrumb"),!!A),"".concat(w,"-has-footer"),!!y),"".concat(w,"-rtl"),"rtl"===s),"".concat(w,"-compact"),i),"".concat(w,"-wide"),"Fixed"===j&&"top"==x),"".concat(w,"-ghost"),!0)),F=function(e,n){var t=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"ltr",r=arguments.length>3?arguments[3]:void 0,o=n.title,i=n.avatar,l=n.subTitle,c=n.tags,u=n.extra,d=n.onBack,s="".concat(e,"-heading"),f=o||l||c||u;if(!f)return null;var g=D(n,t),m=H(e,r,g,d),y=m||i||f;return(0,O.jsxs)("div",{className:s+" "+r,children:[y&&(0,O.jsxs)("div",{className:"".concat(s,"-left ").concat(r).trim(),children:[m,i&&(0,O.jsx)(M.C,(0,a.Z)({className:p()("".concat(s,"-avatar"),r,i.className)},i)),o&&(0,O.jsx)("span",{className:"".concat(s,"-title ").concat(r).trim(),title:"string"==typeof o?o:void 0,children:o}),l&&(0,O.jsx)("span",{className:"".concat(s,"-sub-title ").concat(r).trim(),title:"string"==typeof l?l:void 0,children:l}),c&&(0,O.jsx)("span",{className:"".concat(s,"-tags ").concat(r).trim(),children:c})]}),u&&(0,O.jsx)("span",{className:"".concat(s,"-extra ").concat(r).trim(),children:(0,O.jsx)(T.Z,{children:u})})]})}(w,e,s,k),L=b&&function(e,n,t){return(0,O.jsx)("div",{className:"".concat(e,"-content ").concat(t).trim(),children:n})}(w,b,k),X=function(e,n,t){return n?(0,O.jsx)("div",{className:"".concat(e,"-footer ").concat(t).trim(),children:n}):null}(w,y,k);return A||F||X||L?S((0,O.jsx)(N.Z,{onResize:function(e){var n=e.width;return l(n<768)},children:(0,O.jsxs)("div",{className:z,style:m,children:[A,F,L,X]})})):(0,O.jsx)("div",{className:p()(k,["".concat(w,"-no-children")])})},W=t(83832),A=function(e){var n=(0,h.dQ)().token,t=e.children,r=e.style,o=e.className,i=e.markStyle,l=e.markClassName,c=e.zIndex,u=void 0===c?9:c,s=e.gapX,f=void 0===s?212:s,m=e.gapY,y=void 0===m?222:m,b=e.width,v=void 0===b?120:b,C=e.height,j=void 0===C?64:C,x=e.rotate,w=void 0===x?-22:x,Z=e.image,S=e.offsetLeft,k=e.offsetTop,I=e.fontStyle,_=void 0===I?"normal":I,M=e.fontWeight,T=void 0===M?"normal":M,N=e.fontColor,E=void 0===N?n.colorFill:N,B=e.fontSize,H=void 0===B?16:B,D=e.fontFamily,R=void 0===D?"sans-serif":D,W=e.prefixCls,A=(0,(0,g.useContext)(d.ZP.ConfigContext).getPrefixCls)("pro-layout-watermark",W),z=p()("".concat(A,"-wrapper"),o),F=p()(A,l),L=(0,g.useState)(""),X=(0,P.Z)(L,2),U=X[0],K=X[1];return(0,g.useEffect)((function(){var n=document.createElement("canvas"),t=n.getContext("2d"),r=function(e){if(!e)return 1;var n=e.backingStorePixelRatio||e.webkitBackingStorePixelRatio||e.mozBackingStorePixelRatio||e.msBackingStorePixelRatio||e.oBackingStorePixelRatio||1;return(window.devicePixelRatio||1)/n}(t),o="".concat((f+v)*r,"px"),a="".concat((y+j)*r,"px"),i=S||f/2,l=k||y/2;if(n.setAttribute("width",o),n.setAttribute("height",a),t){t.translate(i*r,l*r),t.rotate(Math.PI/180*Number(w));var c=v*r,u=j*r,d=function(e){var o=arguments.length>1&&void 0!==arguments[1]?arguments[1]:0,a=Number(H)*r;t.font="".concat(_," normal ").concat(T," ").concat(a,"px/").concat(u,"px ").concat(R),t.fillStyle=E,Array.isArray(e)?null==e||e.forEach((function(e,n){return t.fillText(e,0,n*a+o)})):t.fillText(e,0,o?o+a:0),K(n.toDataURL())};if(Z){var s=new Image;return s.crossOrigin="anonymous",s.referrerPolicy="no-referrer",s.src=Z,void(s.onload=function(){t.drawImage(s,0,0,c,u),K(n.toDataURL()),e.content&&d(e.content,s.height+8)})}e.content&&d(e.content)}else console.error("当前环境不支持Canvas")}),[f,y,S,k,w,_,T,v,j,R,E,Z,e.content,H]),(0,O.jsxs)("div",{style:(0,a.Z)({position:"relative"},r),className:z,children:[t,(0,O.jsx)("div",{className:F,style:(0,a.Z)((0,a.Z)({zIndex:u,position:"absolute",left:0,top:0,width:"100%",height:"100%",backgroundSize:"".concat(f+v,"px"),pointerEvents:"none",backgroundRepeat:"repeat"},U?{backgroundImage:"url('".concat(U,"')")}:{}),i)})]})},z=[576,768,992,1200].map((function(e){return"@media (max-width: ".concat(e,"px)")})),F=(0,P.Z)(z,4),L=F[0],X=F[1],U=F[2],K=F[3],V=function(e){var n,t,o,a,i,l,c,u,d,s,f,p,g,m,y,b,v,h;return(0,r.Z)({},e.componentCls,(0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)({position:"relative","&-children-container":{paddingBlockStart:0,paddingBlockEnd:null===(n=e.layout)||void 0===n||null===(n=n.pageContainer)||void 0===n?void 0:n.paddingBlockPageContainerContent,paddingInline:null===(t=e.layout)||void 0===t||null===(t=t.pageContainer)||void 0===t?void 0:t.paddingInlinePageContainerContent},"&-children-container-no-header":{paddingBlockStart:null===(o=e.layout)||void 0===o||null===(o=o.pageContainer)||void 0===o?void 0:o.paddingBlockPageContainerContent},"&-affix":(0,r.Z)({},"".concat(e.antCls,"-affix"),(0,r.Z)({},"".concat(e.componentCls,"-warp"),{backgroundColor:null===(a=e.layout)||void 0===a||null===(a=a.pageContainer)||void 0===a?void 0:a.colorBgPageContainerFixed,transition:"background-color 0.3s",boxShadow:"0 2px 8px #f0f1f2"}))},"& &-warp-page-header",(0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)({paddingBlockStart:(null!==(i=null===(l=e.layout)||void 0===l||null===(l=l.pageContainer)||void 0===l?void 0:l.paddingBlockPageContainerContent)&&void 0!==i?i:40)/4,paddingBlockEnd:(null!==(c=null===(u=e.layout)||void 0===u||null===(u=u.pageContainer)||void 0===u?void 0:u.paddingBlockPageContainerContent)&&void 0!==c?c:40)/2,paddingInlineStart:null===(d=e.layout)||void 0===d||null===(d=d.pageContainer)||void 0===d?void 0:d.paddingInlinePageContainerContent,paddingInlineEnd:null===(s=e.layout)||void 0===s||null===(s=s.pageContainer)||void 0===s?void 0:s.paddingInlinePageContainerContent},"& ~ ".concat(e.proComponentsCls,"-grid-content"),(0,r.Z)({},"".concat(e.proComponentsCls,"-page-container-children-content"),{paddingBlock:(null!==(f=null===(p=e.layout)||void 0===p||null===(p=p.pageContainer)||void 0===p?void 0:p.paddingBlockPageContainerContent)&&void 0!==f?f:24)/3})),"".concat(e.antCls,"-page-header-breadcrumb"),{paddingBlockStart:(null!==(g=null===(m=e.layout)||void 0===m||null===(m=m.pageContainer)||void 0===m?void 0:m.paddingBlockPageContainerContent)&&void 0!==g?g:40)/4+10}),"".concat(e.antCls,"-page-header-heading"),{paddingBlockStart:(null!==(y=null===(b=e.layout)||void 0===b||null===(b=b.pageContainer)||void 0===b?void 0:b.paddingBlockPageContainerContent)&&void 0!==y?y:40)/4}),"".concat(e.antCls,"-page-header-footer"),{marginBlockStart:(null!==(v=null===(h=e.layout)||void 0===h||null===(h=h.pageContainer)||void 0===h?void 0:h.paddingBlockPageContainerContent)&&void 0!==v?v:40)/4})),"&-detail",(0,r.Z)({display:"flex"},L,{display:"block"})),"&-main",{width:"100%"}),"&-row",(0,r.Z)({display:"flex",width:"100%"},X,{display:"block"})),"&-content",{flex:"auto",width:"100%"}),"&-extraContent",(0,r.Z)((0,r.Z)((0,r.Z)((0,r.Z)({flex:"0 1 auto",minWidth:"242px",marginInlineStart:88,textAlign:"end"},K,{marginInlineStart:44}),U,{marginInlineStart:20}),X,{marginInlineStart:0,textAlign:"start"}),L,{marginInlineStart:0})))};var G=t(1977),Y=["title","content","pageHeaderRender","header","prefixedClassName","extraContent","childrenContentStyle","style","prefixCls","hashId","value","breadcrumbRender"],$=["children","loading","className","style","footer","affixProps","token","fixedHeader","breadcrumbRender","footerToolBarProps","childrenContentStyle"];var Q=function(e){var n=e.tabList,t=e.tabActiveKey,r=e.onTabChange,o=e.hashId,i=e.tabBarExtraContent,l=e.tabProps,d=e.prefixedClassName;return Array.isArray(n)||i?(0,O.jsx)(c.Z,(0,a.Z)((0,a.Z)({className:"".concat(d,"-tabs ").concat(o).trim(),activeKey:t,onChange:function(e){r&&r(e)},tabBarExtraContent:i,items:null==n?void 0:n.map((function(e,n){var t;return(0,a.Z)((0,a.Z)({label:e.tab},e),{},{key:(null===(t=e.key)||void 0===t?void 0:t.toString())||(null==n?void 0:n.toString())})}))},l),{},{children:(0,G.n)(u.Z,"4.23.0")<0?null==n?void 0:n.map((function(e,n){return(0,O.jsx)(c.Z.TabPane,(0,a.Z)({tab:e.tab},e),e.key||n)})):null})):null},q=function(e,n,t,r){return e||n?(0,O.jsx)("div",{className:"".concat(t,"-detail ").concat(r).trim(),children:(0,O.jsx)("div",{className:"".concat(t,"-main ").concat(r).trim(),children:(0,O.jsxs)("div",{className:"".concat(t,"-row ").concat(r).trim(),children:[e&&(0,O.jsx)("div",{className:"".concat(t,"-content ").concat(r).trim(),children:e}),n&&(0,O.jsx)("div",{className:"".concat(t,"-extraContent ").concat(r).trim(),children:n})]})})}):null},J=function(e){var n,t=e.title,r=e.content,i=e.pageHeaderRender,l=e.header,c=e.prefixedClassName,u=e.extraContent,d=(e.childrenContentStyle,e.style,e.prefixCls),s=e.hashId,f=e.value,p=e.breadcrumbRender,g=(0,o.Z)(e,Y);if(!1===i)return null;if(i)return(0,O.jsxs)(O.Fragment,{children:[" ",i((0,a.Z)((0,a.Z)({},e),f))]});var m=t;t||!1===t||(m=f.title);var y=(0,a.Z)((0,a.Z)((0,a.Z)({},f),{},{title:m},g),{},{footer:Q((0,a.Z)((0,a.Z)({},g),{},{hashId:s,breadcrumbRender:p,prefixedClassName:c}))},l),b=y.breadcrumb,v=!(b&&(null!=b&&b.itemRender||null!=b&&null!==(n=b.items)&&void 0!==n&&n.length)||p);return["title","subTitle","extra","tags","footer","avatar","backIcon"].every((function(e){return!y[e]}))&&v&&!r&&!u?null:(0,O.jsx)(R,(0,a.Z)((0,a.Z)({},y),{},{className:"".concat(c,"-warp-page-header ").concat(s).trim(),breadcrumb:!1===p?void 0:(0,a.Z)((0,a.Z)({},y.breadcrumb),f.breadcrumbProps),breadcrumbRender:function(){if(p)return p}(),prefixCls:d,children:(null==l?void 0:l.children)||q(r,u,c,s)}))},ee=function(e){var n,t,c=e.children,u=e.loading,f=void 0!==u&&u,y=e.className,b=e.style,v=e.footer,C=e.affixProps,j=e.token,P=e.fixedHeader,Z=e.breadcrumbRender,S=e.footerToolBarProps,k=e.childrenContentStyle,I=(0,o.Z)(e,$),_=(0,g.useContext)(m.X);(0,g.useEffect)((function(){var e;return _&&null!=_&&_.setHasPageContainer?(null==_||null===(e=_.setHasPageContainer)||void 0===e||e.call(_,(function(e){return e+1})),function(){var e;null==_||null===(e=_.setHasPageContainer)||void 0===e||e.call(_,(function(e){return e-1}))}):function(){}}),[]);var M=(0,g.useContext)(l.L_).token,T=(0,g.useContext)(d.ZP.ConfigContext).getPrefixCls,N=e.prefixCls||T("pro"),E="".concat(N,"-page-container"),B=function(e,n){return(0,h.Xj)("ProLayoutPageContainer",(function(t){var r,o=(0,a.Z)((0,a.Z)({},t),{},{componentCls:".".concat(e),layout:(0,a.Z)((0,a.Z)({},null==t?void 0:t.layout),{},{pageContainer:(0,a.Z)((0,a.Z)({},null==t||null===(r=t.layout)||void 0===r?void 0:r.pageContainer),n)})});return[V(o)]}))}(E,j),H=B.wrapSSR,D=B.hashId,R=function(e,n){var t=n.stylish;return(0,h.Xj)("ProLayoutPageContainerStylish",(function(n){var o=(0,a.Z)((0,a.Z)({},n),{},{componentCls:".".concat(e)});return t?[(0,r.Z)({},"div".concat(o.componentCls),null==t?void 0:t(o))]:[]}))}("".concat(E,".").concat(E,"-stylish"),{stylish:e.stylish}),z=(0,g.useMemo)((function(){var e;return 0!=Z&&(Z||(null==I||null===(e=I.header)||void 0===e?void 0:e.breadcrumbRender))}),[Z,null==I||null===(n=I.header)||void 0===n?void 0:n.breadcrumbRender]),F=J((0,a.Z)((0,a.Z)({},I),{},{breadcrumbRender:z,ghost:!0,hashId:D,prefixCls:void 0,prefixedClassName:E,value:_})),L=(0,g.useMemo)((function(){if(g.isValidElement(f))return f;if("boolean"==typeof f&&!f)return null;var e=function(e){return"object"===(0,i.Z)(e)?e:{spinning:e}}(f);return e.spinning?(0,O.jsx)(W.S,(0,a.Z)({},e)):null}),[f]),X=(0,g.useMemo)((function(){return c?(0,O.jsx)(O.Fragment,{children:(0,O.jsx)("div",{className:p()(D,"".concat(E,"-children-container"),(0,r.Z)({},"".concat(E,"-children-container-no-header"),!F)),style:k,children:c})}):null}),[c,E,k,D]),U=(0,g.useMemo)((function(){var n=L||X;if(e.waterMarkProps||_.waterMarkProps){var t=(0,a.Z)((0,a.Z)({},_.waterMarkProps),e.waterMarkProps);return(0,O.jsx)(A,(0,a.Z)((0,a.Z)({},t),{},{children:n}))}return n}),[e.waterMarkProps,_.waterMarkProps,L,X]),K=p()(E,D,y,(0,r.Z)((0,r.Z)((0,r.Z)({},"".concat(E,"-with-footer"),v),"".concat(E,"-with-affix"),P&&F),"".concat(E,"-stylish"),!!I.stylish));return H(R.wrapSSR((0,O.jsxs)(O.Fragment,{children:[(0,O.jsxs)("div",{style:b,className:K,children:[P&&F?(0,O.jsx)(s.Z,(0,a.Z)((0,a.Z)({offsetTop:_.hasHeader&&_.fixedHeader?null===(t=M.layout)||void 0===t||null===(t=t.header)||void 0===t?void 0:t.heightLayoutHeader:1},C),{},{className:"".concat(E,"-affix ").concat(D).trim(),children:(0,O.jsx)("div",{className:"".concat(E,"-warp ").concat(D).trim(),children:F})})):F,U&&(0,O.jsx)(w.f,{children:U})]}),v&&(0,O.jsx)(x,(0,a.Z)((0,a.Z)({stylish:I.footerStylish,prefixCls:N},S),{},{children:v}))]})))},ne=function(e){return(0,O.jsx)(l._Y,{needDeps:!0,children:(0,O.jsx)(ee,(0,a.Z)({},e))})}},3770:function(e,n,t){"use strict";Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var r;const o=(r=t(27863))&&r.__esModule?r:{default:r};n.default=o,e.exports=o},77059:function(e,n,t){"use strict";Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var r;const o=(r=t(21379))&&r.__esModule?r:{default:r};n.default=o,e.exports=o},33046:function(e,n,t){"use strict";Object.defineProperty(n,"__esModule",{value:!0}),Object.defineProperty(n,"default",{enumerable:!0,get:function(){return h}});var r=g(t(67294)),o=f(t(93967)),a=t(34853),i=f(t(61711)),l=f(t(27727)),c=t(26814),u=t(72014);function d(e,n){(null==n||n>e.length)&&(n=e.length);for(var t=0,r=new Array(n);t<n;t++)r[t]=e[t];return r}function s(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function f(e){return e&&e.__esModule?e:{default:e}}function p(e){if("function"!=typeof WeakMap)return null;var n=new WeakMap,t=new WeakMap;return(p=function(e){return e?t:n})(e)}function g(e,n){if(!n&&e&&e.__esModule)return e;if(null===e||"object"!=typeof e&&"function"!=typeof e)return{default:e};var t=p(n);if(t&&t.has(e))return t.get(e);var r={__proto__:null},o=Object.defineProperty&&Object.getOwnPropertyDescriptor;for(var a in e)if("default"!==a&&Object.prototype.hasOwnProperty.call(e,a)){var i=o?Object.getOwnPropertyDescriptor(e,a):null;i&&(i.get||i.set)?Object.defineProperty(r,a,i):r[a]=e[a]}return r.default=e,t&&t.set(e,r),r}function m(e,n){return n=null!=n?n:{},Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):function(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))})),e}function y(e,n){if(null==e)return{};var t,r,o=function(e,n){if(null==e)return{};var t,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||(o[t]=e[t]);return o}(e,n);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}function b(e,n){return function(e){if(Array.isArray(e))return e}(e)||function(e,n){var t=null==e?null:"undefined"!=typeof Symbol&&e[Symbol.iterator]||e["@@iterator"];if(null!=t){var r,o,a=[],i=!0,l=!1;try{for(t=t.call(e);!(i=(r=t.next()).done)&&(a.push(r.value),!n||a.length!==n);i=!0);}catch(e){l=!0,o=e}finally{try{i||null==t.return||t.return()}finally{if(l)throw o}}return a}}(e,n)||function(e,n){if(!e)return;if("string"==typeof e)return d(e,n);var t=Object.prototype.toString.call(e).slice(8,-1);"Object"===t&&e.constructor&&(t=e.constructor.name);if("Map"===t||"Set"===t)return Array.from(t);if("Arguments"===t||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t))return d(e,n)}(e,n)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}(0,c.setTwoToneColor)(a.blue.primary);var v=r.forwardRef((function(e,n){var t,a=e.className,c=e.icon,d=e.spin,f=e.rotate,p=e.tabIndex,g=e.onClick,v=e.twoToneColor,h=y(e,["className","icon","spin","rotate","tabIndex","onClick","twoToneColor"]),C=r.useContext(i.default),O=C.prefixCls,j=void 0===O?"anticon":O,x=C.rootClassName,w=(0,o.default)(x,j,(s(t={},"".concat(j,"-").concat(c.name),!!c.name),s(t,"".concat(j,"-spin"),!!d||"loading"===c.name),t),a),P=p;void 0===P&&g&&(P=-1);var Z=f?{msTransform:"rotate(".concat(f,"deg)"),transform:"rotate(".concat(f,"deg)")}:void 0,S=b((0,u.normalizeTwoToneColors)(v),2),k=S[0],I=S[1];return r.createElement("span",m(function(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{},r=Object.keys(t);"function"==typeof Object.getOwnPropertySymbols&&(r=r.concat(Object.getOwnPropertySymbols(t).filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable})))),r.forEach((function(n){s(e,n,t[n])}))}return e}({role:"img","aria-label":c.name},h),{ref:n,tabIndex:P,onClick:g,className:w}),r.createElement(l.default,{icon:c,primaryColor:k,secondaryColor:I,style:Z}))}));v.displayName="AntdIcon",v.getTwoToneColor=c.getTwoToneColor,v.setTwoToneColor=c.setTwoToneColor;var h=v},61711:function(e,n,t){"use strict";Object.defineProperty(n,"__esModule",{value:!0}),Object.defineProperty(n,"default",{enumerable:!0,get:function(){return r}});var r=(0,t(67294).createContext)({})},27727:function(e,n,t){"use strict";Object.defineProperty(n,"__esModule",{value:!0}),Object.defineProperty(n,"default",{enumerable:!0,get:function(){return p}});var r=l(t(67294)),o=t(72014);function a(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function i(e){if("function"!=typeof WeakMap)return null;var n=new WeakMap,t=new WeakMap;return(i=function(e){return e?t:n})(e)}function l(e,n){if(!n&&e&&e.__esModule)return e;if(null===e||"object"!=typeof e&&"function"!=typeof e)return{default:e};var t=i(n);if(t&&t.has(e))return t.get(e);var r={__proto__:null},o=Object.defineProperty&&Object.getOwnPropertyDescriptor;for(var a in e)if("default"!==a&&Object.prototype.hasOwnProperty.call(e,a)){var l=o?Object.getOwnPropertyDescriptor(e,a):null;l&&(l.get||l.set)?Object.defineProperty(r,a,l):r[a]=e[a]}return r.default=e,t&&t.set(e,r),r}function c(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{},r=Object.keys(t);"function"==typeof Object.getOwnPropertySymbols&&(r=r.concat(Object.getOwnPropertySymbols(t).filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable})))),r.forEach((function(n){a(e,n,t[n])}))}return e}function u(e,n){return n=null!=n?n:{},Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):function(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))})),e}function d(e,n){if(null==e)return{};var t,r,o=function(e,n){if(null==e)return{};var t,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||(o[t]=e[t]);return o}(e,n);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)t=a[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(o[t]=e[t])}return o}var s={primaryColor:"#333",secondaryColor:"#E6E6E6",calculated:!1};var f=function(e){var n=e.icon,t=e.className,a=e.onClick,i=e.style,l=e.primaryColor,f=e.secondaryColor,p=d(e,["icon","className","onClick","style","primaryColor","secondaryColor"]),g=r.useRef(),m=s;if(l&&(m={primaryColor:l,secondaryColor:f||(0,o.getSecondaryColor)(l)}),(0,o.useInsertStyles)(g),(0,o.warning)((0,o.isIconDefinition)(n),"icon should be icon definiton, but got ".concat(n)),!(0,o.isIconDefinition)(n))return null;var y=n;return y&&"function"==typeof y.icon&&(y=u(c({},y),{icon:y.icon(m.primaryColor,m.secondaryColor)})),(0,o.generate)(y.icon,"svg-".concat(y.name),u(c({className:t,onClick:a,style:i,"data-icon":y.name,width:"1em",height:"1em",fill:"currentColor","aria-hidden":"true"},p),{ref:g}))};f.displayName="IconReact",f.getTwoToneColors=function(){return c({},s)},f.setTwoToneColors=function(e){var n=e.primaryColor,t=e.secondaryColor;s.primaryColor=n,s.secondaryColor=t||(0,o.getSecondaryColor)(n),s.calculated=!!t};var p=f},26814:function(e,n,t){"use strict";Object.defineProperty(n,"__esModule",{value:!0}),function(e,n){for(var t in n)Object.defineProperty(e,t,{enumerable:!0,get:n[t]})}(n,{getTwoToneColor:function(){return u},setTwoToneColor:function(){return c}});var r=i(t(27727)),o=t(72014);function a(e,n){(null==n||n>e.length)&&(n=e.length);for(var t=0,r=new Array(n);t<n;t++)r[t]=e[t];return r}function i(e){return e&&e.__esModule?e:{default:e}}function l(e,n){return function(e){if(Array.isArray(e))return e}(e)||function(e,n){var t=null==e?null:"undefined"!=typeof Symbol&&e[Symbol.iterator]||e["@@iterator"];if(null!=t){var r,o,a=[],i=!0,l=!1;try{for(t=t.call(e);!(i=(r=t.next()).done)&&(a.push(r.value),!n||a.length!==n);i=!0);}catch(e){l=!0,o=e}finally{try{i||null==t.return||t.return()}finally{if(l)throw o}}return a}}(e,n)||function(e,n){if(!e)return;if("string"==typeof e)return a(e,n);var t=Object.prototype.toString.call(e).slice(8,-1);"Object"===t&&e.constructor&&(t=e.constructor.name);if("Map"===t||"Set"===t)return Array.from(t);if("Arguments"===t||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t))return a(e,n)}(e,n)||function(){throw new TypeError("Invalid attempt to destructure non-iterable instance.\\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}()}function c(e){var n=l((0,o.normalizeTwoToneColors)(e),2),t=n[0],a=n[1];return r.default.setTwoToneColors({primaryColor:t,secondaryColor:a})}function u(){var e=r.default.getTwoToneColors();return e.calculated?[e.primaryColor,e.secondaryColor]:e.primaryColor}},27863:function(e,n,t){"use strict";Object.defineProperty(n,"__esModule",{value:!0}),Object.defineProperty(n,"default",{enumerable:!0,get:function(){return f}});var r=u(t(67294)),o=l(t(47356)),a=l(t(33046));function i(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function l(e){return e&&e.__esModule?e:{default:e}}function c(e){if("function"!=typeof WeakMap)return null;var n=new WeakMap,t=new WeakMap;return(c=function(e){return e?t:n})(e)}function u(e,n){if(!n&&e&&e.__esModule)return e;if(null===e||"object"!=typeof e&&"function"!=typeof e)return{default:e};var t=c(n);if(t&&t.has(e))return t.get(e);var r={__proto__:null},o=Object.defineProperty&&Object.getOwnPropertyDescriptor;for(var a in e)if("default"!==a&&Object.prototype.hasOwnProperty.call(e,a)){var i=o?Object.getOwnPropertyDescriptor(e,a):null;i&&(i.get||i.set)?Object.defineProperty(r,a,i):r[a]=e[a]}return r.default=e,t&&t.set(e,r),r}function d(e,n){return n=null!=n?n:{},Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):function(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))})),e}var s=function(e,n){return r.createElement(a.default,d(function(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{},r=Object.keys(t);"function"==typeof Object.getOwnPropertySymbols&&(r=r.concat(Object.getOwnPropertySymbols(t).filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable})))),r.forEach((function(n){i(e,n,t[n])}))}return e}({},e),{ref:n,icon:o.default}))};var f=r.forwardRef(s)},21379:function(e,n,t){"use strict";Object.defineProperty(n,"__esModule",{value:!0}),Object.defineProperty(n,"default",{enumerable:!0,get:function(){return f}});var r=u(t(67294)),o=l(t(44149)),a=l(t(33046));function i(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function l(e){return e&&e.__esModule?e:{default:e}}function c(e){if("function"!=typeof WeakMap)return null;var n=new WeakMap,t=new WeakMap;return(c=function(e){return e?t:n})(e)}function u(e,n){if(!n&&e&&e.__esModule)return e;if(null===e||"object"!=typeof e&&"function"!=typeof e)return{default:e};var t=c(n);if(t&&t.has(e))return t.get(e);var r={__proto__:null},o=Object.defineProperty&&Object.getOwnPropertyDescriptor;for(var a in e)if("default"!==a&&Object.prototype.hasOwnProperty.call(e,a)){var i=o?Object.getOwnPropertyDescriptor(e,a):null;i&&(i.get||i.set)?Object.defineProperty(r,a,i):r[a]=e[a]}return r.default=e,t&&t.set(e,r),r}function d(e,n){return n=null!=n?n:{},Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):function(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))})),e}var s=function(e,n){return r.createElement(a.default,d(function(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{},r=Object.keys(t);"function"==typeof Object.getOwnPropertySymbols&&(r=r.concat(Object.getOwnPropertySymbols(t).filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable})))),r.forEach((function(n){i(e,n,t[n])}))}return e}({},e),{ref:n,icon:o.default}))};var f=r.forwardRef(s)},72014:function(e,n,t){"use strict";Object.defineProperty(n,"__esModule",{value:!0}),function(e,n){for(var t in n)Object.defineProperty(e,t,{enumerable:!0,get:n[t]})}(n,{generate:function(){return v},getSecondaryColor:function(){return h},iconStyles:function(){return j},isIconDefinition:function(){return y},normalizeAttrs:function(){return b},normalizeTwoToneColors:function(){return C},svgBaseProps:function(){return O},useInsertStyles:function(){return x},warning:function(){return m}});var r=t(34853),o=t(93399),a=t(63298),i=d(t(45520)),l=f(t(67294)),c=d(t(61711));function u(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function d(e){return e&&e.__esModule?e:{default:e}}function s(e){if("function"!=typeof WeakMap)return null;var n=new WeakMap,t=new WeakMap;return(s=function(e){return e?t:n})(e)}function f(e,n){if(!n&&e&&e.__esModule)return e;if(null===e||"object"!=typeof e&&"function"!=typeof e)return{default:e};var t=s(n);if(t&&t.has(e))return t.get(e);var r={__proto__:null},o=Object.defineProperty&&Object.getOwnPropertyDescriptor;for(var a in e)if("default"!==a&&Object.prototype.hasOwnProperty.call(e,a)){var i=o?Object.getOwnPropertyDescriptor(e,a):null;i&&(i.get||i.set)?Object.defineProperty(r,a,i):r[a]=e[a]}return r.default=e,t&&t.set(e,r),r}function p(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{},r=Object.keys(t);"function"==typeof Object.getOwnPropertySymbols&&(r=r.concat(Object.getOwnPropertySymbols(t).filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable})))),r.forEach((function(n){u(e,n,t[n])}))}return e}function g(e){return e.replace(/-(.)/g,(function(e,n){return n.toUpperCase()}))}function m(e,n){(0,i.default)(e,"[@ant-design/icons] ".concat(n))}function y(e){return"object"==typeof e&&"string"==typeof e.name&&"string"==typeof e.theme&&("object"==typeof e.icon||"function"==typeof e.icon)}function b(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object.keys(e).reduce((function(n,t){var r=e[t];if("class"===t)n.className=r,delete n.class;else delete n[t],n[g(t)]=r;return n}),{})}function v(e,n,t){return t?l.default.createElement(e.tag,p({key:n},b(e.attrs),t),(e.children||[]).map((function(t,r){return v(t,"".concat(n,"-").concat(e.tag,"-").concat(r))}))):l.default.createElement(e.tag,p({key:n},b(e.attrs)),(e.children||[]).map((function(t,r){return v(t,"".concat(n,"-").concat(e.tag,"-").concat(r))})))}function h(e){return(0,r.generate)(e)[0]}function C(e){return e?Array.isArray(e)?e:[e]:[]}var O={width:"1em",height:"1em",fill:"currentColor","aria-hidden":"true",focusable:"false"},j="\n.anticon {\n  display: inline-block;\n  color: inherit;\n  font-style: normal;\n  line-height: 0;\n  text-align: center;\n  text-transform: none;\n  vertical-align: -0.125em;\n  text-rendering: optimizeLegibility;\n  -webkit-font-smoothing: antialiased;\n  -moz-osx-font-smoothing: grayscale;\n}\n\n.anticon > * {\n  line-height: 1;\n}\n\n.anticon svg {\n  display: inline-block;\n}\n\n.anticon::before {\n  display: none;\n}\n\n.anticon .anticon-icon {\n  display: block;\n}\n\n.anticon[tabindex] {\n  cursor: pointer;\n}\n\n.anticon-spin::before,\n.anticon-spin {\n  display: inline-block;\n  -webkit-animation: loadingCircle 1s infinite linear;\n  animation: loadingCircle 1s infinite linear;\n}\n\n@-webkit-keyframes loadingCircle {\n  100% {\n    -webkit-transform: rotate(360deg);\n    transform: rotate(360deg);\n  }\n}\n\n@keyframes loadingCircle {\n  100% {\n    -webkit-transform: rotate(360deg);\n    transform: rotate(360deg);\n  }\n}\n",x=function(e){var n=(0,l.useContext)(c.default),t=n.csp,r=n.prefixCls,i=j;r&&(i=i.replace(/anticon/g,r)),(0,l.useEffect)((function(){var n=e.current,r=(0,a.getShadowRoot)(n);(0,o.updateCSS)(i,"@ant-design-icons",{prepend:!0,csp:t,attachTo:r})}),[])}}}]);