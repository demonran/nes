<template>
  <div>
    <a-button icon="plus" style="background: green">增加</a-button>
    <a-table :columns="columns" :data-source="data">
      <span slot="action" >
        <a-button type="primary" icon="edit"/>
        <a-button type="danger" icon="delete"/>
    </span>
    </a-table>
  </div>
</template>

<script>
const columns = [
  {
    dataIndex: 'ID',
    key: 'ID',
    title: 'ID',
    scopedSlots: { customRender: 'name' },
  },
  {
    title: '公司名称',
    dataIndex: 'company',
    key: 'company',
  },
  {
    title: '客户名称',
    dataIndex: 'name',
    key: 'name',
  },
  {
    title: '手机号',
    key: 'mobile',
    dataIndex: 'mobile',

  },
  {
    title: '操作',
    key: 'action',
    scopedSlots: { customRender: 'action' },
  },
];

import CustomerApi from '@/services/customer.js'
export default {
  name: "Customer",
  data() {
    return {
      data: [],
      columns,
    };
  },

  created() {
    CustomerApi.list().then(res => {
      this.data = res.data
    })
  }
};
</script>

<style scoped>

</style>
