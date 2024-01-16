import Vue from "vue";
import VueRouter from "vue-router";
import demo from "../views/demo.vue"
import Home from "../views/Home.vue";
import stockview from "../views/stockview";
import dashboard from "../views/dashboard.vue"
import loginhistory from "../views/loginHistory.vue"
import adduser from "../views/AddUser.vue";
import stockentry from "../views/StockEntry.vue"
import billentry from "../views/BillEntry.vue"
import salesreport from "../views/salesReport.vue"
import admininventry from "../components/dashboard/adminInventrydashboard.vue"
import manager from "../components/dashboard/managerdashboard.vue"
import salesman from "../components/dashboard/salesmanDashboard.vue"
import dailysales from "../components/apexchart/dailysales"
Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/dailysales",
    name: "dailysales",
    component: dailysales,
  },
  {
    path: "/dashboard",
    name: "dashboard",
    component: dashboard,
  },
  {
    path: "/loginhistory",
    name: "loginhistory",
    component: loginhistory,
  },
  {
    path: "/adduser",
    name: "adduser",
    component: adduser,
  },
  {
    path: "/stockentry",
    name: "stockentry",
    component: stockentry,
  },
  {
    path: "/billentry",
    name: "billentry",
    component: billentry,
  },
  {
    path: "/salesreport",
    name: "salesreport",
    component: salesreport,
  },
  {
    path: "/admininventry",
    name: "admininventry",
    component: admininventry,
  },
  {
    path: "/manager",
    name: "manager",
    component: manager,
  },
  {
    path: "/salesman",
    name: "salesman",
    component: salesman,
  },
  {
    path: "/stockview",
    name: "stockview",
    component: stockview,
  },
  {
    path: "/demo",
    name: "demo",
    component: demo,
  },
];

const router = new VueRouter({
  routes,
});

export default router;
