import{d as D,r as m,k as n,o as $,n as H,l as t,e,u as s,q as w,b as l,w as g,s as x,t as V,a as y,x as O,y as P,g as T,_ as U}from"./index.0d91fd13.js";const j={class:"toolbar"},q=D({__name:"Home",setup(z){let c=m("");const u=a=>{sessionStorage.setItem("activePath",a),c.value=a},B={date:"2016-05-02",name:"Tom",address:"No. 189, Grove St, Los Angeles"},d=m({id:0,username:"\u674E\u534E",admin:0,superAdmin:0,departmentId:0,companyId:0}),p=window.localStorage.getItem("user");if(p!=null){const a=JSON.parse(p);d.value=a}const A=()=>{window.localStorage.removeItem("user"),window.localStorage.removeItem("jwt"),T.push("/login")};return m(Array.from({length:20}).fill(B)),(a,o)=>{const _=n("el-icon"),r=n("el-menu-item"),f=n("el-sub-menu"),b=n("el-menu"),k=n("el-scrollbar"),C=n("el-aside"),E=n("el-dropdown-item"),I=n("el-dropdown-menu"),S=n("el-dropdown"),h=n("el-header"),F=n("router-view"),N=n("el-main"),v=n("el-container");return $(),H(v,{class:"layout-container-demo",style:{height:"800px"}},{default:t(()=>[e(C,{width:"200px"},{default:t(()=>[e(k,null,{default:t(()=>[e(b,{router:"","default-active":s(c)},{default:t(()=>[e(r,{index:"/order",onClick:o[0]||(o[0]=i=>u("/order"))},{default:t(()=>[e(_,null,{default:t(()=>[e(s(w))]),_:1}),l("\u70B9\u9910 ")]),_:1}),g(e(f,{index:"1"},{title:t(()=>[e(_,null,{default:t(()=>[e(s(w))]),_:1}),l("\u70B9\u9910\u7EDF\u8BA1 ")]),default:t(()=>[e(r,{index:"/meal/day",onClick:o[1]||(o[1]=i=>u("/meal/day"))},{default:t(()=>[l("\u5F53\u65E5\u7EDF\u8BA1")]),_:1}),e(r,{index:"/meal/statistics",onClick:o[2]||(o[2]=i=>u("/meal/statistics"))},{default:t(()=>[l("\u5386\u53F2\u7EDF\u8BA1")]),_:1})]),_:1},512),[[x,d.value.admin==1]]),g(e(f,{index:"2"},{title:t(()=>[e(_,null,{default:t(()=>[e(s(V))]),_:1}),l("\u4EBA\u5458\u7BA1\u7406 ")]),default:t(()=>[e(r,{index:"/user/list",onClick:o[3]||(o[3]=i=>u("/user/list"))},{default:t(()=>[l("\u4EBA\u5458\u5217\u8868")]),_:1})]),_:1},512),[[x,d.value.admin==1]])]),_:1},8,["default-active"])]),_:1})]),_:1}),e(v,null,{default:t(()=>[e(h,{style:{"text-align":"right","font-size":"12px"}},{default:t(()=>[y("div",j,[e(S,null,{dropdown:t(()=>[e(I,null,{default:t(()=>[e(E,{onClick:A},{default:t(()=>[l("\u9000\u51FA\u767B\u5F55")]),_:1})]),_:1})]),default:t(()=>[e(_,{style:{"margin-right":"8px","margin-top":"1px"}},{default:t(()=>[e(s(O))]),_:1})]),_:1}),y("span",null,P(d.value.username),1)])]),_:1}),e(N,null,{default:t(()=>[e(F)]),_:1})]),_:1})]),_:1})}}});const J=U(q,[["__scopeId","data-v-fa4ceb8f"]]);export{J as default};