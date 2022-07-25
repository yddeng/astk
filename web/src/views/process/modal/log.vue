<template>
  <a-modal
    :width="1000"
    :title="title"
    centered
    footer=""
    :maskClosable="false"
    :visible="visible"
    @cancel="cancel"
  >
        <div>
          <a-textarea
            :value="context"
            :auto-size="{ minRows: 10, maxRows: 20 }"
            style="background: #0e0e0e; color: white" />
        </div>
  </a-modal>
</template>

<script>
export default {
  name: 'Log',
  props: {
    visible: {
      type: Boolean,
      required: true
    },
    title: {
      type: String,
      default: '程序日志'
    },
    context: {
      type: String,
      default: 'null'
    },
  },
  data () {
    return {
    }
  },
  methods: {
    cancel () {
      this.$emit('cancel')
    }
  }
}
</script>
