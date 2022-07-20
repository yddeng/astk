<template>
  <div>
    <a-card :bordered="false" style="marginBottom:20px;">
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
            {{ tag | statusIn18}}
          </a-checkable-tag>
        </template>
      </div>
    </a-card>
    <a-card :bordered="false" >
      <div style="marginBottom:10px;height: 32px;">
        <a-row justify="space-between" type="flex">
          <a-col> 
            共有<span style="color:darkgoldenrod"> {{data.totalCount}} </span>个进程 &nbsp;&nbsp;
            <a @click="()=>{this.$refs.table.refresh()}" ><a-icon type="sync" :rotate="45" /></a>
          </a-col>
          <a-col style="marginRight:10px;">
            <!-- <a @click="()=>{this.$refs.table.refresh()}" ><a-icon type="sync" :rotate="45" /></a>
            <a-divider type="vertical" /> -->
            <a @click="openEdit(null,'create')" >新增进程</a>
            <a-divider type="vertical" />
            <a-dropdown >
              <a class="ant-dropdown-link" @click="e => e.preventDefault()">
              批量操作<a-icon type="down" />
              </a>
              <a-menu slot="overlay">
                <a-menu-item key="0">
                  <a >全部启动</a>
                </a-menu-item>
                <a-menu-item key="1">
                  <a >全部停止</a>
                </a-menu-item>
              </a-menu>
            </a-dropdown>
          </a-col>
        </a-row>
      </div>
      <s-table
        rowKey="id"
        ref="table"
        size="default"
        data-name="dataList"
        :loading="false"
        :columns="columns"
        :data="loadProcess"
      >
        <template slot="name" slot-scope="text" >
          <a-tooltip v-if="text.length > 10" :title="text">
            {{ text.slice(0, 10) + '...' }}
          </a-tooltip>
          <span v-else>{{ text }}</span>
        </template>
        <template slot="node" slot-scope="text" >
          <a-tooltip v-if="text.length > 10" :title="text">
            {{ text.slice(0, 10) + '...' }}
          </a-tooltip>
          <span v-else>{{ text }}</span>
        </template>
        <template slot="status" slot-scope="text, record" >
          <a-tag :color="tagStatusColor(record.state.status)">{{record.state.status}}</a-tag>
          <span v-if="record.state.status==='running'">Pid:{{ record.state.pid }},Age: {{ record.state.timestamp | showAge }} </span>
        </template>
        <template slot="bell" >
          <a-switch checked-children="开" un-checked-children="关" default-checked />
        </template>
        <template slot="action" slot-scope="text, item">
          <div >
            <a v-show="item.state.status === 'exited' || item.state.status === 'stopped'" @click="startProcess(item.id)">启动</a>
            <a v-show="item.state.status === 'running'" @click="stopProcess(item.id)">停止</a>
            <a-divider type="vertical" />
            <a @click="openEdit(item,'edit')">配置</a>
            <template v-if="item.state.status === 'exited' || item.state.status === 'stopped'">
              <a-divider type="vertical" />
              <a-popconfirm title="确定要删除吗？" @confirm="deleteProcess(item.id)">
                <a-icon slot="icon" type="question-circle-o" style="color: red" />
                <a style="color:red;">删除</a>
              </a-popconfirm>
            </template>
          </div>
        </template>
      </s-table>
    </a-card>


  </div>
</template>
<script>
import { tags, processList, processDelete, processStart, processStop } from '@/api/process'
import STable from '@/components/Table'
import moment from 'moment'
const columns = [
  {
    title: '进程名',
    dataIndex: 'name',
    scopedSlots: { customRender: 'name' },
    width:'10%'
  },
  {
    title: '部署节点',
    dataIndex: 'node',
    scopedSlots: { customRender: 'node' },
    width:'10%'
  },
  {
    title: '状态',
    scopedSlots: { customRender: 'status' },
    width:'40%'
  },
  {
    title: '订阅报警',
    scopedSlots: { customRender: 'bell' },
    width:'10%'
  },
  {
    title: '操作',
    scopedSlots: { customRender: 'action' }
  }
]
export default {
  name: 'ProcessList',
  components:{
    's-table':STable,
  },
  data () {
    return {
      columns,
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
        process: []
      },
      loadInterval:2000,
      ticker: null,

      editVisible:false,

    }
  },
  mounted () {
    console.log(this.$route)
    if (this.$route.params.path) {
      this.path = this.$route.params.path
    }
    this.loadTags()
    // this.ticker = setInterval(() => {
    //   this.$refs.table.refresh()
    // }, 5000)
  },
  filters: {
    showAge (time) {
      // const age = moment().unix() - time
      // return moment.unix(age).format('YYYY-MM-DD hh:mm:ss')
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
        //console.log(res,this.tags);
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

      this.$refs.table.refresh()
    },
    loadProcess (parameter) {
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
      const args = { nodes:nodes,labels:labels,status:status,...parameter}
      return processList(args).then(res => {
        this.data={totalCount:res.totalCount}
        console.log(res)
        return res
      })
    },
    startProcess (id) {
      processStart({ id: id }).then(() => {
        this.$refs.table.refresh()
      })
    },
    stopProcess (id) {
      processStop({ id: id }).then(() => {
        this.$refs.table.refresh()
      })
    },
    deleteProcess (id) {
      processDelete({ id: id }).then(() => {
        this.$refs.table.refresh()
      })
    },
    openEdit (item, option) {
      console.log(option, item)
      this.$router.push({ name: 'pedit', params: { option: option, labels: this.tags.labels, item: { ...item } } })
    },
    startAllProcess () {
      for (const idx in this.status.process) {
        const item = this.status.process[idx]
        if (item.id !== undefined && (item.state.status === 'Stopped' || item.state.status === 'Exited')) {
          this.startProcess(item.id)
        }
      }
    },
    stopAllProcess () {
      for (const idx in this.status.process) {
        const item = this.status.process[idx]
        if (item.id !== undefined && item.state.status === 'Running') {
          this.stopProcess(item.id)
        }
      }
    },
    tagStatusColor(status){
      const m = ['#E0E0E0','#01B468','#1ABB9C','#FF7575','#D0D0D0','#F0F0F0']
      const i = this.tags.status.indexOf(status)
      if (i !== -1){
        return m[i]
      }
      return m[0]
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
  .header-card-value{
    font-size:40px;
    font-weight:bold;
    color:#7B7B7B;
  }

  .header-col{
    padding-left:10px;
  }

  .new-btn {
    background-color: #fff;
    border-radius: 2px;
    width: 100%;
    height: 233px;
  }

  .state_info{
  border: 1px solid #F0F0F0;
  border-radius:5px;
  font:14px;
  align:center;
  padding:0 5px;
  margin-right:5px;
}

.state_desc{
  border: 1px solid #F0F0F0;
  border-radius:5px;
  font:14px;
  align:center;
  padding:0 5px;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

</style>
