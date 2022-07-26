<template>
  <div>
    <a-row :gutter="10">
      <a-col :span="8" >
        <a-card :bordered="false" style="height:200px">
          <a-row style="height:32px;line-height:32px;">
            <a-col :span="5">CPU:</a-col>
            <a-col :span="16"><a-slider :min="0" :max="100" v-model="form.cpu" @afterChange="updateMonitor"/></a-col>
          </a-row>
          <a-row style="height:32px;line-height:32px">
            <a-col :span="5">Mem:</a-col>
            <a-col :span="16"><a-slider :min="0" :max="100" v-model="form.mem" @afterChange="updateMonitor"/></a-col>
          </a-row>
          <a-row style="height:32px;line-height:32px">
            <a-col :span="5">触发间隔:</a-col>
            <a-col :span="18">
            <a-slider :min="2" :max="20" v-model="form.interval" :tip-formatter="v => {return `${v}s`}"  @afterChange="updateMonitor"/>
            </a-col>
          </a-row>
          <a-row style="height:32px;line-height:32px">
            <a-col :span="5">报警间隔:</a-col>
            <a-col :span="18">
            <a-slider :min="600" :max="7200" :step="600" 
              v-model="form.continuityInterval" 
              :tip-formatter="v => {return `${v/60}m`}"
              @afterChange="updateMonitor"/>
              </a-col>
          </a-row>
        </a-card>
      </a-col>
      <a-col :span="16">
        <a-card :bordered="false" >
        3434
        </a-card>
      </a-col>
    </a-row>
    
  </div>
</template>

<script>
import { processMonitor,processMonitorSet } from '@/api/process'
export default {
  data () {
    return {
      form:{
        cpu:90,
        mem:90,
        interval:10,
        continuityInterval:600,
      }
    }
  },
  mounted () {
    this.loadMonitor()
  },
  methods:{
    loadMonitor(){
      processMonitor().then(res =>{
        this.form = res
      })
    },
    updateMonitor(){
      console.log('ser');
      processMonitorSet(this.form).then(()=>{
        this.loadMonitor()
      })
    }
  }
}
</script>