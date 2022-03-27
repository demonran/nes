<template>
  <a-form-model
      id="components-form-demo-validate-other"
      :layout="layout"
      :form="form"
      v-bind="formItemLayout"
  >
    <a-form-model-item label="直播场次" required>
      <a-input-number v-model="form.liveCount" :min="0"/>
      <span class="ant-form-text">
        /月
      </span>
    </a-form-model-item>
    <a-form-model-item label="最大观看终端数" required>
      <a-input-number v-model="form.maxClientCount" :min="0"/>
      <a-tooltip title="只包含需要在公司内部接入直播的终端，如果无法评估就填公司总人数">
        <a-icon type="question-circle"/>
      </a-tooltip>
    </a-form-model-item>
    <a-form-model-item label="直播经验" required>
      <a-switch v-model="form.operationExperience"/>
      <a-tooltip title="直播工具的操作经验（如OBS）">
        <a-icon type="question-circle"/>
      </a-tooltip>
    </a-form-model-item>
    <a-form-model-item label="直播观看对象" required>
      <a-select v-model="form.typesOfViewer"
                mode="multiple"
                placeholder="请选择观看直播的人员"
      >
        <a-select-option v-for="(item,index) in viewers" :value="item.value" :key="index">
          {{ item.name }}
        </a-select-option>
      </a-select>
    </a-form-model-item>
    <a-form-model-item label="供应商支持">
      <a-switch v-model="form.vendorSupport"/>
      <a-tooltip title="是否有供应商支持">
        <a-icon type="question-circle"/>
      </a-tooltip>
    </a-form-model-item>
    <a-form-model-item label="供应商负责的模块" v-if="form.vendorSupport">
      <a-select v-model="form.typesOfViewer"
                mode="multiple"
                placeholder="请选择供应商负责的模块"
      >
        <a-select-option v-for="(item,index) in vendorModules" :value="item.value" :key="index">
          {{ item.name }}
        </a-select-option>
      </a-select>
    </a-form-model-item>
    <a-form-model-item label="外场舞台">
      <a-switch v-model="form.outfieldScene"/>
      <a-tooltip title="开播源是否存在外场舞台的情况">
        <a-icon type="question-circle"/>
      </a-tooltip>
    </a-form-model-item>
    <a-form-model-item label="开播设备">
      <a-select v-model="form.liveDevices"
                mode="multiple"
                placeholder="请选择开播设备类型"
      >
        <a-select-option v-for="(item,index) in deviceTypes" :value="item.value" :key="index">
          {{ item.name }}
        </a-select-option>
      </a-select>
    </a-form-model-item>
    <a-form-model-item label="收看方式">
      <a-select v-model="form.observationStyles"
                mode="multiple"
                placeholder="请选择收看方式"
      >
        <a-select-option v-for="(item,index) in watchWays" :value="item.value" :key="index">
          {{ item.name }}
        </a-select-option>
      </a-select>
    </a-form-model-item>
    <a-form-model-item :wrapper-col="buttonItemLayout.wrapperCol">
      <a-button type="primary" @click="handleSubmit">
        添加
      </a-button>
    </a-form-model-item>
  </a-form-model>
</template>

<script>
export default {
  data: () => ({
    form: {
      liveCount: undefined,
      maxClientCount: undefined,
      operationExperience: false,
      typesOfViewer: undefined,
      vendorSupport: false,
      outfieldScene: undefined,
      vendorSupportModules: undefined,
      liveDevices: undefined,
      observationStyles: undefined
    },
    layout: 'horizontal',
    viewers: [
      {"value": 1, "name": "内部员工"},
      {"value": 2, "name": "外部用户"},
      {"value": 3, "name": "海外员工"},
      {"value": 4, "name": "海外用户"},
    ],
    vendorModules: [
      {"value": "a", "name": "推流"},
      {"value": "b", "name": "线下音频采集"},
      {"value": "c", "name": "线下视频采集"},
      {"value": "d", "name": "推流网络搭建"}
    ],
    deviceTypes: [
      {"value": 1, "name": "Windows PC"},
      {"value": 2, "name": "Mac PC"},
      {"value": 3, "name": "Linux PC"},
      {"value": 4, "name": "摄像机"},
      {"value": 5, "name": "大屏幕"},
      {"value": 6, "name": "导播台"},
      {"value": 7, "name": "麦克风"},
      {"value": 8, "name": "调音台"},
      {"value": 9, "name": "音视频采集卡"}
    ],
    watchWays: [
      {"value": 1, "name": "电脑"},
      {"value": 2, "name": "手机"},
      {"value": 3, "name": "集中单输出源场景（会议室单台电视/投影）"},
      {"value": 4, "name": "集中多输出源场景（餐厅多台电视/投影）"}
    ]
  }),
  computed: {
    formItemLayout() {
      const {layout} = this.form;
      return layout === 'horizontal'
          ? {
            labelCol: {span: 4},
            wrapperCol: {span: 14},
          }
          : {};
    },
    buttonItemLayout() {
      const {layout} = this.form;
      return layout === 'horizontal'
          ? {
            wrapperCol: {span: 14, offset: 4},
          }
          : {};
    },
  },
  beforeCreate() {
    this.form = this.$form.createForm(this, {name: 'validate_other'});
  },
  methods: {
    handleSubmit() {
      console.log(this.form);
    },
  },
};
</script>
<style>
#components-form-demo-validate-other .dropbox {
  height: 180px;
  line-height: 1.5;
}
</style>