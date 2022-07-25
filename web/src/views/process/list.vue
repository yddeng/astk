<template>
  <div>
    <a-card :bordered="false" size="small" style="marginBottom:10px;">
      <div style="marginLeft:6px;marginBottom:10px;font-size: 16px;">
        <span style="marginRight:12px">节点:</span>
        <template v-for="tag in tags.nodes">
          <a-checkable-tag
            :key="tag"
            style="font-size: 14px;"
            :checked="selectedTags.nodes.indexOf(tag) > -1"
            @change="checked => onTagChange('nodes',tag, checked)"
          >
            {{ tag }}
          </a-checkable-tag>
        </template>
      </div>
      <div style="marginLeft:6px;marginBottom:10px;font-size: 16px;">
        <span style="marginRight:12px">标签:</span>
        <template v-for="tag in tags.labels">
          <a-checkable-tag
            :key="tag"
            style="font-size: 14px;"
            :checked="selectedTags.labels.indexOf(tag) > -1"
            @change="checked => onTagChange('labels',tag, checked)"
          >
            {{ tag }}
          </a-checkable-tag>
        </template>
      </div>
      <div style="marginLeft:6px;font-size: 16px;">
        <span style="marginRight:12px">状态:</span>
        <template v-for="tag in tags.status">
          <a-checkable-tag
            :key="tag"
            style="font-size: 14px;"
            :checked="selectedTags.status.indexOf(tag) > -1"
            @change="checked => onTagChange('status',tag, checked)"
          >
            {{ tag }}
          </a-checkable-tag>
        </template>
      </div>
    </a-card>
    <a-card :bordered="false" style="marginBottom:10px;" size="small">
      <a-row justify="space-between" type="flex">
        <a-col>共有<span style="color:darkgoldenrod"> {{data.totalCount}} </span>个进程</a-col>
        <a-col style="marginRight:10px;">
          <!-- <a @click="()=>{this.loadProcess()}" ><a-icon type="sync" :rotate="45" /></a>
          <a-divider type="vertical" /> -->
          <a @click="openEdit(null,'create')" >新增进程</a>
          <a-divider type="vertical" />
          <a-dropdown >
            <a class="ant-dropdown-link" >
            批量操作<a-icon type="down" />
            </a>
            <a-menu slot="overlay">
              <a-menu-item key="0" @click="startAllProcess"> 
                <a >全部启动</a>
              </a-menu-item>
              <a-menu-item key="1" @click="stopAllProcess">
                <a >全部停止</a>
              </a-menu-item>
            </a-menu>
          </a-dropdown>
        </a-col>
      </a-row>
    </a-card>
    <!-- <div style="marginBottom:10px;">
     <a-row justify="space-between" type="flex">
        <a-col>按ID排序</a-col>
        <a-col style="marginRight:10px;">
          <a-radio-group :options="['启动的时间','CPU','Mem']" />
        </a-col>
        </a-row>
    </div>  -->
    <a-list
      :grid="{ gutter: 16, xs: 1, sm: 1, md: 2, lg: 2, xl:3, xxl :4 }"
      :data-source="data.process"
    >
      <a-list-item slot="renderItem" slot-scope="item">
          <a-card :title="item.name" size="small" :bordered="false">
            <template slot="extra" >
              <a-icon type="bell" style="color:#2894FF" @click="processBell" />
            </template>

            <a-row style="height:24px;line-height:24px">
              <a-col :span="3">节点</a-col>
              <a-col :span="20">{{ item.node }}</a-col>
            </a-row>
            <a-row style="height:24px;line-height:24px">
              <a-col :span="3">重启</a-col>
              <a-col :span="20">{{ item.autoStartTimes }}次</a-col>
            </a-row>
            <a-row style="height:24px;line-height:24px">
              <a-col :span="3">状态</a-col>
              <a-col :span="20">
                <a-popconfirm v-if="item.state.status==='exited'" placement="topRight" trigger="click">
                  <span slot="title" style="white-space:pre-wrap;">{{ item.state.exitMsg }}</span>
                  <a-tag :color="tagStatusColor(item.state.status)">{{item.state.status}}</a-tag>
                </a-popconfirm>
                <a-tag v-else :color="tagStatusColor(item.state.status)">{{item.state.status}}</a-tag>
                <span v-if="item.state.status==='running'">
                  Pid:{{ item.state.pid }}, Age: {{ item.state.timestamp | showAge }}
                </span> 
              </a-col>
            </a-row>
            <a-row style="height:24px;line-height:24px">
              <a-col :span="3">CPU</a-col>
              <a-col :span="15">
                <a-progress
                  :stroke-color="progressColor(item.state.cpu)"
                  :percent="parseFloat(item.state.cpu.toFixed(2))" 
                  status="normal"/>
                  
              </a-col>
            </a-row>
            <a-row style="height:24px;line-height:24px">
              <a-col :span="3">内存</a-col>
              <a-col :span="15">
                <a-progress
                  status="normal"
                  :stroke-color="progressColor(item.state.mem)"
                  :percent="parseFloat(item.state.mem.toFixed(2))" />
              </a-col>
            </a-row>
          
            <template slot="actions" >
              <a v-if="item.state.status === 'exited' || item.state.status === 'stopped'" @click="startProcess(item.id)">启动</a>
              <a v-else-if="item.state.status === 'running'" @click="stopProcess(item.id)">停止</a>
              <a @click="openEdit(item,'edit')">配置</a>
              <a v-if="item.state.status === 'running' || item.state.status === 'starting' || item.state.status === 'stopping'" 
                @click="tailLogStart(item.id,item.name)">Tailf</a>
              
              <a-popconfirm
                v-if="item.state.status === 'exited' || item.state.status === 'stopped'"
                title="确定要删除吗？"
                @confirm="deleteProcess(item.id)">
                <a-icon slot="icon" type="question-circle-o" style="color: red" />
                <a href="#"> 删除</a>
              </a-popconfirm>
              <a-dropdown>
                <a href="javascript:;">
                  <a-icon type="ellipsis" />
                </a>
                <a-menu slot="overlay">
                  <a-menu-item>
                    <a @click="openEdit(item,'copy')">复制</a>
                  </a-menu-item>
                </a-menu>
              </a-dropdown>
            </template>
          </a-card>
      </a-list-item>
    </a-list>
    <div >
      <a-row justify="center" type="flex"><a-col>
      <a-pagination
        :current="pageNo"
        :pageSize="pageSize"
        :total="data.totalCount"
        @change="onPageChange"
        hideOnSinglePage
      />
      </a-col></a-row>
    </div>

    <a-log
      :visible="tailVisible"
      :context="tail.context"
      :title="tailTitle"
      @cancel="tailLogCancel"
    >
    </a-log>
  </div>
</template>
<script>
import { tags, processList, processDelete, processStart, processStop,
processBatchStart, processBatchStop,processTail } from '@/api/process'
import moment from 'moment'
import Log from './modal/log'

export default {
  name: 'ProcessList',
  components:{
    'a-log': Log,
  },
  data () {
    return {
      tags: {
        nodes:[],
        labels:[],
        status:['unknown','starting','running','exited','stopping','stopped']
      },
      selectedTags:{
        nodes:[],
        labels:[],
        status:[]
      },
      data: {
        totalCount: 0,
        process:[],
      },
      ticker: null,

      tail:{
        id :0,
        start:0,
        context:'',
      },
      tailTitle:'',
      tailVisible:false,
      tailTicker:null,

      pageNo:1,
      pageSize:8,
    }
  },
  mounted () {
    if (this.$route.params.path) {
      this.path = this.$route.params.path
    }
    this.loadTags()
    this.loadProcess()
    this.ticker = setInterval(() => {
      this.loadProcess()
    }, 2000)
  },
  filters: {
    showAge (time) {
      return moment.unix(time).fromNow(true)
    },
    statusIn18(status){
      const s = ['unknown','starting','running','exited','stopping','stopped']
      const m = ['未知','启动中','运行中','报错','停止中','已停止']
      const i = s.indexOf(status)
      if (i !== -1){
        return m[i]
      }
      return m[0]
    }
  },
  destroyed () {
    console.log('destroyed')
    clearInterval(this.ticker)
    clearInterval(this.tailTicker)
  },
  methods: {
    loadTags () {
      tags()
      .then(res => {
        this.tags.nodes = []
        this.tags.labels = []
        for (let k in res.nodes){
          this.tags.nodes.push(k)
        }
        for (let k in res.labels){
          this.tags.labels.push(k)
        }
        // console.log(res,this.tags);
      })
    },
    onTagChange (tt,tag,checked) {
      if (tt ==='nodes'){
        const nextSelectedTags = checked  ? [...this.selectedTags.nodes, tag] : this.selectedTags.nodes.filter(t => t !== tag);
        this.selectedTags.nodes = nextSelectedTags;
      }else if (tt ==='labels'){
        const nextSelectedTags = checked  ? [...this.selectedTags.labels, tag] : this.selectedTags.labels.filter(t => t !== tag);
        this.selectedTags.labels = nextSelectedTags;
      }else if (tt ==='status'){
        const nextSelectedTags = checked  ? [...this.selectedTags.status, tag] : this.selectedTags.status.filter(t => t !== tag);
        this.selectedTags.status = nextSelectedTags;
      }

      this.loadProcess()
    },
    makeLabelArgs(){
      let nodes = {}
      let labels = {}
      let status = {}
      for (let v of this.selectedTags.nodes){
        nodes[v] = {}
      }
      for (let v of this.selectedTags.labels){
        labels[v] = {}
      }
      for (let v of this.selectedTags.status){
        status[v] = {}
      }
      const args = { nodes:nodes,labels:labels,status:status}
      return args
    },
    loadProcess () {
      const labels = this.makeLabelArgs()
      const args = {...labels,pageNo:this.pageNo,pageSize:this.pageSize}
      return processList(args).then(res => {
        this.data={totalCount:res.totalCount,process:res.dataList}
        // console.log(res,this.data);
        return res
      })
    },
    tailLogStart(id,name){
      this.tail = {id:id,start:0,context:''}
      this.tailTitle=name
      this.tailVisible = true
      this.processTailLoop()
      this.tailTicker = setInterval(() => {
        this.processTailLoop()
      }, 1000)
    },
    processTailLoop(){
      processTail(this.tail).then(res => {
        this.tail.start = res.end
        if (res.context !==''){
          this.tail.context += res.context
        }
        console.log(res,this.tail)
      })
    },
    tailLogCancel(){
      this.tailVisible = false
      this.tail = {id:0,start:0,context:''}
      clearInterval(this.tailTicker)
      console.log('tailLogCancel');
    },
    onPageChange(page){
      this.pageNo = page
      this.loadProcess()
    },
    startProcess (id) {
      processStart({ id: id }).then(() => {
        this.loadProcess()
      })
    },
    stopProcess (id) {
      processStop({ id: id }).then(() => {
        this.loadProcess()
      })
    },
    deleteProcess (id) {
      processDelete({ id: id }).then(() => {
        this.loadProcess()
      })
    },
    openEdit (item, option) {
      console.log(option, item)
      this.$router.push({ name: 'pedit', params: { option: option, labels: this.tags.labels, item: { ...item } } })
    },
    startAllProcess () {
      const args = this.makeLabelArgs()
      processBatchStart(args).then(() => {
        this.loadProcess()
      })
    },
    stopAllProcess () {
      const args = this.makeLabelArgs()
      processBatchStop(args).then(() => {
        this.loadProcess()
      })
    },
    tagStatusColor(status){
      const m = ['#E0E0E0','#01B468','#1ABB9C','#FF7575','#D0D0D0','#D0D0D0']
      const i = this.tags.status.indexOf(status)
      return i !== -1 ? m[i]: m[0]
    },
    progressColor (percent) {
      if (percent >= 80) {
        return 'red'
      } else if (percent >= 50) {
        return '#EAC100'
      }
    },
    processBell () {
      this.$message.info('暂未实现该功能')
    }
  }
}
</script>

<style lang="less" scoped>
</style>
