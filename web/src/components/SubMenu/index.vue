<template>
<a-sub-menu :key="menuInfo.name" v-bind="$props" v-on="$listeners">
  <span slot="title">
    <a-icon v-if="menuInfo.meta && menuInfo.meta.icon!==''" :type="menuInfo.meta.icon" />
    <span>{{menuInfo.name}}</span>
  </span>
  <template v-for="item  in menuInfo.children">
    <template v-if="!item.meta || !item.meta.hidden">
      <a-menu-item v-if="!item.children" :key="item.name" @click="menuClick(item.name)">
        <a-icon v-if="item.meta && item.meta.icon!==''" :type="item.meta.icon" />
      <span>{{item.name}}</span>
      </a-menu-item>
      <sub-menu v-else :menu-info="item" :key="item.name" />
    </template>
  </template>
</a-sub-menu>
</template>



<script>
import { Menu } from 'ant-design-vue';
export default{
  name: 'SubMenu',
  isRootMenu:true,
  props: {
    ...Menu.SubMenu.props,
    menuInfo: {
      type: Object,
      default: () => ({}),
    },
  },
  methods:{
    menuClick(name){
      this.$router.push({ name: name })
    }
  }
}
</script>
