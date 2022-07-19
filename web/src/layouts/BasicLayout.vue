<template>
  <a-layout id="components-layout-demo-custom-trigger">
    <a-layout-sider 
    v-model="collapsed" 
    :trigger="null" 
    :width="width"
    :style="style"
    collapsible>
      <div class="logo">
        <h1>{{title}}</h1>
      </div>
      <a-menu 
      theme="dark" 
      mode="inline" 
      :default-selected-keys="selectedKeys"
      >
        <template v-for="item in menus">
          <template v-if="!item.meta || !item.meta.hidden">
            <a-menu-item v-if="!item.children" :key="item.name" @click="menuClick(item.name)">
              <a-icon v-if="item.meta && item.meta.icon!==''" :type="item.meta.icon" />
              <span>{{item.name}}</span>
            </a-menu-item>
            <sub-menu v-else :key="item.name" :menu-info="item" />
          </template>
        </template>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <a-icon
          class="trigger"
          :type="collapsed ? 'menu-unfold' : 'menu-fold'"
          @click="() => (collapsed = !collapsed)"
        />
      </a-layout-header>
      <a-layout-content
        :style="{ margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' }"
      >
        <router-view/>
      </a-layout-content>
      <a-layout-footer style="text-align: center">
        Ant Design ©2018 Created by Ant UED
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>
<script>
import { asyncRouterMap } from '@/router/'
import SubMenu  from '@/components/SubMenu'
export default {
  components: {
    'sub-menu': SubMenu,
  },
  data() {
    return {
      title:'ASTK',
      collapsed: false,
      width:256,
      style:{
        height: '100vh',
      },
      menus:[],
      selectedKeys:[],
    };
  },
  created(){
    const routes = asyncRouterMap.find((item) => item.path === '/')
    this.menus = (routes && routes.children) || []
    console.log(this.menus,this.$route.name);
    this.selectedKeys = [this.$route.name]
  },
  methods:{
    menuClick(name){
      this.$router.push({ name: name })
    },
  }
};
</script>
<style>
#components-layout-demo-custom-trigger .trigger {
  font-size: 20px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
}

#components-layout-demo-custom-trigger .trigger:hover {
  color: #1890ff;
}

#components-layout-demo-custom-trigger .logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px;
}
</style>