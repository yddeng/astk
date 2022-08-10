<template>
  <div>
  
    <a-card :bordered="false">
      <a-row :style="{ marginBottom: '20px' }" justify="space-between" type="flex">
        <a-col>
          <a-button type="primary" @click="openEdit()" icon="plus">新建</a-button>
        </a-col>
        <a-col :style="{ marginRight: '20px' }">
          <a-button @click="()=>{this.$refs.table.refresh()}"><a-icon type="sync"/></a-button>
        </a-col>
      </a-row>
      
      <s-table
        rowKey="id"
        ref="table"
        size="middle"
        data-name="dataList"
        :isLoading="false"
        :columns="columns"
        :data="loadData"
      >
        <template slot="address" slot-scope="text">
          <a-tooltip v-if="text.length > 30" :title="text">
            {{ text.slice(0, 30) + '...' }}
          </a-tooltip>
          <span v-else>{{ text }}</span>
        </template>
        <template slot="webHook" slot-scope="text">
          <a-tooltip v-if="text.length > 30" :title="text">
            {{ text.slice(0, 30) + '...' }}
          </a-tooltip>
          <span v-else>{{ text }}</span>
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
            <a-popconfirm title="确定要删除吗？" @confirm="gitDelete(record)">
              <a-icon slot="icon" type="question-circle-o" style="color: red" />
              <a style="color:red;">删除</a>
            </a-popconfirm>
          </div>
        </template>
      </s-table>
    </a-card>

    <a-edit
      :visible="editVisible"
      :node-names="nodeNames"
      @success="editSuccess"
      @cancel="editCancel"
    />
  </div>
</template>

<script>
import STable from '@/components/Table'
import { gitCreate, gitList,gitDelete } from '@/api/git'
import { nodeStatus } from '@/api/node'
import Edit from './modal/edit'

export default {
  name: 'CmdList',
  components: {
    STable,
    'a-edit':Edit,
  },
  data () {
    return {
      columns,
      // 查询条件参数
      // name: '',
      // 加载数据方法 必须为 Promise 对象
      loadData: parameter => {
       // const requestParameters = { keyword: this.name, ...parameter } //如果要查询要这样写
        return gitList(parameter).then(res => {
          console.log(res)
          return res
        })
      },
      nodeNames: [],

      openedLoading:false,

      editVisible:false,
      copyTitle: '复制',
    }
  },
  mounted () {
    this.loadNodeStatus()
  },
  methods: {
    loadNodeStatus () {
      nodeStatus()
        .then(res => {
          this.nodeNames = res.all
        })
    },
    gitCreate (record) {
      gitCreate({ id: record.id })
        .then(() => {
          this.$refs.table.refresh()
        })
    },
    gitDelete (record) {
      gitDelete({ id: record.id })
        .then(() => {
          this.$refs.table.refresh()
        })
    },
    openEdit(){
      this.editVisible = true
    },
    editSuccess(){
      this.$refs.table.refresh()
      this.editVisible = false
    },
    editCancel(){
      this.editVisible = false
    },
    copyClick () {
      this.copyTitle = '复制成功'
      setTimeout(() => {
        this.copyTitle = '复制'
      }, 1500)
    }
  }
}
const columns = [
  {title: '名称', dataIndex: 'name',width:'100px'},
  {title: '类型', dataIndex: 'type',width:'150px'},
  {title: '分支', dataIndex: 'branch',width:'150px'},
  {title: '仓库地址', dataIndex: 'address',width:'300px',
    scopedSlots: { customRender: 'address' }},
  {title: 'Hook地址',dataIndex: 'webHook',width:'300px',
    scopedSlots: { customRender: 'webHook' }},
  {title: 'Hook令牌',dataIndex: 'token', width:'150px'},
  {title: '通知类型',dataIndex: 'notify.notifyType', width:'150px'},
  {title: '操作',scopedSlots: { customRender: 'action' }}
]
</script>

<style lang="less" scoped>
</style>
