import Vue from 'vue'
import App from './App.vue'
import Antd from 'ant-design-vue';
import router from './router'
import  storage from 'store'
import 'ant-design-vue/dist/antd.css';

import './global.css'

// http://momentjs.cn/docs/#/use-it/
import moment from 'moment';
moment.locale('zh-cn');

Vue.config.productionTip = false
Vue.use(Antd);


import VueClipboard from 'vue-clipboard2'
Vue.use(VueClipboard)

import ProLayout, { PageHeaderWrapper } from '@ant-design-vue/pro-layout'
Vue.component('pro-layout', ProLayout)
Vue.component('page-container', PageHeaderWrapper)
Vue.component('page-header-wrapper', PageHeaderWrapper)



const allowList = ['login'] // no redirect allowList
const loginRoutePath = '/auth/login'

router.beforeEach((to, from, next) => {
  if (to.meta.title){
    document.title = to.meta.title
  }

  const token = storage.get("Access-Token")
  if (token) {
    next()
  } else {
    if (allowList.includes(to.name)) {
      // 在免登录名单，直接进入
      next()
    } else {
      next({ path: loginRoutePath, query: { redirect: to.fullPath } })
    }
  }
})

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
