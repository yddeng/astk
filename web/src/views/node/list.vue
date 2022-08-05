<template>
  <div >

    <a-card  :bordered="false">
      <div style="margin-bottom:16px">
        <a-row justify="space-between" type="flex">
          <a-col :span="18">
            共有<span style="color:darkgoldenrod"> {{ totalCount }} </span>个服务节点
          </a-col>
          <a-col >
            <a-button icon="bell" size="small" @click="openDrawer">监控</a-button>
          </a-col>
        </a-row>
      </div>
    
    <a-list
      :grid="{ gutter: 16, xs: 1, sm: 1, md: 1, lg: 2, xl:3, xxl :4 }"
      :data-source="nodes"
    >
      <a-list-item slot="renderItem" slot-scope="item">
        <a-card :title="item.name" size="small"  v-if="item.online && item.state">
          <a-row><a-col :span="4" :offset="1">内网</a-col><a-col :span="16" >{{ item.inet }}</a-col></a-row>
          <a-row><a-col :span="4" :offset="1">公网</a-col><a-col :span="16">{{ item.net }}</a-col></a-row>
          <a-row>
            <a-col :span="4" :offset="1">CPU</a-col>
            <a-col :span="16" >
              <a-progress
                status="normal"
                :stroke-color="progressColor(item.state.cpu.usedPercent)"
                :percent="parseFloat(item.state.cpu.usedPercent)" />
            </a-col>
          </a-row>
          <a-row>
            <a-col :span="4" :offset="1">内存</a-col>
            <a-col :span="16" >
              <a-progress
                status="normal"
                :stroke-color="progressColor(item.state.mem.virtualUsedPercent)"
                :percent="parseFloat(item.state.mem.virtualUsedPercent)" />
            </a-col>
          </a-row>
          <a-row>
            <a-col :span="4" :offset="1">网络</a-col>
            <a-col :span="16" >
              <a-icon type="arrow-down" />{{ item.state.net.recentBytesRecv }} <a-icon type="arrow-up" />{{ item.state.net.recentBytesSent }}
            </a-col>
          </a-row>
          <a-row>
            <a-col :span="4" :offset="1">硬盘</a-col>
            <a-col :span="16" >
              <a-progress
                status="normal"
                :stroke-color="progressColor(item.state.disk.usedPercent)"
                :percent="parseFloat(item.state.disk.usedPercent)" />
            </a-col>
          </a-row>
          <a-row><a-col :span="4" :offset="1">状态</a-col><a-col :span="16" ><span style="color:#3CB371">在线</span></a-col></a-row>
          <template slot="extra">
            <a-tooltip
              :arrowPointAtCenter="true"
              placement="bottomLeft">
              <template slot="title" >
                系统：{{ item.state.host.hostname }} [{{ item.state.host.os }}-{{ item.state.host.arch }}]<br/>
                CUP核心数: {{ item.state.cpu.cpuCores }}<br/>
                硬盘'/'：{{ item.state.disk.used }}/{{ item.state.disk.total }}<br/>
                内存：{{ item.state.mem.virtualUsed }}/{{ item.state.mem.virtualTotal }}<br/>
                交换：{{ item.state.mem.swapUsed }}/ {{ item.state.mem.swapTotal }}<br/>
                网络流量：<a-icon type="arrow-down" />{{ item.state.net.totalBytesRecv }} <a-icon type="arrow-up" />{{ item.state.net.totalBytesSent }}<br/>
                网络包：<a-icon type="arrow-down" />{{ item.state.net.totalPacketsRecv }} <a-icon type="arrow-up" />{{ item.state.net.totalPacketsSent }}<br/>
                连接数：TCP {{ item.state.net.tcpConnections }} | UDP {{ item.state.net.udpConnections }}<br/>
              </template>
              <a-icon type="exclamation-circle" />
            </a-tooltip>
          </template>
        </a-card>
        <a-card :title="item.name" size="small" v-else>
          <a-row><a-col :span="4" :offset="1">内网</a-col><a-col :span="16" >{{ item.inet }}</a-col></a-row>
          <a-row><a-col :span="4" :offset="1">公网</a-col><a-col :span="16">{{ item.net }}</a-col></a-row>
          <a-row><a-col :span="4" :offset="1">CPU</a-col><a-col :span="16" ><a-progress :percent="0" /></a-col></a-row>
          <a-row><a-col :span="4" :offset="1">内存</a-col><a-col :span="16" ><a-progress :percent="0" /></a-col></a-row>
          <a-row><a-col :span="4" :offset="1">网络</a-col><a-col :span="16" ><a-icon type="arrow-down" />0B/s<a-icon type="arrow-up" />0B/s</a-col></a-row>
          <a-row><a-col :span="4" :offset="1">硬盘</a-col><a-col :span="16" ><a-progress :percent="0" /></a-col></a-row>
          <a-row><a-col :span="4" :offset="1">状态</a-col><a-col :span="16" ><span style="color:red">离线</span></a-col></a-row>
          <a-popconfirm
            slot="extra"
            title="确定要删除吗？"
            @confirm="nodeRemove(item.name)">
            <a-icon type="delete"/>
          </a-popconfirm>
        </a-card>
      </a-list-item>
    </a-list>

    <a-drawer
      title="监控"
      placement="right"
      width="300"
      :visible="monitorVisible"
      @close="()=>{monitorVisible = false}"
    >
      CPU:<span style="color:darkgoldenrod"> {{monitorInfo.cpu}} </span>
      <a-slider :min="0" :max="100" v-model="monitorInfo.cpu" @afterChange="monitorUpdate"/>
      Mem:<span style="color:darkgoldenrod"> {{monitorInfo.mem}} </span>
      <a-slider :min="0" :max="100" v-model="monitorInfo.mem" @afterChange="monitorUpdate"/>
      Disk:<span style="color:darkgoldenrod"> {{monitorInfo.disk}} </span>
      <a-slider :min="0" :max="100" v-model="monitorInfo.disk" @afterChange="monitorUpdate"/>
      触发间隔:<span style="color:darkgoldenrod"> {{monitorInfo.interval}} </span>秒
      <a-slider :min="2" :max="20" 
      v-model="monitorInfo.interval" 
      :tip-formatter="v => {return `${v}s`}"  
      @afterChange="monitorUpdate"/>
      报警间隔:<span style="color:darkgoldenrod"> {{monitorInfo.continuityInterval/60}} </span>分
      <a-slider :min="600" :max="7200" :step="600" 
      v-model="monitorInfo.continuityInterval" 
      :tip-formatter="v => {return `${v/60}m`}"
      @afterChange="monitorUpdate"/><br/>
      通知方式<br/>
      <a-select v-model="monitorInfo.notify.notifyType" style="min-width:120px">
        <a-select-option value="weixin">企业微信</a-select-option>
        <a-select-option value="callback">自定义</a-select-option>
      </a-select><br/>
      <a-input 
        v-model="monitorInfo.notify.notifyServer" 
        placeholder="http://example.com/posttreceive"/>
      <br/><br/>
      <a-switch 
        checked-children="开启" 
        un-checked-children="关闭" 
        v-model="monitorInfo.opened" 
        :loading="monitorOpenedLoading"
        @click="monitorOpened"/>
    </a-drawer>

    <div>
      <a-row justify="center" type="flex"><a-col>
      <a-pagination
        :current="pageNo"
        :pageSize="pageSize"
        :total="totalCount"
        @change="onPageChange"
        hideOnSinglePage
      />
      </a-col></a-row>
    </div>
    </a-card>
  </div>
</template>

<script>
import { nodeList, nodeRemove } from '@/api/node'
import { monitorInfo, monitorUpdate } from '@/api/monitor'

export default {
  name: 'Node',
  data () {
    return {
      totalCount:0,
      nodes: [],
      ticker: null,
      pageNo:1,
      pageSize:8,

      monitorVisible:false,

      monitorType:'node',
      monitorOpenedLoading:false,
      monitorInfo:{
        cpu:90,
        mem:90,
        disk:90,
        interval:10,
        continuityInterval:600,
        opened:false,
        notify:{
          notifyType:'',
          notifyServer:''
        }
      }
    }
  },
  beforeMount () {
    this.loadNodes()
    this.ticker = setInterval(() => {
      this.loadNodes()
    }, 2000)
  },
  destroyed () {
    clearInterval(this.ticker)
  },
  methods: {
    loadNodes () {
      const args = { pageNo: this.pageNo, pageSize: this.pageSize }
      nodeList(args)
        .then(res => {
          // console.log(res);
          this.totalCount = res.totalCount
          this.nodes = res.dataList
        })
    },
    onPageChange(page){
      this.pageNo = page
      this.loadNodes()
    },
    nodeRemove (name) {
      nodeRemove({ name: name })
        .then(() => {
          this.loadNodes()
        })
    },
    progressColor (percent) {
      if (percent >= '80') {
        return 'red'
      } else if (percent >= '50') {
        return '#EAC100'
      }
    },
    openDrawer(){
      this.loadMonitor()
      this.monitorVisible = true
    },
    loadMonitor(){
      monitorInfo({type:this.monitorType})
      .then(res => {
        //console.log(res);
        this.monitorInfo = res
      })
      .finally(() => {
        this.monitorOpenedLoading = false
      })
    },
    monitorUpdate(){
      const args = {
        type:this.monitorType,
        monitor:this.monitorInfo,
      }
      //console.log(args);
      monitorUpdate(args).then(()=>{
        this.loadMonitor()
      })
    },
    monitorOpened(opened){
      if (opened && (this.monitorInfo.notify.notifyType ==='' || 
        this.monitorInfo.notify.notifyServer ==='')){
        this.$message.error('通知类型及url需填写完整')
        this.monitorInfo.opened = false
        return
      }
      this.monitorOpenedLoading = true
      this.monitorUpdate()
    },
    
  }

}
</script>

<style lang="less" scoped>
</style>
