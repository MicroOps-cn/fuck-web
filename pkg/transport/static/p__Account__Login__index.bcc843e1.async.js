"use strict";(self.webpackChunkfuck_web_ui=self.webpackChunkfuck_web_ui||[]).push([[423],{66515:function(e,n,t){t.r(n),t.d(n,{default:function(){return ae}});var a=t(15009),r=t.n(a),s=t(99289),i=t.n(s),o=t(5574),l=t.n(o),c=t(49677),u=t.n(c),d=t(97857),m=t.n(d),f=t(19632),p=t.n(f),h=t(13769),g=t.n(h),x=t(38437),y=t(2453),v=t(96074),b=t(78957),_=t(55241),w=t(14726),j=t(53220),k=t(50963),T=t(93967),M=t.n(T),B=t(96486),Z=t(38018),P=t(67294),I=t(99090),N=t(43711),z=t(99702),C=t(97374),S=t(8250),F=t(81665),q=t(70960),L=t(41913),A=t(78158),E=t(29177),V=t(94149),O=t(88641),R=t(24454),D=t(26),U=t(5966),G=t(16434),W=t(63434),Q=t(32983),H=t(84059),K=t(79090),$=t(88372),J=t(85893),X=function(e){var n=e.parentIntl,t=void 0===n?new L.f("user.settings.security",(0,I.useIntl)()):n,a=e.token,s=new L.f("mfa",t),o=(0,P.useState)(),c=l()(o,2),u=c[0],d=c[1],m=(0,P.useState)(!1),f=l()(m,2),p=f[0],h=f[1],g=function(){var e=i()(r()().mark((function e(n){return r()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:d(void 0),h(!0),(0,F.oN)({token:n}).then((function(e){d(e.data)})).finally((function(){h(!1)}));case 3:case"end":return e.stop()}}),e)})));return function(n){return e.apply(this,arguments)}}();return(0,P.useEffect)((function(){a&&g(a)}),[a]),(0,J.jsxs)($.f,{children:[(0,J.jsx)("span",{children:(0,J.jsx)(x.Z,{style:{marginBottom:24},message:s.t("code-description","Please obtain two consecutive one-time passwords after scanning and adding MFA and enter them into the input box below."),type:"info",showIcon:!0})}),(0,J.jsxs)("div",{style:{display:"flex"},children:[(0,J.jsxs)("div",{children:[(0,J.jsx)("div",{style:{width:220,height:140},children:null!=u&&u.secret?(0,J.jsx)(H.Qd,{style:{marginLeft:"40px"},value:null==u?void 0:u.secret}):(0,J.jsx)(Q.Z,{description:"",style:{width:200,height:128},image:p?(0,J.jsx)(K.Z,{}):void 0})}),(0,J.jsx)("div",{style:{display:"grid",width:"100%"},children:(0,J.jsx)(w.ZP,{onClick:function(){a&&g(a)},type:"link",children:s.t("refresh","Refresh")})})]}),(0,J.jsxs)("div",{style:{marginLeft:"50px"},children:[(0,J.jsx)(U.Z,{hidden:!0,width:"sm",name:"token",fieldProps:{value:null==u?void 0:u.token}}),(0,J.jsx)(U.Z,{rules:[{pattern:/[0-9]{6}/,required:!0}],width:"sm",name:"firstCode",label:s.t("first-code","First code")}),(0,J.jsx)(U.Z,{rules:[{pattern:/[0-9]{6}/,required:!0}],width:"sm",name:"secondCode",label:s.t("second-code","Second code")})]})]})]})},Y={container:"container___KqtO3",loginTypeTabs:"loginTypeTabs___Q9s6H",lang:"lang___ZsPRm",oauthButton:"oauthButton___RVjpV",oauthLoginTypes:"oauthLoginTypes___qTuQ4",oauthIconButton:"oauthIconButton___dwoAZ",loginByOtherBtn:"loginByOtherBtn___kUaTL"},ee=["children","allows","loginType","hidden"],ne=function(e){var n=e.content;return e.hidden?null:(0,J.jsx)(x.Z,{style:{marginBottom:24},message:n,type:"error",showIcon:!0})},te=function(e){var n=e.children,t=e.allows,a=e.loginType,r=e.hidden,s=g()(e,ee);return new Set([].concat(p()(t),[t.map((function(e){return S.Bz[e]}))])).has(a)?(0,J.jsx)("span",m()(m()({hidden:r},s),{},{children:n})):(0,J.jsx)(J.Fragment,{})},ae=function(e){var n,t,a,s,o;u()(e);var c=(0,Z.parse)(I.history.location.search),d=c.redirect_uri,f=c.redirect,p=(0,P.useState)({success:!0}),h=l()(p,2),g=h[0],T=h[1],Q=(0,P.useState)("normal"),H=l()(Q,2),K=H[0],$=H[1],ee=(0,P.useState)(),ae=l()(ee,2),re=ae[0],se=ae[1],ie=(0,P.useState)(),oe=l()(ie,2),le=oe[0],ce=oe[1],ue=(0,P.useState)(),de=l()(ue,2),me=de[0],fe=de[1],pe=(0,I.useModel)("@@initialState"),he=pe.initialState,ge=pe.setInitialState,xe=new L.f("pages.login",(0,I.useIntl)()),ye=(0,k.cI)(),ve=l()(ye,1)[0],be=["normal","email","sms"],_e=(0,P.useState)(be.map((function(e){return S.Bz[e]}))),we=l()(_e,2),je=we[0],ke=we[1],Te=(0,P.useState)(be.map((function(e){return S.Bz[e]}))),Me=l()(Te,2),Be=Me[0],Ze=Me[1],Pe=(0,P.useState)([]),Ie=l()(Pe,2),Ne=Ie[0],ze=Ie[1],Ce=function(e){(0,B.isNumber)(e)?$(S.Bz[e]):$(e)},Se=function(){var e=i()(r()().mark((function e(){var n,t;return r()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,null==he||null===(n=he.fetchUserInfo)||void 0===n?void 0:n.call(he);case 2:if(!(t=e.sent)){e.next=6;break}return e.next=6,ge((function(e){return m()(m()({},e),{},{currentUser:t})}));case 6:case"end":return e.stop()}}),e)})));return function(){return e.apply(this,arguments)}}(),Fe=(0,P.useState)(!1),qe=l()(Fe,2),Le=qe[0],Ae=qe[1],Ee=function(){var e=i()(r()().mark((function e(n){var t,a,s,i,o,l;return r()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return t=Object.assign({},(u()(n),n)),e.prev=1,e.next=4,(0,F.n$)(m()(m()({},t),{},{type:K,token:re,bindingToken:me}),{skipErrorHandler:!0,ignoreError:!0});case 4:if(s=e.sent,!(null!==(a=s.data)&&void 0!==a&&a.nextMethod&&s.data.nextMethod.length>0)){e.next=13;break}ke(s.data.nextMethod),Ce(s.data.nextMethod[0]),Ae(!0),se(s.data.token),ce(s.data.email),e.next=26;break;case 13:if(!s.success){e.next=25;break}return i=xe.t("success","Login succeeded!"),y.ZP.success(i),e.next=18,Se();case 18:if(!d){e.next=21;break}return window.location.href=(0,B.isArray)(d)?d[0]:d,e.abrupt("return");case 21:return I.history.push(f||"/"),e.abrupt("return");case 25:"E0004"===s.errorCode?(y.ZP.warning(xe.t("".concat(null!==(o=s.errorCode)&&void 0!==o?o:"normal",".errorMessage"),s.errorMessage)),I.history.push("/account/resetPassword?username=".concat(t.username))):s.errorCode&&["E0002","E0005","E0006"].includes(s.errorCode)&&y.ZP.warning(xe.t("".concat(null!==(l=s.errorCode)&&void 0!==l?l:"normal",".errorMessage"),s.errorMessage));case 26:T(s),e.next=32;break;case 29:e.prev=29,e.t0=e.catch(1),T({success:!1});case 32:case"end":return e.stop()}}),e,null,[[1,29]])})));return function(n){return e.apply(this,arguments)}}(),Ve=g.success,Oe=g.errorMessage,Re=g.errorCode,De=(0,P.useState)((0,q.MM)(S.Bz,xe,"loginType",(function(e){return be.includes(e)}))),Ue=l()(De,2),Ge=Ue[0],We=Ue[1],Qe=null!==(n=null==he?void 0:he.globalConfig)&&void 0!==n?n:null;return(0,P.useEffect)((function(){if(Qe){for(var e=0;e<Qe.loginType.length;e++){var n=Qe.loginType[e];if(n.autoRedirect){var t=(0,A.cN)("/api/v1/user/oauth/".concat(n.id));return d&&(t="".concat(t,"?redirect_uri=").concat(encodeURIComponent((0,B.isArray)(d)?d[0]:d))),void(window.location.href=t)}}Qe.defaultLoginType===S.Bz.oauth2?(ke([]),Ze(Qe.loginType.map((function(e){return e.type})).filter((function(e){return e!==S.Bz.oauth2&&void 0!==e}))),Ce(S.Bz.oauth2)):ke(Qe.loginType.map((function(e){return e.type})).filter((function(e){return e!==S.Bz.oauth2&&void 0!==e}))),ze(Qe.loginType.filter((function(e){return e.type===S.Bz.oauth2})))}}),[Qe,d]),console.log(Be,je,Le,K),(0,P.useEffect)((function(){We((0,q.MM)(S.Bz,xe,"loginType",(function(e,n){return je.includes(e)||je.includes(n)}))),je.length>0?(je.includes(K)||$(S.Bz[je[0]]),Ae(!1)):Ae(!0)}),[je]),(0,J.jsxs)("div",{className:Y.container,style:{backgroundColor:"white",height:"100vh",display:"flex",flexFlow:"column"},children:[(0,J.jsx)("div",{className:Y.lang,"data-lang":!0,children:(0,J.jsx)(C.Z,{})}),(0,J.jsxs)(D.B,{logo:null!==(t=null==Qe?void 0:Qe.logo)&&void 0!==t?t:(0,A.Ak)("logo.svg"),title:null!==(a=null==Qe?void 0:Qe.title)&&void 0!==a?a:"IDAS",subTitle:null!==(s=null==Qe?void 0:Qe.subTitle)&&void 0!==s?s:"Identity authentication service",initialValues:{autoLogin:!0},containerStyle:{backgroundColor:"rgba(255,255,255,0.85)",backdropFilter:"blur(4px)",minHeight:"425px",backgroundImage:"unset"},mainStyle:{width:395},backgroundImageUrl:"https://gw.alipayobjects.com/zos/rmsportal/TVYTbAXWheQpRcWDaDMu.svg",style:{background:"#f0f2f5"},form:ve,actions:Ne.length>0?(0,J.jsxs)("div",{className:"ant-pro-form-login-page-other",children:[(0,J.jsx)("div",{children:(0,J.jsx)(v.Z,{plain:!0,children:xe.t("loginWith","Login with")})},"loginWith"),(0,J.jsx)("div",{style:{display:"flex",justifyContent:"center",alignItems:"center",flexDirection:"column"},children:(0,J.jsx)(b.Z,{align:"center",className:Y.oauthLoginTypes,size:24,children:Ne.map((function(e){return(0,J.jsx)(_.Z,{title:xe.t("signInWith","Sign in with {name}","",{name:e.name}),children:(0,J.jsx)("div",{className:M()(Y.oauthIconButton),onClick:function(){var n=(0,A.cN)("/api/v1/user/oauth/".concat(e.id));d&&(n="".concat(n,"?redirect_uri=").concat(encodeURIComponent((0,B.isArray)(d)?d[0]:d))),window.location.href=n},children:(0,J.jsx)("img",{src:e.icon,className:Y.oauthIcon})},e.name)})}))},"loginMethod")}),(0,J.jsx)("div",{className:Y.loginByOtherBtn,hidden:0==(null!==(o=null==Qe?void 0:Qe.loginType.filter((function(e){return e.type!==S.Bz.oauth2})))&&void 0!==o?o:[]).length||(null==Qe?void 0:Qe.defaultLoginType)!==S.Bz.oauth2,children:(0,J.jsx)(w.ZP,{style:{display:0===Be.length?"none":"unset"},type:"link",onClick:function(){Qe&&ke(Qe.loginType.map((function(e){return e.type})).filter((function(e){return e!==S.Bz.oauth2&&void 0!==e})))},children:xe.t("loginByOther","More login methods")})})]}):void 0,submitter:je.length>0&&void 0,onFinish:function(){var e=i()(r()().mark((function e(n){return r()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",Ee(n));case 1:case"end":return e.stop()}}),e)})));return function(n){return e.apply(this,arguments)}}(),children:[(0,J.jsx)(j.Z,{activeKey:K,onChange:function(e){$(e),T({success:!0})},className:Y.loginTypeTabs,items:Ge,centered:!0}),(0,J.jsxs)(te,{hidden:Le,loginType:K,allows:["normal"].concat(["mfa_totp","mfa_email","mfa_sms"]),children:[(0,J.jsx)(ne,{hidden:Ve,content:xe.t("".concat(null!=Re?Re:"normal",".errorMessage"),Oe)}),(0,J.jsx)(U.Z,{name:"username",fieldProps:{size:"large",prefix:(0,J.jsx)(E.Z,{className:Y.prefixIcon}),autoFocus:!0},placeholder:xe.t("username.placeholder","Please enter a username"),rules:[{required:!0,message:xe.t("username.required","Please enter a username!")}]}),(0,J.jsx)(U.Z.Password,{name:"password",fieldProps:{size:"large",prefix:(0,J.jsx)(V.Z,{className:Y.prefixIcon})},placeholder:xe.t("password.placeholder","Please input a password"),rules:[{required:!0,message:xe.t("password.required","Please input a password!")}]})]}),(0,J.jsx)(te,{loginType:K,allows:["mfa_totp"],children:(0,J.jsx)(ne,{hidden:Ve,content:xe.t("totp.errorMessage","verification code error")})}),(0,J.jsxs)(te,{loginType:K,allows:["email","mfa_email"],children:[(0,J.jsx)(ne,{hidden:Ve,content:xe.t("email.errorMessage","Email verification code error")}),(0,J.jsx)(U.Z,{fieldProps:{size:"large",prefix:(0,J.jsx)(O.Z,{className:Y.prefixIcon})},name:"email",placeholder:"".concat(xe.t("email.placeholder","Please enter your email")," ").concat(le?": ".concat(le):""),rules:[{required:!0,message:xe.t("email.required","Please enter your email!")}]})]}),(0,J.jsxs)(te,{loginType:K,allows:["sms","mfa_sms"],children:[(0,J.jsx)(ne,{hidden:Ve,content:xe.t("sms.errorMessage","SMS verification code error")}),(0,J.jsx)(U.Z,{fieldProps:{size:"large",prefix:(0,J.jsx)(R.Z,{className:Y.prefixIcon})},name:"phone",placeholder:xe.t("phoneNumber.placeholder","Please enter your phone number"),rules:[{required:!0,message:xe.t("phoneNumber.required","Please enter your phone number!")},{pattern:/^1\d{10}$/,message:xe.t("phoneNumber.invalid","Mobile phone number format error!")}]})]}),(0,J.jsx)(te,{loginType:K,allows:["enable_mfa_email","enable_mfa_sms","enable_mfa_totp"],style:{width:550,display:"block"},children:(0,J.jsx)(x.Z,{style:{marginBottom:24},message:xe.t("mfa.errorMessage","Because your user is set to enable multiple factor authentication (MFA), you need to enable at least one MFA authentication method."),type:"info",showIcon:!0})}),(0,J.jsx)(te,{loginType:K,allows:["enable_mfa_sms"],style:{width:550,display:"block"},children:(0,J.jsx)(x.Z,{style:{marginBottom:24},message:xe.t("enableMfa.smsMessage","Click Login to automatically enable email as the second authentication factor."),type:"info",showIcon:!0})}),(0,J.jsxs)(te,{loginType:K,allows:["enable_mfa_email"],style:{width:550,display:"block"},children:[(0,J.jsx)(x.Z,{style:{marginBottom:24},message:xe.t("enableMfa.emailMessage","Click Login to automatically enable email as the second authentication factor."),type:"info",showIcon:!0}),(0,J.jsx)(ne,{hidden:Ve,content:xe.t("email.errorMessage","Email verification code error")}),(0,J.jsx)(U.Z,{fieldProps:{size:"large",prefix:(0,J.jsx)(O.Z,{className:Y.prefixIcon})},name:"email",initialValue:le,placeholder:xe.t("email.placeholder","Please enter your email"),rules:[{required:!0,message:xe.t("email.required","Please enter your email!")}]})]}),(0,J.jsx)(te,{loginType:K,allows:["sms","mfa_sms","mfa_email","email","mfa_totp","enable_mfa_email","enable_mfa_sms"],children:(0,J.jsx)(G.Z,{fieldProps:{size:"large",prefix:(0,J.jsx)(V.Z,{className:Y.prefixIcon})},captchaProps:{size:"large",hidden:"mfa_totp"===K},placeholder:xe.t("captcha.placeholder","Please enter the verification code"),captchaTextRender:function(e,n){return e?"".concat(n," ").concat(xe.t("getCaptchaSecondText","Get verification code")):xe.t("phone.getVerificationCode","Get verification code")},name:"code",rules:[{required:!0,message:xe.t("captcha.required","Please enter the verification code!")}],onGetCaptcha:i()(r()().mark((function e(){var n,t,a,s;return r()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:a={type:K},e.t0=K,e.next="enable_mfa_email"===e.t0||"mfa_email"===e.t0?4:"email"===e.t0?6:"enable_mfa_sms"===e.t0||"mfa_sms"===e.t0?9:"sms"===e.t0?11:13;break;case 4:ve.validateFields(["username"]),a.username=ve.getFieldValue("username");case 6:return ve.validateFields(["email"]),a.email=ve.getFieldValue("email"),e.abrupt("break",14);case 9:ve.validateFields(["username"]),a.username=ve.getFieldValue("username");case 11:ve.validateFields(["phone"]),a.phone=ve.getFieldValue("phone");case 13:return e.abrupt("break",14);case 14:return e.next=16,(0,F.Do)(a,{intl:xe});case 16:if((s=e.sent).success){e.next=19;break}return e.abrupt("return");case 19:e.t1=K,e.next="enable_mfa_email"===e.t1||"enable_mfa_sms"===e.t1?22:24;break;case 22:return fe(null===(n=s.data)||void 0===n?void 0:n.token),e.abrupt("break",26);case 24:return se(null===(t=s.data)||void 0===t?void 0:t.token),e.abrupt("break",26);case 26:y.ZP.success(xe.t("captcha.sent","Verification code sent successfully."));case 27:case"end":return e.stop()}}),e)})))})}),(0,J.jsx)(te,{style:{width:550,display:"block"},loginType:K,allows:["enable_mfa_totp"],children:(0,J.jsx)(X,{token:re,setBindingToken:fe})}),(0,J.jsxs)("div",{style:{marginBottom:10,display:je.length>0?"block":"none"},children:[(0,J.jsx)(W.Z,{noStyle:!0,name:"autoLogin",children:xe.t("rememberMe","Automatic login")}),(0,J.jsx)(te,{loginType:K,allows:["normal"],children:(0,J.jsx)(I.Link,{style:{float:"right"},to:N.A,children:xe.t("forgotPassword","Forgot password")})})]})]}),(0,J.jsx)(z.Z,{})]})}},8250:function(e,n,t){t.d(n,{Bz:function(){return r},FI:function(){return a},J0:function(){return s}});var a=function(e){return e[e.unsafe=0]="unsafe",e[e.general=1]="general",e[e.safe=2]="safe",e[e.very_safe=3]="very_safe",e}({}),r=function(e){return e[e.mfa_email=2]="mfa_email",e[e.email=4]="email",e[e.enable_mfa_sms=12]="enable_mfa_sms",e[e.enable_mfa_totp=10]="enable_mfa_totp",e[e.enable_mfa_email=11]="enable_mfa_email",e[e.normal=0]="normal",e[e.mfa_totp=1]="mfa_totp",e[e.mfa_sms=3]="mfa_sms",e[e.sms=5]="sms",e[e.oauth2=6]="oauth2",e}({}),s=function(e){return e[e.disabled=1]="disabled",e[e.user_inactive=2]="user_inactive",e[e.password_expired=4]="password_expired",e[e.normal=0]="normal",e}({})},70960:function(e,n,t){t.d(n,{GG:function(){return r},MM:function(){return a}});var a=function(e,n,t,a){var r=[];for(var s in e)if(Object.prototype.propertyIsEnumerable.call(e,s)&&isNaN(Number(s))){var i=e[s];if(a&&!a(s,i))continue;r.push({label:n.formatMessage({id:"".concat(t,".").concat(s),defaultMessage:s}),key:s,value:i})}return r},r=function(e,n,t,a){var r={};for(var s in e)if(Object.prototype.hasOwnProperty.call(e,s)&&!isNaN(Number(s))){var i,o=e[s];r[s]={text:n.formatMessage({id:"".concat(t,".").concat(o),defaultMessage:o}),status:null!==(i=a[o])&&void 0!==i?i:"Default"}}return r}}}]);