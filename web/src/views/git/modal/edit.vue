<template>
  <a-modal
    :width="600"
    title="新建仓库"
    centered
    :maskClosable="false"
    :visible="visible"
    @ok="submitForm"
    @cancel="cancel"
  >
      <a-form-model
        ref="gitEdit"
        :model="form"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-model-item label="仓库类型" >
          <a-select v-model="form.type" style="width: 100px">
            <a-select-option value="github">github</a-select-option>
            <a-select-option value="gitlab">gitlab</a-select-option>
          </a-select>
        </a-form-model-item>
        <a-form-model-item label="项目名称" >
          <a-input style="width: 200px" v-model="form.name" />
        </a-form-model-item>
        <a-form-model-item label="仓库地址" >
          <a-input style="width: 400px" v-model="form.address" />
        </a-form-model-item>
        <a-form-model-item label="分支" >
          <a-input style="width: 100px" v-model="form.branch" />
        </a-form-model-item>
        <a-form-model-item label="Hook令牌" >
          <a-input style="width: 200px" v-model="form.token" />
        </a-form-model-item>
        <a-form-model-item label="通知方式">
          <a-select v-model="form.notify.notifyType" style="width:120px">
            <a-select-option value="weixin">企业微信</a-select-option>
            <a-select-option value="callback">自定义</a-select-option>
          </a-select>
          <a-input style="width: 400px" 
            placeholder="http://example.com/posttreceive"
            v-model="form.notify.notifyServer" />
          </a-form-model-item>
      </a-form-model>
  </a-modal>
</template>

<script>
import { gitCreate } from '@/api/git'

export default {
  name: 'GitEdit',
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
        type: 'github',
        name: '',
        address:'',
        branch: '',
        token: '',
        notify: {
          notifyType: 'weixin',
          notifyServer:'',
        },
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
        type: 'github',
        name: '',
        address:'',
        branch: '',
        token: '',
        notify: {
          notifyType: 'weixin',
          notifyServer:'',
        },
      }
    },
    submitForm () {
      console.log(this.form)
      gitCreate(this.form).then(() => {
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
