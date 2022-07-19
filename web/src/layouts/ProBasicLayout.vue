<template>
  <pro-layout
    :menus="menus"
    :collapsed="collapsed"
    :mediaQuery="query"
    :isMobile="isMobile"
    :handleMediaQuery="handleMediaQuery"
    :handleCollapse="handleCollapse"
    v-bind="settings"
  >

    <template v-slot:menuHeaderRender>
      <div>
        <h1>{{ title }}</h1>
      </div>
    </template>
    
    <template v-slot:headerContentRender>
      <div>
        <a-tooltip title="刷新页面">
          <a-icon type="reload" style="font-size: 18px;cursor: pointer;" @click="() => { $message.info('只是一个DEMO') }" />
        </a-tooltip>
      </div>
    </template>

    <template v-slot:rightContentRender>
      
    </template>
    
    <router-view />
  </pro-layout>
</template>

<script>
import { asyncRouterMap } from '@/router'
export default {
  name: 'ProBasicLayout',
  data () {
    return {
       // base
      menus: [],
      // 侧栏收起状态
      collapsed: false,
      title: '应用服务器工具集',
      settings: {
        // 布局类型
        layout:  'sidemenu', 
        // CONTENT_WIDTH_TYPE
        contentWidth: 'Fluid',
        // 主题 'dark' | 'light'
        theme: 'dark',
        // 主色调
        primaryColor: '#1890ff',
        fixedHeader: true,
        fixSiderbar: true,
        colorWeak: false,

        hideHintAlert: false,
        hideCopyButton: false
      },
      // 媒体查询
      query: {},

      // 是否手机模式
      isMobile: false
    }
  },
  created () {
    const routes = asyncRouterMap.find((item) => item.path === '/')
    this.menus = (routes && routes.children) || []
  },

  methods: {
    handleMediaQuery (val) {
      this.query = val
      if (this.isMobile && !val['screen-xs']) {
        this.isMobile = false
        return
      }
      if (!this.isMobile && val['screen-xs']) {
        this.isMobile = true
        this.collapsed = true
      }
    },
    handleCollapse (val) {
      this.collapsed = val
    }
  }
}
</script>

<style>

</style>
