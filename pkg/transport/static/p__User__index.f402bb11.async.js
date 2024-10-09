"use strict";(self.webpackChunkfuck_web_ui=self.webpackChunkfuck_web_ui||[]).push([[522],{23051:function(e,t,n){n.r(t),n.d(t,{default:function(){return _e}});var r=n(97857),a=n.n(r),u=n(5574),s=n.n(u),i=n(15009),o=n.n(i),c=n(99289),l=n.n(c),p=n(2453),d=n(17788),f=n(2487),m=n(14726),h=n(98321),v=n(53220),x=n(30381),b=n.n(x),y=n(67294),w=n(99090),k=n(13769),Z=n.n(k),g=n(78957),j=n(85418),P=n(96486),C=n(34804),I=n(85893),S=["success","failed","onClick"],T=function(e){var t=e.success,n=e.failed,r=e.onClick,u=Z()(e,S),i=(0,y.useState)(!1),c=s()(i,2),d=c[0],f=c[1],h=(0,y.useState)(0),v=s()(h,2),x=v[0],b=v[1];return(0,I.jsx)(m.ZP,a()(a()({},u),{},{loading:d,onClick:function(){var e=l()(o()().mark((function e(t){var n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(!r){e.next=16;break}return e.prev=1,f(!0),e.next=5,r(t);case 5:n=e.sent,(0,P.isBoolean)(n)&&!n?b(2):b(1),e.next=13;break;case 9:e.prev=9,e.t0=e.catch(1),p.ZP.error("".concat(e.t0),3),b(2);case 13:return e.prev=13,f(!1),e.finish(13);case 16:case"end":return e.stop()}}),e,null,[[1,9,13,16]])})));return function(t){return e.apply(this,arguments)}}(),children:1===x&&t?t:2===x&&n?n:u.children}))},E=function(e){var t=e.success,n=e.failed,r=e.key,u=e.style,s=e.type,i=void 0===s?"link":s,c=e.onClick,p=e.label;return(0,I.jsx)(T,{success:t,failed:n,style:a()({padding:"4px 0px"},u),type:i,onClick:function(){var e=l()(o()().mark((function e(t){return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",null==c?void 0:c({key:r.toString(),domEvent:t}));case 1:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),children:p},r)},N=function(e){var t=e.maxItems,n=e.items,r=e.moreLabel,u=void 0===r?"More":r,s=n.filter((function(e){return e&&!e.hidden}));if(void 0===t||t>=s.length){var i=s.map(E);return void 0===t?(0,I.jsx)(m.ZP.Group,{children:i}):(0,I.jsx)(g.Z,{children:i})}var o=s.slice(t-1).map((function(e){return a()(a()({},e),{},{label:(0,I.jsx)("div",{style:{padding:"0px 15px"},children:e.label})})}));return(0,I.jsxs)(g.Z,{children:[s.slice(0,t-1).map(E),(0,I.jsx)(j.Z,{menu:{items:o.map((function(e){return a()(a()({},e),{},{type:"item"})}))},trigger:["click"],children:(0,I.jsxs)("a",{onClick:function(e){return e.preventDefault()},style:{gap:3,display:"inline-flex"},children:[u,(0,I.jsx)(C.Z,{})]})})]})},_=n(8250),D=n(78158),F=["id"];function A(){return(A=l()(o()().mark((function e(t,n){return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,D.WY)("/api/v1/sessions",a()({method:"GET",params:a()({},t)},n||{})));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function L(e,t){return O.apply(this,arguments)}function O(){return(O=l()(o()().mark((function e(t,n){var r,u;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return r=t.id,u=Z()(t,F),e.abrupt("return",(0,D.WY)("/api/v1/sessions/".concat(r),a()({method:"DELETE",params:a()({},u)},n||{})));case 2:case"end":return e.stop()}}),e)})))).apply(this,arguments)}var q=["id"];function J(e,t){return W.apply(this,arguments)}function W(){return(W=l()(o()().mark((function e(t,n){return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,D.WY)("/api/v1/users",a()({method:"GET",params:a()({},t)},n||{})));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function M(e,t){return R.apply(this,arguments)}function R(){return(R=l()(o()().mark((function e(t,n){return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,D.WY)("/api/v1/users",a()({method:"POST",headers:{"Content-Type":"application/json"},data:t},n||{})));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function Y(e,t){return G.apply(this,arguments)}function G(){return(G=l()(o()().mark((function e(t,n){return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,D.WY)("/api/v1/users",a()({method:"DELETE",headers:{"Content-Type":"application/json"},data:t},n||{})));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function K(e,t){return V.apply(this,arguments)}function V(){return(V=l()(o()().mark((function e(t,n){return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,D.WY)("/api/v1/users",a()({method:"PATCH",headers:{"Content-Type":"application/json"},data:t},n||{})));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function B(e,t,n){return U.apply(this,arguments)}function U(){return(U=l()(o()().mark((function e(t,n,r){var u,s;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return u=t.id,s=Z()(t,q),e.abrupt("return",(0,D.WY)("/api/v1/users/".concat(u),a()({method:"PUT",headers:{"Content-Type":"application/json"},params:a()({},s),data:n},r||{})));case 2:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function z(e,t){return H.apply(this,arguments)}function H(){return(H=l()(o()().mark((function e(t,n){return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,D.WY)("/api/v1/users/sendActivateMail",a()({method:"POST",headers:{"Content-Type":"application/json"},data:t},n||{})));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}var $=n(70960),Q=n(41913),X=n(78824),ee=n(88284),te=n(11475),ne=n(51042),re=n(57919),ae=n(1413),ue=n(91),se=n(96202),ie=n(66758),oe=n(97658),ce=["fieldProps","children","params","proFieldProps","mode","valueEnum","request","showSearch","options"],le=["fieldProps","children","params","proFieldProps","mode","valueEnum","request","options"],pe=function(e,t){var n=e.fieldProps,r=e.children,a=e.params,u=e.proFieldProps,s=e.mode,i=e.valueEnum,o=e.request,c=e.showSearch,l=e.options,p=(0,ue.Z)(e,ce),d=(0,y.useContext)(ie.Z);return(0,I.jsx)(oe.Z,(0,ae.Z)((0,ae.Z)({valueEnum:(0,se.h)(i),request:o,params:a,valueType:"select",filedConfig:{customLightMode:!0},fieldProps:(0,ae.Z)({options:l,mode:s,showSearch:c,getPopupContainer:d.getPopupContainer},n),ref:t,proFieldProps:u},p),{},{children:r}))},de=y.forwardRef((function(e,t){var n=e.fieldProps,r=e.children,a=e.params,u=e.proFieldProps,s=e.mode,i=e.valueEnum,o=e.request,c=e.options,l=(0,ue.Z)(e,le),p=(0,ae.Z)({options:c,mode:s||"multiple",labelInValue:!0,showSearch:!0,suffixIcon:null,autoClearSearchValue:!0,optionLabelProp:"label"},n),d=(0,y.useContext)(ie.Z);return(0,I.jsx)(oe.Z,(0,ae.Z)((0,ae.Z)({valueEnum:(0,se.h)(i),request:o,params:a,valueType:"select",filedConfig:{customLightMode:!0},fieldProps:(0,ae.Z)({getPopupContainer:d.getPopupContainer},p),ref:t,proFieldProps:u},l),{},{children:r}))})),fe=y.forwardRef(pe);fe.SearchSelect=de,fe.displayName="ProFormComponent";var me=fe,he=n(36541),ve=n(84017),xe=n(55431),be=n(42381),ye=n(82088),we=n(43039),ke=n(16828),Ze=n(37476),ge=n(5966),je=["parentIntl"],Pe=function(e){var t=e.parentIntl,n=Z()(e,je),r=new Q.f("form",t),u=n.values,s=n.title,i=n.modalVisible,c=n.onSubmit,p=n.onCancel;return(0,I.jsxs)(Ze.Y,{labelCol:{span:8},wrapperCol:{span:14},layout:"horizontal",title:r.t("title.basicConfig","Basic"),modalProps:{width:640,bodyStyle:{padding:"32px 40px 48px"},destroyOnClose:!0,title:s,onCancel:function(){d.Z.confirm({title:r.t("cancel?","Cancel editing?"),icon:(0,I.jsx)(te.Z,{}),onOk:function(){p()},maskClosable:!0})}},open:i,initialValues:a()({},u),onFinish:function(){var e=l()(o()().mark((function e(t){return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",c(a()(a()({},t),{},{status:null!=u&&u.status?u.status:_.J0.normal,isDelete:!(null==u||!u.isDelete)&&(null==u?void 0:u.isDelete)})));case 1:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),children:[(0,I.jsx)(ye.HZ,{label:r.t("avatar.label","Avatar"),name:"avatar",request:function(){var e=l()(o()().mark((function e(t,n){var r,a;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return(r=new FormData).append(t,n),e.next=4,(0,we.c)({data:r,requestType:"form"});case 4:if(!(a=e.sent).data){e.next=7;break}return e.abrupt("return",a.data[t]);case 7:return e.abrupt("return","");case 8:case"end":return e.stop()}}),e)})));return function(t,n){return e.apply(this,arguments)}}()}),(0,I.jsx)(ge.Z,{hidden:!0,name:"id"}),(0,I.jsx)(ge.Z,{name:"username",label:r.t("userName.label","Username"),width:"md",rules:[{required:!0,message:r.t("userName.required","Please input username!")},{pattern:/^[-_A-Za-z0-9]+$/,message:r.t("name.invalid","Username format error!")}]}),(0,I.jsx)(ge.Z,{name:"fullName",label:r.t("fullName.label","FullName"),width:"md"}),(0,I.jsx)(ge.Z,{name:"email",label:r.t("email.label","Email"),width:"md"}),(0,I.jsx)(ge.Z,{name:"phoneNumber",label:r.t("phoneNumber.label","Telephone number"),width:"md"}),(0,I.jsx)(me,{name:"roleId",label:r.t("role.label","Role"),width:"md",fieldProps:{optionLabelProp:"name"},request:l()(o()().mark((function e(){var t,n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,(0,ke.F3)({});case 2:if(e.t1=t=e.sent.data,e.t0=null!==e.t1,!e.t0){e.next=6;break}e.t0=void 0!==t;case 6:if(!e.t0){e.next=10;break}e.t2=t,e.next=11;break;case 10:e.t2=[];case 11:return n=e.t2,e.abrupt("return",n.map((function(e){var t=e.id,n=e.name,r=e.describe;return{label:(0,I.jsxs)("div",{children:[n,(0,I.jsx)(be.Z.Text,{style:{width:200,color:"#827d7d",marginLeft:10},ellipsis:!0,children:r})]}),value:t,name:n}})));case 13:case"end":return e.stop()}}),e)})))})]})},Ce=function(){var e=l()(o()().mark((function e(t){var n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return n=p.ZP.loading("Adding ..."),e.prev=1,delete t.id,e.next=5,M(t);case 5:return n(),p.ZP.success("Added successfully"),e.abrupt("return",!0);case 10:return e.prev=10,e.t0=e.catch(1),n(),p.ZP.error("Adding failed, please try again!"),e.abrupt("return",!1);case 15:case"end":return e.stop()}}),e,null,[[1,10]])})));return function(t){return e.apply(this,arguments)}}(),Ie=function(){var e=l()(o()().mark((function e(t){var n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(n=p.ZP.loading("Configuring"),e.prev=1,!t.id){e.next=10;break}return e.next=5,B({id:t.id},t);case 5:return n(),p.ZP.success("update is successful"),e.abrupt("return",!0);case 10:return p.ZP.success("update failed, system error"),e.abrupt("return",!1);case 12:e.next=19;break;case 14:return e.prev=14,e.t0=e.catch(1),n(),p.ZP.error("Configuration failed, please try again!"),e.abrupt("return",!1);case 19:case"end":return e.stop()}}),e,null,[[1,14]])})));return function(t){return e.apply(this,arguments)}}(),Se=function(){var e=l()(o()().mark((function e(t){var n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(n=p.ZP.loading("enabling ..."),t){e.next=3;break}return e.abrupt("return",!0);case 3:return e.prev=3,e.next=6,K(t.map((function(e){return{id:e.id,status:_.J0.normal}})));case 6:return n(),p.ZP.success("Enabled successfully and will refresh soon"),e.abrupt("return",!0);case 11:return e.prev=11,e.t0=e.catch(3),n(),p.ZP.error("Enabled failed, please try again"),e.abrupt("return",!1);case 16:case"end":return e.stop()}}),e,null,[[3,11]])})));return function(t){return e.apply(this,arguments)}}(),Te=function(){var e=l()(o()().mark((function e(t){var n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(n=p.ZP.loading("Disabling ..."),t){e.next=3;break}return e.abrupt("return",!0);case 3:return e.prev=3,e.next=6,K(t.map((function(e){return{id:e.id,status:_.J0.disabled}})));case 6:return n(),p.ZP.success("Disabled successfully and will refresh soon"),e.abrupt("return",!0);case 11:return e.prev=11,e.t0=e.catch(3),n(),p.ZP.error("Disable failed, please try again"),e.abrupt("return",!1);case 16:case"end":return e.stop()}}),e,null,[[3,11]])})));return function(t){return e.apply(this,arguments)}}(),Ee=function(){var e=l()(o()().mark((function e(t){var n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(n=p.ZP.loading("Deleting ..."),t){e.next=3;break}return e.abrupt("return",!0);case 3:return e.prev=3,e.next=6,Y(t.map((function(e){return{id:e.id}})));case 6:return n(),p.ZP.success("Deleted successfully and will refresh soon"),e.abrupt("return",!0);case 11:return e.prev=11,e.t0=e.catch(3),n(),p.ZP.error("Delete failed, please try again: ".concat(e.t0)),e.abrupt("return",!1);case 16:case"end":return e.stop()}}),e,null,[[3,11]])})));return function(t){return e.apply(this,arguments)}}(),Ne=function(){var e=l()(o()().mark((function e(t){var n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(t){e.next=2;break}return e.abrupt("return",!0);case 2:return n=p.ZP.loading("Deleting ..."),e.prev=3,e.next=6,L({id:t});case 6:return n(),p.ZP.success("Deleted successfully and will refresh soon"),e.abrupt("return",!0);case 11:return e.prev=11,e.t0=e.catch(3),n(),p.ZP.error("Delete failed, please try again"),e.abrupt("return",!1);case 16:case"end":return e.stop()}}),e,null,[[3,11]])})));return function(t){return e.apply(this,arguments)}}(),_e=function(){var e=(0,y.useState)(!1),t=s()(e,2),n=t[0],r=t[1],u=(0,y.useState)(!1),i=s()(u,2),c=i[0],x=i[1],k=(0,y.useState)("sessions"),Z=s()(k,2),g=Z[0],j=Z[1],P=(0,y.useRef)(),C=(0,y.useRef)(),S=(0,y.useState)(),T=s()(S,2),E=T[0],D=T[1],F=(0,y.useState)([]),L=s()(F,2),O=L[0],q=L[1],W=(0,y.useState)("all"),M=s()(W,2),R=M[0],Y=M[1],G=(0,y.useState)(),K=s()(G,2),V=K[0],B=K[1],U=new Q.f("pages.users",(0,w.useIntl)()),H=function(){var e=l()(o()().mark((function e(t){var n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return n=a()(a()({},t),{},{status:"all"!==R?R:void 0}),V&&(n=a()({keywords:V},n)),e.abrupt("return",J(n));case 3:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}();(0,y.useEffect)((function(){c&&j("sessions")}),[E,c]);var ae=[{title:U.t("session.title.lastSeen","Last Seen"),dataIndex:"lastSeen",render:function(e,t){return b()(t.lastSeen).locale(U.locale).fromNow()}},{title:U.t("session.title.expiry","Expiry"),dataIndex:"expiry",render:function(e,t){return b()(t.expiry).locale(U.locale).fromNow()}},{title:U.t("session.title.loggedOn","Logged on"),dataIndex:"createTime",render:function(e,t){return b()(t.createTime).locale(U.locale).format("LLL")}},{render:function(e,t){return[(0,I.jsx)("a",{onClick:l()(o()().mark((function e(){var n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,Ne(t.id);case 2:if(!e.sent){e.next=4;break}null===(n=C.current)||void 0===n||n.reload();case 4:case"end":return e.stop()}}),e)}))),children:(0,I.jsx)(X.Z,{})},"delete")]}}],ue=(0,$.GG)(_.J0,U,"status.value",{normal:"Success",disable:"Error",inactive:"Warning"}),se=[{title:U.t("updateForm.userName.nameLabel","User name"),hideInSearch:!0,dataIndex:"username",render:function(e,t){return(0,I.jsx)("a",{onClick:function(){D(t),x(!0)},children:e})}},{title:U.t("title.fullName","FullName"),dataIndex:"fullName",hideInSearch:!0},{title:U.t("title.phoneNumber","Telephone number"),dataIndex:"phoneNumber",hideInSearch:!0},{title:U.t("title.email","Email"),dataIndex:"email",hideInSearch:!0},{title:U.t("title.role","Role"),dataIndex:"role"},{title:U.t("title.status","Status"),dataIndex:"status",hideInForm:!0,valueEnum:ue},{title:U.t("title.updatedTime","Last update time"),dataIndex:"updateTime",valueType:"dateTime",hideInTable:!0,hideInSearch:!0,hideInForm:!0},{title:U.t("title.loginTime","Last login time"),dataIndex:"loginTime",valueType:"dateTime",hideInSearch:!0,hideInForm:!0},{title:U.t("title.createTime","Create time"),dataIndex:"createTime",valueType:"dateTime",hideInSearch:!0,hideInTable:!0,hideInForm:!0},{title:U.t("title.option","Operating"),dataIndex:"option",valueType:"option",render:function(e,t){return[(0,I.jsx)(N,{maxItems:2,moreLabel:U.t("button.more","More"),items:[{key:"edit",label:U.t("button.edit","Edit"),style:{flex:"unset"},onClick:(a=l()(o()().mark((function e(){return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:r(!0),D(t);case 2:case"end":return e.stop()}}),e)}))),function(){return a.apply(this,arguments)})},{key:"activate",label:U.t("button.activate","Activate"),hidden:t.status!==_.J0.user_inactive,success:(0,I.jsx)(ee.Z,{color:"green"}),style:{flex:"unset"},onClick:(n=l()(o()().mark((function e(){return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(t.email){e.next=2;break}throw new Error(U.t("activate.no-email"," The user has no email."));case 2:return e.next=4,z({userId:t.id});case 4:if(!e.sent.success){e.next=9;break}p.ZP.success(U.t("activate.succcess","Email sent successfully.")),e.next=10;break;case 9:throw new Error(U.t("activate.failed","Email sent failed."));case 10:case"end":return e.stop()}}),e)}))),function(){return n.apply(this,arguments)})},{key:"delete",label:U.t("button.delete","Delete"),style:{flex:"unset"},onClick:function(){d.Z.confirm({title:U.t("deleteConfirm","Are you sure you want to delete the following users?            "),icon:(0,I.jsx)(te.Z,{}),onOk:function(){return l()(o()().mark((function e(){return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,Ee([t]);case 2:case"end":return e.stop()}}),e)})))()},content:(0,I.jsx)(f.Z,{dataSource:[t],rowKey:"id",renderItem:function(e){return(0,I.jsxs)(f.Z.Item,{children:[e.username,e.fullName?"(".concat(e.fullName,")"):e.email?"(".concat(e.email,")"):""]})}})})}},{key:"disable",label:U.t("button.disable","Disable"),hidden:t.status===_.J0.disabled,style:{flex:"unset"},onClick:function(){d.Z.confirm({title:U.t("disableConfirm","Are you sure you want to disable the following users?"),icon:(0,I.jsx)(te.Z,{}),onOk:function(){return l()(o()().mark((function e(){var n,r;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,Te([t]);case 2:null===(n=P.current)||void 0===n||null===(r=n.reloadAndRest)||void 0===r||r.call(n);case 3:case"end":return e.stop()}}),e)})))()},content:(0,I.jsx)(f.Z,{dataSource:[t].filter((function(e){return e.status!==_.J0.disabled})),rowKey:"id",renderItem:function(e){return(0,I.jsxs)(f.Z.Item,{children:[e.username,e.fullName?"(".concat(e.fullName,")"):e.email?"(".concat(e.email,")"):""]})}})})}}]},"options")];var n,a}}];return(0,I.jsxs)(ve._z,{children:[(0,I.jsx)(xe.ZP,{actionRef:P,rowKey:"id",search:!1,toolbar:{search:{onSearch:function(e){var t,n;(B(e),P.current)&&(null===(t=(n=P.current).setPageInfo)||void 0===t||t.call(n,a()(a()({},P.current.pageInfo),{},{current:1})),P.current.reload())}},filter:(0,I.jsx)(re.M,{onFinish:function(){var e=l()(o()().mark((function e(t){var n,r,u,s,i;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return r=t.status,Y(null!==(n=r.value)&&void 0!==n?n:"all"),P.current&&(null===(u=(s=P.current).setPageInfo)||void 0===u||u.call(s,a()(a()({},P.current.pageInfo),{},{current:1})),null===(i=P.current)||void 0===i||i.reload()),e.abrupt("return",!0);case 4:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),children:(0,I.jsx)(me,{name:"status",label:U.t("title.status","Status"),initialValue:{value:"all",label:U.t("status.all","All")},fieldProps:{labelInValue:!0},valueEnum:a()({all:U.t("status.all","All")},ue)})}),actions:[(0,I.jsxs)(m.ZP,{type:"primary",onClick:function(){r(!0)},children:[(0,I.jsx)(ne.Z,{}),U.t("button.create","Create")]},"create")]},request:H,columns:se,rowSelection:{onChange:function(e,t){q(t)}}}),(null==O?void 0:O.length)>0&&(0,I.jsxs)(FooterToolbar,{extra:(0,I.jsxs)("div",{children:[U.t("chosen","Chosen")," ",(0,I.jsx)("a",{style:{fontWeight:600},children:O.length})," ",U.t("item","Item(s)")]}),children:[(0,I.jsx)(m.ZP,{danger:!0,onClick:function(){d.Z.confirm({title:U.t("deleteConfirm","Are you sure you want to delete the following users?            "),icon:(0,I.jsx)(te.Z,{}),onOk:function(){return l()(o()().mark((function e(){var t,n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,Ee(O);case 2:q([]),null===(t=P.current)||void 0===t||null===(n=t.reloadAndRest)||void 0===n||n.call(t);case 4:case"end":return e.stop()}}),e)})))()},content:(0,I.jsx)(f.Z,{dataSource:O,rowKey:"id",renderItem:function(e){return(0,I.jsxs)(f.Z.Item,{children:[e.username,e.fullName?"(".concat(e.fullName,")"):e.email?"(".concat(e.email,")"):""]})}})})},children:U.t("batchDeletion","Batch deletion")}),O.filter((function(e){return e.status!==_.J0.disabled})).length>0&&(0,I.jsx)(m.ZP,{onClick:function(){d.Z.confirm({title:U.t("disableConfirm","Are you sure you want to disable the following users?"),icon:(0,I.jsx)(te.Z,{}),onOk:function(){return l()(o()().mark((function e(){var t,n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,Te(O.filter((function(e){return e.status!==_.J0.disabled})));case 2:q([]),null===(t=P.current)||void 0===t||null===(n=t.reloadAndRest)||void 0===n||n.call(t);case 4:case"end":return e.stop()}}),e)})))()},content:(0,I.jsx)(f.Z,{dataSource:O.filter((function(e){return e.status!==_.J0.disabled})),rowKey:"id",renderItem:function(e){return(0,I.jsxs)(f.Z.Item,{children:[e.username,e.fullName?"(".concat(e.fullName,")"):e.email?"(".concat(e.email,")"):""]})}})})},children:U.t("batchDisable","Batch disable")}),O.filter((function(e){return e.status!==_.J0.normal})).length>0&&(0,I.jsx)(m.ZP,{onClick:function(){d.Z.confirm({title:U.t("enableConfirm","Are you sure you want to enable the following users?"),icon:(0,I.jsx)(te.Z,{}),onOk:function(){return l()(o()().mark((function e(){var t,n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,Se(O.filter((function(e){return e.status!==_.J0.normal})));case 2:q([]),null===(t=P.current)||void 0===t||null===(n=t.reloadAndRest)||void 0===n||n.call(t);case 4:case"end":return e.stop()}}),e)})))()},content:(0,I.jsx)(f.Z,{dataSource:O.filter((function(e){return e.status!==_.J0.normal})),rowKey:"id",renderItem:function(e){return(0,I.jsxs)(f.Z.Item,{children:[e.username,e.fullName?"(".concat(e.fullName,")"):e.email?"(".concat(e.email,")"):""]})}})})},children:U.t("batchEnable","Batch enable")})]}),(0,I.jsx)(Pe,{title:U.t(E?"form.title.userUpdate":"form.title.userCreate",E?"Modify user":"Add user"),onSubmit:function(){var e=l()(o()().mark((function e(t){var n;return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,(E?Ie:Ce)(t);case 2:return(n=e.sent)&&(r(!1),D(void 0),P.current&&P.current.reload()),e.abrupt("return",n);case 5:case"end":return e.stop()}}),e)})));return function(t){return e.apply(this,arguments)}}(),onCancel:function(){r(!1),c||D(void 0)},modalVisible:n,values:E,parentIntl:U}),(0,I.jsx)(h.Z,{width:800,open:c,onClose:function(){D(void 0),x(!1)},closable:!1,children:c&&(null==E?void 0:E.username)&&(0,I.jsxs)(I.Fragment,{children:[(0,I.jsx)(he.ZP,{column:2,title:U.t("detail.title","User Details"),request:l()(o()().mark((function e(){return o()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",{data:E||{}});case 1:case"end":return e.stop()}}),e)}))),params:{id:null==E?void 0:E.id},columns:se}),(0,I.jsx)(v.Z,{activeKey:g,onChange:function(e){j(e)},items:[{label:U.t("detail.sessions.title","Sessions"),key:"sessions",children:(0,I.jsx)(xe.ZP,{actionRef:C,toolBarRender:!1,request:function(e){return function(e,t){return A.apply(this,arguments)}(a()(a()({},e),{},{userId:E.id}))},columns:ae,rowKey:"id",search:!1})}]})]})})]})}},8250:function(e,t,n){n.d(t,{Bz:function(){return a},FI:function(){return r},J0:function(){return u}});var r=function(e){return e[e.unsafe=0]="unsafe",e[e.general=1]="general",e[e.safe=2]="safe",e[e.very_safe=3]="very_safe",e}({}),a=function(e){return e[e.mfa_email=2]="mfa_email",e[e.email=4]="email",e[e.enable_mfa_sms=12]="enable_mfa_sms",e[e.enable_mfa_totp=10]="enable_mfa_totp",e[e.enable_mfa_email=11]="enable_mfa_email",e[e.normal=0]="normal",e[e.mfa_totp=1]="mfa_totp",e[e.mfa_sms=3]="mfa_sms",e[e.sms=5]="sms",e[e.oauth2=6]="oauth2",e}({}),u=function(e){return e[e.disabled=1]="disabled",e[e.user_inactive=2]="user_inactive",e[e.password_expired=4]="password_expired",e[e.normal=0]="normal",e}({})},16828:function(e,t,n){n.d(t,{F3:function(){return f},ZV:function(){return x},fA:function(){return h},ul:function(){return y}});var r=n(13769),a=n.n(r),u=n(15009),s=n.n(u),i=n(97857),o=n.n(i),c=n(99289),l=n.n(c),p=n(78158),d=["id"];function f(e,t){return m.apply(this,arguments)}function m(){return(m=l()(s()().mark((function e(t,n){return s()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,p.WY)("/api/v1/roles",o()({method:"GET",params:o()({},t)},n||{})));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function h(e,t){return v.apply(this,arguments)}function v(){return(v=l()(s()().mark((function e(t,n){return s()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,p.WY)("/api/v1/roles",o()({method:"POST",headers:{"Content-Type":"application/json"},data:t},n||{})));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function x(e,t){return b.apply(this,arguments)}function b(){return(b=l()(s()().mark((function e(t,n){return s()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,p.WY)("/api/v1/roles",o()({method:"DELETE",headers:{"Content-Type":"application/json"},data:t},n||{})));case 1:case"end":return e.stop()}}),e)})))).apply(this,arguments)}function y(e,t,n){return w.apply(this,arguments)}function w(){return(w=l()(s()().mark((function e(t,n,r){var u,i;return s()().wrap((function(e){for(;;)switch(e.prev=e.next){case 0:return u=t.id,i=a()(t,d),e.abrupt("return",(0,p.WY)("/api/v1/roles/".concat(u),o()({method:"PUT",headers:{"Content-Type":"application/json"},params:o()({},i),data:n},r||{})));case 2:case"end":return e.stop()}}),e)})))).apply(this,arguments)}},70960:function(e,t,n){n.d(t,{GG:function(){return a},MM:function(){return r}});var r=function(e,t,n,r){var a=[];for(var u in e)if(Object.prototype.propertyIsEnumerable.call(e,u)&&isNaN(Number(u))){var s=e[u];if(r&&!r(u,s))continue;a.push({label:t.formatMessage({id:"".concat(n,".").concat(u),defaultMessage:u}),key:u,value:s})}return a},a=function(e,t,n,r){var a={};for(var u in e)if(Object.prototype.hasOwnProperty.call(e,u)&&!isNaN(Number(u))){var s,i=e[u];a[u]={text:t.formatMessage({id:"".concat(n,".").concat(i),defaultMessage:i}),status:null!==(s=r[i])&&void 0!==s?s:"Default"}}return a}}}]);