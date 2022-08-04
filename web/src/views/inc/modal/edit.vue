<template>
  <a-modal
    :width="600"
    title="新建代理"
    centered
    :maskClosable="false"
    :visible="visible"
    @ok="submitForm"
    @cancel="cancel"
  >
      <a-form-model
        ref="incEdit"
        :model="form"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-model-item label="类型" >
          <a-select v-model="form.type" style="width: 80px">
            <a-select-option value="tcp">tcp</a-select-option>
            <a-select-option value="http">http</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="转发端口" >
          <a-input style="width: 80px" v-model="form.remotePort" />
        </a-form-model-item>
        <a-form-model-item label="描述" >
          <a-input style="width: 400px" v-model="form.desc" />
        </a-form-model-item>
        <a-form-model-item label="目标地址">
          <a-input-group compact>
            <a-input style="width: 200px" v-model="form.localIp" />
            <a-input style="width: 80px" v-model="form.localPort" />
          </a-input-group>
        </a-form-model-item>
        <a-form-model-item label="节点" :wrapper-col="{ span: 6}">
          <a-select placeholder="选择一个节点" v-model="form.node" >
            <template v-for="v in nodeNames" >
              <a-select-option :value="v" :key="v">
                {{ v }}
              </a-select-option>
            </template>
          </a-select>
        </a-form-model-item>
      </a-form-model>
  </a-modal>
</template>

<script>
import { incCreate } from '@/api/inc'

export default {
  name: 'ProcessEdit',
  props: {
    visible: {
      type: Boolean,
      required: true
    },
    nodeNames:{
      type:[],
      default: null
    }
  },
  data () {
    return {
      labelCol: { span: 5 },
      wrapperCol: { span: 18 },
      form: {
        type: 'tcp',
        remotePort: '',
        desc:'',
        localIp: '',
        localPort: '',
        node: '',
      },
    }
  },
  watch:{
    visible(nv){
      if (nv){
        this.onVisible()
      }
    }
  },
  methods: {
    onVisible(){
      this.form = {
        type: 'tcp',
        remotePort: '',
        desc:'',
        localIp: '',
        localPort: '',
        node: '',
      }
    },
    submitForm () {
      console.log(this.form)
      if (this.form.remotePort ==='' ||
      this.form.localIp === '' ||
      this.form.localPort === '' ||
      this.form.node === '') {
          this.$message.info('表单请填写完整')
          return
      }
      incCreate(this.form).then(() => {
        this.$emit('success')// 通知外部页面 添加成功
      })
    },
    cancel(){
      this.$emit('cancel')
    }
  }
}
</script>

<style scoped>

</style>
