import{h as t}from"./http.b92c0875.js";const r=()=>t.get("/order/time"),s=()=>t.get("/order/status"),o=()=>t.post("/order"),a=e=>t.get("/meal/day",e),n=e=>t.get("/meal/statistics",e),d={getOrderTime:r,getOrderStatus:s,postOrder:o,getDayMeal:a,getMealStatistics:n};export{d as o};
