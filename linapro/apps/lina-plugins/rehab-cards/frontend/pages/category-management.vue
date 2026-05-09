<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import { Button, Card, Form, Input, InputNumber, message, Popconfirm, Select, Space, Table } from 'ant-design-vue';
import { categoryCreate, categoryDelete, categoryList, categoryUpdate, type RehabCategory } from './rehab-card-client';

const loading = ref(false);
const items = ref<RehabCategory[]>([]);
const editingId = ref<number | null>(null);
const form = reactive({ name: '', code: '', sort: 0, status: 1, remark: '' });

async function load() {
  loading.value = true;
  try { items.value = await categoryList(); } finally { loading.value = false; }
}

function reset() {
  editingId.value = null;
  Object.assign(form, { name: '', code: '', sort: 0, status: 1, remark: '' });
}

function edit(row: RehabCategory) {
  editingId.value = row.id;
  Object.assign(form, { name: row.name, code: row.code, sort: row.sort, status: row.status, remark: row.remark });
}

async function save() {
  if (editingId.value) await categoryUpdate(editingId.value, form);
  else await categoryCreate(form);
  message.success('已保存');
  reset();
  await load();
}

async function remove(id: number) {
  await categoryDelete(id);
  message.success('已删除');
  await load();
}

onMounted(load);
</script>

<template>
  <Card title="图卡分类管理">
    <Form layout="inline" style="margin-bottom: 16px">
      <Form.Item label="名称"><Input v-model:value="form.name" placeholder="动物" /></Form.Item>
      <Form.Item label="编码"><Input v-model:value="form.code" placeholder="animal" /></Form.Item>
      <Form.Item label="排序"><InputNumber v-model:value="form.sort" /></Form.Item>
      <Form.Item label="状态"><Select v-model:value="form.status" style="width: 100px" :options="[{label:'启用',value:1},{label:'停用',value:0}]" /></Form.Item>
      <Form.Item><Button type="primary" @click="save">{{ editingId ? '更新' : '新增' }}</Button></Form.Item>
      <Form.Item><Button @click="reset">清空</Button></Form.Item>
    </Form>
    <Table :loading="loading" :data-source="items" row-key="id" :pagination="false">
      <Table.Column title="名称" data-index="name" />
      <Table.Column title="编码" data-index="code" />
      <Table.Column title="排序" data-index="sort" />
      <Table.Column title="状态" data-index="status" />
      <Table.Column title="操作">
        <template #default="{ record }"><Space><Button size="small" @click="edit(record)">编辑</Button><Popconfirm title="确认删除？" @confirm="remove(record.id)"><Button size="small" danger>删除</Button></Popconfirm></Space></template>
      </Table.Column>
    </Table>
  </Card>
</template>
