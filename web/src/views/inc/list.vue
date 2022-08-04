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
        <template slot="status" slot-scope="text, record" >
          <a-switch 
          size="small"
          checked-children="开" 
          un-checked-children="关" 
          v-model="record.opened" 
          :loading="openedLoading"
          @click="(checked)=>{incOpened(record,checked)}"/>
        </template>
        
        <template slot="action" slot-scope="text, record">
          <div >
            <a-popconfirm title="确定要删除吗？" @confirm="incDelete(record)">
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
import { incCreate, incList,incDelete,incOpened } from '@/api/inc'
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
        return incList(parameter).then(res => {
          console.log(res)
          return res
        })
      },
      nodeNames: [],

      openedLoading:false,

      editVisible:false,
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
    incCreate (record) {
      incCreate({ id: record.id })
        .then(() => {
          this.$refs.table.refresh()
        })
    },
    incOpened (record,opened) {
      this.openedLoading = true,
      incOpened({ id: record.id,opened:opened })
        .finally(() => {
          this.$refs.table.refresh()
          this.openedLoading = false
        })
    },
    incDelete (record) {
      incDelete({ id: record.id })
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
  }
}
const columns = [
  {title: '类型', dataIndex: 'type',width:'100px'},
  {title: '节点', dataIndex: 'node',width:'200px'},
  {title: '描述', dataIndex: 'desc',width:'400px',
    customRender: (text) => text.length > 20 ? text.slice(0, 20) + '...' :text,},
  {title: '转发端口', dataIndex: 'remotePort', align: 'center',width:'100px'},
  {title: '目标IP',dataIndex: 'localIp',width:'200px'},
  {title: '目标端口',dataIndex: 'localPort', align: 'center',width:'100px'},
  {title: '链接数',dataIndex: 'channel', align: 'center',width:'100px'},
  {title: '状态',scopedSlots: { customRender: 'status' },width:'100px'},
  {title: '操作',scopedSlots: { customRender: 'action' }}
]
</script>

<style lang="less" scoped>
</style>
