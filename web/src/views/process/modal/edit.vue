<template>
  <a-modal
    :width="1000"
    :title="title"
    centered
    :maskClosable="false"
    :visible="visible"
    @ok="submitForm"
    @cancel="cancel"
  >
      <a-form-model
        ref="processEdit"
        :model="form"
        :rules="rules"
        :label-col="labelCol"
        :wrapper-col="wrapperCol"
      >
        <a-form-model-item ref="name" label="名称" prop="name" :wrapper-col="{ span: 6}">
          <a-input v-model="form.name" placeholder="程序名称"/>
        </a-form-model-item>
        <a-form-model-item label="执行目录" :wrapper-col="{ span: 6}">
          <a-input v-model="form.dir" placeholder="非必填，默认为节点启动目录" />
        </a-form-model-item>
        <a-form-model-item
          v-for="(cfg, index) in form.config"
          :key="index"
          :wrapper-col="index === 0 ? wrapperCol:{ span: 14,offset:3}"
          :label="index === 0 ? '启动配置' : ''"
        >
          <a-input
            v-model="cfg.name"
            placeholder="配置名称"
            style="width: 60%; margin-right: 8px"
          />
          <a-input
            v-model="cfg.context"
            placeholder="配置内容"
            style="width: 90%; margin-right: 8px"
            type="textarea"
            :auto-size="{ minRows: 6, maxRows: 16 }"
          />
          <a-icon
            v-if="form.config.length > 0"
            class="dynamic-delete-button"
            type="minus-circle-o"
            @click="removeCfg(cfg)"
          />
        </a-form-model-item>
        <a-form-model-item
          :label="form.config.length === 0 ? '启动配置' : ''"
          :wrapper-col="form.config.length === 0 ? wrapperCol:{ span: 14,offset:3}"
        >
          <a-button type="dashed" style="width:60%" @click="addCfg">
            <a-icon type="plus" /> 添加配置
          </a-button>
        </a-form-model-item>
        <a-form-model-item prop="command">
          <span slot="label">
            启动命令&nbsp;
            <a-tooltip :title="cmd_question_title">
              <a-icon type="question-circle-o" />
            </a-tooltip>
          </span>
          <a-input v-model="form.command" placeholder="启动命令"/>
        </a-form-model-item>
        <a-form-model-item :wrapper-col="{ span: 6}">
          <span slot="label">
            优先级&nbsp;
            <a-tooltip title="子进程启动关闭优先级，优先级低的，最先启动，关闭的时候最后关闭">
              <a-icon type="question-circle-o" />
            </a-tooltip>
          </span>
          <a-input v-model="form.priority"/>
        </a-form-model-item>
        <a-form-model-item :wrapper-col="{ span: 6}">
          <span slot="label">
            启动检测时长&nbsp;
            <a-tooltip title="启动进程一段时间后没有异常退出，就表示进程正常启动了。">
              <a-icon type="question-circle-o" />
            </a-tooltip>
          </span>
          <a-input v-model="form.startSecs"/>
        </a-form-model-item>
        <a-form-model-item :wrapper-col="{ span: 6}">
          <span slot="label">
            停机等待时长&nbsp;
            <a-tooltip title="这个是当我们向子进程发送stop信号后，到系统返回信息所等待的最大时间。超过这个时间会向该子进程发送一个强制kill的信号。">
              <a-icon type="question-circle-o" />
            </a-tooltip>
          </span>
          <a-input v-model="form.stopWaitSecs">
            <span slot="addonAfter" >秒</span>
          </a-input>
        </a-form-model-item>
        <a-form-model-item :wrapper-col="{ span: 6}">
          <span slot="label">
            自动重启次数&nbsp;
            <a-tooltip title="进程状态为 Exited时，自动重启">
              <a-icon type="question-circle-o" />
            </a-tooltip>
          </span>
          <a-input v-model="form.autoStartTimes"/>
        </a-form-model-item>
        <a-form-model-item :wrapper-col="{ span: 14}" prop="labels">
          <span slot="label">
            标签&nbsp;
            <a-tooltip title="程序标签，用于视图过滤">
              <a-icon type="question-circle-o" />
            </a-tooltip>
          </span>
          <template v-for="label in this.allLabels">
            <a-tooltip v-if="label.length > 10" :key="label" :title="label">
              <a-checkable-tag  
              :key="label" 
              style="border: 1px solid #eaeaea;"
              :checked="form.labels.indexOf(label) > -1" 
              @change="checked => handleLableChange(label, checked)"
              >
                {{ `${label.slice(0, 10)}...` }}
              </a-checkable-tag>
            </a-tooltip>
            <a-checkable-tag
              v-else
              :key="label"
              style="border: 1px solid #eaeaea;"
              :checked="form.labels.indexOf(label) > -1"
              @change="checked => handleLableChange(label, checked)"
            >
              {{ label }}
            </a-checkable-tag>
          </template>
          
          <a-input
            v-if="labelInputVisible"
            ref="labelInput"
            :style="{ width: '80px' }"
            size="small"
            v-model="labelInputValue"
            @blur="addLabel"
            @keyup.enter="addLabel"
          />
          <a-tag v-else style="background: #fff; borderStyle: dashed;" @click="showLabelInput">
            <a-icon type="plus" /> New Label
          </a-tag>
        </a-form-model-item>
        <a-form-model-item label="节点" prop="node" :wrapper-col="{ span: 6}">
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
import { processCreate, processUpdate } from '@/api/process'

export default {
  name: 'ProcessEdit',
  props: {
    visible: {
      type: Boolean,
      required: true
    },
    option: {
      type: String,
      default: ''
    },
    labels: {
      type: [],
      default: null
    },
    nodeNames:{
      type:[],
      default: null
    },
    model:{
      type: Object,
      default: null
    }
  },
  data () {
    return {
      labelCol: { span: 4 },
      wrapperCol: { span: 14 },
      cmd_question_title: '命令中存在配置文件时，路径前加上{{path}}，自动填充',
      rules: {
        name: [
          { required: true, message: '请输入名称', trigger: 'blur' },
          { min: 3, message: '名称至少3个字', trigger: 'blur' }
        ],
        command: [
          { required: true, message: '请输入启动命令', trigger: 'blur' }
        ],
        labels: [
          { validator: this.checkLabels, message: '至少设置一个标签', trigger: 'blur' }
        ],
        node: [
          { required: true, message: '选择节点', trigger: 'blur' }
        ]
      },
      form: {
        id: 0,
        name: '',
        dir: '',
        config: [],
        command: '',
        priority: 10,
        startSecs: 3,
        stopWaitSecs: 10,
        autoStartTimes: 3,
        node: '',
        labels: []
      },
      title:'',
      allLabels:[],
      labelInputVisible: false,
      labelInputValue: '',
    }
  },
  watch: {
    model: {
      handler (nv) {
        if (nv) { 
          this.form = { ...nv }
        } else { 
          this.form = {
            id: 0,
            name: '',
            dir: '',
            config: [],
            command: '',
            priority: 10,
            startSecs: 3,
            stopWaitSecs: 10,
            autoStartTimes: 3,
            node: '',
            labels: []
          }
        }
      },
      deep: true,
      immediate: true
    },
    option:{
      handler(nval){
        if (nval === 'edit') {
          this.title = '修改配置'
        }else{
          this.title = '新增配置'
        }
      }
    },
    visible(nv){
      if (nv){
        this.onVisible()
      }
    }
  },
  methods: {
    onVisible(){
      // console.log("visible");
      this.allLabels = this.labels
    },
    removeCfg (item) {
      console.log(item)
      const index = this.form.config.indexOf(item)
      if (index !== -1) {
        this.form.config.splice(index, 1)
      }
    },
    addCfg () {
      this.form.config.push({
        name: '',
        context: ''
      })
    },
    
    handleLableChange(label, checked){
      const nextSelectedTags = checked ? [...this.form.labels, label] : this.form.labels.filter(t => t !== label);
      this.form.labels = nextSelectedTags;
      console.log(this.allLabels,this.form.labels);
    },
    showLabelInput(){
      this.labelInputVisible = true
      this.$nextTick(function() {
        this.$refs.labelInput.focus();
      });
    },
    addLabel () {
      if (this.labelInputValue && this.form.labels.indexOf(this.labelInputValue) === -1) {
        console.log(this.labelInputValue)
        this.form.labels = [...this.form.labels, this.labelInputValue]
        this.allLabels = [...this.allLabels, this.labelInputValue]
      }
      console.log(this.allLabels,this.form.labels)
      this.labelInputVisible = false
      this.labelInputValue = ''
    },
    
    checkLabels () {
      if (this.form.labels.length > 0) {
        return true
      }
      return false
    },
    submitForm () {
       this.$refs.processEdit.validate(valid => {
        if (valid) {
          console.log(this.form)
          const args = { ...this.form }
          args.autoStartTimes = parseInt(this.form.autoStartTimes)
          args.priority = parseInt(this.form.priority)
          args.startSecs = parseInt(this.form.startSecs)
          args.stopWaitSecs = parseInt(this.form.stopWaitSecs)
          if (this.option === 'create' || this.option === 'copy') {
            processCreate(args).then(() => {
              this.$refs.processEdit.resetFields()// 清除表单数据
                this.$emit('success')// 通知外部页面 添加成功
            })
          } else {
            processUpdate(args).then(() => {
              this.$refs.processEdit.resetFields()// 清除表单数据
                this.$emit('success')// 通知外部页面 添加成功
            })
          }
        }
      })
    },
    cancel(){
      this.$refs.processEdit.resetFields()
      this.$emit('cancel')
    }
  }
}
</script>

<style scoped>

</style>
