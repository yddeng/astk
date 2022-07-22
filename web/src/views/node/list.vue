<template>
  <div >
    <a-list
      :grid="{ gutter: 16, xs: 1, sm: 1, md: 1, lg: 2, xl:3, xxl :4 }"
      :data-source="nodes"
    >
      <a-list-item slot="renderItem" slot-scope="item">
        <a-card :title="item.name" size="small" :bordered="false" v-if="item.online">
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
        <a-card :title="item.name" size="small" :bordered="false" v-else>
          <a-row><a-col :span="4" :offset="1">内网</a-col><a-col :span="16" >{{ item.inet }}</a-col></a-row>
          <a-row><a-col :span="4" :offset="1">公网</a-col><a-col :span="16">{{ item.net }}</a-col></a-row>
          <a-row><a-col :span="4" :offset="1">CPU</a-col><a-col :span="16" ><a-progress :percent="0" /></a-col></a-row>
          <a-row><a-col :span="4" :offset="1">内存</a-col><a-col :span="16" ><a-progress :percent="0" /></a-col></a-row>
          <a-row><a-col :span="4" :offset="1">网络</a-col><a-col :span="16" ><a-icon type="arrow-down" />0B/s<a-icon type="arrow-up" />0B/s</a-col></a-row>
          <a-row><a-col :span="4" :offset="1">硬盘</a-col><a-col :span="16" ><a-progress :percent="0" /></a-col></a-row>
          <a-row><a-col :span="4" :offset="1">状态</a-col><a-col :span="16" ><span style="color:red">离线</span></a-col></a-row>
          <a slot="extra" @click="nodeRemove(item.name)"><a-icon style="color:red" type="delete"/></a>
        </a-card>
      </a-list-item>
    </a-list>

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

  </div>
</template>

<script>
import { nodeList, nodeRemove } from '@/api/node'

export default {
  name: 'Home',
  data () {
    return {
      totalCount:0,
      nodes: [],
      ticker: null,
      pageNo:1,
      pageSize:8,
    }
  },
  beforeMount () {
    this.loadNodes()
    this.ticker = setInterval(() => {
      this.loadNodes()
    }, 2000)
  },
  destroyed () {
    console.log('destroyed')
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
    }
  }

}
</script>

<style lang="less" scoped>
</style>
