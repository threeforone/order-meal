import{h as s}from"./http.b92c0875.js";const e=t=>s.post("/login",t),r=t=>s.post("/signUp",t),n=t=>s.get("/user/list",t),u=t=>s.get("/detail/"+t),o=t=>s.put("/user/status",t),p=t=>s.put("/user/admin",t),i=t=>s.put("/user/reset",t),c={login:e,signUp:r,getUserList:n,getUserDetail:u,updateUserStatus:o,updateUserAdmin:p,resetUserPassword:i};export{c as u};