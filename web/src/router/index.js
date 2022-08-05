import Vue from 'vue'
import Router from 'vue-router'

const originalPush = Router.prototype.push
Router.prototype.push = function push (location, onResolve, onReject) {
  if (onResolve || onReject) return originalPush.call(this, location, onResolve, onReject)
  return originalPush.call(this, location).catch(err => err)
}

Vue.use(Router)

/*

{ Route } 对象
参数	说明	类型	默认值
hidden	控制路由和子路由是否显示在 sidebar	boolean	false
redirect	重定向地址, 访问这个路由时,自定进行重定向	string	-
name	路由名称, 必须设置,且不能重名	string	-
meta	路由元信息（路由附带扩展信息）	object	{}
hideChildrenInMenu	强制菜单显示为Item而不是SubItem(配合 meta.hidden)	boolean	-

{ Meta } 路由元信息对象
参数	说明	类型	默认值
title	路由标题, 用于显示面包屑, 页面标题 *推荐设置	string	-
icon	路由在 menu 上显示的图标	[string,svg]	-
keepAlive	缓存该路由 (开启 multi-tab 是默认值为 true)	boolean	false
hiddenHeaderContent	*特殊 隐藏 PageHeader 组件中的页面带的 面包屑和页面标题栏	boolean	false
permission	与项目提供的权限拦截匹配的权限，如果不匹配，则会被禁止访问该路由页面	array	[]
*/

export const asyncRouterMap = [
  {
    path: '/',
    component:() =>import('@/layouts/ProBasicLayout'),
    redirect: '/node',
    children:[
      {
        name:'node',
        path: '/node',
        meta:{title:'节点状态',icon:'cloud-server'},
        component: () => import('@/views/node/list')
      },
      {
        name:'cmdlist',
        path: '/command/list',
        meta:{ title:'计划任务',icon:'code'},
        component: () => import('@/views/command/cmdlist')
      },
      {
        name: 'cmdexec',
        hidden: true,
        meta: { title: '执行命令' },
        component: () => import('@/views/command/cmdexec'),
        path: '/command/exec'
      },
      {
        name: 'cmdlog',
        hidden: true,
        meta: { title: '命令日志' },
        component: () => import('@/views/command/cmdlog'),
        path: '/command/log'
      },
      {
        name:"plist",
        path: '/process/list',
        meta:{title:'应用服务',icon:'project'},
        component: () => import('@/views/process/list')
      },
      {
        name:'git',
        path:'/git',
        meta:{title:'代码发布',icon:'github'},
        component: () => import('@/layouts/BlankLayout'),
      },
      {
        name:'intranet',
        path:'/intranet',
        meta:{title:'内网穿透',icon:'link'},
        component: () => import('@/views/inc/list'),
      }
    ]
  },
  {
    path: '*',
    redirect: '/404'
  }
]

export const constantRouterMap = [
  {
    name:"login",
    path: '/auth/login',
    component: () => import('@/views/Login')
  },
  {
    path: '/404',
    meta:{title:'404'},
    component: () => import('@/views/404')
  }
]

export default new Router({
  mode: 'history',
  routes: constantRouterMap.concat(asyncRouterMap)
})