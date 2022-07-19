<template>

  <page-header-wrapper
    :title="false"
    :breadcrumb="{}"
  >
    <template slot="content">
      <div style="padding: 0px 20px;">
        <a-row >
          <a-col :xs="8" :sm="8" :xl="4" class="header-col">
            <span class="header-card-title"><a-icon type="team"/> 总计</span><br/>
            <span class="header-card-value">{{ status.length }}</span><br/>
          </a-col>
          <a-col :xs="8" :sm="8" :xl="4" class="header-col">
            <span class="header-card-title"><a-icon type="alert" /> 预警中</span><br/>
            <span class="header-card-value">{{ status.alert }}</span><br/>
          </a-col>
          <a-col :xs="8" :sm="8" :xl="4" class="header-col">
            <span class="header-card-title">
              <a-tooltip title="节点离线，服务正处于Starting、Running、Stopping，不确定将要转换的状态">
                <a-icon type="question-circle"/>
              </a-tooltip> 未知
            </span><br/>
            <span class="header-card-value">{{ status.unknown }}</span><br/>
          </a-col>
          <a-col :xs="8" :sm="8" :xl="4" class="header-col">
            <span class="header-card-title"><a-icon type="loading"/> 运行中</span><br/>
            <span class="header-card-value" style="color:#1ABB9C;">{{ status.running }}</span>
            <span v-show="status.starting > 0 "><a-icon type="caret-left" />{{ status.starting }}</span>
          </a-col>
          <a-col :xs="8" :sm="8" :xl="4" class="header-col">
            <span class="header-card-title"><a-icon type="stop"/> 已停止</span><br/>
            <span class="header-card-value">{{ status.stopped }}</span>
            <span v-show="status.stopping > 0 "><a-icon type="caret-left" />{{ status.stopping }}</span>
          </a-col>
          <a-col :xs="8" :sm="8" :xl="4" class="header-col">
            <span class="header-card-title"><a-icon type="close"/> 已失败</span><br/>
            <span class="header-card-value" style="color:red">{{ status.exited }}</span><br/>
          </a-col>
        </a-row>
      </div>
    </template>

    
    <a-card :bordered="false">
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
      <div style="marginLeft:6px;marginBottom:10px;font-size: 16px;">
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
      <div :style="{ marginBottom: '24px' }">
        <a-button
          type="primary"
          value="small"
          @click="openEdit(null,'create')"
          icon="plus"
        >创建命令</a-button>
      </div>
      <s-table
        rowKey="name"
        ref="table"
        size="default"
        data-name="dataList"
        :columns="columns"
        :data="loadProcess"
      >
        <template slot="name" slot-scope="text" >
          <a-tooltip v-if="text.length > 10" :title="text">
            {{ text.slice(0, 10) + '...' }}
          </a-tooltip>
          <span v-else>{{ text }}</span>
        </template>
        <template slot="context" slot-scope="text" >
          <a-tooltip v-if="text.length > 30" :title="text" style="color:#00BFFF">
            {{ text.slice(0, 30) + '...' }}
          </a-tooltip>
          <a-tooltip v-else :title="text" style="color:#00BFFF">
            {{ text }}
          </a-tooltip>
          <a-tooltip >
            <template slot="title">{{ copyTitle }}</template>
            <a-icon
              type="copy"
              v-clipboard:copy="text"
              @click="copyClick"/>
          </a-tooltip>
        </template>

        <template slot="action" slot-scope="text, record">
          <div >
            <a @click="execPage(record)">执行</a>
            <a-divider type="vertical" />
            <a @click="showEdit('edit',record)">修改</a>
            <a-divider type="vertical" />
            <a @click="logPage(record)">日志</a>
            <a-divider type="vertical" />
            <a-popconfirm title="确定要删除吗？" @confirm="deleteCmd(record)">
              <a-icon slot="icon" type="question-circle-o" style="color: red" />
              <a style="color:red;">删除</a>
            </a-popconfirm>
          </div>
        </template>
      </s-table>
    </a-card>
    
  </page-header-wrapper>
</template>
<script>
import { tags, processList, processDelete, processStart, processStop } from '@/api/process'
import STable from '@/components/Table'
import moment from 'moment'

const columns = [
  {
    title: 'ID',
    dataIndex: 'name',
    scopedSlots: { customRender: 'name' }
  },
  {
    title: '命令内容',
    dataIndex: 'context',
    scopedSlots: { customRender: 'context' }
  },
  {
    title: '执行次数',
    align: 'center',
    customRender: (text) => text + ' 次',
    dataIndex: 'call_no'
  },
  {
    title: '创建时间',
    dataIndex: 'create_at',
    customRender: (text) => moment.unix(text).format('YYYY-MM-DD HH:mm:ss')
  },
  {
    title: '最后修改人',
    dataIndex: 'user'
  },
  {
    title: '最后修改时间',
    dataIndex: 'update_at',
    customRender: (text) => moment.unix(text).format('YYYY-MM-DD HH:mm:ss')
  },
  {
    title: '操作',
    scopedSlots: { customRender: 'action' }
  }
]
export default {
  name: 'ProcessList',
  components:{
    STable,
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
      status: {
        alert: 0,
        unknown: 0,
        starting: 0,
        running: 0,
        exited: 0,
        stopping: 0,
        stopped: 0,
        length: 0,
        process: []
      },
      ticker: null
    }
  },
  mounted () {
    console.log(this.$route)
    if (this.$route.params.path) {
      this.path = this.$route.params.path
    }
    this.loadTags()
    this.ticker = setInterval(() => {
      this.$refs.table.refresh()
    }, 2000)
  },
  filters: {
    showAge (time) {
      // const age = moment().unix() - time
      // return moment.unix(age).format('YYYY-MM-DD hh:mm:ss')
      return moment.unix(time).fromNow(true)
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
      for (let v in this.selectedTags.nodes){
        nodes[v] = {}
      }
      for (let v in this.selectedTags.labels){
        labels[v] = {}
      }
      const args = { nodes:nodes,labels:labels,...parameter}
      // console.log(args);
      return processList(args).then(res => {
        this.status = { alert: 0, unknown: 0, starting: 0, running: 0, exited: 0, stopped: 0, stopping: 0, length: 0, process: [] }
        this.status.process.push({})
        // console.log(res)
        for (const v in res.data) {
          console.log(v)
          this.status.length += 1
          this.status.process.push(v)
          if (v.state.status === 'Unknown') {
            this.status.unknown += 1
          } else if (v.state.status === 'Running') {
            this.status.running += 1
          } else if (v.state.status === 'Starting') {
            this.status.starting += 1
          } else if (v.state.status === 'Stopping') {
            this.status.stopping += 1
          } else if (v.state.status === 'Stopped') {
            this.status.stopped += 1
          } else {
            this.status.exited += 1
          }
        }
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
    border-left:2px solid #ADB2B5;
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
