import{h as I}from"./http.b92c0875.js";import{d as b,r as m,i as k,j as V,_ as O,k as u,o as p,c as f,e as s,l as r,F as _,m as C,n as g}from"./index.0d91fd13.js";const D=()=>I.get("/company"),E=n=>I.get("/department",n),h={getCompanyList:D,getDepartmentList:E},L=b({name:"Company",props:{ifInline:{type:Boolean,default:!1}},setup(n,{emit:o}){const c=m(Option[0]),y=async()=>{h.getCompanyList().then(e=>{c.value=e.data.data})},t=m(""),v=()=>{a.value="",d()},l=m(Option[0]),d=async()=>{t.value===""?l.value=[]:h.getDepartmentList({companyId:t.value}).then(e=>{l.value=e.data.data})},a=m(""),i=()=>({companyId:t,departmentId:a});return k(()=>{y()}),V([t,a],()=>{o("selected-options",t.value,a.value)}),{companyUpdate:v,companyId:t,departmentId:a,companyOptions:c,departmentOptions:l,getSelections:i}}});function B(n,o,c,y,t,v){const l=u("el-option"),d=u("el-select"),a=u("el-form-item"),i=u("el-form");return p(),f("div",null,[s(i,{inline:n.ifInline,class:"demo-form-inline"},{default:r(()=>[s(a,{label:"\u516C\u53F8"},{default:r(()=>[s(d,{clearable:"",modelValue:n.companyId,"onUpdate:modelValue":o[0]||(o[0]=e=>n.companyId=e),"value-key":"id",placeholder:"\u516C\u53F8",onChange:n.companyUpdate},{default:r(()=>[(p(!0),f(_,null,C(n.companyOptions,e=>(p(),g(l,{key:e.id,label:e.name,value:e.id},null,8,["label","value"]))),128))]),_:1},8,["modelValue","onChange"])]),_:1}),s(a,{label:"\u90E8\u95E8"},{default:r(()=>[s(d,{modelValue:n.departmentId,"onUpdate:modelValue":o[1]||(o[1]=e=>n.departmentId=e),"value-key":"id",placeholder:"\u90E8\u95E8",clearable:""},{default:r(()=>[(p(!0),f(_,null,C(n.departmentOptions,e=>(p(),g(l,{key:e.id,label:e.name,value:e.id},null,8,["label","value"]))),128))]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["inline"])])}const $=O(L,[["render",B]]);export{$ as C};
