<template>
  <div>
    <a-card 
    :bordered="false" 
    :tab-list="tabList"
    :active-tab-key="activeKey"
    @tabChange="onTabChange"
    style="margin-bottom: 20px;height:320px"
    >
    <a-row :gutter="10">
      <a-col :span="8">
      <a-row style="height:32px;line-height:32px;">
        <a-col :span="5" style="text-align: right;padding-right: 10px;">CPU:</a-col>
        <a-col :span="16"><a-slider :min="0" :max="100" v-model="monitorInfo.cpu" @afterChange="updateMonitorRule"/></a-col>
      </a-row>
      <a-row style="height:32px;line-height:32px">
        <a-col :span="5" style="text-align: right;padding-right: 10px;">Mem:</a-col>
        <a-col :span="16"><a-slider :min="0" :max="100" v-model="monitorInfo.mem" @afterChange="updateMonitorRule"/></a-col>
      </a-row>
      <a-row v-show="this.activeKey==='node'" style="height:32px;line-height:32px">
        <a-col :span="5" style="text-align: right;padding-right: 10px;">Disk:</a-col>
        <a-col :span="16"><a-slider :min="0" :max="100" v-model="monitorInfo.disk" @afterChange="updateMonitorRule"/></a-col>
      </a-row>
      <a-row style="height:32px;line-height:32px">
        <a-col :span="5" style="text-align: right;padding-right: 10px;">触发间隔:</a-col>
        <a-col :span="18">
        <a-slider :min="2" :max="20" v-model="monitorInfo.interval" :tip-formatter="v => {return `${v}s`}"  @afterChange="updateMonitorRule"/>
        </a-col>
      </a-row>
      <a-row style="height:32px;line-height:32px">
        <a-col :span="5" style="text-align: right;padding-right: 10px;">报警间隔:</a-col>
        <a-col :span="18">
        <a-slider :min="600" :max="7200" :step="600" 
          v-model="monitorInfo.continuityInterval" 
          :tip-formatter="v => {return `${v/60}m`}"
          @afterChange="updateMonitorRule"/>
          </a-col>
      </a-row>
      </a-col>
      <a-col :span="16">
        <a-row style="height:40px;line-height:40px;">
          <a-col :span="5" style="text-align: right;padding-right: 10px;">通知类型:</a-col>
          <a-col :span="16">
            <a-select 
            :value="notify.notifyType" 
            @change="selectChange" 
            style="width: 120px">
              <a-select-option value="weixin">企业微信</a-select-option>
              <a-select-option value="callback">自定义</a-select-option>
            </a-select>
          </a-col>
        </a-row>
        <a-row style="height:40px;line-height:40px;">
          <a-col :span="5" style="text-align: right;padding-right: 10px;">挂载URL:</a-col>
          <a-col :span="16">
            <a-input 
            v-model="notify.notifyServer" 
            placeholder="http://example.com/posttreceive"/>
          </a-col>
        </a-row>
        <a-row style="height:32px;line-height:32px;">
          <a-col :span="5"></a-col>
          <a-col :span="16"><a-button  type="primary" @click="updateMonitorNotify">确认</a-button></a-col>
        </a-row>
      </a-col>
    </a-row>
    </a-card>

  </div>
</template>

<script>
import { monitorInfo,monitorRule,monitorNotify } from '@/api/monitor'
export default {
  data () {
    return {
      tabList:[
        {
          key:'node',
          tab:'节点',
        },
        {
          key:'process',
          tab:'程序',
        }
      ],
      activeKey:'node',
      monitorInfo:{
        cpu:90,
        mem:90,
        disk:90,
        interval:10,
        continuityInterval:600,
        notify:{
          notifyType:'weixin',
          notifyServer:''
        }
      },
      notify:{
        notifyType:'weixin',
        notifyServer:''
      }
    }
  },
  mounted () {
    this.loadMonitor()
  },
  methods:{
    onTabChange(key){
      this.activeKey = key
      this.loadMonitor()
    },
    selectChange(value){
      this.notify.notifyType = value
      if (this.monitorInfo.notify && value === this.monitorInfo.notify.notifyType){
        this.notify.notifyServer = this.monitorInfo.notify.notifyServer
      }else{
        this.notify.notifyServer = ''
      }
    },
    loadMonitor(){
      monitorInfo({type:this.activeKey}).then(res =>{
        this.monitorInfo = res
        if (this.monitorInfo.notify){
          this.notify = {
            notifyType:this.monitorInfo.notify.notifyType,
            notifyServer:this.monitorInfo.notify.notifyServer}
        }else{
          this.notify = {
            notifyType:'weixin',
            notifyServer:''
          }
        }
        // console.log(this.monitorInfo,this.notify);
      })
    },
    updateMonitorRule(){
      const args = {
          type:this.activeKey,
          cpu:this.monitorInfo.cpu,
          mem:this.monitorInfo.mem,
          disk:this.monitorInfo.disk,
          interval:this.monitorInfo.interval,
          continuityInterval:this.monitorInfo.continuityInterval,
      }
      monitorRule(args).then(()=>{
          this.loadMonitor()
        })
      
    },
    updateMonitorNotify(){
      const args = {
          type:this.activeKey,
          notifyType:this.notify.notifyType,
          notifyServer:this.notify.notifyServer,
      }
      monitorNotify(args).then(()=>{
          this.$message.info('操作成功')
          this.loadMonitor()
        })
    }
  }
}
</script>